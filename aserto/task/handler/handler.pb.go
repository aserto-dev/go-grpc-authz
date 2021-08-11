// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.0
// 	protoc        v3.17.3
// source: aserto/task/handler/handler.proto

package handler

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type HandleTaskRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Evt *anypb.Any `protobuf:"bytes,1,opt,name=evt,proto3" json:"evt,omitempty"`
}

func (x *HandleTaskRequest) Reset() {
	*x = HandleTaskRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aserto_task_handler_handler_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HandleTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HandleTaskRequest) ProtoMessage() {}

func (x *HandleTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_task_handler_handler_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HandleTaskRequest.ProtoReflect.Descriptor instead.
func (*HandleTaskRequest) Descriptor() ([]byte, []int) {
	return file_aserto_task_handler_handler_proto_rawDescGZIP(), []int{0}
}

func (x *HandleTaskRequest) GetEvt() *anypb.Any {
	if x != nil {
		return x.Evt
	}
	return nil
}

type HandleTaskResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result *emptypb.Empty `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *HandleTaskResponse) Reset() {
	*x = HandleTaskResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aserto_task_handler_handler_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HandleTaskResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HandleTaskResponse) ProtoMessage() {}

func (x *HandleTaskResponse) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_task_handler_handler_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HandleTaskResponse.ProtoReflect.Descriptor instead.
func (*HandleTaskResponse) Descriptor() ([]byte, []int) {
	return file_aserto_task_handler_handler_proto_rawDescGZIP(), []int{1}
}

func (x *HandleTaskResponse) GetResult() *emptypb.Empty {
	if x != nil {
		return x.Result
	}
	return nil
}

type HandleJobRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Evt *anypb.Any `protobuf:"bytes,1,opt,name=evt,proto3" json:"evt,omitempty"`
}

func (x *HandleJobRequest) Reset() {
	*x = HandleJobRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aserto_task_handler_handler_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HandleJobRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HandleJobRequest) ProtoMessage() {}

func (x *HandleJobRequest) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_task_handler_handler_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HandleJobRequest.ProtoReflect.Descriptor instead.
func (*HandleJobRequest) Descriptor() ([]byte, []int) {
	return file_aserto_task_handler_handler_proto_rawDescGZIP(), []int{2}
}

func (x *HandleJobRequest) GetEvt() *anypb.Any {
	if x != nil {
		return x.Evt
	}
	return nil
}

type HandleJobResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result *emptypb.Empty `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *HandleJobResponse) Reset() {
	*x = HandleJobResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aserto_task_handler_handler_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HandleJobResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HandleJobResponse) ProtoMessage() {}

func (x *HandleJobResponse) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_task_handler_handler_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HandleJobResponse.ProtoReflect.Descriptor instead.
func (*HandleJobResponse) Descriptor() ([]byte, []int) {
	return file_aserto_task_handler_handler_proto_rawDescGZIP(), []int{3}
}

func (x *HandleJobResponse) GetResult() *emptypb.Empty {
	if x != nil {
		return x.Result
	}
	return nil
}

var File_aserto_task_handler_handler_proto protoreflect.FileDescriptor

var file_aserto_task_handler_handler_proto_rawDesc = []byte{
	0x0a, 0x21, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x2f, 0x68, 0x61,
	0x6e, 0x64, 0x6c, 0x65, 0x72, 0x2f, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x13, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e, 0x74, 0x61, 0x73, 0x6b,
	0x2e, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3b,
	0x0a, 0x11, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x26, 0x0a, 0x03, 0x65, 0x76, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x03, 0x65, 0x76, 0x74, 0x22, 0x44, 0x0a, 0x12, 0x48,
	0x61, 0x6e, 0x64, 0x6c, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x2e, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x22, 0x3a, 0x0a, 0x10, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x4a, 0x6f, 0x62, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x26, 0x0a, 0x03, 0x65, 0x76, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x03, 0x65, 0x76, 0x74, 0x22, 0x43, 0x0a,
	0x11, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x2e, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x32, 0xf3, 0x01, 0x0a, 0x07, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x12, 0x75,
	0x0a, 0x0a, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x26, 0x2e, 0x61,
	0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x68, 0x61, 0x6e, 0x64, 0x6c,
	0x65, 0x72, 0x2e, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e, 0x74, 0x61,
	0x73, 0x6b, 0x2e, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x2e, 0x48, 0x61, 0x6e, 0x64, 0x6c,
	0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x16, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x10, 0x22, 0x0b, 0x2f, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x74, 0x61,
	0x73, 0x6b, 0x3a, 0x01, 0x2a, 0x12, 0x71, 0x0a, 0x09, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x4a,
	0x6f, 0x62, 0x12, 0x25, 0x2e, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e, 0x74, 0x61, 0x73, 0x6b,
	0x2e, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x2e, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x4a,
	0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x61, 0x73, 0x65, 0x72,
	0x74, 0x6f, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x2e,
	0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x15, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x22, 0x0a, 0x2f, 0x68, 0x61, 0x6e, 0x64,
	0x6c, 0x65, 0x6a, 0x6f, 0x62, 0x3a, 0x01, 0x2a, 0x42, 0x39, 0x5a, 0x37, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2d, 0x64, 0x65,
	0x76, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2f, 0x74,
	0x61, 0x73, 0x6b, 0x2f, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x3b, 0x68, 0x61, 0x6e, 0x64,
	0x6c, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_aserto_task_handler_handler_proto_rawDescOnce sync.Once
	file_aserto_task_handler_handler_proto_rawDescData = file_aserto_task_handler_handler_proto_rawDesc
)

func file_aserto_task_handler_handler_proto_rawDescGZIP() []byte {
	file_aserto_task_handler_handler_proto_rawDescOnce.Do(func() {
		file_aserto_task_handler_handler_proto_rawDescData = protoimpl.X.CompressGZIP(file_aserto_task_handler_handler_proto_rawDescData)
	})
	return file_aserto_task_handler_handler_proto_rawDescData
}

var file_aserto_task_handler_handler_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_aserto_task_handler_handler_proto_goTypes = []interface{}{
	(*HandleTaskRequest)(nil),  // 0: aserto.task.handler.HandleTaskRequest
	(*HandleTaskResponse)(nil), // 1: aserto.task.handler.HandleTaskResponse
	(*HandleJobRequest)(nil),   // 2: aserto.task.handler.HandleJobRequest
	(*HandleJobResponse)(nil),  // 3: aserto.task.handler.HandleJobResponse
	(*anypb.Any)(nil),          // 4: google.protobuf.Any
	(*emptypb.Empty)(nil),      // 5: google.protobuf.Empty
}
var file_aserto_task_handler_handler_proto_depIdxs = []int32{
	4, // 0: aserto.task.handler.HandleTaskRequest.evt:type_name -> google.protobuf.Any
	5, // 1: aserto.task.handler.HandleTaskResponse.result:type_name -> google.protobuf.Empty
	4, // 2: aserto.task.handler.HandleJobRequest.evt:type_name -> google.protobuf.Any
	5, // 3: aserto.task.handler.HandleJobResponse.result:type_name -> google.protobuf.Empty
	0, // 4: aserto.task.handler.Handler.HandleTask:input_type -> aserto.task.handler.HandleTaskRequest
	2, // 5: aserto.task.handler.Handler.HandleJob:input_type -> aserto.task.handler.HandleJobRequest
	1, // 6: aserto.task.handler.Handler.HandleTask:output_type -> aserto.task.handler.HandleTaskResponse
	3, // 7: aserto.task.handler.Handler.HandleJob:output_type -> aserto.task.handler.HandleJobResponse
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_aserto_task_handler_handler_proto_init() }
func file_aserto_task_handler_handler_proto_init() {
	if File_aserto_task_handler_handler_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_aserto_task_handler_handler_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HandleTaskRequest); i {
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
		file_aserto_task_handler_handler_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HandleTaskResponse); i {
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
		file_aserto_task_handler_handler_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HandleJobRequest); i {
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
		file_aserto_task_handler_handler_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HandleJobResponse); i {
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
			RawDescriptor: file_aserto_task_handler_handler_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_aserto_task_handler_handler_proto_goTypes,
		DependencyIndexes: file_aserto_task_handler_handler_proto_depIdxs,
		MessageInfos:      file_aserto_task_handler_handler_proto_msgTypes,
	}.Build()
	File_aserto_task_handler_handler_proto = out.File
	file_aserto_task_handler_handler_proto_rawDesc = nil
	file_aserto_task_handler_handler_proto_goTypes = nil
	file_aserto_task_handler_handler_proto_depIdxs = nil
}
