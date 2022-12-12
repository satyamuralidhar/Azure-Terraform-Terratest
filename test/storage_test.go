package main

import (
	"fmt"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"log"
	"os"
	"testing"
)

// intializing variables
var (
	globalEnvVars = make(map[string]string)
)

//set environmnet varibales

func setEnvVars() (map[string]string, error) {
	/* fetch env varibales form bashprofile */
	TF_VAR_ARM_CLIENT_ID := os.Getenv("TF_VAR_AZURE_CLIENT_ID")
	TF_VAR_ARM_CLIENT_SECRET := os.Getenv("TF_VAR_AZURE_CLIENT_SECRET")
	TF_VAR_ARM_TENANT_ID := os.Getenv("TF_VAR_AZURE_TENANT_ID")
	TF_VAR_ARM_SUBSCRIPTION_ID := os.Getenv("TF_VAR_AZURE_SUBSCRIPTION_ID")

	/* create env vars from globalEnvVars to call Terraform to Terratest */

	if TF_VAR_ARM_CLIENT_ID != "" {
		globalEnvVars["TF_VAR_ARM_CLIENT_ID"] = TF_VAR_ARM_CLIENT_ID
		globalEnvVars["TF_VAR_ARM_CLIENT_SECRET"] = TF_VAR_ARM_CLIENT_SECRET
		globalEnvVars["TF_VAR_ARM_TENANT_ID"] = TF_VAR_ARM_TENANT_ID
		globalEnvVars["TF_VAR_ARM_SUBSCRIPTION_ID"] = TF_VAR_ARM_SUBSCRIPTION_ID
	}

	return globalEnvVars, nil

}

/* storge account terratest */

func TestTerraform_StorageAccount(t *testing.T) {
	t.Parallel()

	/* call Env funtion */
	setEnvVars()
	expectedLocation := "eastus"

	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{

		// Set the path to the Terraform code that will be tested.
		TerraformDir: "../storage",

		// Variables to pass to our Terraform code using -var options
		Vars: map[string]interface{}{

			"location": expectedLocation,
			//"tags":     tags,
		},

		// globalvariables for user account
		EnvVars: globalEnvVars,
		// Disable colors in Terraform commands so its easier to parse stdout/stderr
		NoColor: true,

		// Reconfigure is required if module deployment and go test pipelines are running in one stage
		Reconfigure: true,
	})

	defer terraform.Destroy(t, terraformOptions)

	// This will run `terraform init` and `terraform apply` and fail the test if there are any errors
	terraform.InitAndApply(t, terraformOptions)

	init, err := terraform.InitE(t, terraformOptions)
	if err != nil {
		log.Println(err)
	}
	t.Log(init)

	plan, err := terraform.PlanE(t, terraformOptions)
	if err != nil {
		log.Println(err)
	}
	t.Log(plan)
	//terraform applying
	apply, err := terraform.ApplyE(t, terraformOptions)
	if err != nil {
		log.Println(err)
	}
	t.Log(apply)

	//expected new variable not already declared
	//storagAccountName := terraform.Output(t, terraformOptions, "storage_account_name")

	fmt.Printf("location :: %s\n", expectedLocation)
	//fmt.Printf("storage_account :: %s\n", storagAccountName)
}
