package main

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

type describeInstancesV2Pager interface {
	HasMorePages() bool
	NextPage(context.Context, ...func(*ec2.Options)) (*ec2.DescribeInstancesOutput, error)
}

func describeInstancesV2(ctx context.Context, pager describeInstancesV2Pager) error {
	var err error
	for pager.HasMorePages() {
		_, err = pager.NextPage(ctx)
		if err != nil {
			return err
		}
	}
	return nil
}
