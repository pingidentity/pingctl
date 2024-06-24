package platform_test

import (
	"testing"

	"github.com/pingidentity/pingctl/internal/connector/pingone/platform/resources"
	"github.com/pingidentity/pingctl/internal/testutils/testutils_helpers"
)

// Test --generate-config-out for the Agreement resource
func TestPlatformConnectorTerraformPlanAgreementResource(t *testing.T) {
	// Get an instance of the PingOne SDK Client
	sdkClientInfo := testutils_helpers.GetPingOneSDKClientInfo(t)

	// Create an instance of the Agreement resource
	agreementResource := resources.Agreement(sdkClientInfo)

	// Run terraform plan --generate-config-out on the Agreement resource
	testutils_helpers.TestSingleResourceTerraformPlanGenerateConfigOut(t, agreementResource, sdkClientInfo)
}

// Test --generate-config-out for the AgreementEnable resource
func TestPlatformConnectorTerraformPlanAgreementEnableResource(t *testing.T) {
	// Get an instance of the PingOne SDK Client
	sdkClientInfo := testutils_helpers.GetPingOneSDKClientInfo(t)

	// Create an instance of the AgreementEnable resource
	agreementEnableResource := resources.AgreementEnable(sdkClientInfo)

	// Run terraform plan --generate-config-out on the AgreementEnable resource
	testutils_helpers.TestSingleResourceTerraformPlanGenerateConfigOut(t, agreementEnableResource, sdkClientInfo)
}

// Test --generate-config-out for the AgreementLocalization resource
func TestPlatformConnectorTerraformPlanAgreementLocalizationResource(t *testing.T) {
	// Get an instance of the PingOne SDK Client
	sdkClientInfo := testutils_helpers.GetPingOneSDKClientInfo(t)

	// Create an instance of the AgreementLocalization resource
	agreementLocalizationResource := resources.AgreementLocalization(sdkClientInfo)

	// Run terraform plan --generate-config-out on the AgreementLocalization resource
	testutils_helpers.TestSingleResourceTerraformPlanGenerateConfigOut(t, agreementLocalizationResource, sdkClientInfo)
}

// Test --generate-config-out for the AgreementLocalizationEnable resource
func TestPlatformConnectorTerraformPlanAgreementLocalizationEnableResource(t *testing.T) {
	// Get an instance of the PingOne SDK Client
	sdkClientInfo := testutils_helpers.GetPingOneSDKClientInfo(t)

	// Create an instance of the AgreementLocalizationEnable resource
	agreementLocalizationEnableResource := resources.AgreementLocalizationEnable(sdkClientInfo)

	// Run terraform plan --generate-config-out on the AgreementLocalizationEnable resource
	testutils_helpers.TestSingleResourceTerraformPlanGenerateConfigOut(t, agreementLocalizationEnableResource, sdkClientInfo)
}

// Test --generate-config-out for the AgreementLocalizationRevision resource
func TestPlatformConnectorTerraformPlanAgreementLocalizationRevisionResource(t *testing.T) {
	// Get an instance of the PingOne SDK Client
	sdkClientInfo := testutils_helpers.GetPingOneSDKClientInfo(t)

	// Create an instance of the AgreementLocalizationRevision resource
	agreementLocalizationRevisionResource := resources.AgreementLocalizationRevision(sdkClientInfo)

	// Run terraform plan --generate-config-out on the AgreementLocalizationRevision resource
	testutils_helpers.TestSingleResourceTerraformPlanGenerateConfigOut(t, agreementLocalizationRevisionResource, sdkClientInfo)
}

// Test --generate-config-out for the BrandingSettings resource
func TestPlatformConnectorTerraformPlanBrandingSettingsResource(t *testing.T) {
	// Get an instance of the PingOne SDK Client
	sdkClientInfo := testutils_helpers.GetPingOneSDKClientInfo(t)

	// Create an instance of the BrandingSettings resource
	brandingSettingsResource := resources.BrandingSettings(sdkClientInfo)

	// Run terraform plan --generate-config-out on the BrandingSettings resource
	testutils_helpers.TestSingleResourceTerraformPlanGenerateConfigOut(t, brandingSettingsResource, sdkClientInfo)
}

// Test --generate-config-out for the BrandingTheme resource
func TestPlatformConnectorTerraformPlanBrandingThemeResource(t *testing.T) {
	// Get an instance of the PingOne SDK Client
	sdkClientInfo := testutils_helpers.GetPingOneSDKClientInfo(t)

	// Create an instance of the BrandingTheme resource
	brandingThemeResource := resources.BrandingTheme(sdkClientInfo)

	// Run terraform plan --generate-config-out on the BrandingTheme resource
	testutils_helpers.TestSingleResourceTerraformPlanGenerateConfigOut(t, brandingThemeResource, sdkClientInfo)
}

// Test --generate-config-out for the BrandingThemeDefault resource
func TestPlatformConnectorTerraformPlanBrandingThemeDefaultResource(t *testing.T) {
	// Get an instance of the PingOne SDK Client
	sdkClientInfo := testutils_helpers.GetPingOneSDKClientInfo(t)

	// Create an instance of the BrandingThemeDefault resource
	brandingThemeDefaultResource := resources.BrandingThemeDefault(sdkClientInfo)

	// Run terraform plan --generate-config-out on the BrandingThemeDefault resource
	testutils_helpers.TestSingleResourceTerraformPlanGenerateConfigOut(t, brandingThemeDefaultResource, sdkClientInfo)
}

// Test --generate-config-out for the Certificate resource
func TestPlatformConnectorTerraformPlanCertificateResource(t *testing.T) {
	// Get an instance of the PingOne SDK Client
	sdkClientInfo := testutils_helpers.GetPingOneSDKClientInfo(t)

	// Create an instance of the Certificate resource
	certificateResource := resources.Certificate(sdkClientInfo)

	// Run terraform plan --generate-config-out on the Certificate resource
	testutils_helpers.TestSingleResourceTerraformPlanGenerateConfigOut(t, certificateResource, sdkClientInfo)
}

// Test --generate-config-out for the CustomDomain resource
func TestPlatformConnectorTerraformPlanCustomDomainResource(t *testing.T) {
	// Get an instance of the PingOne SDK Client
	sdkClientInfo := testutils_helpers.GetPingOneSDKClientInfo(t)

	// Create an instance of the CustomDomain resource
	customDomainResource := resources.CustomDomain(sdkClientInfo)

	// Run terraform plan --generate-config-out on the CustomDomain resource
	testutils_helpers.TestSingleResourceTerraformPlanGenerateConfigOut(t, customDomainResource, sdkClientInfo)
}

// Test --generate-config-out for the Environment resource
func TestPlatformConnectorTerraformPlanEnvironmentResource(t *testing.T) {
	// Get an instance of the PingOne SDK Client
	sdkClientInfo := testutils_helpers.GetPingOneSDKClientInfo(t)

	// Create an instance of the Environment resource
	environmentResource := resources.Environment(sdkClientInfo)

	// Run terraform plan --generate-config-out on the Environment resource
	testutils_helpers.TestSingleResourceTerraformPlanGenerateConfigOut(t, environmentResource, sdkClientInfo)
}

// Test --generate-config-out for the Form resource
func TestPlatformConnectorTerraformPlanFormResource(t *testing.T) {
	// Get an instance of the PingOne SDK Client
	sdkClientInfo := testutils_helpers.GetPingOneSDKClientInfo(t)

	// Create an instance of the Form resource
	formResource := resources.Form(sdkClientInfo)

	// Run terraform plan --generate-config-out on the Form resource
	testutils_helpers.TestSingleResourceTerraformPlanGenerateConfigOut(t, formResource, sdkClientInfo)
}

// Test --generate-config-out for the FormRecaptchaV2 resource
func TestPlatformConnectorTerraformPlanFormRecaptchaV2Resource(t *testing.T) {
	// Get an instance of the PingOne SDK Client
	sdkClientInfo := testutils_helpers.GetPingOneSDKClientInfo(t)

	// Create an instance of the FormRecaptchaV2 resource
	formRecaptchaV2Resource := resources.FormRecaptchaV2(sdkClientInfo)

	// Run terraform plan --generate-config-out on the FormRecaptchaV2 resource
	testutils_helpers.TestSingleResourceTerraformPlanGenerateConfigOut(t, formRecaptchaV2Resource, sdkClientInfo)
}

// Test --generate-config-out for the GatewayCredential resource
func TestPlatformConnectorTerraformPlanGatewayCredentialResource(t *testing.T) {
	// Get an instance of the PingOne SDK Client
	sdkClientInfo := testutils_helpers.GetPingOneSDKClientInfo(t)

	// Create an instance of the GatewayCredential resource
	gatewayCredentialResource := resources.GatewayCredential(sdkClientInfo)

	// Run terraform plan --generate-config-out on the GatewayCredential resource
	testutils_helpers.TestSingleResourceTerraformPlanGenerateConfigOut(t, gatewayCredentialResource, sdkClientInfo)
}

// Test --generate-config-out for the GatewayRoleAssignment resource
func TestPlatformConnectorTerraformPlanGatewayRoleAssignmentResource(t *testing.T) {
	// Get an instance of the PingOne SDK Client
	sdkClientInfo := testutils_helpers.GetPingOneSDKClientInfo(t)

	// Create an instance of the GatewayRoleAssignment resource
	gatewayRoleAssignmentResource := resources.GatewayRoleAssignment(sdkClientInfo)

	// Run terraform plan --generate-config-out on the GatewayRoleAssignment resource
	testutils_helpers.TestSingleResourceTerraformPlanGenerateConfigOut(t, gatewayRoleAssignmentResource, sdkClientInfo)
}

// Test --generate-config-out for the IdentityPropagationPlan resource
func TestPlatformConnectorTerraformPlanIdentityPropagationPlanResource(t *testing.T) {
	// Get an instance of the PingOne SDK Client
	sdkClientInfo := testutils_helpers.GetPingOneSDKClientInfo(t)

	// Create an instance of the IdentityPropagationPlan resource
	identityPropagationPlanResource := resources.IdentityPropagationPlan(sdkClientInfo)

	// Run terraform plan --generate-config-out on the IdentityPropagationPlan resource
	testutils_helpers.TestSingleResourceTerraformPlanGenerateConfigOut(t, identityPropagationPlanResource, sdkClientInfo)
}

// Test --generate-config-out for the Key resource
func TestPlatformConnectorTerraformPlanKeyResource(t *testing.T) {
	// Get an instance of the PingOne SDK Client
	sdkClientInfo := testutils_helpers.GetPingOneSDKClientInfo(t)

	// Create an instance of the Key resource
	keyResource := resources.Key(sdkClientInfo)

	// Run terraform plan --generate-config-out on the Key resource
	testutils_helpers.TestSingleResourceTerraformPlanGenerateConfigOut(t, keyResource, sdkClientInfo)
}

// Test --generate-config-out for the KeyRotationPolicy resource
func TestPlatformConnectorTerraformPlanKeyRotationPolicyResource(t *testing.T) {
	// Get an instance of the PingOne SDK Client
	sdkClientInfo := testutils_helpers.GetPingOneSDKClientInfo(t)

	// Create an instance of the KeyRotationPolicy resource
	keyRotationPolicyResource := resources.KeyRotationPolicy(sdkClientInfo)

	// Run terraform plan --generate-config-out on the KeyRotationPolicy resource
	testutils_helpers.TestSingleResourceTerraformPlanGenerateConfigOut(t, keyRotationPolicyResource, sdkClientInfo)
}

// Test --generate-config-out for the Language resource
func TestPlatformConnectorTerraformPlanLanguageResource(t *testing.T) {
	// Get an instance of the PingOne SDK Client
	sdkClientInfo := testutils_helpers.GetPingOneSDKClientInfo(t)

	// Create an instance of the Language resource
	languageResource := resources.Language(sdkClientInfo)

	// Run terraform plan --generate-config-out on the Language resource
	testutils_helpers.TestSingleResourceTerraformPlanGenerateConfigOut(t, languageResource, sdkClientInfo)
}

// Test --generate-config-out for the LanguageUpdate resource
func TestPlatformConnectorTerraformPlanLanguageUpdateResource(t *testing.T) {
	// Get an instance of the PingOne SDK Client
	sdkClientInfo := testutils_helpers.GetPingOneSDKClientInfo(t)

	// Create an instance of the LanguageUpdate resource
	languageUpdateResource := resources.LanguageUpdate(sdkClientInfo)

	// Run terraform plan --generate-config-out on the LanguageUpdate resource
	testutils_helpers.TestSingleResourceTerraformPlanGenerateConfigOut(t, languageUpdateResource, sdkClientInfo)
}

// Test --generate-config-out for the NotificationPolicy resource
func TestPlatformConnectorTerraformPlanNotificationPolicyResource(t *testing.T) {
	// Get an instance of the PingOne SDK Client
	sdkClientInfo := testutils_helpers.GetPingOneSDKClientInfo(t)

	// Create an instance of the NotificationPolicy resource
	notificationPolicyResource := resources.NotificationPolicy(sdkClientInfo)

	// Run terraform plan --generate-config-out on the NotificationPolicy resource
	testutils_helpers.TestSingleResourceTerraformPlanGenerateConfigOut(t, notificationPolicyResource, sdkClientInfo)
}

// Test --generate-config-out for the NotificationSettings resource
func TestPlatformConnectorTerraformPlanNotificationSettingsResource(t *testing.T) {
	// Get an instance of the PingOne SDK Client
	sdkClientInfo := testutils_helpers.GetPingOneSDKClientInfo(t)

	// Create an instance of the NotificationSettings resource
	notificationSettingsResource := resources.NotificationSettings(sdkClientInfo)

	// Run terraform plan --generate-config-out on the NotificationSettings resource
	testutils_helpers.TestSingleResourceTerraformPlanGenerateConfigOut(t, notificationSettingsResource, sdkClientInfo)
}

// Test --generate-config-out for the NotificationSettingsEmail resource
func TestPlatformConnectorTerraformPlanNotificationSettingsEmailResource(t *testing.T) {
	// Get an instance of the PingOne SDK Client
	sdkClientInfo := testutils_helpers.GetPingOneSDKClientInfo(t)

	// Create an instance of the NotificationSettingsEmail resource
	notificationSettingsEmailResource := resources.NotificationSettingsEmail(sdkClientInfo)

	// Run terraform plan --generate-config-out on the NotificationSettingsEmail resource
	testutils_helpers.TestSingleResourceTerraformPlanGenerateConfigOut(t, notificationSettingsEmailResource, sdkClientInfo)
}

// Test --generate-config-out for the NotificationTemplateContent resource
func TestPlatformConnectorTerraformPlanNotificationTemplateContentResource(t *testing.T) {
	// Get an instance of the PingOne SDK Client
	sdkClientInfo := testutils_helpers.GetPingOneSDKClientInfo(t)

	// Create an instance of the NotificationTemplateContent resource
	notificationTemplateContentResource := resources.NotificationTemplateContent(sdkClientInfo)

	// Run terraform plan --generate-config-out on the NotificationTemplateContent resource
	testutils_helpers.TestSingleResourceTerraformPlanGenerateConfigOut(t, notificationTemplateContentResource, sdkClientInfo)
}

// Test --generate-config-out for the PhoneDeliverySettings resource
func TestPlatformConnectorTerraformPlanPhoneDeliverySettingsResource(t *testing.T) {
	// Get an instance of the PingOne SDK Client
	sdkClientInfo := testutils_helpers.GetPingOneSDKClientInfo(t)

	// Create an instance of the PhoneDeliverySettings resource
	phoneDeliverySettingsResource := resources.PhoneDeliverySettings(sdkClientInfo)

	// Run terraform plan --generate-config-out on the PhoneDeliverySettings resource
	testutils_helpers.TestSingleResourceTerraformPlanGenerateConfigOut(t, phoneDeliverySettingsResource, sdkClientInfo)
}

// Test --generate-config-out for the RoleAssignmentUser resource
func TestPlatformConnectorTerraformPlanRoleAssignmentUserResource(t *testing.T) {
	// Get an instance of the PingOne SDK Client
	sdkClientInfo := testutils_helpers.GetPingOneSDKClientInfo(t)

	// Create an instance of the RoleAssignmentUser resource
	roleAssignmentUserResource := resources.RoleAssignmentUser(sdkClientInfo)

	// Run terraform plan --generate-config-out on the RoleAssignmentUser resource
	testutils_helpers.TestSingleResourceTerraformPlanGenerateConfigOut(t, roleAssignmentUserResource, sdkClientInfo)
}

// Test --generate-config-out for the SystemApplication resource
func TestPlatformConnectorTerraformPlanSystemApplicationResource(t *testing.T) {
	// Get an instance of the PingOne SDK Client
	sdkClientInfo := testutils_helpers.GetPingOneSDKClientInfo(t)

	// Create an instance of the SystemApplication resource
	systemApplicationResource := resources.SystemApplication(sdkClientInfo)

	// Run terraform plan --generate-config-out on the SystemApplication resource
	testutils_helpers.TestSingleResourceTerraformPlanGenerateConfigOut(t, systemApplicationResource, sdkClientInfo)
}

// Test --generate-config-out for the TrustedEmailAddress resource
func TestPlatformConnectorTerraformPlanTrustedEmailAddressResource(t *testing.T) {
	// Get an instance of the PingOne SDK Client
	sdkClientInfo := testutils_helpers.GetPingOneSDKClientInfo(t)

	// Create an instance of the TrustedEmailAddress resource
	trustedEmailAddressResource := resources.TrustedEmailAddress(sdkClientInfo)

	// Run terraform plan --generate-config-out on the TrustedEmailAddress resource
	testutils_helpers.TestSingleResourceTerraformPlanGenerateConfigOut(t, trustedEmailAddressResource, sdkClientInfo)
}

// Test --generate-config-out for the TrustedEmailDomain resource
func TestPlatformConnectorTerraformPlanTrustedEmailDomainResource(t *testing.T) {
	// Get an instance of the PingOne SDK Client
	sdkClientInfo := testutils_helpers.GetPingOneSDKClientInfo(t)

	// Create an instance of the TrustedEmailDomain resource
	trustedEmailDomainResource := resources.TrustedEmailDomain(sdkClientInfo)

	// Run terraform plan --generate-config-out on the TrustedEmailDomain resource
	testutils_helpers.TestSingleResourceTerraformPlanGenerateConfigOut(t, trustedEmailDomainResource, sdkClientInfo)
}

// Test --generate-config-out for the Webhook resource
func TestPlatformConnectorTerraformPlanWebhookResource(t *testing.T) {
	// Get an instance of the PingOne SDK Client
	sdkClientInfo := testutils_helpers.GetPingOneSDKClientInfo(t)

	// Create an instance of the Webhook resource
	webhookResource := resources.Webhook(sdkClientInfo)

	// Run terraform plan --generate-config-out on the Webhook resource
	testutils_helpers.TestSingleResourceTerraformPlanGenerateConfigOut(t, webhookResource, sdkClientInfo)
}
