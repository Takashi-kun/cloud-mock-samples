package main

import (
	"context"
	"net"
	"testing"

	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	secretmanagerpb "google.golang.org/genproto/googleapis/cloud/secretmanager/v1"

	"golang.org/x/net/nettest"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type (
	testInternalSMServer struct {
		*secretmanagerpb.UnimplementedSecretManagerServiceServer

		api    *secretManagerAPI
		server *grpc.Server
	}
)

var _ secretmanagerpb.SecretManagerServiceServer = (*testInternalSMServer)(nil)

func (s *testInternalSMServer) ListSecrets(context.Context, *secretmanagerpb.ListSecretsRequest) (*secretmanagerpb.ListSecretsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, errMock.Error())
}

func (s *testInternalSMServer) shutdown() {
	s.api.close()
	s.server.Stop()
}

func prepareServer(t *testing.T, n net.Listener) *testInternalSMServer {
	t.Helper()

	server := grpc.NewServer()
	internal := &testInternalSMServer{
		server: server,
	}
	secretmanagerpb.RegisterSecretManagerServiceServer(server, internal)
	go server.Serve(n)

	return internal
}

func testGRPCClientOptions(t *testing.T, n net.Listener) []option.ClientOption {
	t.Helper()

	conn, err := grpc.Dial(n.Addr().String(), grpc.WithInsecure())
	if err != nil {
		t.Fatal(err)
	}

	return []option.ClientOption{
		internaloption.WithDefaultEndpoint(n.Addr().String()),
		option.WithGRPCConn(conn),
	}
}

func testNewSecretManagerClient(t *testing.T) *testInternalSMServer {
	t.Helper()

	n, err := nettest.NewLocalListener("tcp")
	if err != nil {
		t.Fatal("failed to prepare local listener")
	}

	internal := prepareServer(t, n)
	options := testGRPCClientOptions(t, n)
	client, err := secretmanager.NewClient(context.TODO(), options...)
	if err != nil {
		t.Fatalf("failed to create secretmanager client: %v", err)
	}
	internal.api = &secretManagerAPI{
		project: "p",
		client:  client,
	}
	return internal
}

func TestSecretManagerAPI_listSecretManagers(t *testing.T) {
	internal := testNewSecretManagerClient(t)
	defer internal.shutdown()

	if err := internal.api.listSecretManagers(context.TODO(), &secretmanagerpb.ListSecretsRequest{}); err != nil {
		t.Fatalf("expect no error, got %v", err) // comes here
	}
}
