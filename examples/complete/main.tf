module "example" {
  source = "../.."

  example = coalesce(var.example_input_override, var.example_input)

  context = module.this.context
}
