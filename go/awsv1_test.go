package main

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/ec2/ec2iface"
)

type (
	mockEC2API struct {
		*ec2.EC2
	}
)

var _ ec2iface.EC2API = (*mockEC2API)(nil)

func (m *mockEC2API) DescribeInstancesWithContext(context.Context, *ec2.DescribeInstancesInput, ...request.Option) (*ec2.DescribeInstancesOutput, error) {
	return nil, errMock
}

func Test_describeInstances(t *testing.T) {
	err := describeInstances(context.TODO(), &mockEC2API{}, &ec2.DescribeInstancesInput{})
	if err != nil {
		t.Fatalf("expect no error, got %v", err) // comes here
	}
}
