// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: notes_service/service/v1/notes.proto

package pb_notes_service

import (
	v1 "github.com/almalii/grpc-contracts/gen/go/notes_service/model/v1"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_notes_service_service_v1_notes_proto protoreflect.FileDescriptor

var file_notes_service_service_v1_notes_proto_rawDesc = []byte{
	0x0a, 0x24, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x6e, 0x6f, 0x74, 0x65, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31,
	0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x6e, 0x6f, 0x74,
	0x65, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x2f, 0x76, 0x31, 0x2f, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32,
	0xd8, 0x04, 0x0a, 0x0c, 0x4e, 0x6f, 0x74, 0x65, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x75, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x74, 0x65, 0x12, 0x29,
	0x2e, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x6f,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x6e, 0x6f, 0x74, 0x65,
	0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e,
	0x76, 0x31, 0x2e, 0x4e, 0x6f, 0x74, 0x65, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x14, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x3a, 0x01, 0x2a, 0x22, 0x09, 0x2f, 0x76,
	0x31, 0x2f, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x12, 0x71, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x4e, 0x6f,
	0x74, 0x65, 0x12, 0x25, 0x2e, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x6f, 0x74, 0x65,
	0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x6e, 0x6f, 0x74, 0x65,
	0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x16, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10, 0x12, 0x0e, 0x2f, 0x76, 0x31, 0x2f,
	0x6e, 0x6f, 0x74, 0x65, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x79, 0x0a, 0x08, 0x47, 0x65,
	0x74, 0x4e, 0x6f, 0x74, 0x65, 0x73, 0x12, 0x27, 0x2e, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e,
	0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x28, 0x2e, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x14, 0x12, 0x12, 0x2f, 0x76, 0x31, 0x2f, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x2f, 0x61, 0x6c, 0x6c,
	0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x7e, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4e,
	0x6f, 0x74, 0x65, 0x12, 0x29, 0x2e, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x4e, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a,
	0x2e, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x6f,
	0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x19, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x13, 0x3a, 0x01, 0x2a, 0x1a, 0x0e, 0x2f, 0x76, 0x31, 0x2f, 0x6e, 0x6f, 0x74, 0x65, 0x73,
	0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x63, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4e,
	0x6f, 0x74, 0x65, 0x12, 0x25, 0x2e, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x6f, 0x74,
	0x65, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x22, 0x16, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10, 0x2a, 0x0e, 0x2f, 0x76, 0x31, 0x2f,
	0x6e, 0x6f, 0x74, 0x65, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x42, 0xfc, 0x01, 0x0a, 0x1c, 0x63,
	0x6f, 0x6d, 0x2e, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x42, 0x0a, 0x4e, 0x6f, 0x74,
	0x65, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x52, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x6c, 0x6d, 0x61, 0x6c, 0x69, 0x69, 0x2f, 0x67, 0x72,
	0x70, 0x63, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x73, 0x2f, 0x67, 0x65, 0x6e,
	0x2f, 0x67, 0x6f, 0x2f, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x70, 0x62, 0x5f,
	0x6e, 0x6f, 0x74, 0x65, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0xa2, 0x02, 0x03,
	0x4e, 0x53, 0x58, 0xaa, 0x02, 0x17, 0x4e, 0x6f, 0x74, 0x65, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x17,
	0x4e, 0x6f, 0x74, 0x65, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5c, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x23, 0x4e, 0x6f, 0x74, 0x65, 0x73, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5c, 0x56,
	0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x19,
	0x4e, 0x6f, 0x74, 0x65, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x3a, 0x3a, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var file_notes_service_service_v1_notes_proto_goTypes = []interface{}{
	(*v1.CreateNoteRequest)(nil),  // 0: notes_service.model.v1.CreateNoteRequest
	(*v1.NoteIDRequest)(nil),      // 1: notes_service.model.v1.NoteIDRequest
	(*v1.AuthorIDRequest)(nil),    // 2: notes_service.model.v1.AuthorIDRequest
	(*v1.UpdateNoteRequest)(nil),  // 3: notes_service.model.v1.UpdateNoteRequest
	(*v1.NoteIDResponse)(nil),     // 4: notes_service.model.v1.NoteIDResponse
	(*v1.GetNoteResponse)(nil),    // 5: notes_service.model.v1.GetNoteResponse
	(*v1.NoteResponseList)(nil),   // 6: notes_service.model.v1.NoteResponseList
	(*v1.UpdateNoteResponse)(nil), // 7: notes_service.model.v1.UpdateNoteResponse
	(*emptypb.Empty)(nil),         // 8: google.protobuf.Empty
}
var file_notes_service_service_v1_notes_proto_depIdxs = []int32{
	0, // 0: notes_service.service.v1.NotesService.CreateNote:input_type -> notes_service.model.v1.CreateNoteRequest
	1, // 1: notes_service.service.v1.NotesService.GetNote:input_type -> notes_service.model.v1.NoteIDRequest
	2, // 2: notes_service.service.v1.NotesService.GetNotes:input_type -> notes_service.model.v1.AuthorIDRequest
	3, // 3: notes_service.service.v1.NotesService.UpdateNote:input_type -> notes_service.model.v1.UpdateNoteRequest
	1, // 4: notes_service.service.v1.NotesService.DeleteNote:input_type -> notes_service.model.v1.NoteIDRequest
	4, // 5: notes_service.service.v1.NotesService.CreateNote:output_type -> notes_service.model.v1.NoteIDResponse
	5, // 6: notes_service.service.v1.NotesService.GetNote:output_type -> notes_service.model.v1.GetNoteResponse
	6, // 7: notes_service.service.v1.NotesService.GetNotes:output_type -> notes_service.model.v1.NoteResponseList
	7, // 8: notes_service.service.v1.NotesService.UpdateNote:output_type -> notes_service.model.v1.UpdateNoteResponse
	8, // 9: notes_service.service.v1.NotesService.DeleteNote:output_type -> google.protobuf.Empty
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_notes_service_service_v1_notes_proto_init() }
func file_notes_service_service_v1_notes_proto_init() {
	if File_notes_service_service_v1_notes_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_notes_service_service_v1_notes_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_notes_service_service_v1_notes_proto_goTypes,
		DependencyIndexes: file_notes_service_service_v1_notes_proto_depIdxs,
	}.Build()
	File_notes_service_service_v1_notes_proto = out.File
	file_notes_service_service_v1_notes_proto_rawDesc = nil
	file_notes_service_service_v1_notes_proto_goTypes = nil
	file_notes_service_service_v1_notes_proto_depIdxs = nil
}
