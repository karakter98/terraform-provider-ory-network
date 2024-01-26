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
				resource "orynetwork_project" "test_permissions" {
				  name = "DeleteMe"
				}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("orynetwork_project.test_permissions", "services.permission.config.namespaces.#", "0"),
				),
			},
			// Can import project that was created without specifying settings
			{
				ResourceName:      "orynetwork_project.test_permissions",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Can add permissions for project that was created without specifying them
			{
				Config: `
				resource "orynetwork_project" "test_permissions" {
				  name = "DeleteMe"
				  services = {
				    permission = {
				      config = {
					    namespaces = [{
					      id = 2
					      name = "Test"
					    }]
				      }
				    }
			      }
				}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("orynetwork_project.test_permissions", "services.permission.config.namespaces.#", "1"),
					resource.TestCheckResourceAttr("orynetwork_project.test_permissions", "services.permission.config.namespaces.0.id", "2"),
					resource.TestCheckResourceAttr("orynetwork_project.test_permissions", "services.permission.config.namespaces.0.name", "Test"),
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
				resource "orynetwork_project" "test_permissions" {
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
					resource.TestCheckResourceAttr("orynetwork_project.test_permissions", "services.permission.config.namespaces.#", "1"),
					resource.TestCheckResourceAttr("orynetwork_project.test_permissions", "services.permission.config.namespaces.0.id", "1"),
					resource.TestCheckResourceAttr("orynetwork_project.test_permissions", "services.permission.config.namespaces.0.name", "Test"),
				),
			},
			// Can import project that was created with specifying settings
			{
				ResourceName:      "orynetwork_project.test_permissions",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Can update existing permissions
			{
				Config: `
				resource "orynetwork_project" "test_permissions" {
				  name = "DeleteMe"
				  services = {
				    permission = {
				      config = {
					    namespaces = [
						  {
					        id = 1
					        name = "Test"
					      },
						  {
					        id = 2
					        name = "Test2"
					      }
						]
				      }
				    }
			      }
				}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("orynetwork_project.test_permissions", "services.permission.config.namespaces.#", "2"),
					resource.TestCheckResourceAttr("orynetwork_project.test_permissions", "services.permission.config.namespaces.0.id", "1"),
					resource.TestCheckResourceAttr("orynetwork_project.test_permissions", "services.permission.config.namespaces.0.name", "Test"),
					resource.TestCheckResourceAttr("orynetwork_project.test_permissions", "services.permission.config.namespaces.1.id", "2"),
					resource.TestCheckResourceAttr("orynetwork_project.test_permissions", "services.permission.config.namespaces.1.name", "Test2"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}
