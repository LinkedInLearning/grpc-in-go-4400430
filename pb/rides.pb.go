// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: rides.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type RideType int32

const (
	RideType_USET    RideType = 0
	RideType_REGULAR RideType = 1
	RideType_POOL    RideType = 2
)

// Enum value maps for RideType.
var (
	RideType_name = map[int32]string{
		0: "USET",
		1: "REGULAR",
		2: "POOL",
	}
	RideType_value = map[string]int32{
		"USET":    0,
		"REGULAR": 1,
		"POOL":    2,
	}
)

func (x RideType) Enum() *RideType {
	p := new(RideType)
	*p = x
	return p
}

func (x RideType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RideType) Descriptor() protoreflect.EnumDescriptor {
	return file_rides_proto_enumTypes[0].Descriptor()
}

func (RideType) Type() protoreflect.EnumType {
	return &file_rides_proto_enumTypes[0]
}

func (x RideType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RideType.Descriptor instead.
func (RideType) EnumDescriptor() ([]byte, []int) {
	return file_rides_proto_rawDescGZIP(), []int{0}
}

type Location struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Lat float64 `protobuf:"fixed64,1,opt,name=lat,proto3" json:"lat,omitempty"`
	Lng float64 `protobuf:"fixed64,2,opt,name=lng,proto3" json:"lng,omitempty"`
}

func (x *Location) Reset() {
	*x = Location{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rides_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Location) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Location) ProtoMessage() {}

func (x *Location) ProtoReflect() protoreflect.Message {
	mi := &file_rides_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Location.ProtoReflect.Descriptor instead.
func (*Location) Descriptor() ([]byte, []int) {
	return file_rides_proto_rawDescGZIP(), []int{0}
}

func (x *Location) GetLat() float64 {
	if x != nil {
		return x.Lat
	}
	return 0
}

func (x *Location) GetLng() float64 {
	if x != nil {
		return x.Lng
	}
	return 0
}

type StartRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	DriverId     string                 `protobuf:"bytes,2,opt,name=driver_id,json=driverId,proto3" json:"driver_id,omitempty"`
	Location     *Location              `protobuf:"bytes,3,opt,name=location,proto3" json:"location,omitempty"`
	PassengerIds []string               `protobuf:"bytes,4,rep,name=passenger_ids,json=passengerIds,proto3" json:"passenger_ids,omitempty"`
	Time         *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=time,proto3" json:"time,omitempty"`
	Type         RideType               `protobuf:"varint,6,opt,name=type,proto3,enum=RideType" json:"type,omitempty"`
}

func (x *StartRequest) Reset() {
	*x = StartRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rides_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartRequest) ProtoMessage() {}

func (x *StartRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rides_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartRequest.ProtoReflect.Descriptor instead.
func (*StartRequest) Descriptor() ([]byte, []int) {
	return file_rides_proto_rawDescGZIP(), []int{1}
}

func (x *StartRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *StartRequest) GetDriverId() string {
	if x != nil {
		return x.DriverId
	}
	return ""
}

func (x *StartRequest) GetLocation() *Location {
	if x != nil {
		return x.Location
	}
	return nil
}

func (x *StartRequest) GetPassengerIds() []string {
	if x != nil {
		return x.PassengerIds
	}
	return nil
}

func (x *StartRequest) GetTime() *timestamppb.Timestamp {
	if x != nil {
		return x.Time
	}
	return nil
}

func (x *StartRequest) GetType() RideType {
	if x != nil {
		return x.Type
	}
	return RideType_USET
}

type StartResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *StartResponse) Reset() {
	*x = StartResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rides_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartResponse) ProtoMessage() {}

func (x *StartResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rides_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartResponse.ProtoReflect.Descriptor instead.
func (*StartResponse) Descriptor() ([]byte, []int) {
	return file_rides_proto_rawDescGZIP(), []int{2}
}

func (x *StartResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type EndRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Time     *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=time,proto3" json:"time,omitempty"`
	Distance float64                `protobuf:"fixed64,3,opt,name=distance,proto3" json:"distance,omitempty"`
}

func (x *EndRequest) Reset() {
	*x = EndRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rides_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EndRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EndRequest) ProtoMessage() {}

func (x *EndRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rides_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EndRequest.ProtoReflect.Descriptor instead.
func (*EndRequest) Descriptor() ([]byte, []int) {
	return file_rides_proto_rawDescGZIP(), []int{3}
}

func (x *EndRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *EndRequest) GetTime() *timestamppb.Timestamp {
	if x != nil {
		return x.Time
	}
	return nil
}

func (x *EndRequest) GetDistance() float64 {
	if x != nil {
		return x.Distance
	}
	return 0
}

type EndResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *EndResponse) Reset() {
	*x = EndResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rides_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EndResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EndResponse) ProtoMessage() {}

func (x *EndResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rides_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EndResponse.ProtoReflect.Descriptor instead.
func (*EndResponse) Descriptor() ([]byte, []int) {
	return file_rides_proto_rawDescGZIP(), []int{4}
}

func (x *EndResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_rides_proto protoreflect.FileDescriptor

var file_rides_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x72, 0x69, 0x64, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2e,
	0x0a, 0x08, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x61,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03, 0x6c, 0x61, 0x74, 0x12, 0x10, 0x0a, 0x03,
	0x6c, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03, 0x6c, 0x6e, 0x67, 0x22, 0xd6,
	0x01, 0x0a, 0x0c, 0x53, 0x74, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x1b, 0x0a, 0x09, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x08,
	0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09,
	0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x23, 0x0a, 0x0d, 0x70, 0x61, 0x73, 0x73, 0x65, 0x6e, 0x67, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x61, 0x73, 0x73,
	0x65, 0x6e, 0x67, 0x65, 0x72, 0x49, 0x64, 0x73, 0x12, 0x2e, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x09, 0x2e, 0x52, 0x69, 0x64, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x1f, 0x0a, 0x0d, 0x53, 0x74, 0x61, 0x72, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x68, 0x0a, 0x0a, 0x45, 0x6e, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x2e, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x69, 0x73, 0x74, 0x61, 0x6e,
	0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x64, 0x69, 0x73, 0x74, 0x61, 0x6e,
	0x63, 0x65, 0x22, 0x1d, 0x0a, 0x0b, 0x45, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x2a, 0x2b, 0x0a, 0x08, 0x52, 0x69, 0x64, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x08, 0x0a,
	0x04, 0x55, 0x53, 0x45, 0x54, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x52, 0x45, 0x47, 0x55, 0x4c,
	0x41, 0x52, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x50, 0x4f, 0x4f, 0x4c, 0x10, 0x02, 0x32, 0x55,
	0x0a, 0x05, 0x52, 0x69, 0x64, 0x65, 0x73, 0x12, 0x28, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x72, 0x74,
	0x12, 0x0d, 0x2e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x0e, 0x2e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x22, 0x0a, 0x03, 0x45, 0x6e, 0x64, 0x12, 0x0b, 0x2e, 0x45, 0x6e, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x45, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x22, 0x5a, 0x20, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x33, 0x35, 0x33, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x72, 0x69, 0x64, 0x65, 0x73, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_rides_proto_rawDescOnce sync.Once
	file_rides_proto_rawDescData = file_rides_proto_rawDesc
)

func file_rides_proto_rawDescGZIP() []byte {
	file_rides_proto_rawDescOnce.Do(func() {
		file_rides_proto_rawDescData = protoimpl.X.CompressGZIP(file_rides_proto_rawDescData)
	})
	return file_rides_proto_rawDescData
}

var file_rides_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_rides_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_rides_proto_goTypes = []interface{}{
	(RideType)(0),                 // 0: RideType
	(*Location)(nil),              // 1: Location
	(*StartRequest)(nil),          // 2: StartRequest
	(*StartResponse)(nil),         // 3: StartResponse
	(*EndRequest)(nil),            // 4: EndRequest
	(*EndResponse)(nil),           // 5: EndResponse
	(*timestamppb.Timestamp)(nil), // 6: google.protobuf.Timestamp
}
var file_rides_proto_depIdxs = []int32{
	1, // 0: StartRequest.location:type_name -> Location
	6, // 1: StartRequest.time:type_name -> google.protobuf.Timestamp
	0, // 2: StartRequest.type:type_name -> RideType
	6, // 3: EndRequest.time:type_name -> google.protobuf.Timestamp
	2, // 4: Rides.Start:input_type -> StartRequest
	4, // 5: Rides.End:input_type -> EndRequest
	3, // 6: Rides.Start:output_type -> StartResponse
	5, // 7: Rides.End:output_type -> EndResponse
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_rides_proto_init() }
func file_rides_proto_init() {
	if File_rides_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rides_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Location); i {
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
		file_rides_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartRequest); i {
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
		file_rides_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartResponse); i {
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
		file_rides_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EndRequest); i {
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
		file_rides_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EndResponse); i {
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
			RawDescriptor: file_rides_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_rides_proto_goTypes,
		DependencyIndexes: file_rides_proto_depIdxs,
		EnumInfos:         file_rides_proto_enumTypes,
		MessageInfos:      file_rides_proto_msgTypes,
	}.Build()
	File_rides_proto = out.File
	file_rides_proto_rawDesc = nil
	file_rides_proto_goTypes = nil
	file_rides_proto_depIdxs = nil
}
