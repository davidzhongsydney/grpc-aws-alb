// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.21.12
// source: model/book.proto

package model

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type Point struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Latitude  int32 `protobuf:"varint,1,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude int32 `protobuf:"varint,2,opt,name=longitude,proto3" json:"longitude,omitempty"`
}

func (x *Point) Reset() {
	*x = Point{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_book_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Point) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Point) ProtoMessage() {}

func (x *Point) ProtoReflect() protoreflect.Message {
	mi := &file_model_book_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Point.ProtoReflect.Descriptor instead.
func (*Point) Descriptor() ([]byte, []int) {
	return file_model_book_proto_rawDescGZIP(), []int{0}
}

func (x *Point) GetLatitude() int32 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

func (x *Point) GetLongitude() int32 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

type Feature struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the feature.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The point where the feature is detected.
	Location *Point `protobuf:"bytes,2,opt,name=location,proto3" json:"location,omitempty"`
}

func (x *Feature) Reset() {
	*x = Feature{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_book_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Feature) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Feature) ProtoMessage() {}

func (x *Feature) ProtoReflect() protoreflect.Message {
	mi := &file_model_book_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Feature.ProtoReflect.Descriptor instead.
func (*Feature) Descriptor() ([]byte, []int) {
	return file_model_book_proto_rawDescGZIP(), []int{1}
}

func (x *Feature) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Feature) GetLocation() *Point {
	if x != nil {
		return x.Location
	}
	return nil
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_book_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_model_book_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_model_book_proto_rawDescGZIP(), []int{2}
}

var File_model_book_proto protoreflect.FileDescriptor

var file_model_book_proto_rawDesc = []byte{
	0x0a, 0x10, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x41, 0x0a, 0x05, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x74,
	0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x6c, 0x61, 0x74,
	0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75,
	0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74,
	0x75, 0x64, 0x65, 0x22, 0x41, 0x0a, 0x07, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x22, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x06, 0x2e, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x08, 0x6c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x32,
	0x57, 0x0a, 0x0a, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x47, 0x75, 0x69, 0x64, 0x65, 0x12, 0x20, 0x0a,
	0x0a, 0x47, 0x65, 0x74, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x06, 0x2e, 0x50, 0x6f,
	0x69, 0x6e, 0x74, 0x1a, 0x08, 0x2e, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x22, 0x00, 0x12,
	0x27, 0x0a, 0x05, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x12, 0x06, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x1a, 0x06, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x0e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x08,
	0x12, 0x06, 0x2f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x42, 0x1c, 0x5a, 0x1a, 0x67, 0x6f, 0x6c, 0x61,
	0x6e, 0x67, 0x2e, 0x63, 0x61, 0x66, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_model_book_proto_rawDescOnce sync.Once
	file_model_book_proto_rawDescData = file_model_book_proto_rawDesc
)

func file_model_book_proto_rawDescGZIP() []byte {
	file_model_book_proto_rawDescOnce.Do(func() {
		file_model_book_proto_rawDescData = protoimpl.X.CompressGZIP(file_model_book_proto_rawDescData)
	})
	return file_model_book_proto_rawDescData
}

var file_model_book_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_model_book_proto_goTypes = []interface{}{
	(*Point)(nil),   // 0: Point
	(*Feature)(nil), // 1: Feature
	(*Empty)(nil),   // 2: Empty
}
var file_model_book_proto_depIdxs = []int32{
	0, // 0: Feature.location:type_name -> Point
	0, // 1: RouteGuide.GetFeature:input_type -> Point
	2, // 2: RouteGuide.Check:input_type -> Empty
	1, // 3: RouteGuide.GetFeature:output_type -> Feature
	2, // 4: RouteGuide.Check:output_type -> Empty
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_model_book_proto_init() }
func file_model_book_proto_init() {
	if File_model_book_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_model_book_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Point); i {
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
		file_model_book_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Feature); i {
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
		file_model_book_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
			RawDescriptor: file_model_book_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_model_book_proto_goTypes,
		DependencyIndexes: file_model_book_proto_depIdxs,
		MessageInfos:      file_model_book_proto_msgTypes,
	}.Build()
	File_model_book_proto = out.File
	file_model_book_proto_rawDesc = nil
	file_model_book_proto_goTypes = nil
	file_model_book_proto_depIdxs = nil
}