// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package delegatedidentityv1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// DelegatedIdentityClient is the client API for DelegatedIdentity service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DelegatedIdentityClient interface {
	// Subscribe to get X.509-SVIDs for workloads that match the given selectors.
	// The lifetime of the subscription aligns to the lifetime of the stream.
	SubscribeToX509SVIDs(ctx context.Context, in *SubscribeToX509SVIDsRequest, opts ...grpc.CallOption) (DelegatedIdentity_SubscribeToX509SVIDsClient, error)
	// Subscribe to get X.509-SVIDs for workloads that match the given selectors.
	// The lifetime of the subscription aligns to the lifetime of the stream.
	//
	// Subscribe to get local and all federated bundles.
	// The lifetime of the subscription aligns to the lifetime of the stream.
	SubscribeToX509Bundles(ctx context.Context, in *SubscribeToX509BundlesRequest, opts ...grpc.CallOption) (DelegatedIdentity_SubscribeToX509BundlesClient, error)
	// Fetch JWT-SVIDs for workloads that match the given selectors, and
	// for the requested audience.
	FetchJWTSVIDs(ctx context.Context, in *FetchJWTSVIDsRequest, opts ...grpc.CallOption) (*FetchJWTSVIDsResponse, error)
	// Subscribe to get local and all federated JWKS bundles.
	// The lifetime of the subscription aligns to the lifetime of the stream.
	SubscribeToJWTBundles(ctx context.Context, in *SubscribeToJWTBundlesRequest, opts ...grpc.CallOption) (DelegatedIdentity_SubscribeToJWTBundlesClient, error)
}

type delegatedIdentityClient struct {
	cc grpc.ClientConnInterface
}

func NewDelegatedIdentityClient(cc grpc.ClientConnInterface) DelegatedIdentityClient {
	return &delegatedIdentityClient{cc}
}

func (c *delegatedIdentityClient) SubscribeToX509SVIDs(ctx context.Context, in *SubscribeToX509SVIDsRequest, opts ...grpc.CallOption) (DelegatedIdentity_SubscribeToX509SVIDsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_DelegatedIdentity_serviceDesc.Streams[0], "/spire.api.agent.delegatedidentity.v1.DelegatedIdentity/SubscribeToX509SVIDs", opts...)
	if err != nil {
		return nil, err
	}
	x := &delegatedIdentitySubscribeToX509SVIDsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type DelegatedIdentity_SubscribeToX509SVIDsClient interface {
	Recv() (*SubscribeToX509SVIDsResponse, error)
	grpc.ClientStream
}

type delegatedIdentitySubscribeToX509SVIDsClient struct {
	grpc.ClientStream
}

func (x *delegatedIdentitySubscribeToX509SVIDsClient) Recv() (*SubscribeToX509SVIDsResponse, error) {
	m := new(SubscribeToX509SVIDsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *delegatedIdentityClient) SubscribeToX509Bundles(ctx context.Context, in *SubscribeToX509BundlesRequest, opts ...grpc.CallOption) (DelegatedIdentity_SubscribeToX509BundlesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_DelegatedIdentity_serviceDesc.Streams[1], "/spire.api.agent.delegatedidentity.v1.DelegatedIdentity/SubscribeToX509Bundles", opts...)
	if err != nil {
		return nil, err
	}
	x := &delegatedIdentitySubscribeToX509BundlesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type DelegatedIdentity_SubscribeToX509BundlesClient interface {
	Recv() (*SubscribeToX509BundlesResponse, error)
	grpc.ClientStream
}

type delegatedIdentitySubscribeToX509BundlesClient struct {
	grpc.ClientStream
}

func (x *delegatedIdentitySubscribeToX509BundlesClient) Recv() (*SubscribeToX509BundlesResponse, error) {
	m := new(SubscribeToX509BundlesResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *delegatedIdentityClient) FetchJWTSVIDs(ctx context.Context, in *FetchJWTSVIDsRequest, opts ...grpc.CallOption) (*FetchJWTSVIDsResponse, error) {
	out := new(FetchJWTSVIDsResponse)
	err := c.cc.Invoke(ctx, "/spire.api.agent.delegatedidentity.v1.DelegatedIdentity/FetchJWTSVIDs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *delegatedIdentityClient) SubscribeToJWTBundles(ctx context.Context, in *SubscribeToJWTBundlesRequest, opts ...grpc.CallOption) (DelegatedIdentity_SubscribeToJWTBundlesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_DelegatedIdentity_serviceDesc.Streams[2], "/spire.api.agent.delegatedidentity.v1.DelegatedIdentity/SubscribeToJWTBundles", opts...)
	if err != nil {
		return nil, err
	}
	x := &delegatedIdentitySubscribeToJWTBundlesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type DelegatedIdentity_SubscribeToJWTBundlesClient interface {
	Recv() (*SubscribeToJWTBundlesResponse, error)
	grpc.ClientStream
}

type delegatedIdentitySubscribeToJWTBundlesClient struct {
	grpc.ClientStream
}

func (x *delegatedIdentitySubscribeToJWTBundlesClient) Recv() (*SubscribeToJWTBundlesResponse, error) {
	m := new(SubscribeToJWTBundlesResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// DelegatedIdentityServer is the server API for DelegatedIdentity service.
// All implementations must embed UnimplementedDelegatedIdentityServer
// for forward compatibility
type DelegatedIdentityServer interface {
	// Subscribe to get X.509-SVIDs for workloads that match the given selectors.
	// The lifetime of the subscription aligns to the lifetime of the stream.
	SubscribeToX509SVIDs(*SubscribeToX509SVIDsRequest, DelegatedIdentity_SubscribeToX509SVIDsServer) error
	// Subscribe to get X.509-SVIDs for workloads that match the given selectors.
	// The lifetime of the subscription aligns to the lifetime of the stream.
	//
	// Subscribe to get local and all federated bundles.
	// The lifetime of the subscription aligns to the lifetime of the stream.
	SubscribeToX509Bundles(*SubscribeToX509BundlesRequest, DelegatedIdentity_SubscribeToX509BundlesServer) error
	// Fetch JWT-SVIDs for workloads that match the given selectors, and
	// for the requested audience.
	FetchJWTSVIDs(context.Context, *FetchJWTSVIDsRequest) (*FetchJWTSVIDsResponse, error)
	// Subscribe to get local and all federated JWKS bundles.
	// The lifetime of the subscription aligns to the lifetime of the stream.
	SubscribeToJWTBundles(*SubscribeToJWTBundlesRequest, DelegatedIdentity_SubscribeToJWTBundlesServer) error
	mustEmbedUnimplementedDelegatedIdentityServer()
}

// UnimplementedDelegatedIdentityServer must be embedded to have forward compatible implementations.
type UnimplementedDelegatedIdentityServer struct {
}

func (UnimplementedDelegatedIdentityServer) SubscribeToX509SVIDs(*SubscribeToX509SVIDsRequest, DelegatedIdentity_SubscribeToX509SVIDsServer) error {
	return status.Errorf(codes.Unimplemented, "method SubscribeToX509SVIDs not implemented")
}
func (UnimplementedDelegatedIdentityServer) SubscribeToX509Bundles(*SubscribeToX509BundlesRequest, DelegatedIdentity_SubscribeToX509BundlesServer) error {
	return status.Errorf(codes.Unimplemented, "method SubscribeToX509Bundles not implemented")
}
func (UnimplementedDelegatedIdentityServer) FetchJWTSVIDs(context.Context, *FetchJWTSVIDsRequest) (*FetchJWTSVIDsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchJWTSVIDs not implemented")
}
func (UnimplementedDelegatedIdentityServer) SubscribeToJWTBundles(*SubscribeToJWTBundlesRequest, DelegatedIdentity_SubscribeToJWTBundlesServer) error {
	return status.Errorf(codes.Unimplemented, "method SubscribeToJWTBundles not implemented")
}
func (UnimplementedDelegatedIdentityServer) mustEmbedUnimplementedDelegatedIdentityServer() {}

// UnsafeDelegatedIdentityServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DelegatedIdentityServer will
// result in compilation errors.
type UnsafeDelegatedIdentityServer interface {
	mustEmbedUnimplementedDelegatedIdentityServer()
}

func RegisterDelegatedIdentityServer(s grpc.ServiceRegistrar, srv DelegatedIdentityServer) {
	s.RegisterService(&_DelegatedIdentity_serviceDesc, srv)
}

func _DelegatedIdentity_SubscribeToX509SVIDs_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SubscribeToX509SVIDsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DelegatedIdentityServer).SubscribeToX509SVIDs(m, &delegatedIdentitySubscribeToX509SVIDsServer{stream})
}

type DelegatedIdentity_SubscribeToX509SVIDsServer interface {
	Send(*SubscribeToX509SVIDsResponse) error
	grpc.ServerStream
}

type delegatedIdentitySubscribeToX509SVIDsServer struct {
	grpc.ServerStream
}

func (x *delegatedIdentitySubscribeToX509SVIDsServer) Send(m *SubscribeToX509SVIDsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _DelegatedIdentity_SubscribeToX509Bundles_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SubscribeToX509BundlesRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DelegatedIdentityServer).SubscribeToX509Bundles(m, &delegatedIdentitySubscribeToX509BundlesServer{stream})
}

type DelegatedIdentity_SubscribeToX509BundlesServer interface {
	Send(*SubscribeToX509BundlesResponse) error
	grpc.ServerStream
}

type delegatedIdentitySubscribeToX509BundlesServer struct {
	grpc.ServerStream
}

func (x *delegatedIdentitySubscribeToX509BundlesServer) Send(m *SubscribeToX509BundlesResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _DelegatedIdentity_FetchJWTSVIDs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchJWTSVIDsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DelegatedIdentityServer).FetchJWTSVIDs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spire.api.agent.delegatedidentity.v1.DelegatedIdentity/FetchJWTSVIDs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DelegatedIdentityServer).FetchJWTSVIDs(ctx, req.(*FetchJWTSVIDsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DelegatedIdentity_SubscribeToJWTBundles_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SubscribeToJWTBundlesRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DelegatedIdentityServer).SubscribeToJWTBundles(m, &delegatedIdentitySubscribeToJWTBundlesServer{stream})
}

type DelegatedIdentity_SubscribeToJWTBundlesServer interface {
	Send(*SubscribeToJWTBundlesResponse) error
	grpc.ServerStream
}

type delegatedIdentitySubscribeToJWTBundlesServer struct {
	grpc.ServerStream
}

func (x *delegatedIdentitySubscribeToJWTBundlesServer) Send(m *SubscribeToJWTBundlesResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _DelegatedIdentity_serviceDesc = grpc.ServiceDesc{
	ServiceName: "spire.api.agent.delegatedidentity.v1.DelegatedIdentity",
	HandlerType: (*DelegatedIdentityServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FetchJWTSVIDs",
			Handler:    _DelegatedIdentity_FetchJWTSVIDs_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SubscribeToX509SVIDs",
			Handler:       _DelegatedIdentity_SubscribeToX509SVIDs_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "SubscribeToX509Bundles",
			Handler:       _DelegatedIdentity_SubscribeToX509Bundles_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "SubscribeToJWTBundles",
			Handler:       _DelegatedIdentity_SubscribeToJWTBundles_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "spire/api/agent/delegatedidentity/v1/delegatedidentity.proto",
}
