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
	v1 "google.golang.org/genproto/googleapis/iam/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type (
	testInternalSMServer struct {
		api    *secretManagerAPI
		server *grpc.Server
	}
)

var _ secretmanagerpb.SecretManagerServiceServer = (*testInternalSMServer)(nil)

func (s *testInternalSMServer) ListSecrets(context.Context, *secretmanagerpb.ListSecretsRequest) (*secretmanagerpb.ListSecretsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "not implemented yet")
}
func (s *testInternalSMServer) CreateSecret(context.Context, *secretmanagerpb.CreateSecretRequest) (*secretmanagerpb.Secret, error) {
	return nil, status.Errorf(codes.Unimplemented, "not implemented yet")
}
func (s *testInternalSMServer) AddSecretVersion(context.Context, *secretmanagerpb.AddSecretVersionRequest) (*secretmanagerpb.SecretVersion, error) {
	return nil, status.Errorf(codes.Unimplemented, "not implemented yet")
}
func (s *testInternalSMServer) GetSecret(context.Context, *secretmanagerpb.GetSecretRequest) (*secretmanagerpb.Secret, error) {
	return nil, status.Errorf(codes.Unimplemented, "not implemented yet")
}
func (s *testInternalSMServer) UpdateSecret(context.Context, *secretmanagerpb.UpdateSecretRequest) (*secretmanagerpb.Secret, error) {
	return nil, status.Errorf(codes.Unimplemented, "not implemented yet")
}
func (s *testInternalSMServer) DeleteSecret(context.Context, *secretmanagerpb.DeleteSecretRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "not implemented yet")
}
func (s *testInternalSMServer) ListSecretVersions(context.Context, *secretmanagerpb.ListSecretVersionsRequest) (*secretmanagerpb.ListSecretVersionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "not implemented yet")
}
func (s *testInternalSMServer) GetSecretVersion(context.Context, *secretmanagerpb.GetSecretVersionRequest) (*secretmanagerpb.SecretVersion, error) {
	return nil, status.Errorf(codes.Unimplemented, "not implemented yet")
}
func (s *testInternalSMServer) AccessSecretVersion(context.Context, *secretmanagerpb.AccessSecretVersionRequest) (*secretmanagerpb.AccessSecretVersionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "not implemented yet")
}
func (s *testInternalSMServer) DisableSecretVersion(context.Context, *secretmanagerpb.DisableSecretVersionRequest) (*secretmanagerpb.SecretVersion, error) {
	return nil, status.Errorf(codes.Unimplemented, "not implemented yet")
}
func (s *testInternalSMServer) EnableSecretVersion(context.Context, *secretmanagerpb.EnableSecretVersionRequest) (*secretmanagerpb.SecretVersion, error) {
	return nil, status.Errorf(codes.Unimplemented, "not implemented yet")
}
func (s *testInternalSMServer) DestroySecretVersion(context.Context, *secretmanagerpb.DestroySecretVersionRequest) (*secretmanagerpb.SecretVersion, error) {
	return nil, status.Errorf(codes.Unimplemented, "not implemented yet")
}
func (s *testInternalSMServer) SetIamPolicy(context.Context, *v1.SetIamPolicyRequest) (*v1.Policy, error) {
	return nil, status.Errorf(codes.Unimplemented, "not implemented yet")
}
func (s *testInternalSMServer) GetIamPolicy(context.Context, *v1.GetIamPolicyRequest) (*v1.Policy, error) {
	return nil, status.Errorf(codes.Unimplemented, "not implemented yet")
}
func (s *testInternalSMServer) TestIamPermissions(context.Context, *v1.TestIamPermissionsRequest) (*v1.TestIamPermissionsResponse, error) {
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
