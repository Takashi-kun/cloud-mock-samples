package main

import (
	"context"
	"testing"

	"github.com/Azure/azure-sdk-for-go/profiles/latest/compute/mgmt/compute/computeapi"
	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2021-07-01/compute"
)

var _ (computeapi.VirtualMachinesClientAPI) = new(mockVMClient)

type (
	mockVMClient struct {
		compute.VirtualMachinesClient
	}
)

func (m mockVMClient) List(ctx context.Context, resourceGroupName string) (result compute.VirtualMachineListResultPage, err error) {
	return compute.VirtualMachineListResultPage{}, errMock
}

func Test_listAzureVM(t *testing.T) {
	_, err := listAzureVM(context.TODO(), new(mockVMClient))
	if err != nil {
		t.Fatalf("failed to getAzureVM: %v", err) // comes here
	}
}
