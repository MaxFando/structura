// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: structura/v1/structura.proto

package structurav1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetDomainInfoRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Domain        string                 `protobuf:"bytes,1,opt,name=domain,proto3" json:"domain,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetDomainInfoRequest) Reset() {
	*x = GetDomainInfoRequest{}
	mi := &file_structura_v1_structura_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetDomainInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDomainInfoRequest) ProtoMessage() {}

func (x *GetDomainInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_structura_v1_structura_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDomainInfoRequest.ProtoReflect.Descriptor instead.
func (*GetDomainInfoRequest) Descriptor() ([]byte, []int) {
	return file_structura_v1_structura_proto_rawDescGZIP(), []int{0}
}

func (x *GetDomainInfoRequest) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

type GetDomainInfoResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	PagesCount    int32                  `protobuf:"varint,1,opt,name=pages_count,json=pagesCount,proto3" json:"pages_count,omitempty"`
	Domain        string                 `protobuf:"bytes,2,opt,name=domain,proto3" json:"domain,omitempty"`
	Pages         []*Page                `protobuf:"bytes,3,rep,name=pages,proto3" json:"pages,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetDomainInfoResponse) Reset() {
	*x = GetDomainInfoResponse{}
	mi := &file_structura_v1_structura_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetDomainInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDomainInfoResponse) ProtoMessage() {}

func (x *GetDomainInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_structura_v1_structura_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDomainInfoResponse.ProtoReflect.Descriptor instead.
func (*GetDomainInfoResponse) Descriptor() ([]byte, []int) {
	return file_structura_v1_structura_proto_rawDescGZIP(), []int{1}
}

func (x *GetDomainInfoResponse) GetPagesCount() int32 {
	if x != nil {
		return x.PagesCount
	}
	return 0
}

func (x *GetDomainInfoResponse) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

func (x *GetDomainInfoResponse) GetPages() []*Page {
	if x != nil {
		return x.Pages
	}
	return nil
}

type GetPageBlocksRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Url           string                 `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetPageBlocksRequest) Reset() {
	*x = GetPageBlocksRequest{}
	mi := &file_structura_v1_structura_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetPageBlocksRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPageBlocksRequest) ProtoMessage() {}

func (x *GetPageBlocksRequest) ProtoReflect() protoreflect.Message {
	mi := &file_structura_v1_structura_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPageBlocksRequest.ProtoReflect.Descriptor instead.
func (*GetPageBlocksRequest) Descriptor() ([]byte, []int) {
	return file_structura_v1_structura_proto_rawDescGZIP(), []int{2}
}

func (x *GetPageBlocksRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type GetPageBlocksResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Blocks        []*Block               `protobuf:"bytes,1,rep,name=blocks,proto3" json:"blocks,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetPageBlocksResponse) Reset() {
	*x = GetPageBlocksResponse{}
	mi := &file_structura_v1_structura_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetPageBlocksResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPageBlocksResponse) ProtoMessage() {}

func (x *GetPageBlocksResponse) ProtoReflect() protoreflect.Message {
	mi := &file_structura_v1_structura_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPageBlocksResponse.ProtoReflect.Descriptor instead.
func (*GetPageBlocksResponse) Descriptor() ([]byte, []int) {
	return file_structura_v1_structura_proto_rawDescGZIP(), []int{3}
}

func (x *GetPageBlocksResponse) GetBlocks() []*Block {
	if x != nil {
		return x.Blocks
	}
	return nil
}

type Block struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Index         int32                  `protobuf:"varint,2,opt,name=index,proto3" json:"index,omitempty"`
	Type          string                 `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Block) Reset() {
	*x = Block{}
	mi := &file_structura_v1_structura_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Block) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Block) ProtoMessage() {}

func (x *Block) ProtoReflect() protoreflect.Message {
	mi := &file_structura_v1_structura_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Block.ProtoReflect.Descriptor instead.
func (*Block) Descriptor() ([]byte, []int) {
	return file_structura_v1_structura_proto_rawDescGZIP(), []int{4}
}

func (x *Block) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Block) GetIndex() int32 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *Block) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

type Page struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Endpoint      string                 `protobuf:"bytes,2,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
	Blocks        []*Block               `protobuf:"bytes,3,rep,name=blocks,proto3" json:"blocks,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Page) Reset() {
	*x = Page{}
	mi := &file_structura_v1_structura_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Page) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Page) ProtoMessage() {}

func (x *Page) ProtoReflect() protoreflect.Message {
	mi := &file_structura_v1_structura_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Page.ProtoReflect.Descriptor instead.
func (*Page) Descriptor() ([]byte, []int) {
	return file_structura_v1_structura_proto_rawDescGZIP(), []int{5}
}

func (x *Page) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Page) GetEndpoint() string {
	if x != nil {
		return x.Endpoint
	}
	return ""
}

func (x *Page) GetBlocks() []*Block {
	if x != nil {
		return x.Blocks
	}
	return nil
}

var File_structura_v1_structura_proto protoreflect.FileDescriptor

const file_structura_v1_structura_proto_rawDesc = "" +
	"\n" +
	"\x1cstructura/v1/structura.proto\x12\fstructura.v1\x1a\x1cgoogle/api/annotations.proto\".\n" +
	"\x14GetDomainInfoRequest\x12\x16\n" +
	"\x06domain\x18\x01 \x01(\tR\x06domain\"z\n" +
	"\x15GetDomainInfoResponse\x12\x1f\n" +
	"\vpages_count\x18\x01 \x01(\x05R\n" +
	"pagesCount\x12\x16\n" +
	"\x06domain\x18\x02 \x01(\tR\x06domain\x12(\n" +
	"\x05pages\x18\x03 \x03(\v2\x12.structura.v1.PageR\x05pages\"(\n" +
	"\x14GetPageBlocksRequest\x12\x10\n" +
	"\x03url\x18\x01 \x01(\tR\x03url\"D\n" +
	"\x15GetPageBlocksResponse\x12+\n" +
	"\x06blocks\x18\x01 \x03(\v2\x13.structura.v1.BlockR\x06blocks\"A\n" +
	"\x05Block\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x14\n" +
	"\x05index\x18\x02 \x01(\x05R\x05index\x12\x12\n" +
	"\x04type\x18\x03 \x01(\tR\x04type\"_\n" +
	"\x04Page\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x1a\n" +
	"\bendpoint\x18\x02 \x01(\tR\bendpoint\x12+\n" +
	"\x06blocks\x18\x03 \x03(\v2\x13.structura.v1.BlockR\x06blocks2\x88\x02\n" +
	"\x10StructuraService\x12s\n" +
	"\rGetDomainInfo\x12\".structura.v1.GetDomainInfoRequest\x1a#.structura.v1.GetDomainInfoResponse\"\x19\x82\xd3\xe4\x93\x02\x13\x12\x11/api/v1/structura\x12\x7f\n" +
	"\rGetPageBlocks\x12\".structura.v1.GetPageBlocksRequest\x1a#.structura.v1.GetPageBlocksResponse\"%\x82\xd3\xe4\x93\x02\x1f\x12\x1d/api/v1/structura/page/blocksB\xbb\x01\n" +
	"\x10com.structura.v1B\x0eStructuraProtoP\x01ZFgithub.com/MaxFando/structura/api/grpc/gen/go/structura/v1;structurav1\xa2\x02\x03SXX\xaa\x02\fStructura.V1\xca\x02\fStructura\\V1\xe2\x02\x18Structura\\V1\\GPBMetadata\xea\x02\rStructura::V1b\x06proto3"

var (
	file_structura_v1_structura_proto_rawDescOnce sync.Once
	file_structura_v1_structura_proto_rawDescData []byte
)

func file_structura_v1_structura_proto_rawDescGZIP() []byte {
	file_structura_v1_structura_proto_rawDescOnce.Do(func() {
		file_structura_v1_structura_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_structura_v1_structura_proto_rawDesc), len(file_structura_v1_structura_proto_rawDesc)))
	})
	return file_structura_v1_structura_proto_rawDescData
}

var file_structura_v1_structura_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_structura_v1_structura_proto_goTypes = []any{
	(*GetDomainInfoRequest)(nil),  // 0: structura.v1.GetDomainInfoRequest
	(*GetDomainInfoResponse)(nil), // 1: structura.v1.GetDomainInfoResponse
	(*GetPageBlocksRequest)(nil),  // 2: structura.v1.GetPageBlocksRequest
	(*GetPageBlocksResponse)(nil), // 3: structura.v1.GetPageBlocksResponse
	(*Block)(nil),                 // 4: structura.v1.Block
	(*Page)(nil),                  // 5: structura.v1.Page
}
var file_structura_v1_structura_proto_depIdxs = []int32{
	5, // 0: structura.v1.GetDomainInfoResponse.pages:type_name -> structura.v1.Page
	4, // 1: structura.v1.GetPageBlocksResponse.blocks:type_name -> structura.v1.Block
	4, // 2: structura.v1.Page.blocks:type_name -> structura.v1.Block
	0, // 3: structura.v1.StructuraService.GetDomainInfo:input_type -> structura.v1.GetDomainInfoRequest
	2, // 4: structura.v1.StructuraService.GetPageBlocks:input_type -> structura.v1.GetPageBlocksRequest
	1, // 5: structura.v1.StructuraService.GetDomainInfo:output_type -> structura.v1.GetDomainInfoResponse
	3, // 6: structura.v1.StructuraService.GetPageBlocks:output_type -> structura.v1.GetPageBlocksResponse
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_structura_v1_structura_proto_init() }
func file_structura_v1_structura_proto_init() {
	if File_structura_v1_structura_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_structura_v1_structura_proto_rawDesc), len(file_structura_v1_structura_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_structura_v1_structura_proto_goTypes,
		DependencyIndexes: file_structura_v1_structura_proto_depIdxs,
		MessageInfos:      file_structura_v1_structura_proto_msgTypes,
	}.Build()
	File_structura_v1_structura_proto = out.File
	file_structura_v1_structura_proto_goTypes = nil
	file_structura_v1_structura_proto_depIdxs = nil
}
