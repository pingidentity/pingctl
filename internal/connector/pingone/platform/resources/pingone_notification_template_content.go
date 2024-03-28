package resources

import (
	"fmt"

	"github.com/patrickcping/pingone-go-sdk-v2/management"
	"github.com/pingidentity/pingctl/internal/connector"
	"github.com/pingidentity/pingctl/internal/connector/common"
	"github.com/pingidentity/pingctl/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingoneNotificationTemplateContentResource{}
)

type PingoneNotificationTemplateContentResource struct {
	clientInfo *connector.SDKClientInfo
}

// Utility method for creating a PingoneNotificationTemplateContentResource
func NotificationTemplateContent(clientInfo *connector.SDKClientInfo) *PingoneNotificationTemplateContentResource {
	return &PingoneNotificationTemplateContentResource{
		clientInfo: clientInfo,
	}
}

func (r *PingoneNotificationTemplateContentResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()

	l.Debug().Msgf("Fetching all %s resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	l.Debug().Msgf("Generating Import Blocks for all %s resources...", r.ResourceType())

	// This is weird... the provider mentions many possible template types,
	// but pingone console and the API only support the following types:
	validTemplateNames := []management.EnumTemplateName{
		management.ENUMTEMPLATENAME_DEVICE_PAIRING,
		management.ENUMTEMPLATENAME_EMAIL_VERIFICATION_ADMIN,
		management.ENUMTEMPLATENAME_EMAIL_VERIFICATION_USER,
		management.ENUMTEMPLATENAME_GENERAL,
		management.ENUMTEMPLATENAME_NEW_DEVICE_PAIRED,
		management.ENUMTEMPLATENAME_STRONG_AUTHENTICATION,
		management.ENUMTEMPLATENAME_TRANSACTION,
		management.ENUMTEMPLATENAME_VERIFICATION_CODE_TEMPLATE,
	}

	// TODO: When the above hard-coded values are fixed, use the following
	// for _, templateNameEnum := range management.AllowedEnumTemplateNameEnumValues {
	for _, templateNameEnum := range validTemplateNames {
		apiExecuteFunc := r.clientInfo.ApiClient.ManagementAPIClient.NotificationsTemplatesApi.ReadAllTemplateContents(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID, templateNameEnum).Execute
		apiFunctionName := "ReadAllTemplateContents"

		embedded, err := common.GetManagementEmbedded(apiExecuteFunc, apiFunctionName, r.ResourceType())
		if err != nil {
			return nil, err
		}

		for _, templateContents := range embedded.GetContents() {
			var (
				templateContentsId       *string
				templateContentsIdOk     bool
				templateDeliveryMethod   *management.EnumTemplateContentDeliveryMethod
				templateDeliveryMethodOk bool
				templateLocale           *string
				templateLocaleOk         bool
				templateVariant          *string
				templateVariantOk        bool
			)

			switch {
			case templateContents.TemplateContentEmail != nil:
				templateContentsId, templateContentsIdOk = templateContents.TemplateContentEmail.GetIdOk()
				templateDeliveryMethod, templateDeliveryMethodOk = templateContents.TemplateContentEmail.GetDeliveryMethodOk()
				templateLocale, templateLocaleOk = templateContents.TemplateContentEmail.GetLocaleOk()
				templateVariant, templateVariantOk = templateContents.TemplateContentEmail.GetVariantOk()
			case templateContents.TemplateContentPush != nil:
				templateContentsId, templateContentsIdOk = templateContents.TemplateContentPush.GetIdOk()
				templateDeliveryMethod, templateDeliveryMethodOk = templateContents.TemplateContentPush.GetDeliveryMethodOk()
				templateLocale, templateLocaleOk = templateContents.TemplateContentPush.GetLocaleOk()
				templateVariant, templateVariantOk = templateContents.TemplateContentPush.GetVariantOk()
			case templateContents.TemplateContentSMS != nil:
				templateContentsId, templateContentsIdOk = templateContents.TemplateContentSMS.GetIdOk()
				templateDeliveryMethod, templateDeliveryMethodOk = templateContents.TemplateContentSMS.GetDeliveryMethodOk()
				templateLocale, templateLocaleOk = templateContents.TemplateContentSMS.GetLocaleOk()
				templateVariant, templateVariantOk = templateContents.TemplateContentSMS.GetVariantOk()
			case templateContents.TemplateContentVoice != nil:
				templateContentsId, templateContentsIdOk = templateContents.TemplateContentVoice.GetIdOk()
				templateDeliveryMethod, templateDeliveryMethodOk = templateContents.TemplateContentVoice.GetDeliveryMethodOk()
				templateLocale, templateLocaleOk = templateContents.TemplateContentVoice.GetLocaleOk()
				templateVariant, templateVariantOk = templateContents.TemplateContentVoice.GetVariantOk()
			default:
				continue
			}

			// This variable handles the case where template type, locale,
			// and delivery method are the same across two content instances
			// Append it to the ResourceName if present from SDK
			if templateVariantOk {
				*templateVariant = "_" + *templateVariant
			} else {
				emptyString := ""
				templateVariant = &emptyString
			}

			if templateContentsIdOk && templateDeliveryMethodOk && templateLocaleOk {
				importBlocks = append(importBlocks, connector.ImportBlock{
					ResourceType: r.ResourceType(),
					ResourceName: fmt.Sprintf("%s_%s_%s%s", templateNameEnum, *templateDeliveryMethod, *templateLocale, *templateVariant),
					ResourceID:   fmt.Sprintf("%s/%s/%s", r.clientInfo.ExportEnvironmentID, templateNameEnum, *templateContentsId),
				})
			}
		}

	}

	return &importBlocks, nil
}

func (r *PingoneNotificationTemplateContentResource) ResourceType() string {
	return "pingone_notification_template_content"
}
