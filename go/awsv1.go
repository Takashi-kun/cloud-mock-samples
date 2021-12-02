package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/efs"
	"github.com/aws/aws-sdk-go/service/efs/efsiface"
)

func describeFileSystems(svc efsiface.EFSAPI) error {
	req := &efs.DescribeFileSystemsInput{}
	for {
		res, err := svc.DescribeFileSystems(req)
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
