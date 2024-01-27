// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccProjectResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create testing
			{
				Config: `
					resource "orynetwork_project" "test_project" {
					  name = "DeleteMe"
					  services = {
						permission = {
						  config = jsonencode({
							namespaces = [{
							  id = 1
							  name = "Test"
							}]
						  })
						}
						identity = {
						  config = jsonencode({
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
							}
						  })
						}
					  }
					  cors_admin = {
						enabled = true
						origins = ["https://google.com"]
					  }
					}
					`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("orynetwork_project.test_project", "id"),
					resource.TestCheckResourceAttr("orynetwork_project.test_project", "name", "DeleteMe"),
					resource.TestCheckResourceAttrSet("orynetwork_project.test_project", "services.permission.config"),
					resource.TestCheckResourceAttrSet("orynetwork_project.test_project", "services.identity.config"),
					resource.TestCheckResourceAttrSet("orynetwork_project.test_project", "services.oauth2.config"),
					resource.TestCheckResourceAttr("orynetwork_project.test_project", "cors_admin.origins.#", "1"),
					resource.TestCheckResourceAttr("orynetwork_project.test_project", "cors_admin.origins.0", "https://google.com"),
				),
			},
			// Import testing
			{
				ResourceName:            "orynetwork_project.test_project",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"services.identity.config", "services.oauth2.config", "services.permission.config"},
			},
			// Update testing
			{
				Config: `
					resource "orynetwork_project" "test_project" {
					  name = "DeleteMe"
					  services = {
						permission = {
						  config = jsonencode({})
						}
						identity = {
						  config = jsonencode({
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
							  default_browser_return_url = "https://stackoverflow.com"
							}
						  })
						}
					  }
					  cors_admin = {
						enabled = true
						origins = ["https://stackoverflow.com"]
					  }
					}
					`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("orynetwork_project.test_project", "id"),
					resource.TestCheckResourceAttr("orynetwork_project.test_project", "name", "DeleteMe"),
					resource.TestCheckResourceAttrSet("orynetwork_project.test_project", "services.permission.config"),
					resource.TestCheckResourceAttrSet("orynetwork_project.test_project", "services.identity.config"),
					resource.TestCheckResourceAttrSet("orynetwork_project.test_project", "services.oauth2.config"),
					resource.TestCheckResourceAttr("orynetwork_project.test_project", "cors_admin.origins.#", "1"),
					resource.TestCheckResourceAttr("orynetwork_project.test_project", "cors_admin.origins.0", "https://stackoverflow.com"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}
