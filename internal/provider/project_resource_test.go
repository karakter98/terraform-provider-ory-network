// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"os"
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
				Config: testAccProjectResourceConfigWithSettings,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("orynetwork_project.test_with_settings", "id"),
					resource.TestCheckResourceAttr("orynetwork_project.test_with_settings", "name", os.Getenv("TF_VAR_TEST_ORY_NETWORK_PROJECT_NAME")),
					resource.TestCheckResourceAttr("orynetwork_project.test_with_settings", "services.permission.config.namespaces.0.id", "1"),
					resource.TestCheckResourceAttr("orynetwork_project.test_with_settings", "services.permission.config.namespaces.0.name", "Test"),
					resource.TestCheckResourceAttr("orynetwork_project.test_with_settings", "cors_admin.origins.#", "1"),
					resource.TestCheckResourceAttr("orynetwork_project.test_with_settings", "cors_admin.origins.0", "https://google.com"),
				),
			},
			// Create testing
			{
				Config: testAccProjectResourceConfigNoSettings,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("orynetwork_project.test_no_settings", "id"),
					resource.TestCheckResourceAttr("orynetwork_project.test_no_settings", "name", os.Getenv("TF_VAR_TEST_ORY_NETWORK_PROJECT_NAME")),
					resource.TestCheckResourceAttr("orynetwork_project.test_no_settings", "services.permission.config.namespaces.#", "0"),
				),
			},
			// Import testing
			{
				ResourceName:      "orynetwork_project.test_no_settings",
				ImportState:       true,
				ImportStateVerify: true,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("orynetwork_project.test_no_settings", "id"),
					resource.TestCheckResourceAttr("orynetwork_project.test_no_settings", "name", os.Getenv("TF_VAR_TEST_ORY_NETWORK_PROJECT_NAME")),
					resource.TestCheckResourceAttr("orynetwork_project.test_no_settings", "services.permission.config", `{"namespaces":[]}`),
				),
			},
			// Update testing
			{
				Config: testAccProjectResourceConfigUpdateSettings,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("orynetwork_project.test_no_settings", "id"),
					resource.TestCheckResourceAttr("orynetwork_project.test_no_settings", "name", os.Getenv("TF_VAR_TEST_ORY_NETWORK_PROJECT_NAME")),
					resource.TestCheckResourceAttr("orynetwork_project.test_no_settings", "services.permission.config.namespaces.0.id", "2"),
					resource.TestCheckResourceAttr("orynetwork_project.test_no_settings", "services.permission.config.namespaces.0.name", "Test2"),
					resource.TestCheckResourceAttr("orynetwork_project.test_no_settings", "cors_admin.origins.#", "1"),
					resource.TestCheckResourceAttr("orynetwork_project.test_no_settings", "cors_admin.origins.0", "https://google.com"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

const testAccProjectResourceConfigWithSettings = `
variable "TEST_ORY_NETWORK_PROJECT_NAME" {
  type = string
}
resource "orynetwork_project" "test_with_settings" {
  name = var.TEST_ORY_NETWORK_PROJECT_NAME
  services = {
	permission = {
	  config = {
		namespaces = [{
		  id = 1
		  name = "Test"
		}]
      }
	}
  }
  cors_admin = {
	enabled = true
	origins = ["https://google.com"]
  }
}
`
const testAccProjectResourceConfigNoSettings = `
variable "TEST_ORY_NETWORK_PROJECT_NAME" {
  type = string
}

resource "orynetwork_project" "test_no_settings" {
  name = var.TEST_ORY_NETWORK_PROJECT_NAME
}
`

const testAccProjectResourceConfigUpdateSettings = `
variable "TEST_ORY_NETWORK_PROJECT_NAME" {
  type = string
}

resource "orynetwork_project" "test_no_settings" {
  name = var.TEST_ORY_NETWORK_PROJECT_NAME
  services = {
	permission = {
	  config = {
		namespaces = [{
		  id = 2
		  name = "Test2"
		}]
      }
	}
  }
  cors_admin = {
	enabled = true
	origins = ["https://google.com"]
  }
}
`
