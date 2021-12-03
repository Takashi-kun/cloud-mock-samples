package main

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/ec2/ec2iface"
)

func describeInstances(ctx context.Context, svc ec2iface.EC2API, req *ec2.DescribeInstancesInput) error {
	for {
		res, err := svc.DescribeInstancesWithContext(ctx, req)
		if err != nil {
			return err
		}
		for _, rs := range res.Reservations {
			for _, instance := range rs.Instances {
				println(aws.StringValue(instance.InstanceId))
			}
		}
		if res.NextToken == nil {
			break
		}

		req = req.SetNextToken(aws.StringValue(res.NextToken))
	}
	return nil
}
