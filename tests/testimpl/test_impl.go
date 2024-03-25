package common

import (
	"context"
	"strings"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	armNetwork "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v5"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/launchbynttdata/lcaf-component-terratest/lib/azure/login"
	"github.com/launchbynttdata/lcaf-component-terratest/types"
	"github.com/stretchr/testify/assert"
)

func TestNsg(t *testing.T, ctx types.TestContext) {

	envVarMap := login.GetEnvironmentVariables()
	subscriptionID := envVarMap["subscriptionID"]

	credential, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		t.Fatalf("Unable to get credentials: %e\n", err)
	}

	options := arm.ClientOptions{
		ClientOptions: azcore.ClientOptions{
			Cloud: cloud.AzurePublic,
		},
	}

	nsgClient, err := armNetwork.NewSecurityGroupsClient(subscriptionID, credential, &options)
	if err != nil {
		t.Fatalf("Error getting NSG client: %v", err)
	}

	t.Run("doesNsgExist", func(t *testing.T) {
		resourceGroupName := terraform.Output(t, ctx.TerratestTerraformOptions(), "resource_group_name")
		nsgName := terraform.Output(t, ctx.TerratestTerraformOptions(), "nsg_name")
		nsgId := terraform.Output(t, ctx.TerratestTerraformOptions(), "network_security_group_id")

		nsg, err := nsgClient.Get(context.Background(), resourceGroupName, nsgName, nil)
		if err != nil {
			t.Fatalf("Error getting NSG: %v", err)
		}
		if nsg.Name == nil {
			t.Fatalf("NSG does not exist")
		}

		assert.Equal(t, getNsgName(*nsg.ID), strings.Trim(getNsgName(nsgId), "]"))
		assert.NotEmpty(t, nsg.Properties.SecurityRules)
	})
}

func getNsgName(input string) string {
	parts := strings.Split(input, "/")
	return parts[len(parts)-1]
}
