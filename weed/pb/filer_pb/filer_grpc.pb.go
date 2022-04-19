// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package filer_pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// SeaweedFilerClient is the client API for SeaweedFiler service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SeaweedFilerClient interface {
	LookupDirectoryEntry(ctx context.Context, in *LookupDirectoryEntryRequest, opts ...grpc.CallOption) (*LookupDirectoryEntryResponse, error)
	ListEntries(ctx context.Context, in *ListEntriesRequest, opts ...grpc.CallOption) (SeaweedFiler_ListEntriesClient, error)
	CreateEntry(ctx context.Context, in *CreateEntryRequest, opts ...grpc.CallOption) (*CreateEntryResponse, error)
	UpdateEntry(ctx context.Context, in *UpdateEntryRequest, opts ...grpc.CallOption) (*UpdateEntryResponse, error)
	AppendToEntry(ctx context.Context, in *AppendToEntryRequest, opts ...grpc.CallOption) (*AppendToEntryResponse, error)
	DeleteEntry(ctx context.Context, in *DeleteEntryRequest, opts ...grpc.CallOption) (*DeleteEntryResponse, error)
	AtomicRenameEntry(ctx context.Context, in *AtomicRenameEntryRequest, opts ...grpc.CallOption) (*AtomicRenameEntryResponse, error)
	StreamRenameEntry(ctx context.Context, in *StreamRenameEntryRequest, opts ...grpc.CallOption) (SeaweedFiler_StreamRenameEntryClient, error)
	AssignVolume(ctx context.Context, in *AssignVolumeRequest, opts ...grpc.CallOption) (*AssignVolumeResponse, error)
	LookupVolume(ctx context.Context, in *LookupVolumeRequest, opts ...grpc.CallOption) (*LookupVolumeResponse, error)
	CollectionList(ctx context.Context, in *CollectionListRequest, opts ...grpc.CallOption) (*CollectionListResponse, error)
	DeleteCollection(ctx context.Context, in *DeleteCollectionRequest, opts ...grpc.CallOption) (*DeleteCollectionResponse, error)
	Statistics(ctx context.Context, in *StatisticsRequest, opts ...grpc.CallOption) (*StatisticsResponse, error)
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error)
	GetFilerConfiguration(ctx context.Context, in *GetFilerConfigurationRequest, opts ...grpc.CallOption) (*GetFilerConfigurationResponse, error)
	SubscribeMetadata(ctx context.Context, in *SubscribeMetadataRequest, opts ...grpc.CallOption) (SeaweedFiler_SubscribeMetadataClient, error)
	SubscribeLocalMetadata(ctx context.Context, in *SubscribeMetadataRequest, opts ...grpc.CallOption) (SeaweedFiler_SubscribeLocalMetadataClient, error)
	KeepConnected(ctx context.Context, opts ...grpc.CallOption) (SeaweedFiler_KeepConnectedClient, error)
	LocateBroker(ctx context.Context, in *LocateBrokerRequest, opts ...grpc.CallOption) (*LocateBrokerResponse, error)
	KvGet(ctx context.Context, in *KvGetRequest, opts ...grpc.CallOption) (*KvGetResponse, error)
	KvPut(ctx context.Context, in *KvPutRequest, opts ...grpc.CallOption) (*KvPutResponse, error)
	CacheRemoteObjectToLocalCluster(ctx context.Context, in *CacheRemoteObjectToLocalClusterRequest, opts ...grpc.CallOption) (*CacheRemoteObjectToLocalClusterResponse, error)
}

type seaweedFilerClient struct {
	cc grpc.ClientConnInterface
}

func NewSeaweedFilerClient(cc grpc.ClientConnInterface) SeaweedFilerClient {
	return &seaweedFilerClient{cc}
}

func (c *seaweedFilerClient) LookupDirectoryEntry(ctx context.Context, in *LookupDirectoryEntryRequest, opts ...grpc.CallOption) (*LookupDirectoryEntryResponse, error) {
	out := new(LookupDirectoryEntryResponse)
	err := c.cc.Invoke(ctx, "/filer_pb.SeaweedFiler/LookupDirectoryEntry", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *seaweedFilerClient) ListEntries(ctx context.Context, in *ListEntriesRequest, opts ...grpc.CallOption) (SeaweedFiler_ListEntriesClient, error) {
	stream, err := c.cc.NewStream(ctx, &SeaweedFiler_ServiceDesc.Streams[0], "/filer_pb.SeaweedFiler/ListEntries", opts...)
	if err != nil {
		return nil, err
	}
	x := &seaweedFilerListEntriesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type SeaweedFiler_ListEntriesClient interface {
	Recv() (*ListEntriesResponse, error)
	grpc.ClientStream
}

type seaweedFilerListEntriesClient struct {
	grpc.ClientStream
}

func (x *seaweedFilerListEntriesClient) Recv() (*ListEntriesResponse, error) {
	m := new(ListEntriesResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *seaweedFilerClient) CreateEntry(ctx context.Context, in *CreateEntryRequest, opts ...grpc.CallOption) (*CreateEntryResponse, error) {
	out := new(CreateEntryResponse)
	err := c.cc.Invoke(ctx, "/filer_pb.SeaweedFiler/CreateEntry", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *seaweedFilerClient) UpdateEntry(ctx context.Context, in *UpdateEntryRequest, opts ...grpc.CallOption) (*UpdateEntryResponse, error) {
	out := new(UpdateEntryResponse)
	err := c.cc.Invoke(ctx, "/filer_pb.SeaweedFiler/UpdateEntry", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *seaweedFilerClient) AppendToEntry(ctx context.Context, in *AppendToEntryRequest, opts ...grpc.CallOption) (*AppendToEntryResponse, error) {
	out := new(AppendToEntryResponse)
	err := c.cc.Invoke(ctx, "/filer_pb.SeaweedFiler/AppendToEntry", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *seaweedFilerClient) DeleteEntry(ctx context.Context, in *DeleteEntryRequest, opts ...grpc.CallOption) (*DeleteEntryResponse, error) {
	out := new(DeleteEntryResponse)
	err := c.cc.Invoke(ctx, "/filer_pb.SeaweedFiler/DeleteEntry", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *seaweedFilerClient) AtomicRenameEntry(ctx context.Context, in *AtomicRenameEntryRequest, opts ...grpc.CallOption) (*AtomicRenameEntryResponse, error) {
	out := new(AtomicRenameEntryResponse)
	err := c.cc.Invoke(ctx, "/filer_pb.SeaweedFiler/AtomicRenameEntry", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *seaweedFilerClient) StreamRenameEntry(ctx context.Context, in *StreamRenameEntryRequest, opts ...grpc.CallOption) (SeaweedFiler_StreamRenameEntryClient, error) {
	stream, err := c.cc.NewStream(ctx, &SeaweedFiler_ServiceDesc.Streams[1], "/filer_pb.SeaweedFiler/StreamRenameEntry", opts...)
	if err != nil {
		return nil, err
	}
	x := &seaweedFilerStreamRenameEntryClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type SeaweedFiler_StreamRenameEntryClient interface {
	Recv() (*StreamRenameEntryResponse, error)
	grpc.ClientStream
}

type seaweedFilerStreamRenameEntryClient struct {
	grpc.ClientStream
}

func (x *seaweedFilerStreamRenameEntryClient) Recv() (*StreamRenameEntryResponse, error) {
	m := new(StreamRenameEntryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *seaweedFilerClient) AssignVolume(ctx context.Context, in *AssignVolumeRequest, opts ...grpc.CallOption) (*AssignVolumeResponse, error) {
	out := new(AssignVolumeResponse)
	err := c.cc.Invoke(ctx, "/filer_pb.SeaweedFiler/AssignVolume", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *seaweedFilerClient) LookupVolume(ctx context.Context, in *LookupVolumeRequest, opts ...grpc.CallOption) (*LookupVolumeResponse, error) {
	out := new(LookupVolumeResponse)
	err := c.cc.Invoke(ctx, "/filer_pb.SeaweedFiler/LookupVolume", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *seaweedFilerClient) CollectionList(ctx context.Context, in *CollectionListRequest, opts ...grpc.CallOption) (*CollectionListResponse, error) {
	out := new(CollectionListResponse)
	err := c.cc.Invoke(ctx, "/filer_pb.SeaweedFiler/CollectionList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *seaweedFilerClient) DeleteCollection(ctx context.Context, in *DeleteCollectionRequest, opts ...grpc.CallOption) (*DeleteCollectionResponse, error) {
	out := new(DeleteCollectionResponse)
	err := c.cc.Invoke(ctx, "/filer_pb.SeaweedFiler/DeleteCollection", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *seaweedFilerClient) Statistics(ctx context.Context, in *StatisticsRequest, opts ...grpc.CallOption) (*StatisticsResponse, error) {
	out := new(StatisticsResponse)
	err := c.cc.Invoke(ctx, "/filer_pb.SeaweedFiler/Statistics", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *seaweedFilerClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error) {
	out := new(PingResponse)
	err := c.cc.Invoke(ctx, "/filer_pb.SeaweedFiler/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *seaweedFilerClient) GetFilerConfiguration(ctx context.Context, in *GetFilerConfigurationRequest, opts ...grpc.CallOption) (*GetFilerConfigurationResponse, error) {
	out := new(GetFilerConfigurationResponse)
	err := c.cc.Invoke(ctx, "/filer_pb.SeaweedFiler/GetFilerConfiguration", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *seaweedFilerClient) SubscribeMetadata(ctx context.Context, in *SubscribeMetadataRequest, opts ...grpc.CallOption) (SeaweedFiler_SubscribeMetadataClient, error) {
	stream, err := c.cc.NewStream(ctx, &SeaweedFiler_ServiceDesc.Streams[2], "/filer_pb.SeaweedFiler/SubscribeMetadata", opts...)
	if err != nil {
		return nil, err
	}
	x := &seaweedFilerSubscribeMetadataClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type SeaweedFiler_SubscribeMetadataClient interface {
	Recv() (*SubscribeMetadataResponse, error)
	grpc.ClientStream
}

type seaweedFilerSubscribeMetadataClient struct {
	grpc.ClientStream
}

func (x *seaweedFilerSubscribeMetadataClient) Recv() (*SubscribeMetadataResponse, error) {
	m := new(SubscribeMetadataResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *seaweedFilerClient) SubscribeLocalMetadata(ctx context.Context, in *SubscribeMetadataRequest, opts ...grpc.CallOption) (SeaweedFiler_SubscribeLocalMetadataClient, error) {
	stream, err := c.cc.NewStream(ctx, &SeaweedFiler_ServiceDesc.Streams[3], "/filer_pb.SeaweedFiler/SubscribeLocalMetadata", opts...)
	if err != nil {
		return nil, err
	}
	x := &seaweedFilerSubscribeLocalMetadataClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type SeaweedFiler_SubscribeLocalMetadataClient interface {
	Recv() (*SubscribeMetadataResponse, error)
	grpc.ClientStream
}

type seaweedFilerSubscribeLocalMetadataClient struct {
	grpc.ClientStream
}

func (x *seaweedFilerSubscribeLocalMetadataClient) Recv() (*SubscribeMetadataResponse, error) {
	m := new(SubscribeMetadataResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *seaweedFilerClient) KeepConnected(ctx context.Context, opts ...grpc.CallOption) (SeaweedFiler_KeepConnectedClient, error) {
	stream, err := c.cc.NewStream(ctx, &SeaweedFiler_ServiceDesc.Streams[4], "/filer_pb.SeaweedFiler/KeepConnected", opts...)
	if err != nil {
		return nil, err
	}
	x := &seaweedFilerKeepConnectedClient{stream}
	return x, nil
}

type SeaweedFiler_KeepConnectedClient interface {
	Send(*KeepConnectedRequest) error
	Recv() (*KeepConnectedResponse, error)
	grpc.ClientStream
}

type seaweedFilerKeepConnectedClient struct {
	grpc.ClientStream
}

func (x *seaweedFilerKeepConnectedClient) Send(m *KeepConnectedRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *seaweedFilerKeepConnectedClient) Recv() (*KeepConnectedResponse, error) {
	m := new(KeepConnectedResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *seaweedFilerClient) LocateBroker(ctx context.Context, in *LocateBrokerRequest, opts ...grpc.CallOption) (*LocateBrokerResponse, error) {
	out := new(LocateBrokerResponse)
	err := c.cc.Invoke(ctx, "/filer_pb.SeaweedFiler/LocateBroker", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *seaweedFilerClient) KvGet(ctx context.Context, in *KvGetRequest, opts ...grpc.CallOption) (*KvGetResponse, error) {
	out := new(KvGetResponse)
	err := c.cc.Invoke(ctx, "/filer_pb.SeaweedFiler/KvGet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *seaweedFilerClient) KvPut(ctx context.Context, in *KvPutRequest, opts ...grpc.CallOption) (*KvPutResponse, error) {
	out := new(KvPutResponse)
	err := c.cc.Invoke(ctx, "/filer_pb.SeaweedFiler/KvPut", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *seaweedFilerClient) CacheRemoteObjectToLocalCluster(ctx context.Context, in *CacheRemoteObjectToLocalClusterRequest, opts ...grpc.CallOption) (*CacheRemoteObjectToLocalClusterResponse, error) {
	out := new(CacheRemoteObjectToLocalClusterResponse)
	err := c.cc.Invoke(ctx, "/filer_pb.SeaweedFiler/CacheRemoteObjectToLocalCluster", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SeaweedFilerServer is the server API for SeaweedFiler service.
// All implementations must embed UnimplementedSeaweedFilerServer
// for forward compatibility
type SeaweedFilerServer interface {
	LookupDirectoryEntry(context.Context, *LookupDirectoryEntryRequest) (*LookupDirectoryEntryResponse, error)
	ListEntries(*ListEntriesRequest, SeaweedFiler_ListEntriesServer) error
	CreateEntry(context.Context, *CreateEntryRequest) (*CreateEntryResponse, error)
	UpdateEntry(context.Context, *UpdateEntryRequest) (*UpdateEntryResponse, error)
	AppendToEntry(context.Context, *AppendToEntryRequest) (*AppendToEntryResponse, error)
	DeleteEntry(context.Context, *DeleteEntryRequest) (*DeleteEntryResponse, error)
	AtomicRenameEntry(context.Context, *AtomicRenameEntryRequest) (*AtomicRenameEntryResponse, error)
	StreamRenameEntry(*StreamRenameEntryRequest, SeaweedFiler_StreamRenameEntryServer) error
	AssignVolume(context.Context, *AssignVolumeRequest) (*AssignVolumeResponse, error)
	LookupVolume(context.Context, *LookupVolumeRequest) (*LookupVolumeResponse, error)
	CollectionList(context.Context, *CollectionListRequest) (*CollectionListResponse, error)
	DeleteCollection(context.Context, *DeleteCollectionRequest) (*DeleteCollectionResponse, error)
	Statistics(context.Context, *StatisticsRequest) (*StatisticsResponse, error)
	Ping(context.Context, *PingRequest) (*PingResponse, error)
	GetFilerConfiguration(context.Context, *GetFilerConfigurationRequest) (*GetFilerConfigurationResponse, error)
	SubscribeMetadata(*SubscribeMetadataRequest, SeaweedFiler_SubscribeMetadataServer) error
	SubscribeLocalMetadata(*SubscribeMetadataRequest, SeaweedFiler_SubscribeLocalMetadataServer) error
	KeepConnected(SeaweedFiler_KeepConnectedServer) error
	LocateBroker(context.Context, *LocateBrokerRequest) (*LocateBrokerResponse, error)
	KvGet(context.Context, *KvGetRequest) (*KvGetResponse, error)
	KvPut(context.Context, *KvPutRequest) (*KvPutResponse, error)
	CacheRemoteObjectToLocalCluster(context.Context, *CacheRemoteObjectToLocalClusterRequest) (*CacheRemoteObjectToLocalClusterResponse, error)
	mustEmbedUnimplementedSeaweedFilerServer()
}

// UnimplementedSeaweedFilerServer must be embedded to have forward compatible implementations.
type UnimplementedSeaweedFilerServer struct {
}

func (UnimplementedSeaweedFilerServer) LookupDirectoryEntry(context.Context, *LookupDirectoryEntryRequest) (*LookupDirectoryEntryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LookupDirectoryEntry not implemented")
}
func (UnimplementedSeaweedFilerServer) ListEntries(*ListEntriesRequest, SeaweedFiler_ListEntriesServer) error {
	return status.Errorf(codes.Unimplemented, "method ListEntries not implemented")
}
func (UnimplementedSeaweedFilerServer) CreateEntry(context.Context, *CreateEntryRequest) (*CreateEntryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateEntry not implemented")
}
func (UnimplementedSeaweedFilerServer) UpdateEntry(context.Context, *UpdateEntryRequest) (*UpdateEntryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateEntry not implemented")
}
func (UnimplementedSeaweedFilerServer) AppendToEntry(context.Context, *AppendToEntryRequest) (*AppendToEntryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AppendToEntry not implemented")
}
func (UnimplementedSeaweedFilerServer) DeleteEntry(context.Context, *DeleteEntryRequest) (*DeleteEntryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteEntry not implemented")
}
func (UnimplementedSeaweedFilerServer) AtomicRenameEntry(context.Context, *AtomicRenameEntryRequest) (*AtomicRenameEntryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AtomicRenameEntry not implemented")
}
func (UnimplementedSeaweedFilerServer) StreamRenameEntry(*StreamRenameEntryRequest, SeaweedFiler_StreamRenameEntryServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamRenameEntry not implemented")
}
func (UnimplementedSeaweedFilerServer) AssignVolume(context.Context, *AssignVolumeRequest) (*AssignVolumeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AssignVolume not implemented")
}
func (UnimplementedSeaweedFilerServer) LookupVolume(context.Context, *LookupVolumeRequest) (*LookupVolumeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LookupVolume not implemented")
}
func (UnimplementedSeaweedFilerServer) CollectionList(context.Context, *CollectionListRequest) (*CollectionListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CollectionList not implemented")
}
func (UnimplementedSeaweedFilerServer) DeleteCollection(context.Context, *DeleteCollectionRequest) (*DeleteCollectionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCollection not implemented")
}
func (UnimplementedSeaweedFilerServer) Statistics(context.Context, *StatisticsRequest) (*StatisticsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Statistics not implemented")
}
func (UnimplementedSeaweedFilerServer) Ping(context.Context, *PingRequest) (*PingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedSeaweedFilerServer) GetFilerConfiguration(context.Context, *GetFilerConfigurationRequest) (*GetFilerConfigurationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFilerConfiguration not implemented")
}
func (UnimplementedSeaweedFilerServer) SubscribeMetadata(*SubscribeMetadataRequest, SeaweedFiler_SubscribeMetadataServer) error {
	return status.Errorf(codes.Unimplemented, "method SubscribeMetadata not implemented")
}
func (UnimplementedSeaweedFilerServer) SubscribeLocalMetadata(*SubscribeMetadataRequest, SeaweedFiler_SubscribeLocalMetadataServer) error {
	return status.Errorf(codes.Unimplemented, "method SubscribeLocalMetadata not implemented")
}
func (UnimplementedSeaweedFilerServer) KeepConnected(SeaweedFiler_KeepConnectedServer) error {
	return status.Errorf(codes.Unimplemented, "method KeepConnected not implemented")
}
func (UnimplementedSeaweedFilerServer) LocateBroker(context.Context, *LocateBrokerRequest) (*LocateBrokerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LocateBroker not implemented")
}
func (UnimplementedSeaweedFilerServer) KvGet(context.Context, *KvGetRequest) (*KvGetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method KvGet not implemented")
}
func (UnimplementedSeaweedFilerServer) KvPut(context.Context, *KvPutRequest) (*KvPutResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method KvPut not implemented")
}
func (UnimplementedSeaweedFilerServer) CacheRemoteObjectToLocalCluster(context.Context, *CacheRemoteObjectToLocalClusterRequest) (*CacheRemoteObjectToLocalClusterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CacheRemoteObjectToLocalCluster not implemented")
}
func (UnimplementedSeaweedFilerServer) mustEmbedUnimplementedSeaweedFilerServer() {}

// UnsafeSeaweedFilerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SeaweedFilerServer will
// result in compilation errors.
type UnsafeSeaweedFilerServer interface {
	mustEmbedUnimplementedSeaweedFilerServer()
}

func RegisterSeaweedFilerServer(s grpc.ServiceRegistrar, srv SeaweedFilerServer) {
	s.RegisterService(&SeaweedFiler_ServiceDesc, srv)
}

func _SeaweedFiler_LookupDirectoryEntry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LookupDirectoryEntryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SeaweedFilerServer).LookupDirectoryEntry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/filer_pb.SeaweedFiler/LookupDirectoryEntry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SeaweedFilerServer).LookupDirectoryEntry(ctx, req.(*LookupDirectoryEntryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SeaweedFiler_ListEntries_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListEntriesRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SeaweedFilerServer).ListEntries(m, &seaweedFilerListEntriesServer{stream})
}

type SeaweedFiler_ListEntriesServer interface {
	Send(*ListEntriesResponse) error
	grpc.ServerStream
}

type seaweedFilerListEntriesServer struct {
	grpc.ServerStream
}

func (x *seaweedFilerListEntriesServer) Send(m *ListEntriesResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _SeaweedFiler_CreateEntry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateEntryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SeaweedFilerServer).CreateEntry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/filer_pb.SeaweedFiler/CreateEntry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SeaweedFilerServer).CreateEntry(ctx, req.(*CreateEntryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SeaweedFiler_UpdateEntry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateEntryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SeaweedFilerServer).UpdateEntry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/filer_pb.SeaweedFiler/UpdateEntry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SeaweedFilerServer).UpdateEntry(ctx, req.(*UpdateEntryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SeaweedFiler_AppendToEntry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AppendToEntryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SeaweedFilerServer).AppendToEntry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/filer_pb.SeaweedFiler/AppendToEntry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SeaweedFilerServer).AppendToEntry(ctx, req.(*AppendToEntryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SeaweedFiler_DeleteEntry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteEntryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SeaweedFilerServer).DeleteEntry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/filer_pb.SeaweedFiler/DeleteEntry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SeaweedFilerServer).DeleteEntry(ctx, req.(*DeleteEntryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SeaweedFiler_AtomicRenameEntry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AtomicRenameEntryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SeaweedFilerServer).AtomicRenameEntry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/filer_pb.SeaweedFiler/AtomicRenameEntry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SeaweedFilerServer).AtomicRenameEntry(ctx, req.(*AtomicRenameEntryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SeaweedFiler_StreamRenameEntry_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StreamRenameEntryRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SeaweedFilerServer).StreamRenameEntry(m, &seaweedFilerStreamRenameEntryServer{stream})
}

type SeaweedFiler_StreamRenameEntryServer interface {
	Send(*StreamRenameEntryResponse) error
	grpc.ServerStream
}

type seaweedFilerStreamRenameEntryServer struct {
	grpc.ServerStream
}

func (x *seaweedFilerStreamRenameEntryServer) Send(m *StreamRenameEntryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _SeaweedFiler_AssignVolume_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AssignVolumeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SeaweedFilerServer).AssignVolume(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/filer_pb.SeaweedFiler/AssignVolume",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SeaweedFilerServer).AssignVolume(ctx, req.(*AssignVolumeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SeaweedFiler_LookupVolume_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LookupVolumeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SeaweedFilerServer).LookupVolume(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/filer_pb.SeaweedFiler/LookupVolume",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SeaweedFilerServer).LookupVolume(ctx, req.(*LookupVolumeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SeaweedFiler_CollectionList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CollectionListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SeaweedFilerServer).CollectionList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/filer_pb.SeaweedFiler/CollectionList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SeaweedFilerServer).CollectionList(ctx, req.(*CollectionListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SeaweedFiler_DeleteCollection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCollectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SeaweedFilerServer).DeleteCollection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/filer_pb.SeaweedFiler/DeleteCollection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SeaweedFilerServer).DeleteCollection(ctx, req.(*DeleteCollectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SeaweedFiler_Statistics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StatisticsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SeaweedFilerServer).Statistics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/filer_pb.SeaweedFiler/Statistics",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SeaweedFilerServer).Statistics(ctx, req.(*StatisticsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SeaweedFiler_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SeaweedFilerServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/filer_pb.SeaweedFiler/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SeaweedFilerServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SeaweedFiler_GetFilerConfiguration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFilerConfigurationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SeaweedFilerServer).GetFilerConfiguration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/filer_pb.SeaweedFiler/GetFilerConfiguration",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SeaweedFilerServer).GetFilerConfiguration(ctx, req.(*GetFilerConfigurationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SeaweedFiler_SubscribeMetadata_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SubscribeMetadataRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SeaweedFilerServer).SubscribeMetadata(m, &seaweedFilerSubscribeMetadataServer{stream})
}

type SeaweedFiler_SubscribeMetadataServer interface {
	Send(*SubscribeMetadataResponse) error
	grpc.ServerStream
}

type seaweedFilerSubscribeMetadataServer struct {
	grpc.ServerStream
}

func (x *seaweedFilerSubscribeMetadataServer) Send(m *SubscribeMetadataResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _SeaweedFiler_SubscribeLocalMetadata_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SubscribeMetadataRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SeaweedFilerServer).SubscribeLocalMetadata(m, &seaweedFilerSubscribeLocalMetadataServer{stream})
}

type SeaweedFiler_SubscribeLocalMetadataServer interface {
	Send(*SubscribeMetadataResponse) error
	grpc.ServerStream
}

type seaweedFilerSubscribeLocalMetadataServer struct {
	grpc.ServerStream
}

func (x *seaweedFilerSubscribeLocalMetadataServer) Send(m *SubscribeMetadataResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _SeaweedFiler_KeepConnected_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SeaweedFilerServer).KeepConnected(&seaweedFilerKeepConnectedServer{stream})
}

type SeaweedFiler_KeepConnectedServer interface {
	Send(*KeepConnectedResponse) error
	Recv() (*KeepConnectedRequest, error)
	grpc.ServerStream
}

type seaweedFilerKeepConnectedServer struct {
	grpc.ServerStream
}

func (x *seaweedFilerKeepConnectedServer) Send(m *KeepConnectedResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *seaweedFilerKeepConnectedServer) Recv() (*KeepConnectedRequest, error) {
	m := new(KeepConnectedRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _SeaweedFiler_LocateBroker_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LocateBrokerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SeaweedFilerServer).LocateBroker(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/filer_pb.SeaweedFiler/LocateBroker",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SeaweedFilerServer).LocateBroker(ctx, req.(*LocateBrokerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SeaweedFiler_KvGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KvGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SeaweedFilerServer).KvGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/filer_pb.SeaweedFiler/KvGet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SeaweedFilerServer).KvGet(ctx, req.(*KvGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SeaweedFiler_KvPut_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KvPutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SeaweedFilerServer).KvPut(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/filer_pb.SeaweedFiler/KvPut",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SeaweedFilerServer).KvPut(ctx, req.(*KvPutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SeaweedFiler_CacheRemoteObjectToLocalCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CacheRemoteObjectToLocalClusterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SeaweedFilerServer).CacheRemoteObjectToLocalCluster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/filer_pb.SeaweedFiler/CacheRemoteObjectToLocalCluster",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SeaweedFilerServer).CacheRemoteObjectToLocalCluster(ctx, req.(*CacheRemoteObjectToLocalClusterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SeaweedFiler_ServiceDesc is the grpc.ServiceDesc for SeaweedFiler service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SeaweedFiler_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "filer_pb.SeaweedFiler",
	HandlerType: (*SeaweedFilerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "LookupDirectoryEntry",
			Handler:    _SeaweedFiler_LookupDirectoryEntry_Handler,
		},
		{
			MethodName: "CreateEntry",
			Handler:    _SeaweedFiler_CreateEntry_Handler,
		},
		{
			MethodName: "UpdateEntry",
			Handler:    _SeaweedFiler_UpdateEntry_Handler,
		},
		{
			MethodName: "AppendToEntry",
			Handler:    _SeaweedFiler_AppendToEntry_Handler,
		},
		{
			MethodName: "DeleteEntry",
			Handler:    _SeaweedFiler_DeleteEntry_Handler,
		},
		{
			MethodName: "AtomicRenameEntry",
			Handler:    _SeaweedFiler_AtomicRenameEntry_Handler,
		},
		{
			MethodName: "AssignVolume",
			Handler:    _SeaweedFiler_AssignVolume_Handler,
		},
		{
			MethodName: "LookupVolume",
			Handler:    _SeaweedFiler_LookupVolume_Handler,
		},
		{
			MethodName: "CollectionList",
			Handler:    _SeaweedFiler_CollectionList_Handler,
		},
		{
			MethodName: "DeleteCollection",
			Handler:    _SeaweedFiler_DeleteCollection_Handler,
		},
		{
			MethodName: "Statistics",
			Handler:    _SeaweedFiler_Statistics_Handler,
		},
		{
			MethodName: "Ping",
			Handler:    _SeaweedFiler_Ping_Handler,
		},
		{
			MethodName: "GetFilerConfiguration",
			Handler:    _SeaweedFiler_GetFilerConfiguration_Handler,
		},
		{
			MethodName: "LocateBroker",
			Handler:    _SeaweedFiler_LocateBroker_Handler,
		},
		{
			MethodName: "KvGet",
			Handler:    _SeaweedFiler_KvGet_Handler,
		},
		{
			MethodName: "KvPut",
			Handler:    _SeaweedFiler_KvPut_Handler,
		},
		{
			MethodName: "CacheRemoteObjectToLocalCluster",
			Handler:    _SeaweedFiler_CacheRemoteObjectToLocalCluster_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListEntries",
			Handler:       _SeaweedFiler_ListEntries_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "StreamRenameEntry",
			Handler:       _SeaweedFiler_StreamRenameEntry_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "SubscribeMetadata",
			Handler:       _SeaweedFiler_SubscribeMetadata_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "SubscribeLocalMetadata",
			Handler:       _SeaweedFiler_SubscribeLocalMetadata_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "KeepConnected",
			Handler:       _SeaweedFiler_KeepConnected_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "filer.proto",
}
