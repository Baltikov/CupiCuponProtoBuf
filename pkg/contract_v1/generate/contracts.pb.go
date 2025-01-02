// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.27.1
// source: contracts.proto

package contracts_v1

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

type Contract struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title      string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Decription string `protobuf:"bytes,3,opt,name=decription,proto3" json:"decription,omitempty"`
}

func (x *Contract) Reset() {
	*x = Contract{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contracts_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Contract) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Contract) ProtoMessage() {}

func (x *Contract) ProtoReflect() protoreflect.Message {
	mi := &file_contracts_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Contract.ProtoReflect.Descriptor instead.
func (*Contract) Descriptor() ([]byte, []int) {
	return file_contracts_proto_rawDescGZIP(), []int{0}
}

func (x *Contract) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Contract) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Contract) GetDecription() string {
	if x != nil {
		return x.Decription
	}
	return ""
}

type CreateContractRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Contract *Contract `protobuf:"bytes,1,opt,name=contract,proto3" json:"contract,omitempty"`
}

func (x *CreateContractRequest) Reset() {
	*x = CreateContractRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contracts_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateContractRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateContractRequest) ProtoMessage() {}

func (x *CreateContractRequest) ProtoReflect() protoreflect.Message {
	mi := &file_contracts_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateContractRequest.ProtoReflect.Descriptor instead.
func (*CreateContractRequest) Descriptor() ([]byte, []int) {
	return file_contracts_proto_rawDescGZIP(), []int{1}
}

func (x *CreateContractRequest) GetContract() *Contract {
	if x != nil {
		return x.Contract
	}
	return nil
}

type CreateContractResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CreateContractResponse) Reset() {
	*x = CreateContractResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contracts_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateContractResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateContractResponse) ProtoMessage() {}

func (x *CreateContractResponse) ProtoReflect() protoreflect.Message {
	mi := &file_contracts_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateContractResponse.ProtoReflect.Descriptor instead.
func (*CreateContractResponse) Descriptor() ([]byte, []int) {
	return file_contracts_proto_rawDescGZIP(), []int{2}
}

func (x *CreateContractResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetContractRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetContractRequest) Reset() {
	*x = GetContractRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contracts_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetContractRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetContractRequest) ProtoMessage() {}

func (x *GetContractRequest) ProtoReflect() protoreflect.Message {
	mi := &file_contracts_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetContractRequest.ProtoReflect.Descriptor instead.
func (*GetContractRequest) Descriptor() ([]byte, []int) {
	return file_contracts_proto_rawDescGZIP(), []int{3}
}

func (x *GetContractRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetContractResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Contract *Contract `protobuf:"bytes,1,opt,name=contract,proto3" json:"contract,omitempty"`
}

func (x *GetContractResponse) Reset() {
	*x = GetContractResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contracts_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetContractResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetContractResponse) ProtoMessage() {}

func (x *GetContractResponse) ProtoReflect() protoreflect.Message {
	mi := &file_contracts_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetContractResponse.ProtoReflect.Descriptor instead.
func (*GetContractResponse) Descriptor() ([]byte, []int) {
	return file_contracts_proto_rawDescGZIP(), []int{4}
}

func (x *GetContractResponse) GetContract() *Contract {
	if x != nil {
		return x.Contract
	}
	return nil
}

var File_contracts_proto protoreflect.FileDescriptor

var file_contracts_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x50, 0x0a, 0x08, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x65, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x65, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x22, 0x3e, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e,
	0x74, 0x72, 0x61, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x08,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09,
	0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x52, 0x08, 0x63, 0x6f, 0x6e, 0x74, 0x72,
	0x61, 0x63, 0x74, 0x22, 0x28, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e,
	0x74, 0x72, 0x61, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x24, 0x0a,
	0x12, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x02, 0x69, 0x64, 0x22, 0x3c, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61,
	0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x08, 0x63, 0x6f,
	0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x43,
	0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x52, 0x08, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63,
	0x74, 0x32, 0x88, 0x01, 0x0a, 0x09, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x73, 0x12,
	0x41, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63,
	0x74, 0x12, 0x16, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61,
	0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x38, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63,
	0x74, 0x12, 0x13, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x74,
	0x72, 0x61, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x52, 0x5a, 0x50,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x42, 0x61, 0x6c, 0x74, 0x69,
	0x6b, 0x6f, 0x76, 0x2f, 0x43, 0x75, 0x70, 0x69, 0x43, 0x75, 0x70, 0x6f, 0x6e, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x42, 0x75, 0x66, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x63,
	0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x5f, 0x76, 0x31, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72,
	0x61, 0x74, 0x65, 0x3b, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x73, 0x5f, 0x76, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_contracts_proto_rawDescOnce sync.Once
	file_contracts_proto_rawDescData = file_contracts_proto_rawDesc
)

func file_contracts_proto_rawDescGZIP() []byte {
	file_contracts_proto_rawDescOnce.Do(func() {
		file_contracts_proto_rawDescData = protoimpl.X.CompressGZIP(file_contracts_proto_rawDescData)
	})
	return file_contracts_proto_rawDescData
}

var file_contracts_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_contracts_proto_goTypes = []interface{}{
	(*Contract)(nil),               // 0: Contract
	(*CreateContractRequest)(nil),  // 1: CreateContractRequest
	(*CreateContractResponse)(nil), // 2: CreateContractResponse
	(*GetContractRequest)(nil),     // 3: GetContractRequest
	(*GetContractResponse)(nil),    // 4: GetContractResponse
}
var file_contracts_proto_depIdxs = []int32{
	0, // 0: CreateContractRequest.contract:type_name -> Contract
	0, // 1: GetContractResponse.contract:type_name -> Contract
	1, // 2: Contracts.CreateContract:input_type -> CreateContractRequest
	3, // 3: Contracts.GetContract:input_type -> GetContractRequest
	2, // 4: Contracts.CreateContract:output_type -> CreateContractResponse
	4, // 5: Contracts.GetContract:output_type -> GetContractResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_contracts_proto_init() }
func file_contracts_proto_init() {
	if File_contracts_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_contracts_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Contract); i {
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
		file_contracts_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateContractRequest); i {
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
		file_contracts_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateContractResponse); i {
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
		file_contracts_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetContractRequest); i {
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
		file_contracts_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetContractResponse); i {
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
			RawDescriptor: file_contracts_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_contracts_proto_goTypes,
		DependencyIndexes: file_contracts_proto_depIdxs,
		MessageInfos:      file_contracts_proto_msgTypes,
	}.Build()
	File_contracts_proto = out.File
	file_contracts_proto_rawDesc = nil
	file_contracts_proto_goTypes = nil
	file_contracts_proto_depIdxs = nil
}
