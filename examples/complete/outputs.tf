output "id" {
  description = "ID of the created example"
  value       = module.example.id
}

output "example" {
  value       = module.example.example
}

output "random" {
  value       = module.example.random
}
