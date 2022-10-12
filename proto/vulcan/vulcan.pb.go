// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: vulcan.proto

package vulcan

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

type ResponseSelectMobileAppByAppID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status    string            `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Message   string            `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	APP_ID    string            `protobuf:"bytes,3,opt,name=APP_ID,json=APPID,proto3" json:"APP_ID,omitempty"`
	Title     string            `protobuf:"bytes,4,opt,name=title,proto3" json:"title,omitempty"`
	TableName string            `protobuf:"bytes,5,opt,name=table_name,json=tableName,proto3" json:"table_name,omitempty"`
	Result    map[string]string `protobuf:"bytes,6,rep,name=result,proto3" json:"result,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *ResponseSelectMobileAppByAppID) Reset() {
	*x = ResponseSelectMobileAppByAppID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vulcan_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseSelectMobileAppByAppID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseSelectMobileAppByAppID) ProtoMessage() {}

func (x *ResponseSelectMobileAppByAppID) ProtoReflect() protoreflect.Message {
	mi := &file_vulcan_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseSelectMobileAppByAppID.ProtoReflect.Descriptor instead.
func (*ResponseSelectMobileAppByAppID) Descriptor() ([]byte, []int) {
	return file_vulcan_proto_rawDescGZIP(), []int{0}
}

func (x *ResponseSelectMobileAppByAppID) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *ResponseSelectMobileAppByAppID) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ResponseSelectMobileAppByAppID) GetAPP_ID() string {
	if x != nil {
		return x.APP_ID
	}
	return ""
}

func (x *ResponseSelectMobileAppByAppID) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *ResponseSelectMobileAppByAppID) GetTableName() string {
	if x != nil {
		return x.TableName
	}
	return ""
}

func (x *ResponseSelectMobileAppByAppID) GetResult() map[string]string {
	if x != nil {
		return x.Result
	}
	return nil
}

type RequestSelectMobileAppByAppID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	APP_ID string `protobuf:"bytes,1,opt,name=APP_ID,json=APPID,proto3" json:"APP_ID,omitempty"`
}

func (x *RequestSelectMobileAppByAppID) Reset() {
	*x = RequestSelectMobileAppByAppID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vulcan_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestSelectMobileAppByAppID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestSelectMobileAppByAppID) ProtoMessage() {}

func (x *RequestSelectMobileAppByAppID) ProtoReflect() protoreflect.Message {
	mi := &file_vulcan_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestSelectMobileAppByAppID.ProtoReflect.Descriptor instead.
func (*RequestSelectMobileAppByAppID) Descriptor() ([]byte, []int) {
	return file_vulcan_proto_rawDescGZIP(), []int{1}
}

func (x *RequestSelectMobileAppByAppID) GetAPP_ID() string {
	if x != nil {
		return x.APP_ID
	}
	return ""
}

var File_vulcan_proto protoreflect.FileDescriptor

var file_vulcan_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x76, 0x75, 0x6c, 0x63, 0x61, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9e,
	0x02, 0x0a, 0x1e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x53, 0x65, 0x6c, 0x65, 0x63,
	0x74, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x41, 0x70, 0x70, 0x42, 0x79, 0x41, 0x70, 0x70, 0x49,
	0x44, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x15, 0x0a, 0x06, 0x41, 0x50, 0x50, 0x5f, 0x49, 0x44, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x41, 0x50, 0x50, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x43, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x2b, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74,
	0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x41, 0x70, 0x70, 0x42, 0x79, 0x41, 0x70, 0x70, 0x49, 0x44,
	0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x1a, 0x39, 0x0a, 0x0b, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22,
	0x36, 0x0a, 0x1d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74,
	0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x41, 0x70, 0x70, 0x42, 0x79, 0x41, 0x70, 0x70, 0x49, 0x44,
	0x12, 0x15, 0x0a, 0x06, 0x41, 0x50, 0x50, 0x5f, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x41, 0x50, 0x50, 0x49, 0x44, 0x32, 0x69, 0x0a, 0x0a, 0x4d, 0x6f, 0x62, 0x69, 0x6c,
	0x65, 0x41, 0x70, 0x70, 0x73, 0x12, 0x5b, 0x0a, 0x16, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x4d,
	0x6f, 0x62, 0x69, 0x6c, 0x65, 0x41, 0x70, 0x70, 0x42, 0x79, 0x41, 0x70, 0x70, 0x49, 0x44, 0x12,
	0x1e, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x4d,
	0x6f, 0x62, 0x69, 0x6c, 0x65, 0x41, 0x70, 0x70, 0x42, 0x79, 0x41, 0x70, 0x70, 0x49, 0x44, 0x1a,
	0x1f, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74,
	0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x41, 0x70, 0x70, 0x42, 0x79, 0x41, 0x70, 0x70, 0x49, 0x44,
	0x22, 0x00, 0x42, 0x18, 0x5a, 0x16, 0x67, 0x6f, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x32, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x75, 0x6c, 0x63, 0x61, 0x6e, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_vulcan_proto_rawDescOnce sync.Once
	file_vulcan_proto_rawDescData = file_vulcan_proto_rawDesc
)

func file_vulcan_proto_rawDescGZIP() []byte {
	file_vulcan_proto_rawDescOnce.Do(func() {
		file_vulcan_proto_rawDescData = protoimpl.X.CompressGZIP(file_vulcan_proto_rawDescData)
	})
	return file_vulcan_proto_rawDescData
}

var file_vulcan_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_vulcan_proto_goTypes = []interface{}{
	(*ResponseSelectMobileAppByAppID)(nil), // 0: ResponseSelectMobileAppByAppID
	(*RequestSelectMobileAppByAppID)(nil),  // 1: RequestSelectMobileAppByAppID
	nil,                                    // 2: ResponseSelectMobileAppByAppID.ResultEntry
}
var file_vulcan_proto_depIdxs = []int32{
	2, // 0: ResponseSelectMobileAppByAppID.result:type_name -> ResponseSelectMobileAppByAppID.ResultEntry
	1, // 1: MobileApps.SelectMobileAppByAppID:input_type -> RequestSelectMobileAppByAppID
	0, // 2: MobileApps.SelectMobileAppByAppID:output_type -> ResponseSelectMobileAppByAppID
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_vulcan_proto_init() }
func file_vulcan_proto_init() {
	if File_vulcan_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_vulcan_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseSelectMobileAppByAppID); i {
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
		file_vulcan_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestSelectMobileAppByAppID); i {
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
			RawDescriptor: file_vulcan_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_vulcan_proto_goTypes,
		DependencyIndexes: file_vulcan_proto_depIdxs,
		MessageInfos:      file_vulcan_proto_msgTypes,
	}.Build()
	File_vulcan_proto = out.File
	file_vulcan_proto_rawDesc = nil
	file_vulcan_proto_goTypes = nil
	file_vulcan_proto_depIdxs = nil
}