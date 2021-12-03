package main

// Import key modules.
import (
	"context"

	"github.com/Azure/azure-sdk-for-go/profiles/latest/compute/mgmt/compute/computeapi"
	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2021-07-01/compute"
)

func listAzureVM(ctx context.Context, client computeapi.VirtualMachinesClientAPI) (compute.VirtualMachineListResultPage, error) {
	return client.List(ctx, "resource-group")
}
