resource "orynetwork_project" "test_identity" {
  name = "DeleteMe"
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
            link = {
              enabled = true
              config = {
                lifespan = "15m0s"
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
                haveibeenpwned_enabled = true
                max_breaches = 1
                ignore_network_errors = true
                min_password_length = 8
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
                  id = "Test"
                  display_name = "Test"
                }
              }
            }
            oidc = {
              enabled = true
              config = {
                base_redirect_uri = "https://google.com"
                providers = [{
                  id = "Test"
                  provider = "google"
                  client_id = "Test"
                  client_secret = "Test"
                  mapper_url = "https://storage.googleapis.com/bac-gcs-production/365a484e56a5caad3c2524c80cc2df4b4fa1f21ba9ae8ac654c9d6f666d912858ff619c2e4ac9fcb7f5786d3e6a44747c9d881c3b7ad9eabe722798887238ace.jsonnet"
                  scope = ["email"]
                }]
              }
            }
          }
          flows = {
            logout = {
              after = {
                default_browser_return_url = "https://google.com"
              }
            }
            error = {
              ui_url = "https://google.com"
            }
            registration = {
              login_hints = true
              ui_url = "https://google.com"
              lifespan = "5m0s"
              enabled = true
            }
            login = {
              ui_url = "https://google.com"
              lifespan = "5m0s"
            }
            verification = {
              ui_url = "https://google.com"
              lifespan = "5m0s"
              use = "code"
              notify_unknown_recipients = true
              enabled = true
            }
            recovery = {
              ui_url = "https://google.com"
              lifespan = "5m0s"
              use = "code"
              notify_unknown_recipients = true
              enabled = true
            }
            settings = {
              ui_url = "https://google.com"
              lifespan = "5m0s"
              privileged_session_max_age = "5m0s"
              required_aal = "highest_available"
            }
          }
        }
      }
    }
  }
}
