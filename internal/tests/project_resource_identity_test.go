package tests

import (
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
				Config: `
				resource "orynetwork_project" "test_identity" {
				  name = "DeleteMe"
				  services = {
				    identity = {
					  config = {
						identity = {
						  default_schema_id = "preset://username"
						  schemas = [
							{
							  id  = "preset://username",
							  url = "base64://ewogICIkaWQiOiAiaHR0cHM6Ly9zY2hlbWFzLm9yeS5zaC9wcmVzZXRzL2tyYXRvcy9pZGVudGl0eS51c2VybmFtZS5zY2hlbWEuanNvbiIsCiAgIiRzY2hlbWEiOiAiaHR0cDovL2pzb24tc2NoZW1hLm9yZy9kcmFmdC0wNy9zY2hlbWEjIiwKICAidGl0bGUiOiAiUGVyc29uIiwKICAidHlwZSI6ICJvYmplY3QiLAogICJwcm9wZXJ0aWVzIjogewogICAgInRyYWl0cyI6IHsKICAgICAgInR5cGUiOiAib2JqZWN0IiwKICAgICAgInByb3BlcnRpZXMiOiB7CiAgICAgICAgInVzZXJuYW1lIjogewogICAgICAgICAgInR5cGUiOiAic3RyaW5nIiwKICAgICAgICAgICJ0aXRsZSI6ICJVc2VybmFtZSIsCiAgICAgICAgICAibWF4TGVuZ3RoIjogMTAwLAogICAgICAgICAgIm9yeS5zaC9rcmF0b3MiOiB7CiAgICAgICAgICAgICJjcmVkZW50aWFscyI6IHsKICAgICAgICAgICAgICAicGFzc3dvcmQiOiB7CiAgICAgICAgICAgICAgICAiaWRlbnRpZmllciI6IHRydWUKICAgICAgICAgICAgICB9LAogICAgICAgICAgICAgICJ3ZWJhdXRobiI6IHsKICAgICAgICAgICAgICAgICJpZGVudGlmaWVyIjogdHJ1ZQogICAgICAgICAgICAgIH0sCiAgICAgICAgICAgICAgInRvdHAiOiB7CiAgICAgICAgICAgICAgICAiYWNjb3VudF9uYW1lIjogdHJ1ZQogICAgICAgICAgICAgIH0KICAgICAgICAgICAgfQogICAgICAgICAgfQogICAgICAgIH0KICAgICAgfSwKICAgICAgInJlcXVpcmVkIjogWwogICAgICAgICJ1c2VybmFtZSIKICAgICAgXSwKICAgICAgImFkZGl0aW9uYWxQcm9wZXJ0aWVzIjogZmFsc2UKICAgIH0KICB9Cn0K"
							}
						  ]
						}
						selfservice = {
						  default_browser_return_url = "https://google.com"
						  methods = {
							code = {
							  config = {
								lifespan = "15m0s"
							  }
							  enabled              = true
							  passwordless_enabled = true
							}
							password = {
							  enabled = false
							}
						  }
						}
					  }
					}
			      }
				}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("orynetwork_project.test_identity", "services.identity.config.identity.schemas.#", "1"),
					resource.TestCheckResourceAttr("orynetwork_project.test_identity", "services.identity.config.identity.default_schema_id", "preset://username"),
					resource.TestCheckResourceAttr("orynetwork_project.test_identity", "services.identity.config.selfservice.methods.password.enabled", "false"),
					resource.TestCheckResourceAttr("orynetwork_project.test_identity", "services.identity.config.selfservice.methods.code.enabled", "true"),
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
				Config: `
				resource "orynetwork_project" "test_identity" {
				  name = "DeleteMe"
				  services = {
				    identity = {
					  config = {
						identity = {
						  default_schema_id = "preset://username"
						  schemas = [
							{
							  id  = "preset://username",
							  url = "base64://ewogICIkaWQiOiAiaHR0cHM6Ly9zY2hlbWFzLm9yeS5zaC9wcmVzZXRzL2tyYXRvcy9pZGVudGl0eS51c2VybmFtZS5zY2hlbWEuanNvbiIsCiAgIiRzY2hlbWEiOiAiaHR0cDovL2pzb24tc2NoZW1hLm9yZy9kcmFmdC0wNy9zY2hlbWEjIiwKICAidGl0bGUiOiAiUGVyc29uIiwKICAidHlwZSI6ICJvYmplY3QiLAogICJwcm9wZXJ0aWVzIjogewogICAgInRyYWl0cyI6IHsKICAgICAgInR5cGUiOiAib2JqZWN0IiwKICAgICAgInByb3BlcnRpZXMiOiB7CiAgICAgICAgInVzZXJuYW1lIjogewogICAgICAgICAgInR5cGUiOiAic3RyaW5nIiwKICAgICAgICAgICJ0aXRsZSI6ICJVc2VybmFtZSIsCiAgICAgICAgICAibWF4TGVuZ3RoIjogMTAwLAogICAgICAgICAgIm9yeS5zaC9rcmF0b3MiOiB7CiAgICAgICAgICAgICJjcmVkZW50aWFscyI6IHsKICAgICAgICAgICAgICAicGFzc3dvcmQiOiB7CiAgICAgICAgICAgICAgICAiaWRlbnRpZmllciI6IHRydWUKICAgICAgICAgICAgICB9LAogICAgICAgICAgICAgICJ3ZWJhdXRobiI6IHsKICAgICAgICAgICAgICAgICJpZGVudGlmaWVyIjogdHJ1ZQogICAgICAgICAgICAgIH0sCiAgICAgICAgICAgICAgInRvdHAiOiB7CiAgICAgICAgICAgICAgICAiYWNjb3VudF9uYW1lIjogdHJ1ZQogICAgICAgICAgICAgIH0KICAgICAgICAgICAgfQogICAgICAgICAgfQogICAgICAgIH0KICAgICAgfSwKICAgICAgInJlcXVpcmVkIjogWwogICAgICAgICJ1c2VybmFtZSIKICAgICAgXSwKICAgICAgImFkZGl0aW9uYWxQcm9wZXJ0aWVzIjogZmFsc2UKICAgIH0KICB9Cn0K"
							}
						  ]
						}
						selfservice = {
						  default_browser_return_url = "https://google.com"
						  methods = {
							code = {
							  config = {
								lifespan = "15m0s"
							  }
							  enabled              = false
							  passwordless_enabled = true
							}
							password = {
							  enabled = true
							}
						  }
						}
					  }
					}
			      }
				}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("orynetwork_project.test_identity", "services.identity.config.selfservice.methods.password.enabled", "true"),
					resource.TestCheckResourceAttr("orynetwork_project.test_identity", "services.identity.config.selfservice.methods.code.enabled", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}
