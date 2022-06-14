package test

import (
  "os"
  "strings"
  "testing"

  "github.com/gruntwork-io/terratest/modules/random"
  "github.com/gruntwork-io/terratest/modules/terraform"
  testStructure "github.com/gruntwork-io/terratest/modules/test-structure"
  "github.com/stretchr/testify/assert"
)

func cleanup(t *testing.T, terraformOptions *terraform.Options, tempTestFolder string) {
  terraform.Destroy(t, terraformOptions)
  os.RemoveAll(tempTestFolder)
}

// Test the Terraform module in examples/complete using Terratest.
func TestExamplesComplete(t *testing.T) {
  t.Parallel()
  randID := strings.ToLower(random.UniqueId())
  attributes := []string{randID}

  rootFolder := "../../"
  terraformFolderRelativeToRoot := "examples/complete"
  varFiles := []string{"fixtures.us-east-2.tfvars"}

  tempTestFolder := testStructure.CopyTerraformFolderToTemp(t, rootFolder, terraformFolderRelativeToRoot)

  terraformOptions := &terraform.Options{
    // The path to where our Terraform code is located
    TerraformDir: tempTestFolder,
    Upgrade:      true,
    // Variables to pass to our Terraform code using -var-file options
    VarFiles: varFiles,
    Vars: map[string]interface{}{
      "attributes": attributes,
    },
  }

  // At the end of the test, run `terraform destroy` to clean up any resources that were created
  defer cleanup(t, terraformOptions, tempTestFolder)

  // This will run `terraform init` and `terraform apply` and fail the test if there are any errors
  terraform.InitAndApply(t, terraformOptions)

  expectedExampleInput := "Hello, world!"

  // Run `terraform output` to get the value of an output variable
  id := terraform.Output(t, terraformOptions, "id")
  example := terraform.Output(t, terraformOptions, "example")
  random := terraform.Output(t, terraformOptions, "random")

  // Verify we're getting back the outputs we expect
  // Ensure we get a random number appended
  assert.Equal(t, expectedExampleInput+" "+random, example)
  // Ensure we get the attribute included in the ID
  assert.Equal(t, "eg-ue2-test-example-"+randID, id)

  // ************************************************************************
  // This steps below are unusual, not generally part of the testing
  // but included here as an example of testing this specific module.
  // This module has a random number that is supposed to change
  // only when the example changes. So we run it again to ensure
  // it does not change.

  // This will run `terraform apply` a second time and fail the test if there are any errors
  terraform.Apply(t, terraformOptions)

  id2 := terraform.Output(t, terraformOptions, "id")
  example2 := terraform.Output(t, terraformOptions, "example")
  random2 := terraform.Output(t, terraformOptions, "random")

  assert.Equal(t, id, id2, "Expected `id` to be stable")
  assert.Equal(t, example, example2, "Expected `example` to be stable")
  assert.Equal(t, random, random2, "Expected `random` to be stable")

  // Then we run change the example and run it a third time and
  // verify that the random number changed
  newExample := "Goodbye"
  terraformOptions.Vars["example_input_override"] = newExample
  terraform.Apply(t, terraformOptions)

  example3 := terraform.Output(t, terraformOptions, "example")
  random3 := terraform.Output(t, terraformOptions, "random")

  assert.NotEqual(t, random, random3, "Expected `random` to change when `example` changed")
  assert.Equal(t, newExample+" "+random3, example3, "Expected `example` to use new random number")

}

func TestExamplesCompleteDisabled(t *testing.T) {
  t.Parallel()
  randID := strings.ToLower(random.UniqueId())
  attributes := []string{randID}

  rootFolder := "../../"
  terraformFolderRelativeToRoot := "examples/complete"
  varFiles := []string{"fixtures.us-east-2.tfvars"}

  tempTestFolder := testStructure.CopyTerraformFolderToTemp(t, rootFolder, terraformFolderRelativeToRoot)

  terraformOptions := &terraform.Options{
    // The path to where our Terraform code is located
    TerraformDir: tempTestFolder,
    Upgrade:      true,
    // Variables to pass to our Terraform code using -var-file options
    VarFiles: varFiles,
    Vars: map[string]interface{}{
      "attributes": attributes,
      "enabled":    "false",
    },
  }

  // At the end of the test, run `terraform destroy` to clean up any resources that were created
  defer cleanup(t, terraformOptions, tempTestFolder)

  // This will run `terraform init` and `terraform apply` and fail the test if there are any errors
  results := terraform.InitAndApply(t, terraformOptions)

  // Should complete successfully without creating or changing any resources
  assert.Contains(t, results, "Resources: 0 added, 0 changed, 0 destroyed.")
}
