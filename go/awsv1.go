package main

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/efs"
	"github.com/aws/aws-sdk-go/service/efs/efsiface"
)

func describeFileSystems(ctx context.Context, svc efsiface.EFSAPI, req *efs.DescribeFileSystemsInput) error {
	for {
		res, err := svc.DescribeFileSystemsWithContext(ctx, req)
		if err != nil {
			return err
		}
		for _, fs := range res.FileSystems {
			println(aws.StringValue(fs.Name))
		}
		if res.NextMarker == nil {
			break
		}

		req = req.SetMarker(*res.NextMarker)
	}
	return nil
}
