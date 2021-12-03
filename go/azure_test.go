package main

import (
	"context"
	"errors"
	"testing"

	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2021-07-01/compute"
)

type (
	mockVMClient struct {
		compute.VirtualMachinesClient
	}
)

func (m mockVMClient) Get(ctx context.Context, resourceGroupName string, VMName string, expand compute.InstanceViewTypes) (result compute.VirtualMachine, err error) {
	return compute.VirtualMachine{}, errors.New("not implemented yet")
}

func Test_getAzureVM(t *testing.T) {
	_, err := getAzureVM(mockVMClient{})
	if err != nil {
		t.Fatalf("failed to getAzureVM: %v", err) // comes here
	}
}
