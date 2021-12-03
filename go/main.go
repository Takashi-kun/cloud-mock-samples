package main

import (
	"context"
	"fmt"
	"log"

	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	secretmanagerpb "google.golang.org/genproto/googleapis/cloud/secretmanager/v1"

	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2021-07-01/compute"
	"github.com/Azure/go-autorest/autorest/to"

	cfgV2 "github.com/aws/aws-sdk-go-v2/config"
	credV2 "github.com/aws/aws-sdk-go-v2/credentials"
	efsV2 "github.com/aws/aws-sdk-go-v2/service/efs"

	awsV1 "github.com/aws/aws-sdk-go/aws"
	credV1 "github.com/aws/aws-sdk-go/aws/credentials"
	sessV1 "github.com/aws/aws-sdk-go/aws/session"
	efsV1 "github.com/aws/aws-sdk-go/service/efs"
)

func sampleEFSV1() {
	sess, err := sessV1.NewSession(&awsV1.Config{
		Credentials: credV1.NewStaticCredentials("AWS_ACCESS_KEY_ID", "AWS_SECRET_ACCESS_KEY", "TOKEN"),
		Region:      awsV1.String("us-west-1"),
	})
	if err != nil {
		log.Fatalf("failed to create session: %v", err)
	}
	describeFileSystems(efsV1.New(sess), &efsV1.DescribeFileSystemsInput{
		MaxItems: awsV1.Int64(100),
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

	pager := efsV2.NewDescribeFileSystemsPaginator(efsV2.NewFromConfig(cfg), &efsV2.DescribeFileSystemsInput{
		MaxItems: awsV1.Int32(100),
	})
	describeFileSystemsV2(context.TODO(), pager)
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
	sm.createSecretManager(context.TODO(), &secretmanagerpb.CreateSecretRequest{
		SecretId: "sample",
	})
}

func sampleAzureVM() {
	pager, err := listAzureVM(compute.NewVirtualMachinesClient("SUBSCRIPTION_ID"))
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
