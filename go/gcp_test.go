package main

import (
	"context"
	"net"
	"testing"

	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	"golang.org/x/net/nettest"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	secretmanagerpb "google.golang.org/genproto/googleapis/cloud/secretmanager/v1"
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

func (s *testInternalSMServer) CreateSecret(context.Context, *secretmanagerpb.CreateSecretRequest) (*secretmanagerpb.Secret, error) {
	return nil, status.Errorf(codes.Unimplemented, "not implemented yet")
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

func TestSecretManagerAPI_CreateSecretManager(t *testing.T) {
	internal := testNewSecretManagerClient(t)
	defer internal.shutdown()

	_, err := internal.api.createSecretManager(context.TODO(), &secretmanagerpb.CreateSecretRequest{})
	if err != nil {
		t.Fatalf("expect no error, got %v", err) // comes here
	}
}
