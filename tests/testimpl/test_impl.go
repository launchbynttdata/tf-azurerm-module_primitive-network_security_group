package common

import (
	"context"
	"strings"
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/nexient-llc/lcaf-component-terratest-common/lib/azure/configure"
	"github.com/nexient-llc/lcaf-component-terratest-common/lib/azure/login"
	"github.com/nexient-llc/lcaf-component-terratest-common/lib/azure/network"
	"github.com/nexient-llc/lcaf-component-terratest-common/types"
	"github.com/stretchr/testify/assert"
)

const terraformDir string = "../../examples/nsg"
const varFile string = "test.tfvars"

func TestNsg(t *testing.T, ctx types.TestContext) {

	envVarMap := login.GetEnvironmentVariables()
	clientID := envVarMap["clientID"]
	clientSecret := envVarMap["clientSecret"]
	tenantID := envVarMap["tenantID"]
	subscriptionID := envVarMap["subscriptionID"]

	// Create an authorizer from env vars or Azure Managed Service Idenity
	spt, err := login.GetServicePrincipalToken(clientID, clientSecret, tenantID)
	if err != nil {
		t.Fatalf("Error getting Service Principal Token: %v", err)
	}

	// Create network security group client
	nsgClient := network.GetNsgClient(spt, subscriptionID)

	terraformOptions := configure.ConfigureTerraform(terraformDir, []string{terraformDir + "/" + varFile})
	t.Run("doesNsgExist", func(t *testing.T) {
		resourceGroupName := terraform.Output(t, terraformOptions, "resource_group_name")
		nsgName := terraform.Output(t, terraformOptions, "nsg_name")
		nsgId := terraform.Output(t, terraformOptions, "network_security_group_id")

		nsg, err := nsgClient.Get(context.Background(), resourceGroupName, nsgName, "")
		if err != nil {
			t.Fatalf("Error getting NSG: %v", err)
		}
		if nsg.Name == nil {
			t.Fatalf("NSG does not exist")
		}

		assert.Equal(t, getNsgName(*nsg.ID), strings.Trim(getNsgName(nsgId), "]"))
	})
}

func getNsgName(input string) string {
	parts := strings.Split(input, "/")
	return parts[len(parts)-1]
}
