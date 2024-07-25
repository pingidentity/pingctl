package platform_test

import (
	"testing"

	"github.com/pingidentity/pingctl/internal/connector"
	"github.com/pingidentity/pingctl/internal/connector/pingone/platform/resources"
	"github.com/pingidentity/pingctl/internal/testing/testutils"
	"github.com/pingidentity/pingctl/internal/testing/testutils_terraform"
)

func TestPlatformTerraformPlan(t *testing.T) {
	sdkClientInfo := testutils.GetPingOneSDKClientInfo(t)

	testutils_terraform.InitTerraform(t)

	testCases := []struct {
		name          string
		resource      connector.ExportableResource
		ignoredErrors []string
	}{
		{
			name:          "Agreement",
			resource:      resources.Agreement(sdkClientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "AgreementEnable",
			resource:      resources.AgreementEnable(sdkClientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "AgreementLocalization",
			resource:      resources.AgreementLocalization(sdkClientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "AgreementLocalizationEnable",
			resource:      resources.AgreementLocalizationEnable(sdkClientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "AgreementLocalizationRevision",
			resource:      resources.AgreementLocalizationRevision(sdkClientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "BrandingSettings",
			resource:      resources.BrandingSettings(sdkClientInfo),
			ignoredErrors: nil,
		},
		{
			name:     "BrandingTheme",
			resource: resources.BrandingTheme(sdkClientInfo),
			ignoredErrors: []string{
				"Error: Invalid Attribute Combination",
			},
		},
		{
			name:          "BrandingThemeDefault",
			resource:      resources.BrandingThemeDefault(sdkClientInfo),
			ignoredErrors: nil,
		},
		{
			name:     "Certificate",
			resource: resources.Certificate(sdkClientInfo),
			ignoredErrors: []string{
				"Error: Invalid combination of arguments",
			},
		},
		{
			name:          "CustomDomain",
			resource:      resources.CustomDomain(sdkClientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "Environment",
			resource:      resources.Environment(sdkClientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "Form",
			resource:      resources.Form(sdkClientInfo),
			ignoredErrors: nil,
		},
		{
			name:     "FormRecaptchaV2",
			resource: resources.FormRecaptchaV2(sdkClientInfo),
			ignoredErrors: []string{
				"Error: Missing Configuration for Required Attribute",
			},
		},
		{
			name:     "Gateway",
			resource: resources.Gateway(sdkClientInfo),
			ignoredErrors: []string{
				"Error: Invalid Attribute Combination",
				"Error: Missing required argument",
			},
		},
		{
			name:          "GatewayCredential",
			resource:      resources.GatewayCredential(sdkClientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "GatewayRoleAssignment",
			resource:      resources.GatewayRoleAssignment(sdkClientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "IdentityPropagationPlan",
			resource:      resources.IdentityPropagationPlan(sdkClientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "Key",
			resource:      resources.Key(sdkClientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "KeyRotationPolicy",
			resource:      resources.KeyRotationPolicy(sdkClientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "Language",
			resource:      resources.Language(sdkClientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "LanguageUpdate",
			resource:      resources.LanguageUpdate(sdkClientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "NotificationPolicy",
			resource:      resources.NotificationPolicy(sdkClientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "NotificationSettings",
			resource:      resources.NotificationSettings(sdkClientInfo),
			ignoredErrors: nil,
		},
		{
			name:     "NotificationSettingsEmail",
			resource: resources.NotificationSettingsEmail(sdkClientInfo),
			ignoredErrors: []string{
				"Error: Missing Configuration for Required Attribute",
			},
		},
		{
			name:          "NotificationTemplateContent",
			resource:      resources.NotificationTemplateContent(sdkClientInfo),
			ignoredErrors: nil,
		},
		{
			name:     "PhoneDeliverySettings",
			resource: resources.PhoneDeliverySettings(sdkClientInfo),
			ignoredErrors: []string{
				"Error: Missing required argument",
				"Error: Missing Configuration for Required Attribute",
			},
		},
		{
			name:          "SystemApplication",
			resource:      resources.SystemApplication(sdkClientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "TrustedEmailAddress",
			resource:      resources.TrustedEmailAddress(sdkClientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "TrustedEmailDomain",
			resource:      resources.TrustedEmailDomain(sdkClientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "Webhook",
			resource:      resources.Webhook(sdkClientInfo),
			ignoredErrors: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			testutils_terraform.ValidateTerraformPlan(t, tc.resource, tc.ignoredErrors)
		})
	}
}
