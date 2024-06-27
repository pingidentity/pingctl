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

	// TODO - Remove this ignore error.
	// This test is failing due to a bug where computed values are failing
	// config generation as they are treated as required attributes.
	ignoreErrors := []string{
		"Error: Missing Configuration for Required Attribute",
	}

	testutils_helpers.ValidateTerraformPlan(t, agreementLocalizationRevisionResource, ignoreErrors)
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

	// TODO - Remove this ignore error.
	// This test is failing due to a bug where computed values are failing
	// config generation as they are treated as required attributes.
	ignoreErrors := []string{
		"Error: Invalid Attribute Combination",
	}

	testutils_helpers.ValidateTerraformPlan(t, brandingThemeResource, ignoreErrors)
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

	// TODO - Remove this ignore error.
	// This test is failing due to a bug where computed values are failing
	// config generation as they are treated as required attributes.
	ignoreErrors := []string{
		`Error: attribute "components": attribute "fields": incorrect set element type: attributes "other_option_attribute_disabled", "other_option_enabled", "other_option_input_label", "other_option_key", and "other_option_label" are required`,
	}

	testutils_helpers.ValidateTerraformPlan(t, formResource, ignoreErrors)
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

	// TODO - Remove this ignore error.
	// This test is failing due to a bug where computed values are failing
	// config generation as they are treated as required attributes.
	ignoreErrors := []string{
		`Error: expected locale to be one of ["af" "af-ZA" "ar" "ar-AE" "ar-BH" "ar-DZ" "ar-EG" "ar-IQ" "ar-JO" "ar-KW" "ar-LB" "ar-LY" "ar-MA" "ar-OM" "ar-QA" "ar-SA" "ar-SY" "ar-TN" "ar-YE" "az" "az-AZ" "be" "be-BY" "bg" "bg-BG" "bs-BA" "ca" "ca-ES" "cs-CZ" "cy" "cy-GB" "da" "da-DK" "de-AT" "de-CH" "de-DE" "de-LI" "de-LU" "dv" "dv-MV" "el" "el-GR" "en-AU" "en-BZ" "en-CA" "en-CB" "en-GB" "en-IE" "en-JM" "en-NZ" "en-PH" "en-TT" "en-US" "en-ZA" "en-ZW" "eo" "es-AR" "es-BO" "es-CL" "es-CO" "es-CR" "es-DO" "es-EC" "es-ES" "es-GT" "es-HN" "es-MX" "es-NI" "es-PA" "es-PE" "es-PR" "es-PY" "es-SV" "es-UY" "es-VE" "et" "et-EE" "eu" "eu-ES" "fa" "fa-IR" "fi" "fi-FI" "fo" "fo-FO" "fr-BE" "fr-CH" "fr-FR" "fr-LU" "fr-MC" "gl" "gl-ES" "gu" "gu-IN" "he" "he-IL" "hi" "hi-IN" "hr" "hr-BA" "hr-HR" "hu-HU" "hy" "hy-AM" "id" "id-ID" "is" "is-IS" "it-CH" "it-IT" "ja-JP" "ka" "ka-GE" "kk" "kk-KZ" "kn" "kn-IN" "ko-KR" "kok" "kok-IN" "ky" "ky-KG" "lt" "lt-LT" "lv" "lv-LV" "mi" "mi-NZ" "mk" "mk-MK" "mn" "mn-MN" "mr" "mr-IN" "ms" "ms-BN" "ms-MY" "mt" "mt-MT" "nb" "nb-NO" "nl-BE" "nl-NL" "nn-NO" "ns" "ns-ZA" "pa" "pa-IN" "pl-PL" "ps" "ps-AR" "pt-BR" "pt-PT" "qu" "qu-BO" "qu-EC" "qu-PE" "ro" "ro-RO" "ru-RU" "sa" "sa-IN" "se" "se-FI" "se-NO" "se-SE" "sk" "sk-SK" "sl" "sl-SI" "sq" "sq-AL" "sr-BA" "sr-SP" "sv" "sv-FI" "sv-SE" "sw" "sw-KE" "syr" "syr-SY" "ta" "ta-IN" "te" "te-IN" "th-TH" "tl" "tl-PH" "tn" "tn-ZA" "tr-TR" "tt" "tt-RU" "ts" "uk" "uk-UA" "ur" "ur-PK" "uz" "uz-UZ" "vi" "vi-VN" "xh" "xh-ZA" "zh-CN" "zh-HK" "zh-MO" "zh-SG" "zh-TW" "zu" "zu-ZA"], got aa`,
	}

	testutils_helpers.ValidateTerraformPlan(t, languageResource, ignoreErrors)
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

// Test --generate-config-out for the Webhook resource
func TestWebhookTerraformPlan(t *testing.T) {
	sdkClientInfo := testutils_helpers.GetPingOneSDKClientInfo(t)
	webhookResource := resources.Webhook(sdkClientInfo)
	testutils_helpers.ValidateTerraformPlan(t, webhookResource, nil)
}
