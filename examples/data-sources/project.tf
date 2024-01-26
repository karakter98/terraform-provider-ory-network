terraform {
  required_providers {
    orynetwork = {
      source = "hashicorp.com/karakter98/ory-network"
    }
  }
}

provider "orynetwork" {}

data "orynetwork_project" "project" {
  id = "YOUR PROJECT ID"
}

output "project_name" {
  value = data.orynetwork_project.project.name
}
