// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        (unknown)
// source: py-rpc/proto/filestream.proto

package file_stream

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type InputFrame struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileName string `protobuf:"bytes,1,opt,name=file_name,json=fileName,proto3" json:"file_name,omitempty"`
	FileType string `protobuf:"bytes,2,opt,name=file_type,json=fileType,proto3" json:"file_type,omitempty"`
	Data     []byte `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	UserId   string `protobuf:"bytes,5,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *InputFrame) Reset() {
	*x = InputFrame{}
	if protoimpl.UnsafeEnabled {
		mi := &file_py_rpc_proto_filestream_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InputFrame) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InputFrame) ProtoMessage() {}

func (x *InputFrame) ProtoReflect() protoreflect.Message {
	mi := &file_py_rpc_proto_filestream_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InputFrame.ProtoReflect.Descriptor instead.
func (*InputFrame) Descriptor() ([]byte, []int) {
	return file_py_rpc_proto_filestream_proto_rawDescGZIP(), []int{0}
}

func (x *InputFrame) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

func (x *InputFrame) GetFileType() string {
	if x != nil {
		return x.FileType
	}
	return ""
}

func (x *InputFrame) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *InputFrame) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type OutputFrame struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ValidData   bool   `protobuf:"varint,1,opt,name=valid_data,json=validData,proto3" json:"valid_data,omitempty"`
	Message     string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	FilePath    string `protobuf:"bytes,3,opt,name=file_path,json=filePath,proto3" json:"file_path,omitempty"`
	FileEncrypt string `protobuf:"bytes,4,opt,name=file_encrypt,json=fileEncrypt,proto3" json:"file_encrypt,omitempty"`
	Rows        int64  `protobuf:"varint,5,opt,name=rows,proto3" json:"rows,omitempty"`
	Cols        int64  `protobuf:"varint,6,opt,name=cols,proto3" json:"cols,omitempty"`
}

func (x *OutputFrame) Reset() {
	*x = OutputFrame{}
	if protoimpl.UnsafeEnabled {
		mi := &file_py_rpc_proto_filestream_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OutputFrame) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OutputFrame) ProtoMessage() {}

func (x *OutputFrame) ProtoReflect() protoreflect.Message {
	mi := &file_py_rpc_proto_filestream_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OutputFrame.ProtoReflect.Descriptor instead.
func (*OutputFrame) Descriptor() ([]byte, []int) {
	return file_py_rpc_proto_filestream_proto_rawDescGZIP(), []int{1}
}

func (x *OutputFrame) GetValidData() bool {
	if x != nil {
		return x.ValidData
	}
	return false
}

func (x *OutputFrame) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *OutputFrame) GetFilePath() string {
	if x != nil {
		return x.FilePath
	}
	return ""
}

func (x *OutputFrame) GetFileEncrypt() string {
	if x != nil {
		return x.FileEncrypt
	}
	return ""
}

func (x *OutputFrame) GetRows() int64 {
	if x != nil {
		return x.Rows
	}
	return 0
}

func (x *OutputFrame) GetCols() int64 {
	if x != nil {
		return x.Cols
	}
	return 0
}

var File_py_rpc_proto_filestream_proto protoreflect.FileDescriptor

var file_py_rpc_proto_filestream_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x70, 0x79, 0x2d, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x66,
	0x69, 0x6c, 0x65, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0b, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x22, 0x74, 0x0a, 0x0b,
	0x69, 0x6e, 0x70, 0x75, 0x74, 0x5f, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x66,
	0x69, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x22, 0xaf, 0x01, 0x0a, 0x0c, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x5f, 0x66, 0x72,
	0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x5f, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x44, 0x61,
	0x74, 0x61, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1b, 0x0a, 0x09,
	0x66, 0x69, 0x6c, 0x65, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x66, 0x69, 0x6c, 0x65, 0x50, 0x61, 0x74, 0x68, 0x12, 0x21, 0x0a, 0x0c, 0x66, 0x69, 0x6c,
	0x65, 0x5f, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x66, 0x69, 0x6c, 0x65, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x72, 0x6f, 0x77, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x72, 0x6f, 0x77, 0x73,
	0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x6c, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04,
	0x63, 0x6f, 0x6c, 0x73, 0x32, 0x59, 0x0a, 0x0c, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f, 0x69,
	0x6e, 0x70, 0x75, 0x74, 0x12, 0x49, 0x0a, 0x10, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x44,
	0x61, 0x74, 0x61, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x5f,
	0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x5f, 0x66, 0x72, 0x61,
	0x6d, 0x65, 0x1a, 0x19, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x2e, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x5f, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x22, 0x00, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_py_rpc_proto_filestream_proto_rawDescOnce sync.Once
	file_py_rpc_proto_filestream_proto_rawDescData = file_py_rpc_proto_filestream_proto_rawDesc
)

func file_py_rpc_proto_filestream_proto_rawDescGZIP() []byte {
	file_py_rpc_proto_filestream_proto_rawDescOnce.Do(func() {
		file_py_rpc_proto_filestream_proto_rawDescData = protoimpl.X.CompressGZIP(file_py_rpc_proto_filestream_proto_rawDescData)
	})
	return file_py_rpc_proto_filestream_proto_rawDescData
}

var file_py_rpc_proto_filestream_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_py_rpc_proto_filestream_proto_goTypes = []interface{}{
	(*InputFrame)(nil),  // 0: file_stream.input_frame
	(*OutputFrame)(nil), // 1: file_stream.output_frame
}
var file_py_rpc_proto_filestream_proto_depIdxs = []int32{
	0, // 0: file_stream.stream_input.ConvertDataframe:input_type -> file_stream.input_frame
	1, // 1: file_stream.stream_input.ConvertDataframe:output_type -> file_stream.output_frame
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_py_rpc_proto_filestream_proto_init() }
func file_py_rpc_proto_filestream_proto_init() {
	if File_py_rpc_proto_filestream_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_py_rpc_proto_filestream_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InputFrame); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_py_rpc_proto_filestream_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OutputFrame); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_py_rpc_proto_filestream_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_py_rpc_proto_filestream_proto_goTypes,
		DependencyIndexes: file_py_rpc_proto_filestream_proto_depIdxs,
		MessageInfos:      file_py_rpc_proto_filestream_proto_msgTypes,
	}.Build()
	File_py_rpc_proto_filestream_proto = out.File
	file_py_rpc_proto_filestream_proto_rawDesc = nil
	file_py_rpc_proto_filestream_proto_goTypes = nil
	file_py_rpc_proto_filestream_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// StreamInputClient is the client API for StreamInput service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StreamInputClient interface {
	ConvertDataframe(ctx context.Context, in *InputFrame, opts ...grpc.CallOption) (*OutputFrame, error)
}

type streamInputClient struct {
	cc grpc.ClientConnInterface
}

func NewStreamInputClient(cc grpc.ClientConnInterface) StreamInputClient {
	return &streamInputClient{cc}
}

func (c *streamInputClient) ConvertDataframe(ctx context.Context, in *InputFrame, opts ...grpc.CallOption) (*OutputFrame, error) {
	out := new(OutputFrame)
	err := c.cc.Invoke(ctx, "/file_stream.stream_input/ConvertDataframe", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StreamInputServer is the server API for StreamInput service.
type StreamInputServer interface {
	ConvertDataframe(context.Context, *InputFrame) (*OutputFrame, error)
}

// UnimplementedStreamInputServer can be embedded to have forward compatible implementations.
type UnimplementedStreamInputServer struct {
}

func (*UnimplementedStreamInputServer) ConvertDataframe(context.Context, *InputFrame) (*OutputFrame, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConvertDataframe not implemented")
}

func RegisterStreamInputServer(s *grpc.Server, srv StreamInputServer) {
	s.RegisterService(&_StreamInput_serviceDesc, srv)
}

func _StreamInput_ConvertDataframe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InputFrame)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StreamInputServer).ConvertDataframe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/file_stream.stream_input/ConvertDataframe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StreamInputServer).ConvertDataframe(ctx, req.(*InputFrame))
	}
	return interceptor(ctx, in, info, handler)
}

var _StreamInput_serviceDesc = grpc.ServiceDesc{
	ServiceName: "file_stream.stream_input",
	HandlerType: (*StreamInputServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ConvertDataframe",
			Handler:    _StreamInput_ConvertDataframe_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "py-rpc/proto/filestream.proto",
}
