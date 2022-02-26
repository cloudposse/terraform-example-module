variable "region" {
  type = string
}

variable "example_input" {
  type        = string
  description = "The value which will be passed to the example module"
}

variable "example_input_override" {
  type        = string
  default     = null
  description = "Override for `example_input`"
}
