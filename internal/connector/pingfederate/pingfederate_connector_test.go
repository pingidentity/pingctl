package pingfederate_test

import (
	"testing"

	"github.com/pingidentity/pingctl/internal/connector"
	"github.com/pingidentity/pingctl/internal/connector/pingfederate/resources"
	"github.com/pingidentity/pingctl/internal/testing/testutils"
	"github.com/pingidentity/pingctl/internal/testing/testutils_terraform"
)

func TestPingFederateTerraformPlan(t *testing.T) {
	PingFederateClientInfo := testutils.GetPingFederateClientInfo(t)

	testutils_terraform.InitPingFederateTerraform(t)

	testCases := []struct {
		name          string
		resource      connector.ExportableResource
		ignoredErrors []string
	}{
		{
			name:     "PingFederateAuthenticationApiApplication",
			resource: resources.AuthenticationApiApplication(PingFederateClientInfo),
			ignoredErrors: []string{
				"Error: Invalid Attribute Value", // TODO - Remove with PDI-1925 fix
			},
		},
		{
			name:          "PingFederateAuthenticationApiSettings",
			resource:      resources.AuthenticationApiSettings(PingFederateClientInfo),
			ignoredErrors: nil,
		},
		{
			name:     "PingFederateAuthenticationPolicies",
			resource: resources.AuthenticationPolicies(PingFederateClientInfo),
			ignoredErrors: []string{
				"Error: Plugin did not respond",
				"Error: Request cancelled",
			},
		},
		{
			name:     "PingFederateAuthenticationPoliciesFragment",
			resource: resources.AuthenticationPoliciesFragment(PingFederateClientInfo),
			ignoredErrors: []string{
				"Error: Reference to undeclared resource",
			},
		},
		{
			name:          "PingFederateAuthenticationPoliciesSettings",
			resource:      resources.AuthenticationPoliciesSettings(PingFederateClientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "PingFederateAuthenticationPolicyContract",
			resource:      resources.AuthenticationPolicyContract(PingFederateClientInfo),
			ignoredErrors: nil,
		},
		{
			name:     "PingFederateAuthenticationSelector",
			resource: resources.AuthenticationSelector(PingFederateClientInfo),
			ignoredErrors: []string{
				"Error: Plugin did not respond",
				"Error: Request cancelled",
			},
		},
		{
			name:     "PingFederateCertificateCA",
			resource: resources.CertificateCA(PingFederateClientInfo),
			ignoredErrors: []string{
				"Error: Invalid Attribute Value Length",
			},
		},
		{
			name:     "PingFederateDataStore",
			resource: resources.DataStore(PingFederateClientInfo),
			ignoredErrors: []string{
				"Error: 'password' and 'user_dn' must be set together",
			},
		},
		{
			name:     "PingFederateExtendedProperties",
			resource: resources.ExtendedProperties(PingFederateClientInfo),
			ignoredErrors: []string{
				"Error: Plugin did not respond",
				"Error: Request cancelled",
			},
		},
		{
			name:     "PingFederateIDPAdapter",
			resource: resources.IDPAdapter(PingFederateClientInfo),
			ignoredErrors: []string{
				"Error: Missing Configuration for Required Attribute",
				"Error: Reference to undeclared resource",
			},
		},
		{
			name:          "PingFederateIDPDefaultURLs",
			resource:      resources.IDPDefaultURLs(PingFederateClientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "PingFederateIDPSPConnection",
			resource:      resources.IDPSPConnection(PingFederateClientInfo),
			ignoredErrors: nil,
		},
		{
			name:     "PingFederateIncomingProxySettings",
			resource: resources.IncomingProxySettings(PingFederateClientInfo),
			ignoredErrors: []string{
				"Error: Plugin did not respond",
				"Error: Request cancelled",
			},
		},
		{
			name:     "PingFederateKerberosRealm",
			resource: resources.KerberosRealm(PingFederateClientInfo),
			ignoredErrors: []string{
				"Error: Property Required:",
			},
		},
		{
			name:          "PingFederateLocalIdentityIdentityProfile",
			resource:      resources.LocalIdentityIdentityProfile(PingFederateClientInfo),
			ignoredErrors: nil,
		},
		{
			name:     "PingFederateNotificationPublishersSettings",
			resource: resources.NotificationPublishersSettings(PingFederateClientInfo),
			ignoredErrors: []string{
				"Error: Plugin did not respond",
				"Error: Request cancelled",
			},
		},
		{
			name:          "PingFederateOAuthAccessTokenManager",
			resource:      resources.OAuthAccessTokenManager(PingFederateClientInfo),
			ignoredErrors: nil,
		},
		{
			name:     "PingFederateOAuthAccessTokenMapping",
			resource: resources.OAuthAccessTokenMapping(PingFederateClientInfo),
			ignoredErrors: []string{
				"Error: Plugin did not respond",
				"Error: Request cancelled",
			},
		},
		{
			name:          "PingFederateOAuthCIBAServerPolicySettings",
			resource:      resources.OAuthCIBAServerPolicySettings(PingFederateClientInfo),
			ignoredErrors: nil,
		},
		{
			name:     "PingFederateOAuthClient",
			resource: resources.OAuthClient(PingFederateClientInfo),
			ignoredErrors: []string{
				"Error: persistent_grant_expiration_type must be configured to \"OVERRIDE_SERVER_DEFAULT\" to modify the other persistent_grant_expiration values.",
				"Error: client_auth.secret cannot be empty when \"CLIENT_CREDENTIALS\" is included in grant_types.",
				"Error: Invalid Attribute Value",
			},
		},
		{
			name:          "PingFederateOAuthIssuer",
			resource:      resources.OAuthIssuer(PingFederateClientInfo),
			ignoredErrors: nil,
		},
		{
			name:     "PingFederateOpenIDConnectSettings",
			resource: resources.OpenIDConnectSettings(PingFederateClientInfo),
			ignoredErrors: []string{
				"Error: Plugin did not respond",
				"Error: Request cancelled",
			},
		},
		{
			name:     "PingFederatePasswordCredentialValidator",
			resource: resources.PasswordCredentialValidator(PingFederateClientInfo),
			ignoredErrors: []string{
				"Error: The \"LDAP Datastore\" field is required for the LDAP Username Password Credential Validator",
				"Error: The \"Search Base\" field is required for the LDAP Username Password Credential Validator",
				"Error: The \"Search Filter\" field is required for the LDAP Username Password Credential Validator",
			},
		},
		{
			name:          "PingFederateRedirectValidation",
			resource:      resources.RedirectValidation(PingFederateClientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "PingFederateServerSettings",
			resource:      resources.ServerSettings(PingFederateClientInfo),
			ignoredErrors: nil,
		},
		{
			name:     "PingFederateServerSettingsSystemKeys",
			resource: resources.ServerSettingsSystemKeys(PingFederateClientInfo),
			ignoredErrors: []string{
				"Error: Invalid Attribute Combination",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			testutils_terraform.ValidateTerraformPlan(t, tc.resource, tc.ignoredErrors)
		})
	}
}
