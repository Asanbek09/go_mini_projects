// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.3
// source: habit.proto

package api

import (
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

type Habit struct {
	state           protoimpl.MessageState `protogen:"open.v1"`
	Name            string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	WeeklyFrequency int32                  `protobuf:"varint,2,opt,name=weekly_frequency,json=weeklyFrequency,proto3" json:"weekly_frequency,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *Habit) Reset() {
	*x = Habit{}
	mi := &file_habit_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Habit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Habit) ProtoMessage() {}

func (x *Habit) ProtoReflect() protoreflect.Message {
	mi := &file_habit_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Habit.ProtoReflect.Descriptor instead.
func (*Habit) Descriptor() ([]byte, []int) {
	return file_habit_proto_rawDescGZIP(), []int{0}
}

func (x *Habit) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Habit) GetWeeklyFrequency() int32 {
	if x != nil {
		return x.WeeklyFrequency
	}
	return 0
}

var File_habit_proto protoreflect.FileDescriptor

var file_habit_proto_rawDesc = string([]byte{
	0x0a, 0x0b, 0x68, 0x61, 0x62, 0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x68,
	0x61, 0x62, 0x69, 0x74, 0x73, 0x22, 0x46, 0x0a, 0x05, 0x48, 0x61, 0x62, 0x69, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x29, 0x0a, 0x10, 0x77, 0x65, 0x65, 0x6b, 0x6c, 0x79, 0x5f, 0x66, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0f, 0x77, 0x65,
	0x65, 0x6b, 0x6c, 0x79, 0x46, 0x72, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x79, 0x42, 0x0c, 0x5a,
	0x0a, 0x68, 0x61, 0x62, 0x69, 0x74, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
})

var (
	file_habit_proto_rawDescOnce sync.Once
	file_habit_proto_rawDescData []byte
)

func file_habit_proto_rawDescGZIP() []byte {
	file_habit_proto_rawDescOnce.Do(func() {
		file_habit_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_habit_proto_rawDesc), len(file_habit_proto_rawDesc)))
	})
	return file_habit_proto_rawDescData
}

var file_habit_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_habit_proto_goTypes = []any{
	(*Habit)(nil), // 0: habits.Habit
}
var file_habit_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_habit_proto_init() }
func file_habit_proto_init() {
	if File_habit_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_habit_proto_rawDesc), len(file_habit_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_habit_proto_goTypes,
		DependencyIndexes: file_habit_proto_depIdxs,
		MessageInfos:      file_habit_proto_msgTypes,
	}.Build()
	File_habit_proto = out.File
	file_habit_proto_goTypes = nil
	file_habit_proto_depIdxs = nil
}
