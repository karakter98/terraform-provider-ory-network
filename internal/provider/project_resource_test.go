// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccProjectResource(t *testing.T) {
	permissionJson, _ := json.Marshal(map[string]interface{}{
		"namespaces": []map[string]interface{}{{"id": 1, "name": "Test"}},
	})
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create testing
			{
				Config: testAccProjectResourceConfigWithSettings,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("orynetwork_project.test_with_settings", "name", os.Getenv("TF_VAR_TEST_ORY_NETWORK_PROJECT_NAME")),
					resource.TestCheckResourceAttr("orynetwork_project.test_with_settings", "services.permission.config", string(permissionJson)),
					resource.TestCheckResourceAttrSet("orynetwork_project.test_with_settings", "services.identity.config"),
					resource.TestCheckResourceAttrSet("orynetwork_project.test_with_settings", "services.oauth2.config"),
					resource.TestCheckResourceAttr("orynetwork_project.test_with_settings", "cors_admin.origins.#", "1"),
					resource.TestCheckResourceAttr("orynetwork_project.test_with_settings", "cors_admin.origins.0", "https://google.com"),
				),
			},
			// Create testing
			{
				Config: testAccProjectResourceConfigNoSettings,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("orynetwork_project.test_no_settings", "name", os.Getenv("TF_VAR_TEST_ORY_NETWORK_PROJECT_NAME")),
					resource.TestCheckResourceAttr("orynetwork_project.test_no_settings", "services.permission.config", `{"namespaces":[]}`),
				),
			},
			// Import testing
			{
				ResourceName:      "orynetwork_project.test_no_settings",
				ImportState:       true,
				ImportStateVerify: true,
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
	  config = jsonencode({
		namespaces = [{
		  id = 1
		  name = "Test"
		}]
      })
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
