// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: spacemesh/v1/node.proto

package v1

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_spacemesh_v1_node_proto protoreflect.FileDescriptor

var file_spacemesh_v1_node_proto_rawDesc = []byte{
	0x0a, 0x17, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x76, 0x31, 0x2f, 0x6e,
	0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x2f,
	0x76, 0x31, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x32, 0xbc, 0x05, 0x0a, 0x0b, 0x4e, 0x6f, 0x64, 0x65, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x3d, 0x0a, 0x04, 0x45, 0x63, 0x68, 0x6f, 0x12, 0x19, 0x2e, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x63, 0x68, 0x6f, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65,
	0x73, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x63, 0x68, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x40, 0x0a, 0x07, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1d, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73,
	0x68, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c, 0x0a, 0x05, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x12, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1b, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73,
	0x68, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x43, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1b, 0x2e, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4c, 0x0a, 0x09, 0x53, 0x79, 0x6e, 0x63, 0x53,
	0x74, 0x61, 0x72, 0x74, 0x12, 0x1e, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x79, 0x6e, 0x63, 0x53, 0x74, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x79, 0x6e, 0x63, 0x53, 0x74, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x49, 0x0a, 0x08, 0x53, 0x68, 0x75, 0x74, 0x64, 0x6f, 0x77,
	0x6e, 0x12, 0x1d, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x31,
	0x2e, 0x53, 0x68, 0x75, 0x74, 0x64, 0x6f, 0x77, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1e, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x31, 0x2e,
	0x53, 0x68, 0x75, 0x74, 0x64, 0x6f, 0x77, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x61, 0x0a, 0x10, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x65, 0x74, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x12, 0x25, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68,
	0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x65, 0x74, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x50, 0x6f, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x57, 0x0a, 0x0c, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x53, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x12, 0x21, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x2e,
	0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65,
	0x73, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x53, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x12, 0x54, 0x0a, 0x0b,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x20, 0x2e, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x30, 0x01, 0x42, 0x34, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x73, 0x68, 0x6f, 0x73, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x2f, 0x67, 0x6f, 0x2f, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_spacemesh_v1_node_proto_goTypes = []interface{}{
	(*EchoRequest)(nil),              // 0: spacemesh.v1.EchoRequest
	(*empty.Empty)(nil),              // 1: google.protobuf.Empty
	(*StatusRequest)(nil),            // 2: spacemesh.v1.StatusRequest
	(*SyncStartRequest)(nil),         // 3: spacemesh.v1.SyncStartRequest
	(*ShutdownRequest)(nil),          // 4: spacemesh.v1.ShutdownRequest
	(*UpdatePoetServerRequest)(nil),  // 5: spacemesh.v1.UpdatePoetServerRequest
	(*StatusStreamRequest)(nil),      // 6: spacemesh.v1.StatusStreamRequest
	(*ErrorStreamRequest)(nil),       // 7: spacemesh.v1.ErrorStreamRequest
	(*EchoResponse)(nil),             // 8: spacemesh.v1.EchoResponse
	(*VersionResponse)(nil),          // 9: spacemesh.v1.VersionResponse
	(*BuildResponse)(nil),            // 10: spacemesh.v1.BuildResponse
	(*StatusResponse)(nil),           // 11: spacemesh.v1.StatusResponse
	(*SyncStartResponse)(nil),        // 12: spacemesh.v1.SyncStartResponse
	(*ShutdownResponse)(nil),         // 13: spacemesh.v1.ShutdownResponse
	(*UpdatePoetServerResponse)(nil), // 14: spacemesh.v1.UpdatePoetServerResponse
	(*StatusStreamResponse)(nil),     // 15: spacemesh.v1.StatusStreamResponse
	(*ErrorStreamResponse)(nil),      // 16: spacemesh.v1.ErrorStreamResponse
}
var file_spacemesh_v1_node_proto_depIdxs = []int32{
	0,  // 0: spacemesh.v1.NodeService.Echo:input_type -> spacemesh.v1.EchoRequest
	1,  // 1: spacemesh.v1.NodeService.Version:input_type -> google.protobuf.Empty
	1,  // 2: spacemesh.v1.NodeService.Build:input_type -> google.protobuf.Empty
	2,  // 3: spacemesh.v1.NodeService.Status:input_type -> spacemesh.v1.StatusRequest
	3,  // 4: spacemesh.v1.NodeService.SyncStart:input_type -> spacemesh.v1.SyncStartRequest
	4,  // 5: spacemesh.v1.NodeService.Shutdown:input_type -> spacemesh.v1.ShutdownRequest
	5,  // 6: spacemesh.v1.NodeService.UpdatePoetServer:input_type -> spacemesh.v1.UpdatePoetServerRequest
	6,  // 7: spacemesh.v1.NodeService.StatusStream:input_type -> spacemesh.v1.StatusStreamRequest
	7,  // 8: spacemesh.v1.NodeService.ErrorStream:input_type -> spacemesh.v1.ErrorStreamRequest
	8,  // 9: spacemesh.v1.NodeService.Echo:output_type -> spacemesh.v1.EchoResponse
	9,  // 10: spacemesh.v1.NodeService.Version:output_type -> spacemesh.v1.VersionResponse
	10, // 11: spacemesh.v1.NodeService.Build:output_type -> spacemesh.v1.BuildResponse
	11, // 12: spacemesh.v1.NodeService.Status:output_type -> spacemesh.v1.StatusResponse
	12, // 13: spacemesh.v1.NodeService.SyncStart:output_type -> spacemesh.v1.SyncStartResponse
	13, // 14: spacemesh.v1.NodeService.Shutdown:output_type -> spacemesh.v1.ShutdownResponse
	14, // 15: spacemesh.v1.NodeService.UpdatePoetServer:output_type -> spacemesh.v1.UpdatePoetServerResponse
	15, // 16: spacemesh.v1.NodeService.StatusStream:output_type -> spacemesh.v1.StatusStreamResponse
	16, // 17: spacemesh.v1.NodeService.ErrorStream:output_type -> spacemesh.v1.ErrorStreamResponse
	9,  // [9:18] is the sub-list for method output_type
	0,  // [0:9] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_spacemesh_v1_node_proto_init() }
func file_spacemesh_v1_node_proto_init() {
	if File_spacemesh_v1_node_proto != nil {
		return
	}
	file_spacemesh_v1_node_types_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_spacemesh_v1_node_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_spacemesh_v1_node_proto_goTypes,
		DependencyIndexes: file_spacemesh_v1_node_proto_depIdxs,
	}.Build()
	File_spacemesh_v1_node_proto = out.File
	file_spacemesh_v1_node_proto_rawDesc = nil
	file_spacemesh_v1_node_proto_goTypes = nil
	file_spacemesh_v1_node_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// NodeServiceClient is the client API for NodeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NodeServiceClient interface {
	// A simple test endpoint
	Echo(ctx context.Context, in *EchoRequest, opts ...grpc.CallOption) (*EchoResponse, error)
	// Returns the version of the node software as a semver string
	Version(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*VersionResponse, error)
	// Returns the github commit hash used to build the node
	Build(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*BuildResponse, error)
	// Current node status (net and sync)
	Status(ctx context.Context, in *StatusRequest, opts ...grpc.CallOption) (*StatusResponse, error)
	// Request that the node start syncing the mesh
	SyncStart(ctx context.Context, in *SyncStartRequest, opts ...grpc.CallOption) (*SyncStartResponse, error)
	// Request that the node initiate graceful shutdown
	Shutdown(ctx context.Context, in *ShutdownRequest, opts ...grpc.CallOption) (*ShutdownResponse, error)
	// UpdatePoetServer updates poet server
	UpdatePoetServer(ctx context.Context, in *UpdatePoetServerRequest, opts ...grpc.CallOption) (*UpdatePoetServerResponse, error)
	// Node status events (sync and net)
	StatusStream(ctx context.Context, in *StatusStreamRequest, opts ...grpc.CallOption) (NodeService_StatusStreamClient, error)
	// Node error events
	ErrorStream(ctx context.Context, in *ErrorStreamRequest, opts ...grpc.CallOption) (NodeService_ErrorStreamClient, error)
}

type nodeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewNodeServiceClient(cc grpc.ClientConnInterface) NodeServiceClient {
	return &nodeServiceClient{cc}
}

func (c *nodeServiceClient) Echo(ctx context.Context, in *EchoRequest, opts ...grpc.CallOption) (*EchoResponse, error) {
	out := new(EchoResponse)
	err := c.cc.Invoke(ctx, "/spacemesh.v1.NodeService/Echo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeServiceClient) Version(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*VersionResponse, error) {
	out := new(VersionResponse)
	err := c.cc.Invoke(ctx, "/spacemesh.v1.NodeService/Version", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeServiceClient) Build(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*BuildResponse, error) {
	out := new(BuildResponse)
	err := c.cc.Invoke(ctx, "/spacemesh.v1.NodeService/Build", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeServiceClient) Status(ctx context.Context, in *StatusRequest, opts ...grpc.CallOption) (*StatusResponse, error) {
	out := new(StatusResponse)
	err := c.cc.Invoke(ctx, "/spacemesh.v1.NodeService/Status", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeServiceClient) SyncStart(ctx context.Context, in *SyncStartRequest, opts ...grpc.CallOption) (*SyncStartResponse, error) {
	out := new(SyncStartResponse)
	err := c.cc.Invoke(ctx, "/spacemesh.v1.NodeService/SyncStart", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeServiceClient) Shutdown(ctx context.Context, in *ShutdownRequest, opts ...grpc.CallOption) (*ShutdownResponse, error) {
	out := new(ShutdownResponse)
	err := c.cc.Invoke(ctx, "/spacemesh.v1.NodeService/Shutdown", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeServiceClient) UpdatePoetServer(ctx context.Context, in *UpdatePoetServerRequest, opts ...grpc.CallOption) (*UpdatePoetServerResponse, error) {
	out := new(UpdatePoetServerResponse)
	err := c.cc.Invoke(ctx, "/spacemesh.v1.NodeService/UpdatePoetServer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeServiceClient) StatusStream(ctx context.Context, in *StatusStreamRequest, opts ...grpc.CallOption) (NodeService_StatusStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_NodeService_serviceDesc.Streams[0], "/spacemesh.v1.NodeService/StatusStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &nodeServiceStatusStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type NodeService_StatusStreamClient interface {
	Recv() (*StatusStreamResponse, error)
	grpc.ClientStream
}

type nodeServiceStatusStreamClient struct {
	grpc.ClientStream
}

func (x *nodeServiceStatusStreamClient) Recv() (*StatusStreamResponse, error) {
	m := new(StatusStreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *nodeServiceClient) ErrorStream(ctx context.Context, in *ErrorStreamRequest, opts ...grpc.CallOption) (NodeService_ErrorStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_NodeService_serviceDesc.Streams[1], "/spacemesh.v1.NodeService/ErrorStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &nodeServiceErrorStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type NodeService_ErrorStreamClient interface {
	Recv() (*ErrorStreamResponse, error)
	grpc.ClientStream
}

type nodeServiceErrorStreamClient struct {
	grpc.ClientStream
}

func (x *nodeServiceErrorStreamClient) Recv() (*ErrorStreamResponse, error) {
	m := new(ErrorStreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// NodeServiceServer is the server API for NodeService service.
type NodeServiceServer interface {
	// A simple test endpoint
	Echo(context.Context, *EchoRequest) (*EchoResponse, error)
	// Returns the version of the node software as a semver string
	Version(context.Context, *empty.Empty) (*VersionResponse, error)
	// Returns the github commit hash used to build the node
	Build(context.Context, *empty.Empty) (*BuildResponse, error)
	// Current node status (net and sync)
	Status(context.Context, *StatusRequest) (*StatusResponse, error)
	// Request that the node start syncing the mesh
	SyncStart(context.Context, *SyncStartRequest) (*SyncStartResponse, error)
	// Request that the node initiate graceful shutdown
	Shutdown(context.Context, *ShutdownRequest) (*ShutdownResponse, error)
	// UpdatePoetServer updates poet server
	UpdatePoetServer(context.Context, *UpdatePoetServerRequest) (*UpdatePoetServerResponse, error)
	// Node status events (sync and net)
	StatusStream(*StatusStreamRequest, NodeService_StatusStreamServer) error
	// Node error events
	ErrorStream(*ErrorStreamRequest, NodeService_ErrorStreamServer) error
}

// UnimplementedNodeServiceServer can be embedded to have forward compatible implementations.
type UnimplementedNodeServiceServer struct {
}

func (*UnimplementedNodeServiceServer) Echo(context.Context, *EchoRequest) (*EchoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Echo not implemented")
}
func (*UnimplementedNodeServiceServer) Version(context.Context, *empty.Empty) (*VersionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Version not implemented")
}
func (*UnimplementedNodeServiceServer) Build(context.Context, *empty.Empty) (*BuildResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Build not implemented")
}
func (*UnimplementedNodeServiceServer) Status(context.Context, *StatusRequest) (*StatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Status not implemented")
}
func (*UnimplementedNodeServiceServer) SyncStart(context.Context, *SyncStartRequest) (*SyncStartResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SyncStart not implemented")
}
func (*UnimplementedNodeServiceServer) Shutdown(context.Context, *ShutdownRequest) (*ShutdownResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Shutdown not implemented")
}
func (*UnimplementedNodeServiceServer) UpdatePoetServer(context.Context, *UpdatePoetServerRequest) (*UpdatePoetServerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePoetServer not implemented")
}
func (*UnimplementedNodeServiceServer) StatusStream(*StatusStreamRequest, NodeService_StatusStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method StatusStream not implemented")
}
func (*UnimplementedNodeServiceServer) ErrorStream(*ErrorStreamRequest, NodeService_ErrorStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method ErrorStream not implemented")
}

func RegisterNodeServiceServer(s *grpc.Server, srv NodeServiceServer) {
	s.RegisterService(&_NodeService_serviceDesc, srv)
}

func _NodeService_Echo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EchoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServiceServer).Echo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spacemesh.v1.NodeService/Echo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServiceServer).Echo(ctx, req.(*EchoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeService_Version_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServiceServer).Version(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spacemesh.v1.NodeService/Version",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServiceServer).Version(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeService_Build_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServiceServer).Build(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spacemesh.v1.NodeService/Build",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServiceServer).Build(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeService_Status_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServiceServer).Status(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spacemesh.v1.NodeService/Status",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServiceServer).Status(ctx, req.(*StatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeService_SyncStart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SyncStartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServiceServer).SyncStart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spacemesh.v1.NodeService/SyncStart",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServiceServer).SyncStart(ctx, req.(*SyncStartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeService_Shutdown_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShutdownRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServiceServer).Shutdown(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spacemesh.v1.NodeService/Shutdown",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServiceServer).Shutdown(ctx, req.(*ShutdownRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeService_UpdatePoetServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePoetServerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServiceServer).UpdatePoetServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spacemesh.v1.NodeService/UpdatePoetServer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServiceServer).UpdatePoetServer(ctx, req.(*UpdatePoetServerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeService_StatusStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StatusStreamRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(NodeServiceServer).StatusStream(m, &nodeServiceStatusStreamServer{stream})
}

type NodeService_StatusStreamServer interface {
	Send(*StatusStreamResponse) error
	grpc.ServerStream
}

type nodeServiceStatusStreamServer struct {
	grpc.ServerStream
}

func (x *nodeServiceStatusStreamServer) Send(m *StatusStreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _NodeService_ErrorStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ErrorStreamRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(NodeServiceServer).ErrorStream(m, &nodeServiceErrorStreamServer{stream})
}

type NodeService_ErrorStreamServer interface {
	Send(*ErrorStreamResponse) error
	grpc.ServerStream
}

type nodeServiceErrorStreamServer struct {
	grpc.ServerStream
}

func (x *nodeServiceErrorStreamServer) Send(m *ErrorStreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _NodeService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "spacemesh.v1.NodeService",
	HandlerType: (*NodeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Echo",
			Handler:    _NodeService_Echo_Handler,
		},
		{
			MethodName: "Version",
			Handler:    _NodeService_Version_Handler,
		},
		{
			MethodName: "Build",
			Handler:    _NodeService_Build_Handler,
		},
		{
			MethodName: "Status",
			Handler:    _NodeService_Status_Handler,
		},
		{
			MethodName: "SyncStart",
			Handler:    _NodeService_SyncStart_Handler,
		},
		{
			MethodName: "Shutdown",
			Handler:    _NodeService_Shutdown_Handler,
		},
		{
			MethodName: "UpdatePoetServer",
			Handler:    _NodeService_UpdatePoetServer_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StatusStream",
			Handler:       _NodeService_StatusStream_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ErrorStream",
			Handler:       _NodeService_ErrorStream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "spacemesh/v1/node.proto",
}
