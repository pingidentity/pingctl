package common

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"

	"github.com/pingidentity/pingctl/internal/connector"
	"github.com/pingidentity/pingctl/internal/logger"
)

func WriteFiles(exportableResources []connector.ExportableResource, format, outputDir string, service string, overwriteExport bool) error {
	l := logger.Get()

	// Check if the output directory is empty
	// If not, default behavior is to exit and not overwrite.
	// This can be changed with the --overwrite export parameter
	if !overwriteExport {
		dirEntries, err := os.ReadDir(outputDir)
		if err != nil {
			return fmt.Errorf("failed to read export directory %q. err: %s", outputDir, err.Error())
		}

		if len(dirEntries) > 0 {
			return fmt.Errorf("export directory %q is not empty. Use --overwrite to overwrite existing export data", outputDir)
		}
	}

	// Make subdirectory for exported service if there are resources to export
	if len(exportableResources) > 0 {
		err := os.MkdirAll(filepath.Join(outputDir, service), os.ModePerm)
		if err != nil {
			return fmt.Errorf("failed to create %s export subfolder. err: %s", service, err.Error())
		}
	} else {
		return nil
	}

	// Parse the HCL import block template
	hclImportBlockTemplate, err := template.New("HCLImportBlock").Parse(connector.HCLImportBlockTemplate)
	if err != nil {
		return fmt.Errorf("failed to parse HCL import block template. err: %s", err.Error())
	}

	for _, exportableResource := range exportableResources {
		importBlocks, err := exportableResource.ExportAll()
		if err != nil {
			return fmt.Errorf("failed to export resource %s. err: %s", exportableResource.ResourceType(), err.Error())
		}

		if len(*importBlocks) == 0 {
			// No resources exported. Avoid creating an empty import.tf file
			l.Debug().Msgf("Nothing exported for resource %s. Skipping import file generation...", exportableResource.ResourceType())
			continue
		}

		l.Debug().Msgf("Generating import file for %s resource...", exportableResource.ResourceType())

		outputFileName := fmt.Sprintf("%s.tf", exportableResource.ResourceType())
		outputFilePath := filepath.Join(outputDir, service, filepath.Base(outputFileName))

		// Check to see if outputFile already exists.
		// If so, default behavior is to exit and not overwrite.
		// This can be changed with the --overwrite export parameter
		_, err = os.Stat(outputFilePath)
		if err == nil && !overwriteExport {
			return fmt.Errorf("generated import file for %q already exists. Use --overwrite to overwrite existing export data", outputFileName)
		}

		outputFile, err := os.Create(outputFilePath)
		if err != nil {
			return fmt.Errorf("failed to create export file %q. err: %s", outputFilePath, err.Error())
		}
		defer outputFile.Close()

		for _, importBlock := range *importBlocks {
			// Sanitize import block "to". Make lowercase, remove special chars, convert space to underscore
			importBlock.Sanitize()

			switch format {
			case connector.ENUMEXPORTFORMAT_HCL:
				err := hclImportBlockTemplate.Execute(outputFile, importBlock)
				if err != nil {
					return fmt.Errorf("failed to write import block template to file %q. err: %s", outputFilePath, err.Error())
				}
				// default:
				// Note that this default case is already handled in export.go, and should never be called.
			}
		}
	}
	return nil
}
