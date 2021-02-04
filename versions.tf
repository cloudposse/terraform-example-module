terraform {
  required_version = ">= 0.13"

  required_providers {
    # Update these to reflect the actual requirements of your module
    local = {
      source  = "hashicorp/local"
      version = ">= 1.2"
    }
  }
}
