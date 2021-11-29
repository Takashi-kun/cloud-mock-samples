package main

import (
	"context"
	"errors"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/efs"
	"github.com/aws/aws-sdk-go-v2/service/efs/types"
	"github.com/aws/aws-sdk-go/aws"
)

type (
	mockDescribeFileSystemsV2Pager struct {
		PageNum int
		Pages   []*efs.DescribeFileSystemsOutput
	}
)

func (m *mockDescribeFileSystemsV2Pager) HasMorePages() bool {
	return m.PageNum < len(m.Pages)
}

func (m *mockDescribeFileSystemsV2Pager) NextPage(ctx context.Context, opts ...func(*efs.Options)) (*efs.DescribeFileSystemsOutput, error) {
	var output *efs.DescribeFileSystemsOutput
	if m.PageNum >= len(m.Pages) {
		return nil, errors.New("no more pages")
	}
	output = m.Pages[m.PageNum]
	m.PageNum++

	if m.PageNum > 2 {
		return nil, errors.New("dummy error")
	}
	return output, nil
}

func Test_describeInstances(t *testing.T) {
	pager := &mockDescribeFileSystemsV2Pager{
		Pages: []*efs.DescribeFileSystemsOutput{
			{
				FileSystems: []types.FileSystemDescription{
					{Name: aws.String("foo")},
				},
			},
			{
				FileSystems: []types.FileSystemDescription{
					{Name: aws.String("bar")},
				},
			},
			{
				FileSystems: []types.FileSystemDescription{
					{Name: aws.String("baz")},
				},
			},
		},
	}
	err := describeFileSystems(context.TODO(), pager)
	if err != nil {
		t.Fatalf("expect no error, got %v", err) // comes here
	}

}
