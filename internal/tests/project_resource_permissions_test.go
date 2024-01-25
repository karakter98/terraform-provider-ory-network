package tests

import (
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"testing"
)

func TestAccProjectResourcePermissions(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Project can be created without specifying permissions config
			{
				Config: `
				resource "orynetwork_project" "test_permissions_no_settings" {
				  name = "DeleteMe"
				}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("orynetwork_project.test_permissions_no_settings", "services.permission.config.namespaces.#", "0"),
				),
			},
			// Can import project that was created without specifying settings
			{
				ResourceName:      "orynetwork_project.test_permissions_no_settings",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Can add permissions for project that was created without specifying them
			{
				Config: `
				resource "orynetwork_project" "test_permissions_no_settings" {
				  name = "DeleteMe"
				  services = {
				    permission = {
				      config = {
					    namespaces = [{
					      id = 2
					      name = "TestNoSettings"
					    }]
				      }
				    }
			      }
				}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("orynetwork_project.test_permissions_no_settings", "services.permission.config.namespaces.#", "1"),
					resource.TestCheckResourceAttr("orynetwork_project.test_permissions_no_settings", "services.permission.config.namespaces.0.id", "2"),
					resource.TestCheckResourceAttr("orynetwork_project.test_permissions_no_settings", "services.permission.config.namespaces.0.name", "TestNoSettings"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Project can be created with specifying permissions config
			{
				Config: `
				resource "orynetwork_project" "test_permissions_with_settings" {
				  name = "DeleteMe"
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
				}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("orynetwork_project.test_permissions_with_settings", "services.permission.config.namespaces.#", "1"),
					resource.TestCheckResourceAttr("orynetwork_project.test_permissions_with_settings", "services.permission.config.namespaces.0.id", "1"),
					resource.TestCheckResourceAttr("orynetwork_project.test_permissions_with_settings", "services.permission.config.namespaces.0.name", "Test"),
				),
			},
			// Can import project that was created with specifying settings
			{
				ResourceName:      "orynetwork_project.test_permissions_with_settings",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}
