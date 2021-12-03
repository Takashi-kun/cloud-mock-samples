package main

import (
	"context"
	"errors"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/aws-sdk-go/aws"
)

type (
	mockDescribeInstancesV2Pager struct {
		PageNum int
		Pages   []*ec2.DescribeInstancesOutput
	}
)

func (m *mockDescribeInstancesV2Pager) HasMorePages() bool {
	return m.PageNum < len(m.Pages)
}

func (m *mockDescribeInstancesV2Pager) NextPage(ctx context.Context, opts ...func(*ec2.Options)) (*ec2.DescribeInstancesOutput, error) {
	var output *ec2.DescribeInstancesOutput
	if m.PageNum >= len(m.Pages) {
		return nil, errors.New("no more pages")
	}
	output = m.Pages[m.PageNum]
	m.PageNum++

	if m.PageNum > 2 {
		return nil, errMock
	}
	return output, nil
}

func Test_describeInstancesV2(t *testing.T) {
	pager := &mockDescribeInstancesV2Pager{
		Pages: []*ec2.DescribeInstancesOutput{
			{
				Reservations: []types.Reservation{
					{
						Instances: []types.Instance{
							{InstanceId: aws.String("foo_1")},
							{InstanceId: aws.String("bar_1")},
						},
					},
				},
			},
			{
				Reservations: []types.Reservation{
					{
						Instances: []types.Instance{
							{InstanceId: aws.String("foo_2")},
							{InstanceId: aws.String("bar_2")},
							{InstanceId: aws.String("baz_2")},
						},
					},
				},
			},
			{
				Reservations: []types.Reservation{
					{
						Instances: []types.Instance{
							{InstanceId: aws.String("foo_3")},
							{InstanceId: aws.String("bar_3")},
							{InstanceId: aws.String("baz_3")},
						},
					},
				},
			},
		},
	}
	err := describeInstancesV2(context.TODO(), pager)
	if err != nil {
		t.Fatalf("expect no error, got %v", err) // comes here
	}

}
