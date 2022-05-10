// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.1
// source: user-service/userService.proto

package api

import (
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

type LoginUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Money    int32  `protobuf:"varint,2,opt,name=money,proto3" json:"money,omitempty"`
}

func (x *LoginUserRequest) Reset() {
	*x = LoginUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_service_userService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginUserRequest) ProtoMessage() {}

func (x *LoginUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_service_userService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginUserRequest.ProtoReflect.Descriptor instead.
func (*LoginUserRequest) Descriptor() ([]byte, []int) {
	return file_user_service_userService_proto_rawDescGZIP(), []int{0}
}

func (x *LoginUserRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *LoginUserRequest) GetMoney() int32 {
	if x != nil {
		return x.Money
	}
	return 0
}

type LoginUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *LoginUserResponse) Reset() {
	*x = LoginUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_service_userService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginUserResponse) ProtoMessage() {}

func (x *LoginUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_service_userService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginUserResponse.ProtoReflect.Descriptor instead.
func (*LoginUserResponse) Descriptor() ([]byte, []int) {
	return file_user_service_userService_proto_rawDescGZIP(), []int{1}
}

type DecreaseMoneyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username    string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	MoneyAmount int32  `protobuf:"varint,2,opt,name=moneyAmount,proto3" json:"moneyAmount,omitempty"`
}

func (x *DecreaseMoneyRequest) Reset() {
	*x = DecreaseMoneyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_service_userService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DecreaseMoneyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DecreaseMoneyRequest) ProtoMessage() {}

func (x *DecreaseMoneyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_service_userService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DecreaseMoneyRequest.ProtoReflect.Descriptor instead.
func (*DecreaseMoneyRequest) Descriptor() ([]byte, []int) {
	return file_user_service_userService_proto_rawDescGZIP(), []int{2}
}

func (x *DecreaseMoneyRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *DecreaseMoneyRequest) GetMoneyAmount() int32 {
	if x != nil {
		return x.MoneyAmount
	}
	return 0
}

type DecreaseMoneyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DecreaseMoneyResponse) Reset() {
	*x = DecreaseMoneyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_service_userService_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DecreaseMoneyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DecreaseMoneyResponse) ProtoMessage() {}

func (x *DecreaseMoneyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_service_userService_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DecreaseMoneyResponse.ProtoReflect.Descriptor instead.
func (*DecreaseMoneyResponse) Descriptor() ([]byte, []int) {
	return file_user_service_userService_proto_rawDescGZIP(), []int{3}
}

var File_user_service_userService_proto protoreflect.FileDescriptor

var file_user_service_userService_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x75, 0x73, 0x65, 0x72, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x75,
	0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0c, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x44,
	0x0a, 0x10, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x6d, 0x6f, 0x6e, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6d,
	0x6f, 0x6e, 0x65, 0x79, 0x22, 0x13, 0x0a, 0x11, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x54, 0x0a, 0x14, 0x44, 0x65, 0x63,
	0x72, 0x65, 0x61, 0x73, 0x65, 0x4d, 0x6f, 0x6e, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a,
	0x0b, 0x6d, 0x6f, 0x6e, 0x65, 0x79, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0b, 0x6d, 0x6f, 0x6e, 0x65, 0x79, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22,
	0x17, 0x0a, 0x15, 0x44, 0x65, 0x63, 0x72, 0x65, 0x61, 0x73, 0x65, 0x4d, 0x6f, 0x6e, 0x65, 0x79,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xb2, 0x01, 0x0a, 0x04, 0x55, 0x73, 0x65,
	0x72, 0x12, 0x4e, 0x0a, 0x09, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1e,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4c, 0x6f,
	0x67, 0x69, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4c, 0x6f,
	0x67, 0x69, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x5a, 0x0a, 0x0d, 0x44, 0x65, 0x63, 0x72, 0x65, 0x61, 0x73, 0x65, 0x4d, 0x6f, 0x6e,
	0x65, 0x79, 0x12, 0x22, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x44, 0x65, 0x63, 0x72, 0x65, 0x61, 0x73, 0x65, 0x4d, 0x6f, 0x6e, 0x65, 0x79, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x44, 0x65, 0x63, 0x72, 0x65, 0x61, 0x73, 0x65, 0x4d, 0x6f,
	0x6e, 0x65, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x24, 0x5a,
	0x22, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x61, 0x74, 0x66,
	0x2d, 0x70, 0x70, 0x2f, 0x32, 0x30, 0x32, 0x32, 0x5f, 0x4d, 0x41, 0x54, 0x44, 0x41, 0x51, 0x2f,
	0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_user_service_userService_proto_rawDescOnce sync.Once
	file_user_service_userService_proto_rawDescData = file_user_service_userService_proto_rawDesc
)

func file_user_service_userService_proto_rawDescGZIP() []byte {
	file_user_service_userService_proto_rawDescOnce.Do(func() {
		file_user_service_userService_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_service_userService_proto_rawDescData)
	})
	return file_user_service_userService_proto_rawDescData
}

var file_user_service_userService_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_user_service_userService_proto_goTypes = []interface{}{
	(*LoginUserRequest)(nil),      // 0: user_service.LoginUserRequest
	(*LoginUserResponse)(nil),     // 1: user_service.LoginUserResponse
	(*DecreaseMoneyRequest)(nil),  // 2: user_service.DecreaseMoneyRequest
	(*DecreaseMoneyResponse)(nil), // 3: user_service.DecreaseMoneyResponse
}
var file_user_service_userService_proto_depIdxs = []int32{
	0, // 0: user_service.User.LoginUser:input_type -> user_service.LoginUserRequest
	2, // 1: user_service.User.DecreaseMoney:input_type -> user_service.DecreaseMoneyRequest
	1, // 2: user_service.User.LoginUser:output_type -> user_service.LoginUserResponse
	3, // 3: user_service.User.DecreaseMoney:output_type -> user_service.DecreaseMoneyResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_user_service_userService_proto_init() }
func file_user_service_userService_proto_init() {
	if File_user_service_userService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_user_service_userService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginUserRequest); i {
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
		file_user_service_userService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginUserResponse); i {
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
		file_user_service_userService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DecreaseMoneyRequest); i {
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
		file_user_service_userService_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DecreaseMoneyResponse); i {
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
			RawDescriptor: file_user_service_userService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_user_service_userService_proto_goTypes,
		DependencyIndexes: file_user_service_userService_proto_depIdxs,
		MessageInfos:      file_user_service_userService_proto_msgTypes,
	}.Build()
	File_user_service_userService_proto = out.File
	file_user_service_userService_proto_rawDesc = nil
	file_user_service_userService_proto_goTypes = nil
	file_user_service_userService_proto_depIdxs = nil
}
