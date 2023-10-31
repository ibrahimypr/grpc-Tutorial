// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.4
// source: proto/weather.proto

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

type WeatherForecast struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data        string  `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Description string  `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Temperature float32 `protobuf:"fixed32,3,opt,name=temperature,proto3" json:"temperature,omitempty"`
}

func (x *WeatherForecast) Reset() {
	*x = WeatherForecast{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_weather_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WeatherForecast) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WeatherForecast) ProtoMessage() {}

func (x *WeatherForecast) ProtoReflect() protoreflect.Message {
	mi := &file_proto_weather_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WeatherForecast.ProtoReflect.Descriptor instead.
func (*WeatherForecast) Descriptor() ([]byte, []int) {
	return file_proto_weather_proto_rawDescGZIP(), []int{0}
}

func (x *WeatherForecast) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

func (x *WeatherForecast) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *WeatherForecast) GetTemperature() float32 {
	if x != nil {
		return x.Temperature
	}
	return 0
}

type CityRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CityName string `protobuf:"bytes,1,opt,name=city_name,json=cityName,proto3" json:"city_name,omitempty"`
}

func (x *CityRequest) Reset() {
	*x = CityRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_weather_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CityRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CityRequest) ProtoMessage() {}

func (x *CityRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_weather_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CityRequest.ProtoReflect.Descriptor instead.
func (*CityRequest) Descriptor() ([]byte, []int) {
	return file_proto_weather_proto_rawDescGZIP(), []int{1}
}

func (x *CityRequest) GetCityName() string {
	if x != nil {
		return x.CityName
	}
	return ""
}

var File_proto_weather_proto protoreflect.FileDescriptor

var file_proto_weather_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x77, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x77, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x22, 0x69,
	0x0a, 0x0f, 0x57, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x46, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x73,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x74, 0x65, 0x6d, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0b, 0x74, 0x65,
	0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65, 0x22, 0x2a, 0x0a, 0x0b, 0x43, 0x69, 0x74,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x69, 0x74, 0x79,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x69, 0x74,
	0x79, 0x4e, 0x61, 0x6d, 0x65, 0x32, 0x57, 0x0a, 0x0e, 0x57, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x45, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x57, 0x65,
	0x65, 0x6b, 0x6c, 0x79, 0x46, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x73, 0x74, 0x12, 0x14, 0x2e, 0x77,
	0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x2e, 0x43, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x18, 0x2e, 0x77, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x2e, 0x57, 0x65, 0x61,
	0x74, 0x68, 0x65, 0x72, 0x46, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x73, 0x74, 0x30, 0x01, 0x42, 0x09,
	0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_proto_weather_proto_rawDescOnce sync.Once
	file_proto_weather_proto_rawDescData = file_proto_weather_proto_rawDesc
)

func file_proto_weather_proto_rawDescGZIP() []byte {
	file_proto_weather_proto_rawDescOnce.Do(func() {
		file_proto_weather_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_weather_proto_rawDescData)
	})
	return file_proto_weather_proto_rawDescData
}

var file_proto_weather_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_weather_proto_goTypes = []interface{}{
	(*WeatherForecast)(nil), // 0: weather.WeatherForecast
	(*CityRequest)(nil),     // 1: weather.CityRequest
}
var file_proto_weather_proto_depIdxs = []int32{
	1, // 0: weather.WeatherService.GetWeeklyForecast:input_type -> weather.CityRequest
	0, // 1: weather.WeatherService.GetWeeklyForecast:output_type -> weather.WeatherForecast
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_weather_proto_init() }
func file_proto_weather_proto_init() {
	if File_proto_weather_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_weather_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WeatherForecast); i {
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
		file_proto_weather_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CityRequest); i {
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
			RawDescriptor: file_proto_weather_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_weather_proto_goTypes,
		DependencyIndexes: file_proto_weather_proto_depIdxs,
		MessageInfos:      file_proto_weather_proto_msgTypes,
	}.Build()
	File_proto_weather_proto = out.File
	file_proto_weather_proto_rawDesc = nil
	file_proto_weather_proto_goTypes = nil
	file_proto_weather_proto_depIdxs = nil
}
