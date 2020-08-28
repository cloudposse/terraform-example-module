output "id" {
  description = "ID of the created example"
  value       = module.this.enabled ? module.this.id : null
}

output "example" {
  description = "Example output"
  value       = module.this.enabled ? local.example : null
}

output "random" {
  description = "Stable random number for this example"
  value       = module.this.enabled ? join("", random_integer.example[*].result) : null
}
