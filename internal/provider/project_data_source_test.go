// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccProjectDataSource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: testAccProjectDataSourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.orynetwork_project.test", "id", os.Getenv("TF_VAR_TEST_ORY_NETWORK_PROJECT_ID")),
					resource.TestCheckResourceAttr("data.orynetwork_project.test", "name", "Test"),
					resource.TestCheckResourceAttrSet("data.orynetwork_project.test", "slug"),
					resource.TestCheckResourceAttrSet("data.orynetwork_project.test", "services.identity.config"),
					resource.TestCheckResourceAttrSet("data.orynetwork_project.test", "services.oauth2.config"),
					resource.TestCheckResourceAttrSet("data.orynetwork_project.test", "services.permission.config"),
				),
			},
		},
	})
}

const testAccProjectDataSourceConfig = `
variable "TEST_ORY_NETWORK_PROJECT_ID" {
  type = string
}
data "orynetwork_project" "test" {
  id = var.TEST_ORY_NETWORK_PROJECT_ID
}
`
