package tests

import (
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"testing"
)

func TestAccProjectResourceCors(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Project can be created with specifying cors config
			{
				Config: `
				resource "orynetwork_project" "test_cors" {
				  name = "DeleteMe"
				  cors_admin = {
					enabled = true
					origins = ["https://google.com"]
				  }
				  cors_public = {
					enabled = true
					origins = ["https://google.com"]
				  }
				}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("orynetwork_project.test_cors", "cors_admin.origins.#", "1"),
					resource.TestCheckResourceAttr("orynetwork_project.test_cors", "cors_admin.enabled", "true"),
					resource.TestCheckResourceAttr("orynetwork_project.test_cors", "cors_public.origins.#", "1"),
					resource.TestCheckResourceAttr("orynetwork_project.test_cors", "cors_public.enabled", "true"),
				),
			},
			// Can import project that was created with specifying settings
			{
				ResourceName:      "orynetwork_project.test_cors",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Can update existing cors
			{
				Config: `
				resource "orynetwork_project" "test_cors" {
				  name = "DeleteMe"
				  cors_admin = {
					enabled = true
					origins = ["https://google.com", "https://stackoverflow.com"]
				  }
				  cors_public = {
					enabled = true
					origins = ["https://google.com", "https://stackoverflow.com"]
				  }
				}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("orynetwork_project.test_cors", "cors_admin.origins.#", "2"),
					resource.TestCheckResourceAttr("orynetwork_project.test_cors", "cors_admin.enabled", "true"),
					resource.TestCheckResourceAttr("orynetwork_project.test_cors", "cors_public.origins.#", "2"),
					resource.TestCheckResourceAttr("orynetwork_project.test_cors", "cors_public.enabled", "true"),
				),
			},
			// Can disable existing cors
			{
				Config: `
				resource "orynetwork_project" "test_cors" {
				  name = "DeleteMe"
				  cors_admin = {
					enabled = false
				  }
				  cors_public = {
					enabled = false
				  }
				}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("orynetwork_project.test_cors", "cors_admin.origins.#", "0"),
					resource.TestCheckResourceAttr("orynetwork_project.test_cors", "cors_admin.enabled", "false"),
					resource.TestCheckResourceAttr("orynetwork_project.test_cors", "cors_public.origins.#", "0"),
					resource.TestCheckResourceAttr("orynetwork_project.test_cors", "cors_public.enabled", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}
