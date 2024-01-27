terraform {
  required_providers {
    orynetwork = {
      source = "hashicorp.com/karakter98/ory-network"
    }
  }
}

provider "orynetwork" {}

resource "orynetwork_project" "project" {
  name = "Test Project"
  services = {
    identity = {
      config = {
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
          methods = {
            code = {
              config = {
                lifespan = "15m0s"
              }
              enabled              = true
              passwordless_enabled = true
            }
            password = {
              enabled = false
            }
          }
        }
      }
    }
  }
}
