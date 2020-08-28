output "example" {
  description = "Example output"
  value       = module.this.enabled ? local.example : null
}

output "random" {
  description = "Stable random number for this example"
  value       = module.this.enabled ? join("", random_integer.example[*].result) : null
}
