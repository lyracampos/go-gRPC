// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: proto/server.proto

package pb

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

type CreateCategoryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *CreateCategoryRequest) Reset() {
	*x = CreateCategoryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_server_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCategoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCategoryRequest) ProtoMessage() {}

func (x *CreateCategoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_server_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCategoryRequest.ProtoReflect.Descriptor instead.
func (*CreateCategoryRequest) Descriptor() ([]byte, []int) {
	return file_proto_server_proto_rawDescGZIP(), []int{0}
}

func (x *CreateCategoryRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateCategoryRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type CreateCategoryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Category *Category `protobuf:"bytes,1,opt,name=category,proto3" json:"category,omitempty"`
}

func (x *CreateCategoryResponse) Reset() {
	*x = CreateCategoryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_server_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCategoryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCategoryResponse) ProtoMessage() {}

func (x *CreateCategoryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_server_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCategoryResponse.ProtoReflect.Descriptor instead.
func (*CreateCategoryResponse) Descriptor() ([]byte, []int) {
	return file_proto_server_proto_rawDescGZIP(), []int{1}
}

func (x *CreateCategoryResponse) GetCategory() *Category {
	if x != nil {
		return x.Category
	}
	return nil
}

type ListCategoryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListCategoryRequest) Reset() {
	*x = ListCategoryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_server_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListCategoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCategoryRequest) ProtoMessage() {}

func (x *ListCategoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_server_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCategoryRequest.ProtoReflect.Descriptor instead.
func (*ListCategoryRequest) Descriptor() ([]byte, []int) {
	return file_proto_server_proto_rawDescGZIP(), []int{2}
}

type ListCategoryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Categories []*Category `protobuf:"bytes,1,rep,name=categories,proto3" json:"categories,omitempty"`
	Total      int32       `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *ListCategoryResponse) Reset() {
	*x = ListCategoryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_server_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListCategoryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCategoryResponse) ProtoMessage() {}

func (x *ListCategoryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_server_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCategoryResponse.ProtoReflect.Descriptor instead.
func (*ListCategoryResponse) Descriptor() ([]byte, []int) {
	return file_proto_server_proto_rawDescGZIP(), []int{3}
}

func (x *ListCategoryResponse) GetCategories() []*Category {
	if x != nil {
		return x.Categories
	}
	return nil
}

func (x *ListCategoryResponse) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

type GetCategoryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetCategoryRequest) Reset() {
	*x = GetCategoryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_server_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCategoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCategoryRequest) ProtoMessage() {}

func (x *GetCategoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_server_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCategoryRequest.ProtoReflect.Descriptor instead.
func (*GetCategoryRequest) Descriptor() ([]byte, []int) {
	return file_proto_server_proto_rawDescGZIP(), []int{4}
}

func (x *GetCategoryRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetCategoryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Category *Category `protobuf:"bytes,1,opt,name=category,proto3" json:"category,omitempty"`
}

func (x *GetCategoryResponse) Reset() {
	*x = GetCategoryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_server_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCategoryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCategoryResponse) ProtoMessage() {}

func (x *GetCategoryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_server_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCategoryResponse.ProtoReflect.Descriptor instead.
func (*GetCategoryResponse) Descriptor() ([]byte, []int) {
	return file_proto_server_proto_rawDescGZIP(), []int{5}
}

func (x *GetCategoryResponse) GetCategory() *Category {
	if x != nil {
		return x.Category
	}
	return nil
}

type Category struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *Category) Reset() {
	*x = Category{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_server_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Category) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Category) ProtoMessage() {}

func (x *Category) ProtoReflect() protoreflect.Message {
	mi := &file_proto_server_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Category.ProtoReflect.Descriptor instead.
func (*Category) Descriptor() ([]byte, []int) {
	return file_proto_server_proto_rawDescGZIP(), []int{6}
}

func (x *Category) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Category) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Category) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

var File_proto_server_proto protoreflect.FileDescriptor

var file_proto_server_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0x4d, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x42, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x28, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x22, 0x15, 0x0a, 0x13, 0x4c,
	0x69, 0x73, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x22, 0x5a, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x0a, 0x63, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c,
	0x2e, 0x70, 0x62, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x0a, 0x63, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0x24,
	0x0a, 0x12, 0x47, 0x65, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x22, 0x3f, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x08, 0x63,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e,
	0x70, 0x62, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x08, 0x63, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x22, 0x50, 0x0a, 0x08, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x32, 0xe3, 0x01, 0x0a, 0x0f, 0x43, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x49, 0x0a, 0x0e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x19, 0x2e,
	0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x17, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x18, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x40, 0x0a, 0x0b, 0x47,
	0x65, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x16, 0x2e, 0x70, 0x62, 0x2e,
	0x47, 0x65, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x17, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0d, 0x5a,
	0x0b, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_server_proto_rawDescOnce sync.Once
	file_proto_server_proto_rawDescData = file_proto_server_proto_rawDesc
)

func file_proto_server_proto_rawDescGZIP() []byte {
	file_proto_server_proto_rawDescOnce.Do(func() {
		file_proto_server_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_server_proto_rawDescData)
	})
	return file_proto_server_proto_rawDescData
}

var file_proto_server_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_proto_server_proto_goTypes = []interface{}{
	(*CreateCategoryRequest)(nil),  // 0: pb.CreateCategoryRequest
	(*CreateCategoryResponse)(nil), // 1: pb.CreateCategoryResponse
	(*ListCategoryRequest)(nil),    // 2: pb.ListCategoryRequest
	(*ListCategoryResponse)(nil),   // 3: pb.ListCategoryResponse
	(*GetCategoryRequest)(nil),     // 4: pb.GetCategoryRequest
	(*GetCategoryResponse)(nil),    // 5: pb.GetCategoryResponse
	(*Category)(nil),               // 6: pb.Category
}
var file_proto_server_proto_depIdxs = []int32{
	6, // 0: pb.CreateCategoryResponse.category:type_name -> pb.Category
	6, // 1: pb.ListCategoryResponse.categories:type_name -> pb.Category
	6, // 2: pb.GetCategoryResponse.category:type_name -> pb.Category
	0, // 3: pb.CategoryService.CreateCategory:input_type -> pb.CreateCategoryRequest
	2, // 4: pb.CategoryService.ListCategory:input_type -> pb.ListCategoryRequest
	4, // 5: pb.CategoryService.GetCategory:input_type -> pb.GetCategoryRequest
	1, // 6: pb.CategoryService.CreateCategory:output_type -> pb.CreateCategoryResponse
	3, // 7: pb.CategoryService.ListCategory:output_type -> pb.ListCategoryResponse
	5, // 8: pb.CategoryService.GetCategory:output_type -> pb.GetCategoryResponse
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_proto_server_proto_init() }
func file_proto_server_proto_init() {
	if File_proto_server_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_server_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCategoryRequest); i {
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
		file_proto_server_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCategoryResponse); i {
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
		file_proto_server_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListCategoryRequest); i {
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
		file_proto_server_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListCategoryResponse); i {
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
		file_proto_server_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCategoryRequest); i {
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
		file_proto_server_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCategoryResponse); i {
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
		file_proto_server_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Category); i {
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
			RawDescriptor: file_proto_server_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_server_proto_goTypes,
		DependencyIndexes: file_proto_server_proto_depIdxs,
		MessageInfos:      file_proto_server_proto_msgTypes,
	}.Build()
	File_proto_server_proto = out.File
	file_proto_server_proto_rawDesc = nil
	file_proto_server_proto_goTypes = nil
	file_proto_server_proto_depIdxs = nil
}
