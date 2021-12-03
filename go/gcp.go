package main

import (
	"context"
	"fmt"

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

func (s *secretManagerAPI) listSecretManagers(ctx context.Context, req *secretmanagerpb.ListSecretsRequest) error {
	iter := s.client.ListSecrets(ctx, req)
	for {
		secretManagers, token, err := iter.InternalFetch(100, req.PageToken)
		if err != nil {
			return err
		}

		for _, sm := range secretManagers {
			fmt.Println(sm.Name)
		}
		if token == "" {
			break
		}
		req.PageToken = token
	}
	return nil
}
