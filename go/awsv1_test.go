package main

import (
	"context"
	"errors"
	"testing"

	"github.com/aws/aws-sdk-go/service/efs"
	"github.com/aws/aws-sdk-go/service/efs/efsiface"
)

type (
	mockEFSAPI struct {
		*efs.EFS
	}
)

var _ efsiface.EFSAPI = (*mockEFSAPI)(nil)

func (m *mockEFSAPI) DescribeFileSystems(*efs.DescribeFileSystemsInput) (*efs.DescribeFileSystemsOutput, error) {
	return nil, errors.New("not implemented yet")
}

func Test_describeFileSystems(t *testing.T) {
	err := describeFileSystems(context.TODO(), &mockEFSAPI{}, &efs.DescribeFileSystemsInput{})
	if err != nil {
		t.Fatalf("expect no error, got %v", err) // comes here
	}
}
