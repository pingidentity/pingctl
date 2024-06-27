package platform_test

import (
	"testing"

	"github.com/pingidentity/pingctl/internal/connector/pingone/platform/resources"
	"github.com/pingidentity/pingctl/internal/testutils/testutils_helpers"
)

// Test --generate-config-out for the Agreement resource
func TestAgreementTerraformPlan(t *testing.T) {
	sdkClientInfo := testutils_helpers.GetPingOneSDKClientInfo(t)
	agreementResource := resources.Agreement(sdkClientInfo)
	testutils_helpers.ValidateTerraformPlan(t, agreementResource, nil)
}

// Test --generate-config-out for the AgreementEnable resource
func TestAgreementEnableTerraformPlan(t *testing.T) {
	sdkClientInfo := testutils_helpers.GetPingOneSDKClientInfo(t)
	agreementEnableResource := resources.AgreementEnable(sdkClientInfo)
	testutils_helpers.ValidateTerraformPlan(t, agreementEnableResource, nil)
}

// Test --generate-config-out for the AgreementLocalization resource
func TestAgreementLocalizationTerraformPlan(t *testing.T) {
	sdkClientInfo := testutils_helpers.GetPingOneSDKClientInfo(t)
	agreementLocalizationResource := resources.AgreementLocalization(sdkClientInfo)
	testutils_helpers.ValidateTerraformPlan(t, agreementLocalizationResource, nil)
}

// Test --generate-config-out for the AgreementLocalizationEnable resource
func TestAgreementLocalizationEnableTerraformPlan(t *testing.T) {
	sdkClientInfo := testutils_helpers.GetPingOneSDKClientInfo(t)
	agreementLocalizationEnableResource := resources.AgreementLocalizationEnable(sdkClientInfo)
	testutils_helpers.ValidateTerraformPlan(t, agreementLocalizationEnableResource, nil)
}

// Test --generate-config-out for the AgreementLocalizationRevision resource
func TestAgreementLocalizationRevisionTerraformPlan(t *testing.T) {
	sdkClientInfo := testutils_helpers.GetPingOneSDKClientInfo(t)
	agreementLocalizationRevisionResource := resources.AgreementLocalizationRevision(sdkClientInfo)
	testutils_helpers.ValidateTerraformPlan(t, agreementLocalizationRevisionResource, nil)
}

// Test --generate-config-out for the BrandingSettings resource
func TestBrandingSettingsTerraformPlan(t *testing.T) {
	sdkClientInfo := testutils_helpers.GetPingOneSDKClientInfo(t)
	brandingSettingsResource := resources.BrandingSettings(sdkClientInfo)
	testutils_helpers.ValidateTerraformPlan(t, brandingSettingsResource, nil)
}

// Test --generate-config-out for the BrandingTheme resource
func TestBrandingThemeTerraformPlan(t *testing.T) {
	sdkClientInfo := testutils_helpers.GetPingOneSDKClientInfo(t)
	brandingThemeResource := resources.BrandingTheme(sdkClientInfo)
	testutils_helpers.ValidateTerraformPlan(t, brandingThemeResource, nil)
}

// Test --generate-config-out for the BrandingThemeDefault resource
func TestBrandingThemeDefaultTerraformPlan(t *testing.T) {
	sdkClientInfo := testutils_helpers.GetPingOneSDKClientInfo(t)
	brandingThemeDefaultResource := resources.BrandingThemeDefault(sdkClientInfo)
	testutils_helpers.ValidateTerraformPlan(t, brandingThemeDefaultResource, nil)
}

// Test --generate-config-out for the Certificate resource
func TestCertificateTerraformPlan(t *testing.T) {
	sdkClientInfo := testutils_helpers.GetPingOneSDKClientInfo(t)
	certificateResource := resources.Certificate(sdkClientInfo)
	ignoreErrors := []string{
		"Error: Invalid combination of arguments",
	}
	testutils_helpers.ValidateTerraformPlan(t, certificateResource, ignoreErrors)
}

// Test --generate-config-out for the CustomDomain resource
func TestCustomDomainTerraformPlan(t *testing.T) {
	sdkClientInfo := testutils_helpers.GetPingOneSDKClientInfo(t)
	customDomainResource := resources.CustomDomain(sdkClientInfo)
	testutils_helpers.ValidateTerraformPlan(t, customDomainResource, nil)
}

// Test --generate-config-out for the Environment resource
func TestEnvironmentTerraformPlan(t *testing.T) {
	sdkClientInfo := testutils_helpers.GetPingOneSDKClientInfo(t)
	environmentResource := resources.Environment(sdkClientInfo)
	testutils_helpers.ValidateTerraformPlan(t, environmentResource, nil)
}

// Test --generate-config-out for the Form resource
func TestFormTerraformPlan(t *testing.T) {
	sdkClientInfo := testutils_helpers.GetPingOneSDKClientInfo(t)
	formResource := resources.Form(sdkClientInfo)
	testutils_helpers.ValidateTerraformPlan(t, formResource, nil)
}

// Test --generate-config-out for the FormRecaptchaV2 resource
func TestFormRecaptchaV2TerraformPlan(t *testing.T) {
	sdkClientInfo := testutils_helpers.GetPingOneSDKClientInfo(t)
	formRecaptchaV2Resource := resources.FormRecaptchaV2(sdkClientInfo)
	ignoreErrors := []string{
		"Error: Missing Configuration for Required Attribute",
	}
	testutils_helpers.ValidateTerraformPlan(t, formRecaptchaV2Resource, ignoreErrors)
}

// Test --generate-config-out for the Gateway resource
func TestGatewayTerraformPlan(t *testing.T) {
	sdkClientInfo := testutils_helpers.GetPingOneSDKClientInfo(t)
	gatewayResource := resources.Gateway(sdkClientInfo)
	ignoreErrors := []string{
		"Error: Invalid Attribute Combination",
		"Error: Missing required argument",
	}
	testutils_helpers.ValidateTerraformPlan(t, gatewayResource, ignoreErrors)
}

// Test --generate-config-out for the GatewayCredential resource
func TestGatewayCredentialTerraformPlan(t *testing.T) {
	sdkClientInfo := testutils_helpers.GetPingOneSDKClientInfo(t)
	gatewayCredentialResource := resources.GatewayCredential(sdkClientInfo)
	testutils_helpers.ValidateTerraformPlan(t, gatewayCredentialResource, nil)
}

// Test --generate-config-out for the GatewayRoleAssignment resource
func TestGatewayRoleAssignmentTerraformPlan(t *testing.T) {
	sdkClientInfo := testutils_helpers.GetPingOneSDKClientInfo(t)
	gatewayRoleAssignmentResource := resources.GatewayRoleAssignment(sdkClientInfo)
	testutils_helpers.ValidateTerraformPlan(t, gatewayRoleAssignmentResource, nil)
}

// Test --generate-config-out for the IdentityPropagationPlan resource
func TestIdentityPropagationPlanTerraformPlan(t *testing.T) {
	sdkClientInfo := testutils_helpers.GetPingOneSDKClientInfo(t)
	identityPropagationPlanResource := resources.IdentityPropagationPlan(sdkClientInfo)
	testutils_helpers.ValidateTerraformPlan(t, identityPropagationPlanResource, nil)
}

// Test --generate-config-out for the Key resource
func TestKeyTerraformPlan(t *testing.T) {
	sdkClientInfo := testutils_helpers.GetPingOneSDKClientInfo(t)
	keyResource := resources.Key(sdkClientInfo)
	testutils_helpers.ValidateTerraformPlan(t, keyResource, nil)
}

// Test --generate-config-out for the KeyRotationPolicy resource
func TestKeyRotationPolicyTerraformPlan(t *testing.T) {
	sdkClientInfo := testutils_helpers.GetPingOneSDKClientInfo(t)
	keyRotationPolicyResource := resources.KeyRotationPolicy(sdkClientInfo)
	testutils_helpers.ValidateTerraformPlan(t, keyRotationPolicyResource, nil)
}

// Test --generate-config-out for the Language resource
func TestLanguageTerraformPlan(t *testing.T) {
	sdkClientInfo := testutils_helpers.GetPingOneSDKClientInfo(t)
	languageResource := resources.Language(sdkClientInfo)
	testutils_helpers.ValidateTerraformPlan(t, languageResource, nil)
}

// Test --generate-config-out for the LanguageUpdate resource
func TestLanguageUpdateTerraformPlan(t *testing.T) {
	sdkClientInfo := testutils_helpers.GetPingOneSDKClientInfo(t)
	languageUpdateResource := resources.LanguageUpdate(sdkClientInfo)
	testutils_helpers.ValidateTerraformPlan(t, languageUpdateResource, nil)
}

// Test --generate-config-out for the NotificationPolicy resource
func TestNotificationPolicyTerraformPlan(t *testing.T) {
	sdkClientInfo := testutils_helpers.GetPingOneSDKClientInfo(t)
	notificationPolicyResource := resources.NotificationPolicy(sdkClientInfo)
	testutils_helpers.ValidateTerraformPlan(t, notificationPolicyResource, nil)
}

// Test --generate-config-out for the NotificationSettings resource
func TestNotificationSettingsTerraformPlan(t *testing.T) {
	sdkClientInfo := testutils_helpers.GetPingOneSDKClientInfo(t)
	notificationSettingsResource := resources.NotificationSettings(sdkClientInfo)
	testutils_helpers.ValidateTerraformPlan(t, notificationSettingsResource, nil)
}

// Test --generate-config-out for the NotificationSettingsEmail resource
func TestNotificationSettingsEmailTerraformPlan(t *testing.T) {
	sdkClientInfo := testutils_helpers.GetPingOneSDKClientInfo(t)
	notificationSettingsEmailResource := resources.NotificationSettingsEmail(sdkClientInfo)
	ignoreErrors := []string{
		"Error: Missing Configuration for Required Attribute",
	}
	testutils_helpers.ValidateTerraformPlan(t, notificationSettingsEmailResource, ignoreErrors)
}

// Test --generate-config-out for the NotificationTemplateContent resource
func TestNotificationTemplateContentTerraformPlan(t *testing.T) {
	sdkClientInfo := testutils_helpers.GetPingOneSDKClientInfo(t)
	notificationTemplateContentResource := resources.NotificationTemplateContent(sdkClientInfo)
	testutils_helpers.ValidateTerraformPlan(t, notificationTemplateContentResource, nil)
}

// Test --generate-config-out for the PhoneDeliverySettings resource
func TestPhoneDeliverySettingsTerraformPlan(t *testing.T) {
	sdkClientInfo := testutils_helpers.GetPingOneSDKClientInfo(t)
	phoneDeliverySettingsResource := resources.PhoneDeliverySettings(sdkClientInfo)
	ignoreErrors := []string{
		"Error: Missing required argument",
		"Error: Missing Configuration for Required Attribute",
	}
	testutils_helpers.ValidateTerraformPlan(t, phoneDeliverySettingsResource, ignoreErrors)
}

// Test --generate-config-out for the SystemApplication resource
func TestSystemApplicationTerraformPlan(t *testing.T) {
	sdkClientInfo := testutils_helpers.GetPingOneSDKClientInfo(t)
	systemApplicationResource := resources.SystemApplication(sdkClientInfo)
	testutils_helpers.ValidateTerraformPlan(t, systemApplicationResource, nil)
}

// Test --generate-config-out for the TrustedEmailAddress resource
func TestTrustedEmailAddressTerraformPlan(t *testing.T) {
	sdkClientInfo := testutils_helpers.GetPingOneSDKClientInfo(t)
	trustedEmailAddressResource := resources.TrustedEmailAddress(sdkClientInfo)
	testutils_helpers.ValidateTerraformPlan(t, trustedEmailAddressResource, nil)
}

// Test --generate-config-out for the TrustedEmailDomain resource
func TestTrustedEmailDomainTerraformPlan(t *testing.T) {
	sdkClientInfo := testutils_helpers.GetPingOneSDKClientInfo(t)
	trustedEmailDomainResource := resources.TrustedEmailDomain(sdkClientInfo)
	testutils_helpers.ValidateTerraformPlan(t, trustedEmailDomainResource, nil)
}

// Test --generate-config-out for the UserRoleAssignment resource
func TestUserRoleAssignmentTerraformPlan(t *testing.T) {
	sdkClientInfo := testutils_helpers.GetPingOneSDKClientInfo(t)
	userRoleAssignmentResource := resources.UserRoleAssignment(sdkClientInfo)
	testutils_helpers.ValidateTerraformPlan(t, userRoleAssignmentResource, nil)
}

// Test --generate-config-out for the Webhook resource
func TestWebhookTerraformPlan(t *testing.T) {
	sdkClientInfo := testutils_helpers.GetPingOneSDKClientInfo(t)
	webhookResource := resources.Webhook(sdkClientInfo)
	testutils_helpers.ValidateTerraformPlan(t, webhookResource, nil)
}
