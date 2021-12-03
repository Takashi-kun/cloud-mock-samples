package main

// Import key modules.
import (
	"context"

	"github.com/Azure/azure-sdk-for-go/profiles/latest/compute/mgmt/compute/computeapi"
	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2021-07-01/compute"
)

func getAzureVM(client computeapi.VirtualMachinesClientAPI) (compute.VirtualMachine, error) {
	return client.Get(context.TODO(), "resource-group", "sample-vm", compute.InstanceViewTypesInstanceView)
}
