terraform {
  required_providers {
    orynetwork = {
      source = "hashicorp.com/karakter98/ory-network"
    }
  }
}

variable "ory_account_email" {
  type = string
}

variable "ory_account_password" {
  type      = string
  sensitive = true
}

provider "orynetwork" {
  email    = var.ory_account_email
  password = var.ory_account_password
}