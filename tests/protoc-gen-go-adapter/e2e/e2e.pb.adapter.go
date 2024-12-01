//Code generated by protoc-gen-go-adapter. DO NOT EDIT.
//versions:
//- protoc-gen-go-adapter v1.0.13
//- protoc 5.28.3
//source: e2e.proto

package e2e

import (
	context "context"
	annotations "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
)

// GreeterAdapter  is the adapter for  Greeter  server into  Greeter  client interface
type GreeterAdapter struct {
	in GreeterServer
}

var _ GreeterClient = (*GreeterAdapter)(nil)

func NewGreeterAdapter(in GreeterServer) GreeterClient {
	return &GreeterAdapter{in: in}
}

func (a *GreeterAdapter) SayHello(ctx context.Context, req *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	return a.in.SayHello(ctx, req)
}
func (a *GreeterAdapter) SayHelloPost(ctx context.Context, req *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	return a.in.SayHelloPost(ctx, req)
}
func (a *GreeterAdapter) SayHttp(ctx context.Context, req *HelloRequest, opts ...grpc.CallOption) (*annotations.Http, error) {
	return a.in.SayHttp(ctx, req)
}

// GreeterGatewayClientAdapter  is the adapter for  Greeter  server into  Greeter  client interface
type GreeterGatewayClientAdapter struct {
	in GreeterGatewayClient
}

var _ GreeterClient = (*GreeterGatewayClientAdapter)(nil)

func NewGreeterGatewayClientAdapter(in GreeterGatewayClient) GreeterClient {
	return &GreeterGatewayClientAdapter{in: in}
}

func (a *GreeterGatewayClientAdapter) SayHello(ctx context.Context, req *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	return a.in.SayHello(ctx, req)
}
func (a *GreeterGatewayClientAdapter) SayHelloPost(ctx context.Context, req *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	return a.in.SayHelloPost(ctx, req)
}
func (a *GreeterGatewayClientAdapter) SayHttp(ctx context.Context, req *HelloRequest, opts ...grpc.CallOption) (*annotations.Http, error) {
	return a.in.SayHttp(ctx, req)
}