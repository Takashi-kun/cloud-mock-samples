package main

import (
	"context"
	"errors"
	"fmt"
	"log"

	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2021-07-01/compute"
	"github.com/Azure/go-autorest/autorest/to"
	secretmanagerpb "google.golang.org/genproto/googleapis/cloud/secretmanager/v1"

	cfgV2 "github.com/aws/aws-sdk-go-v2/config"
	credV2 "github.com/aws/aws-sdk-go-v2/credentials"
	ec2V2 "github.com/aws/aws-sdk-go-v2/service/ec2"

	awsV1 "github.com/aws/aws-sdk-go/aws"
	credV1 "github.com/aws/aws-sdk-go/aws/credentials"
	sessV1 "github.com/aws/aws-sdk-go/aws/session"
	ec2v1 "github.com/aws/aws-sdk-go/service/ec2"
)

var (
	errMock = errors.New("mock error")
)

func sampleEC2V1() {
	sess, err := sessV1.NewSession(&awsV1.Config{
		Credentials: credV1.NewStaticCredentials("AWS_ACCESS_KEY_ID", "AWS_SECRET_ACCESS_KEY", "TOKEN"),
		Region:      awsV1.String("us-west-1"),
	})
	if err != nil {
		log.Fatalf("failed to create session: %v", err)
	}
	describeInstances(context.TODO(), ec2v1.New(sess), &ec2v1.DescribeInstancesInput{
		MaxResults: awsV1.Int64(100),
	})
}

func sampleEFSV2() {
	cfg, err := cfgV2.LoadDefaultConfig(
		context.Background(),
		cfgV2.WithRegion("us-west-2"),
		cfgV2.WithCredentialsProvider(
			credV2.NewStaticCredentialsProvider("AWS_ACCESS_KEY_ID", "AWS_SECRET_ACCESS_KEY", "TOKEN"),
		),
	)
	if err != nil {
		log.Fatalf("unable to load config, %v", err)
	}

	pager := ec2V2.NewDescribeInstancesPaginator(ec2V2.NewFromConfig(cfg), &ec2V2.DescribeInstancesInput{
		MaxResults: awsV1.Int32(100),
	})
	describeInstancesV2(context.TODO(), pager)
}

func sampleGCPSecretManager() {
	client, err := secretmanager.NewClient(context.TODO())
	if err != nil {
		log.Fatalf("unable to load config, %v", err)
	}
	sm := &secretManagerAPI{
		project: "cloud-mock-samples",
		client:  client,
	}
	if err := sm.listSecretManagers(context.TODO(), &secretmanagerpb.ListSecretsRequest{}); err != nil {
		log.Fatalf("failed to list: %v", err)
	}
}

func sampleAzureVM() {
	pager, err := listAzureVM(context.TODO(), compute.NewVirtualMachinesClient("SUBSCRIPTION_ID"))
	if err != nil {
		log.Fatalf("failed to list Azure VM: %v", err)
	}

	for pager.NotDone() {
		for _, vm := range pager.Values() {
			fmt.Println(to.String(vm.Name))
		}

		if err := pager.NextWithContext(context.TODO()); err != nil {
			log.Fatalf("failed to get next: %v", err)
		}
	}
}
