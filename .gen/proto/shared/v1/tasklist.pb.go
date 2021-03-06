// Copyright (c) 2020 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: uber/cadence/shared/v1/tasklist.proto

package sharedv1

import (
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// TaskSource is the source from which a task was produced.
type TaskSource int32

const (
	TaskSource_TASK_SOURCE_INVALID TaskSource = 0
	// Task produced by history service.
	TaskSource_TASK_SOURCE_HISTORY TaskSource = 1
	// Task produced from matching db backlog.
	TaskSource_TASK_SOURCE_DB_BACKLOG TaskSource = 2
)

// Enum value maps for TaskSource.
var (
	TaskSource_name = map[int32]string{
		0: "TASK_SOURCE_INVALID",
		1: "TASK_SOURCE_HISTORY",
		2: "TASK_SOURCE_DB_BACKLOG",
	}
	TaskSource_value = map[string]int32{
		"TASK_SOURCE_INVALID":    0,
		"TASK_SOURCE_HISTORY":    1,
		"TASK_SOURCE_DB_BACKLOG": 2,
	}
)

func (x TaskSource) Enum() *TaskSource {
	p := new(TaskSource)
	*p = x
	return p
}

func (x TaskSource) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TaskSource) Descriptor() protoreflect.EnumDescriptor {
	return file_uber_cadence_shared_v1_tasklist_proto_enumTypes[0].Descriptor()
}

func (TaskSource) Type() protoreflect.EnumType {
	return &file_uber_cadence_shared_v1_tasklist_proto_enumTypes[0]
}

func (x TaskSource) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TaskSource.Descriptor instead.
func (TaskSource) EnumDescriptor() ([]byte, []int) {
	return file_uber_cadence_shared_v1_tasklist_proto_rawDescGZIP(), []int{0}
}

var File_uber_cadence_shared_v1_tasklist_proto protoreflect.FileDescriptor

var file_uber_cadence_shared_v1_tasklist_proto_rawDesc = []byte{
	0x0a, 0x25, 0x75, 0x62, 0x65, 0x72, 0x2f, 0x63, 0x61, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x2f, 0x73,
	0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x6c, 0x69, 0x73,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x75, 0x62, 0x65, 0x72, 0x2e, 0x63, 0x61,
	0x64, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x76, 0x31, 0x2a,
	0x5a, 0x0a, 0x0a, 0x54, 0x61, 0x73, 0x6b, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x17, 0x0a,
	0x13, 0x54, 0x41, 0x53, 0x4b, 0x5f, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x49, 0x4e, 0x56,
	0x41, 0x4c, 0x49, 0x44, 0x10, 0x00, 0x12, 0x17, 0x0a, 0x13, 0x54, 0x41, 0x53, 0x4b, 0x5f, 0x53,
	0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x48, 0x49, 0x53, 0x54, 0x4f, 0x52, 0x59, 0x10, 0x01, 0x12,
	0x1a, 0x0a, 0x16, 0x54, 0x41, 0x53, 0x4b, 0x5f, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x44,
	0x42, 0x5f, 0x42, 0x41, 0x43, 0x4b, 0x4c, 0x4f, 0x47, 0x10, 0x02, 0x42, 0x37, 0x5a, 0x35, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x75, 0x62, 0x65, 0x72, 0x2f, 0x63,
	0x61, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x2f, 0x2e, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x76, 0x31, 0x3b, 0x73, 0x68, 0x61, 0x72,
	0x65, 0x64, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_uber_cadence_shared_v1_tasklist_proto_rawDescOnce sync.Once
	file_uber_cadence_shared_v1_tasklist_proto_rawDescData = file_uber_cadence_shared_v1_tasklist_proto_rawDesc
)

func file_uber_cadence_shared_v1_tasklist_proto_rawDescGZIP() []byte {
	file_uber_cadence_shared_v1_tasklist_proto_rawDescOnce.Do(func() {
		file_uber_cadence_shared_v1_tasklist_proto_rawDescData = protoimpl.X.CompressGZIP(file_uber_cadence_shared_v1_tasklist_proto_rawDescData)
	})
	return file_uber_cadence_shared_v1_tasklist_proto_rawDescData
}

var file_uber_cadence_shared_v1_tasklist_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_uber_cadence_shared_v1_tasklist_proto_goTypes = []interface{}{
	(TaskSource)(0), // 0: uber.cadence.shared.v1.TaskSource
}
var file_uber_cadence_shared_v1_tasklist_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_uber_cadence_shared_v1_tasklist_proto_init() }
func file_uber_cadence_shared_v1_tasklist_proto_init() {
	if File_uber_cadence_shared_v1_tasklist_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_uber_cadence_shared_v1_tasklist_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_uber_cadence_shared_v1_tasklist_proto_goTypes,
		DependencyIndexes: file_uber_cadence_shared_v1_tasklist_proto_depIdxs,
		EnumInfos:         file_uber_cadence_shared_v1_tasklist_proto_enumTypes,
	}.Build()
	File_uber_cadence_shared_v1_tasklist_proto = out.File
	file_uber_cadence_shared_v1_tasklist_proto_rawDesc = nil
	file_uber_cadence_shared_v1_tasklist_proto_goTypes = nil
	file_uber_cadence_shared_v1_tasklist_proto_depIdxs = nil
}
