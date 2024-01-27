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
  cors_admin = {
    enabled = true
    origins = ["https://google.com"]
  }
  cors_public = {
    enabled = true
    origins = ["https://google.com"]
  }
  services = {
    identity = {
      config = jsonencode({
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
            link = {
              enabled = true
              config = {
                lifespan = "15m05"
                base_url = "https://google.com"
              }
            }
            code = {
              config = {
                lifespan = "15m0s"
              }
              enabled              = true
              passwordless_enabled = true
            }
            password = {
              enabled = true
              config = {
                haveibeenpwned_enabled              = true
                max_breaches                        = 1
                ignore_network_errors               = true
                min_password_length                 = 8
                identifier_similarity_check_enabled = true
              }
            }
            totp = {
              enabled = true
              config = {
                issuer = "https://google.com"
              }
            }
            lookup_secret = {
              enabled = true
            }
            profile = {
              enabled = true
            }
            webauthn = {
              enabled = true
              config = {
                passwordless = true
                rp = {
                  id           = "Test"
                  display_name = "Test"
                }
              }
            }
            oidc = {
              enabled = true
              config = {
                base_redirect_uri = "https://google.com"
                providers = [{
                  id                   = "Test"
                  provider             = "google"
                  client_id            = "Test"
                  mapper_url           = "https://google.com"
                  client_secret        = "Test"
                  issuer_url           = "https://google.com"
                  auth_url             = "https://google.com"
                  token_url            = "https://google.com"
                  scope                = ["profile"]
                  microsoft_tenant     = "Test"
                  subject_source       = "Test"
                  apple_team_id        = "Test"
                  apple_private_key_id = "Test"
                  apple_private_key    = "Test"
                  requested_claims = {
                    id_token : ["profile"]
                  }
                  organization_id               = "Test"
                  label                         = "Test"
                  additional_id_token_audiences = ["Test"]
                }]
              }
            }
          }
        }
      })
    }
  }
}