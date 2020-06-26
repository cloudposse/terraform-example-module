resource "random_integer" "example" {
  min = 1
  max = 50000
  keepers = {
    example = var.example
  }
}

locals {
  example = format("%v %d", var.example, random_integer.example.result)
}
