package main

import (
	"context"

	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	secretmanagerpb "google.golang.org/genproto/googleapis/cloud/secretmanager/v1"
)

type (
	secretManagerAPI struct {
		project string
		client  *secretmanager.Client
	}
)

func (s *secretManagerAPI) close() error {
	return nil
}

func (s *secretManagerAPI) createSecretManager(ctx context.Context, req *secretmanagerpb.CreateSecretRequest) (*secretmanagerpb.Secret, error) {
	return s.client.CreateSecret(ctx, req)
}
