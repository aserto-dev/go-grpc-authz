// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.0
// 	protoc        v3.17.3
// source: aserto/api/invite.proto

package api

import (
	_ "github.com/aserto-dev/go-grpc-server/aserto/options"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type InviteStatus int32

const (
	InviteStatus_INVITE_STATUS_UNKNOWN InviteStatus = 0 // Unknown invite status
	InviteStatus_ACCEPTED              InviteStatus = 1 // Normal attribute
	InviteStatus_DECLINED              InviteStatus = 2 // User declined the invitation
	InviteStatus_EXPIRED               InviteStatus = 3 // Invitation expired
	InviteStatus_CANCELED              InviteStatus = 4 // Invitation was canceled
)

// Enum value maps for InviteStatus.
var (
	InviteStatus_name = map[int32]string{
		0: "INVITE_STATUS_UNKNOWN",
		1: "ACCEPTED",
		2: "DECLINED",
		3: "EXPIRED",
		4: "CANCELED",
	}
	InviteStatus_value = map[string]int32{
		"INVITE_STATUS_UNKNOWN": 0,
		"ACCEPTED":              1,
		"DECLINED":              2,
		"EXPIRED":               3,
		"CANCELED":              4,
	}
)

func (x InviteStatus) Enum() *InviteStatus {
	p := new(InviteStatus)
	*p = x
	return p
}

func (x InviteStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (InviteStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_aserto_api_invite_proto_enumTypes[0].Descriptor()
}

func (InviteStatus) Type() protoreflect.EnumType {
	return &file_aserto_api_invite_proto_enumTypes[0]
}

func (x InviteStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use InviteStatus.Descriptor instead.
func (InviteStatus) EnumDescriptor() ([]byte, []int) {
	return file_aserto_api_invite_proto_rawDescGZIP(), []int{0}
}

type Invite struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`                                       // Unique ID of the invite
	Email       string                 `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`                                 // Email of the invitee
	Status      InviteStatus           `protobuf:"varint,3,opt,name=status,proto3,enum=aserto.api.InviteStatus" json:"status,omitempty"` // Whether the invite was accepted/declined
	CreatedAt   *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`        // When was the invite created?
	RespondedAt *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=responded_at,json=respondedAt,proto3" json:"responded_at,omitempty"`  // When was the invite accepted/declined?
	InvitedBy   string                 `protobuf:"bytes,6,opt,name=invited_by,json=invitedBy,proto3" json:"invited_by,omitempty"`        // Account ID for the inviter
	Role        string                 `protobuf:"bytes,7,opt,name=role,proto3" json:"role,omitempty"`                                   // Role assumed by the invitee on acceptance
}

func (x *Invite) Reset() {
	*x = Invite{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aserto_api_invite_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Invite) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Invite) ProtoMessage() {}

func (x *Invite) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_api_invite_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Invite.ProtoReflect.Descriptor instead.
func (*Invite) Descriptor() ([]byte, []int) {
	return file_aserto_api_invite_proto_rawDescGZIP(), []int{0}
}

func (x *Invite) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Invite) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *Invite) GetStatus() InviteStatus {
	if x != nil {
		return x.Status
	}
	return InviteStatus_INVITE_STATUS_UNKNOWN
}

func (x *Invite) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Invite) GetRespondedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.RespondedAt
	}
	return nil
}

func (x *Invite) GetInvitedBy() string {
	if x != nil {
		return x.InvitedBy
	}
	return ""
}

func (x *Invite) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

var File_aserto_api_invite_proto protoreflect.FileDescriptor

var file_aserto_api_invite_proto_rawDesc = []byte{
	0x0a, 0x17, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x69, 0x6e, 0x76,
	0x69, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x61, 0x73, 0x65, 0x72, 0x74,
	0x6f, 0x2e, 0x61, 0x70, 0x69, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2f, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x69, 0x64, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x93, 0x02, 0x0a, 0x06, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x12, 0x14, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x04, 0x90, 0xb5, 0x18, 0x08, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x30, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x3d, 0x0a, 0x0c, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x64, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x64, 0x65,
	0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x64, 0x5f, 0x62,
	0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x64,
	0x42, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x2a, 0x60, 0x0a, 0x0c, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x19, 0x0a, 0x15, 0x49, 0x4e, 0x56, 0x49, 0x54, 0x45,
	0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10,
	0x00, 0x12, 0x0c, 0x0a, 0x08, 0x41, 0x43, 0x43, 0x45, 0x50, 0x54, 0x45, 0x44, 0x10, 0x01, 0x12,
	0x0c, 0x0a, 0x08, 0x44, 0x45, 0x43, 0x4c, 0x49, 0x4e, 0x45, 0x44, 0x10, 0x02, 0x12, 0x0b, 0x0a,
	0x07, 0x45, 0x58, 0x50, 0x49, 0x52, 0x45, 0x44, 0x10, 0x03, 0x12, 0x0c, 0x0a, 0x08, 0x43, 0x41,
	0x4e, 0x43, 0x45, 0x4c, 0x45, 0x44, 0x10, 0x04, 0x42, 0x39, 0x5a, 0x2a, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2d, 0x64, 0x65,
	0x76, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2f, 0x61,
	0x70, 0x69, 0x3b, 0x61, 0x70, 0x69, 0xaa, 0x02, 0x0a, 0x41, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e,
	0x41, 0x50, 0x49, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_aserto_api_invite_proto_rawDescOnce sync.Once
	file_aserto_api_invite_proto_rawDescData = file_aserto_api_invite_proto_rawDesc
)

func file_aserto_api_invite_proto_rawDescGZIP() []byte {
	file_aserto_api_invite_proto_rawDescOnce.Do(func() {
		file_aserto_api_invite_proto_rawDescData = protoimpl.X.CompressGZIP(file_aserto_api_invite_proto_rawDescData)
	})
	return file_aserto_api_invite_proto_rawDescData
}

var file_aserto_api_invite_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_aserto_api_invite_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_aserto_api_invite_proto_goTypes = []interface{}{
	(InviteStatus)(0),             // 0: aserto.api.InviteStatus
	(*Invite)(nil),                // 1: aserto.api.Invite
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_aserto_api_invite_proto_depIdxs = []int32{
	0, // 0: aserto.api.Invite.status:type_name -> aserto.api.InviteStatus
	2, // 1: aserto.api.Invite.created_at:type_name -> google.protobuf.Timestamp
	2, // 2: aserto.api.Invite.responded_at:type_name -> google.protobuf.Timestamp
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_aserto_api_invite_proto_init() }
func file_aserto_api_invite_proto_init() {
	if File_aserto_api_invite_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_aserto_api_invite_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Invite); i {
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
			RawDescriptor: file_aserto_api_invite_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_aserto_api_invite_proto_goTypes,
		DependencyIndexes: file_aserto_api_invite_proto_depIdxs,
		EnumInfos:         file_aserto_api_invite_proto_enumTypes,
		MessageInfos:      file_aserto_api_invite_proto_msgTypes,
	}.Build()
	File_aserto_api_invite_proto = out.File
	file_aserto_api_invite_proto_rawDesc = nil
	file_aserto_api_invite_proto_goTypes = nil
	file_aserto_api_invite_proto_depIdxs = nil
}
