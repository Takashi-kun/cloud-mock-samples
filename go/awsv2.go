package main

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/efs"
)

type describeFileSystemsV2Pager interface {
	HasMorePages() bool
	NextPage(context.Context, ...func(*efs.Options)) (*efs.DescribeFileSystemsOutput, error)
}

func describeFileSystemsV2(ctx context.Context, pager describeFileSystemsV2Pager) error {
	var err error
	for pager.HasMorePages() {
		_, err = pager.NextPage(ctx)
		if err != nil {
			return err
		}
	}
	return nil
}
