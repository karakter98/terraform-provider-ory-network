package tests

import (
	"github.com/hashicorp/terraform-plugin-testing/config"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"testing"
)

func TestAccProjectResourceIdentity(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Project can be created with specifying identity config
			{
				ConfigFile: config.StaticFile("resource_identity_test_create.tf"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("orynetwork_project.test_identity", "services.identity.config.identity.schemas.#", "1"),
					resource.TestCheckResourceAttr("orynetwork_project.test_identity", "services.identity.config.identity.schemas.0.id", "preset://username"),
					resource.TestCheckResourceAttrSet("orynetwork_project.test_identity", "services.identity.config.identity.schemas.0.url"),
					resource.TestCheckResourceAttr("orynetwork_project.test_identity", "services.identity.config.identity.default_schema_id", "preset://username"),

					resource.TestCheckResourceAttr("orynetwork_project.test_identity", "services.identity.config.selfservice.default_browser_return_url", "https://google.com"),

					resource.TestCheckResourceAttr("orynetwork_project.test_identity", "services.identity.config.selfservice.methods.link.enabled", "true"),
					resource.TestCheckResourceAttr("orynetwork_project.test_identity", "services.identity.config.selfservice.methods.link.config.lifespan", "15m0s"),
					resource.TestCheckResourceAttr("orynetwork_project.test_identity", "services.identity.config.selfservice.methods.link.config.base_url", "https://google.com"),

					resource.TestCheckResourceAttr("orynetwork_project.test_identity", "services.identity.config.selfservice.methods.code.enabled", "true"),
					resource.TestCheckResourceAttr("orynetwork_project.test_identity", "services.identity.config.selfservice.methods.code.passwordless_enabled", "true"),
					resource.TestCheckResourceAttr("orynetwork_project.test_identity", "services.identity.config.selfservice.methods.code.config.lifespan", "15m0s"),

					resource.TestCheckResourceAttr("orynetwork_project.test_identity", "services.identity.config.selfservice.methods.password.enabled", "true"),
					resource.TestCheckResourceAttr("orynetwork_project.test_identity", "services.identity.config.selfservice.methods.password.config.haveibeenpwned_enabled", "true"),
					resource.TestCheckResourceAttr("orynetwork_project.test_identity", "services.identity.config.selfservice.methods.password.config.max_breaches", "1"),
					resource.TestCheckResourceAttr("orynetwork_project.test_identity", "services.identity.config.selfservice.methods.password.config.ignore_network_errors", "true"),
					resource.TestCheckResourceAttr("orynetwork_project.test_identity", "services.identity.config.selfservice.methods.password.config.min_password_length", "8"),
					resource.TestCheckResourceAttr("orynetwork_project.test_identity", "services.identity.config.selfservice.methods.password.config.identifier_similarity_check_enabled", "true"),

					resource.TestCheckResourceAttr("orynetwork_project.test_identity", "services.identity.config.selfservice.methods.oidc.enabled", "true"),
					resource.TestCheckResourceAttr("orynetwork_project.test_identity", "services.identity.config.selfservice.methods.oidc.config.base_redirect_uri", "https://google.com"),
					resource.TestCheckResourceAttr("orynetwork_project.test_identity", "services.identity.config.selfservice.methods.oidc.config.providers.#", "1"),
					resource.TestCheckResourceAttr("orynetwork_project.test_identity", "services.identity.config.selfservice.methods.oidc.config.providers.0.id", "Test"),
					resource.TestCheckResourceAttr("orynetwork_project.test_identity", "services.identity.config.selfservice.methods.oidc.config.providers.0.provider", "google"),
					resource.TestCheckResourceAttr("orynetwork_project.test_identity", "services.identity.config.selfservice.methods.oidc.config.providers.0.client_id", "Test"),
					resource.TestCheckResourceAttr("orynetwork_project.test_identity", "services.identity.config.selfservice.methods.oidc.config.providers.0.client_secret", "Test"),
					resource.TestCheckResourceAttrSet("orynetwork_project.test_identity", "services.identity.config.selfservice.methods.oidc.config.providers.0.mapper_url"),
					resource.TestCheckResourceAttr("orynetwork_project.test_identity", "services.identity.config.selfservice.methods.oidc.config.providers.0.scope.0", "email"),

					resource.TestCheckResourceAttr("orynetwork_project.test_identity", "services.identity.config.selfservice.flows.logout.after.default_browser_return_url", "https://google.com"),
					resource.TestCheckResourceAttr("orynetwork_project.test_identity", "services.identity.config.selfservice.flows.error.ui_url", "https://google.com"),
					resource.TestCheckResourceAttr("orynetwork_project.test_identity", "services.identity.config.selfservice.flows.registration.login_hints", "true"),
					resource.TestCheckResourceAttr("orynetwork_project.test_identity", "services.identity.config.selfservice.flows.registration.ui_url", "https://google.com"),
					resource.TestCheckResourceAttr("orynetwork_project.test_identity", "services.identity.config.selfservice.flows.registration.lifespan", "5m0s"),
					resource.TestCheckResourceAttr("orynetwork_project.test_identity", "services.identity.config.selfservice.flows.registration.enabled", "true"),
					resource.TestCheckResourceAttr("orynetwork_project.test_identity", "services.identity.config.selfservice.flows.login.ui_url", "https://google.com"),
					resource.TestCheckResourceAttr("orynetwork_project.test_identity", "services.identity.config.selfservice.flows.login.lifespan", "5m0s"),
					resource.TestCheckResourceAttr("orynetwork_project.test_identity", "services.identity.config.selfservice.flows.verification.ui_url", "https://google.com"),
					resource.TestCheckResourceAttr("orynetwork_project.test_identity", "services.identity.config.selfservice.flows.verification.lifespan", "5m0s"),
					resource.TestCheckResourceAttr("orynetwork_project.test_identity", "services.identity.config.selfservice.flows.verification.use", "code"),
					resource.TestCheckResourceAttr("orynetwork_project.test_identity", "services.identity.config.selfservice.flows.verification.notify_unknown_recipients", "true"),
					resource.TestCheckResourceAttr("orynetwork_project.test_identity", "services.identity.config.selfservice.flows.verification.enabled", "true"),
					resource.TestCheckResourceAttr("orynetwork_project.test_identity", "services.identity.config.selfservice.flows.recovery.ui_url", "https://google.com"),
					resource.TestCheckResourceAttr("orynetwork_project.test_identity", "services.identity.config.selfservice.flows.recovery.lifespan", "5m0s"),
					resource.TestCheckResourceAttr("orynetwork_project.test_identity", "services.identity.config.selfservice.flows.recovery.use", "code"),
					resource.TestCheckResourceAttr("orynetwork_project.test_identity", "services.identity.config.selfservice.flows.recovery.notify_unknown_recipients", "true"),
					resource.TestCheckResourceAttr("orynetwork_project.test_identity", "services.identity.config.selfservice.flows.recovery.enabled", "true"),
					resource.TestCheckResourceAttr("orynetwork_project.test_identity", "services.identity.config.selfservice.flows.settings.ui_url", "https://google.com"),
					resource.TestCheckResourceAttr("orynetwork_project.test_identity", "services.identity.config.selfservice.flows.settings.lifespan", "5m0s"),
					resource.TestCheckResourceAttr("orynetwork_project.test_identity", "services.identity.config.selfservice.flows.settings.privileged_session_max_age", "5m0s"),
					resource.TestCheckResourceAttr("orynetwork_project.test_identity", "services.identity.config.selfservice.flows.settings.required_aal", "highest_available"),
				),
			},
			// Can import project that was created with specifying settings
			{
				ResourceName:      "orynetwork_project.test_identity",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Can update existing identity
			{
				ConfigFile: config.StaticFile("resource_identity_test_update.tf"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("orynetwork_project.test_identity", "services.identity.config.selfservice.methods.link.enabled", "false"),
					resource.TestCheckResourceAttr("orynetwork_project.test_identity", "services.identity.config.selfservice.methods.code.enabled", "true"),
					resource.TestCheckResourceAttr("orynetwork_project.test_identity", "services.identity.config.selfservice.methods.code.passwordless_enabled", "true"),
					resource.TestCheckResourceAttr("orynetwork_project.test_identity", "services.identity.config.selfservice.methods.password.enabled", "true"),
					resource.TestCheckResourceAttr("orynetwork_project.test_identity", "services.identity.config.selfservice.methods.oidc.enabled", "false"),
					resource.TestCheckResourceAttr("orynetwork_project.test_identity", "services.identity.config.selfservice.flows.registration.enabled", "false"),
					resource.TestCheckResourceAttr("orynetwork_project.test_identity", "services.identity.config.selfservice.flows.verification.enabled", "false"),
					resource.TestCheckResourceAttr("orynetwork_project.test_identity", "services.identity.config.selfservice.flows.recovery.enabled", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}
