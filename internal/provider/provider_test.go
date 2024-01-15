// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
)

// testAccProtoV6ProviderFactories are used to instantiate a provider during
// acceptance testing. The factory function will be invoked for every Terraform
// CLI command executed to create a provider server to which the CLI can
// reattach.
var testAccProtoV6ProviderFactories = map[string]func() (tfprotov6.ProviderServer, error){
	"orynetwork": providerserver.NewProtocol6WithError(New("test")()),
}

func testAccPreCheck(t *testing.T) {
	// You can add code here to run prior to any test case execution, for example assertions
	// about the appropriate environment variables being set are common to see in a pre-check
	// function.
	if os.Getenv("ORY_NETWORK_EMAIL") == "" {
		t.Fatal("ORY_NETWORK_EMAIL must be set for acceptance tests")
	}
	if os.Getenv("ORY_NETWORK_PASSWORD") == "" {
		t.Fatal("ORY_NETWORK_PASSWORD must be set for acceptance tests")
	}
	if os.Getenv("TF_VAR_TEST_ORY_NETWORK_PROJECT_ID") == "" {
		t.Fatal("TF_VAR_TEST_ORY_NETWORK_PROJECT_ID must be set for acceptance tests")
	}
}
