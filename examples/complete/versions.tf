terraform {
  required_version = ">= 0.13.0"

  required_providers {
    local = {
      source  = "hashicorp/local"
      version = "2.1.0"
    }
  }
}
