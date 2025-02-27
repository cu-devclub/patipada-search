// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.1
// source: mlGateway.proto

package proto

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

type Text2VecRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *Text2VecRequest) Reset() {
	*x = Text2VecRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mlGateway_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Text2VecRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Text2VecRequest) ProtoMessage() {}

func (x *Text2VecRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mlGateway_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Text2VecRequest.ProtoReflect.Descriptor instead.
func (*Text2VecRequest) Descriptor() ([]byte, []int) {
	return file_mlGateway_proto_rawDescGZIP(), []int{0}
}

func (x *Text2VecRequest) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

type Result struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string    `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Embedding []float32 `protobuf:"fixed32,2,rep,packed,name=embedding,proto3" json:"embedding,omitempty"`
	Score     float32   `protobuf:"fixed32,3,opt,name=score,proto3" json:"score,omitempty"`
}

func (x *Result) Reset() {
	*x = Result{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mlGateway_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Result) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Result) ProtoMessage() {}

func (x *Result) ProtoReflect() protoreflect.Message {
	mi := &file_mlGateway_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Result.ProtoReflect.Descriptor instead.
func (*Result) Descriptor() ([]byte, []int) {
	return file_mlGateway_proto_rawDescGZIP(), []int{1}
}

func (x *Result) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Result) GetEmbedding() []float32 {
	if x != nil {
		return x.Embedding
	}
	return nil
}

func (x *Result) GetScore() float32 {
	if x != nil {
		return x.Score
	}
	return 0
}

type Text2VecResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Results []*Result `protobuf:"bytes,2,rep,name=results,proto3" json:"results,omitempty"`
}

func (x *Text2VecResponse) Reset() {
	*x = Text2VecResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mlGateway_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Text2VecResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Text2VecResponse) ProtoMessage() {}

func (x *Text2VecResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mlGateway_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Text2VecResponse.ProtoReflect.Descriptor instead.
func (*Text2VecResponse) Descriptor() ([]byte, []int) {
	return file_mlGateway_proto_rawDescGZIP(), []int{2}
}

func (x *Text2VecResponse) GetResults() []*Result {
	if x != nil {
		return x.Results
	}
	return nil
}

var File_mlGateway_proto protoreflect.FileDescriptor

var file_mlGateway_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x6d, 0x6c, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x25, 0x0a, 0x0f, 0x54, 0x65, 0x78, 0x74,
	0x32, 0x56, 0x65, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74,
	0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x22,
	0x50, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a,
	0x09, 0x65, 0x6d, 0x62, 0x65, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x03, 0x28, 0x02,
	0x52, 0x09, 0x65, 0x6d, 0x62, 0x65, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x14, 0x0a, 0x05, 0x73,
	0x63, 0x6f, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x72,
	0x65, 0x22, 0x3b, 0x0a, 0x10, 0x54, 0x65, 0x78, 0x74, 0x32, 0x56, 0x65, 0x63, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x32, 0x4f,
	0x0a, 0x10, 0x4d, 0x6c, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x3b, 0x0a, 0x08, 0x54, 0x65, 0x78, 0x74, 0x32, 0x56, 0x65, 0x63, 0x12, 0x16,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x65, 0x78, 0x74, 0x32, 0x56, 0x65, 0x63, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54,
	0x65, 0x78, 0x74, 0x32, 0x56, 0x65, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x08, 0x5a, 0x06, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_mlGateway_proto_rawDescOnce sync.Once
	file_mlGateway_proto_rawDescData = file_mlGateway_proto_rawDesc
)

func file_mlGateway_proto_rawDescGZIP() []byte {
	file_mlGateway_proto_rawDescOnce.Do(func() {
		file_mlGateway_proto_rawDescData = protoimpl.X.CompressGZIP(file_mlGateway_proto_rawDescData)
	})
	return file_mlGateway_proto_rawDescData
}

var file_mlGateway_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_mlGateway_proto_goTypes = []interface{}{
	(*Text2VecRequest)(nil),  // 0: proto.Text2VecRequest
	(*Result)(nil),           // 1: proto.Result
	(*Text2VecResponse)(nil), // 2: proto.Text2VecResponse
}
var file_mlGateway_proto_depIdxs = []int32{
	1, // 0: proto.Text2VecResponse.results:type_name -> proto.Result
	0, // 1: proto.MlGatewayService.Text2Vec:input_type -> proto.Text2VecRequest
	2, // 2: proto.MlGatewayService.Text2Vec:output_type -> proto.Text2VecResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_mlGateway_proto_init() }
func file_mlGateway_proto_init() {
	if File_mlGateway_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mlGateway_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Text2VecRequest); i {
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
		file_mlGateway_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Result); i {
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
		file_mlGateway_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Text2VecResponse); i {
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
			RawDescriptor: file_mlGateway_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_mlGateway_proto_goTypes,
		DependencyIndexes: file_mlGateway_proto_depIdxs,
		MessageInfos:      file_mlGateway_proto_msgTypes,
	}.Build()
	File_mlGateway_proto = out.File
	file_mlGateway_proto_rawDesc = nil
	file_mlGateway_proto_goTypes = nil
	file_mlGateway_proto_depIdxs = nil
}
