// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.0
// source: service.proto

package api

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

// CreateHabitRequest is the message sent to create a habit.
type CreateHabitRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of the new habit. Cannot be empty.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Frequency of the new habit. Defaults to once per week.
	WeeklyFrequency *int32 `protobuf:"varint,2,opt,name=weekly_frequency,json=weeklyFrequency,proto3,oneof" json:"weekly_frequency,omitempty"`
}

func (x *CreateHabitRequest) Reset() {
	*x = CreateHabitRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateHabitRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateHabitRequest) ProtoMessage() {}

func (x *CreateHabitRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateHabitRequest.ProtoReflect.Descriptor instead.
func (*CreateHabitRequest) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{0}
}

func (x *CreateHabitRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateHabitRequest) GetWeeklyFrequency() int32 {
	if x != nil && x.WeeklyFrequency != nil {
		return *x.WeeklyFrequency
	}
	return 0
}

// CreateHabitResponse is the response of the create endpoint.
type CreateHabitResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Habit *Habit `protobuf:"bytes,1,opt,name=habit,proto3" json:"habit,omitempty"`
}

func (x *CreateHabitResponse) Reset() {
	*x = CreateHabitResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateHabitResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateHabitResponse) ProtoMessage() {}

func (x *CreateHabitResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateHabitResponse.ProtoReflect.Descriptor instead.
func (*CreateHabitResponse) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{1}
}

func (x *CreateHabitResponse) GetHabit() *Habit {
	if x != nil {
		return x.Habit
	}
	return nil
}

var File_service_proto protoreflect.FileDescriptor

var file_service_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x06, 0x68, 0x61, 0x62, 0x69, 0x74, 0x73, 0x1a, 0x0b, 0x68, 0x61, 0x62, 0x69, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6d, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x48, 0x61,
	0x62, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2e,
	0x0a, 0x10, 0x77, 0x65, 0x65, 0x6b, 0x6c, 0x79, 0x5f, 0x66, 0x72, 0x65, 0x71, 0x75, 0x65, 0x6e,
	0x63, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00, 0x52, 0x0f, 0x77, 0x65, 0x65, 0x6b,
	0x6c, 0x79, 0x46, 0x72, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x79, 0x88, 0x01, 0x01, 0x42, 0x13,
	0x0a, 0x11, 0x5f, 0x77, 0x65, 0x65, 0x6b, 0x6c, 0x79, 0x5f, 0x66, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x6e, 0x63, 0x79, 0x22, 0x3a, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x48, 0x61, 0x62,
	0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a, 0x05, 0x68, 0x61,
	0x62, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x68, 0x61, 0x62, 0x69,
	0x74, 0x73, 0x2e, 0x48, 0x61, 0x62, 0x69, 0x74, 0x52, 0x05, 0x68, 0x61, 0x62, 0x69, 0x74, 0x32,
	0x50, 0x0a, 0x06, 0x48, 0x61, 0x62, 0x69, 0x74, 0x73, 0x12, 0x46, 0x0a, 0x0b, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x48, 0x61, 0x62, 0x69, 0x74, 0x12, 0x1a, 0x2e, 0x68, 0x61, 0x62, 0x69, 0x74,
	0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x48, 0x61, 0x62, 0x69, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x68, 0x61, 0x62, 0x69, 0x74, 0x73, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x48, 0x61, 0x62, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0x1c, 0x5a, 0x1a, 0x6c, 0x65, 0x61, 0x72, 0x6e, 0x67, 0x6f, 0x2d, 0x70, 0x6f, 0x63,
	0x6b, 0x65, 0x74, 0x73, 0x2f, 0x68, 0x61, 0x62, 0x69, 0x74, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_service_proto_rawDescOnce sync.Once
	file_service_proto_rawDescData = file_service_proto_rawDesc
)

func file_service_proto_rawDescGZIP() []byte {
	file_service_proto_rawDescOnce.Do(func() {
		file_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_proto_rawDescData)
	})
	return file_service_proto_rawDescData
}

var file_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_service_proto_goTypes = []interface{}{
	(*CreateHabitRequest)(nil),  // 0: habits.CreateHabitRequest
	(*CreateHabitResponse)(nil), // 1: habits.CreateHabitResponse
	(*Habit)(nil),               // 2: habits.Habit
}
var file_service_proto_depIdxs = []int32{
	2, // 0: habits.CreateHabitResponse.habit:type_name -> habits.Habit
	0, // 1: habits.Habits.CreateHabit:input_type -> habits.CreateHabitRequest
	1, // 2: habits.Habits.CreateHabit:output_type -> habits.CreateHabitResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_service_proto_init() }
func file_service_proto_init() {
	if File_service_proto != nil {
		return
	}
	file_habit_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateHabitRequest); i {
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
		file_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateHabitResponse); i {
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
	file_service_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_proto_goTypes,
		DependencyIndexes: file_service_proto_depIdxs,
		MessageInfos:      file_service_proto_msgTypes,
	}.Build()
	File_service_proto = out.File
	file_service_proto_rawDesc = nil
	file_service_proto_goTypes = nil
	file_service_proto_depIdxs = nil
}