// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.26.1
// source: srv.proto

package __

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_srv_proto protoreflect.FileDescriptor

var file_srv_proto_rawDesc = []byte{
	0x0a, 0x09, 0x73, 0x72, 0x76, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x61, 0x76, 0x65,
	0x72, 0x61, 0x67, 0x65, 0x1a, 0x09, 0x6d, 0x73, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32,
	0x3f, 0x0a, 0x0e, 0x41, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x2d, 0x0a, 0x07, 0x41, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x12, 0x0f, 0x2e, 0x61,
	0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x41, 0x76, 0x67, 0x52, 0x65, 0x71, 0x1a, 0x0f, 0x2e,
	0x61, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x41, 0x76, 0x67, 0x52, 0x65, 0x73, 0x28, 0x01,
	0x42, 0x04, 0x5a, 0x02, 0x2e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_srv_proto_goTypes = []interface{}{
	(*AvgReq)(nil), // 0: average.AvgReq
	(*AvgRes)(nil), // 1: average.AvgRes
}
var file_srv_proto_depIdxs = []int32{
	0, // 0: average.AverageService.Average:input_type -> average.AvgReq
	1, // 1: average.AverageService.Average:output_type -> average.AvgRes
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_srv_proto_init() }
func file_srv_proto_init() {
	if File_srv_proto != nil {
		return
	}
	file_msg_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_srv_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_srv_proto_goTypes,
		DependencyIndexes: file_srv_proto_depIdxs,
	}.Build()
	File_srv_proto = out.File
	file_srv_proto_rawDesc = nil
	file_srv_proto_goTypes = nil
	file_srv_proto_depIdxs = nil
}