// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0
// 	protoc        v3.6.1
// source: protos/service_mesh.proto

package smp

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type ServiceMesh_Type int32

const (
	ServiceMesh_INVALID_MESH         ServiceMesh_Type = 0
	ServiceMesh_APP_MESH             ServiceMesh_Type = 1
	ServiceMesh_CONSUL               ServiceMesh_Type = 2
	ServiceMesh_ISTIO                ServiceMesh_Type = 3
	ServiceMesh_KUMA                 ServiceMesh_Type = 4
	ServiceMesh_LINKERD              ServiceMesh_Type = 5
	ServiceMesh_MAESH                ServiceMesh_Type = 6
	ServiceMesh_OCTARINE             ServiceMesh_Type = 7
	ServiceMesh_NETWORK_SERVICE_MESH ServiceMesh_Type = 8
	ServiceMesh_TANZU                ServiceMesh_Type = 9
	ServiceMesh_OPEN_SERVICE_MESH    ServiceMesh_Type = 10
	ServiceMesh_NGINX_SERVICE_MESH   ServiceMesh_Type = 11
)

// Enum value maps for ServiceMesh_Type.
var (
	ServiceMesh_Type_name = map[int32]string{
		0:  "INVALID_MESH",
		1:  "APP_MESH",
		2:  "CONSUL",
		3:  "ISTIO",
		4:  "KUMA",
		5:  "LINKERD",
		6:  "TRAEFIK_MESH",
		7:  "OCTARINE",
		8:  "NETWORK_SERVICE_MESH",
		9:  "TANZU",
		10: "OPEN_SERVICE_MESH",
		11: "NGINX_SERVICE_MESH",
	}
	ServiceMesh_Type_value = map[string]int32{
		"INVALID_MESH":         0,
		"APP_MESH":             1,
		"CONSUL":               2,
		"ISTIO":                3,
		"KUMA":                 4,
		"LINKERD":              5,
		"TRAEFIK_MESH":                6,
		"OCTARINE":             7,
		"NETWORK_SERVICE_MESH": 8,
		"TANZU":                9,
		"OPEN_SERVICE_MESH":    10,
		"NGINX_SERVICE_MESH":   11,
	}
)

func (x ServiceMesh_Type) Enum() *ServiceMesh_Type {
	p := new(ServiceMesh_Type)
	*p = x
	return p
}

func (x ServiceMesh_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ServiceMesh_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_protos_service_mesh_proto_enumTypes[0].Descriptor()
}

func (ServiceMesh_Type) Type() protoreflect.EnumType {
	return &file_protos_service_mesh_proto_enumTypes[0]
}

func (x ServiceMesh_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ServiceMesh_Type.Descriptor instead.
func (ServiceMesh_Type) EnumDescriptor() ([]byte, []int) {
	return file_protos_service_mesh_proto_rawDescGZIP(), []int{0, 0}
}

type ServiceMesh struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type        ServiceMesh_Type  `protobuf:"varint,1,opt,name=type,proto3,enum=smp.ServiceMesh_Type" json:"type,omitempty"`
	Version     string            `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	Labels      map[string]string `protobuf:"bytes,3,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Annotations map[string]string `protobuf:"bytes,4,rep,name=annotations,proto3" json:"annotations,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *ServiceMesh) Reset() {
	*x = ServiceMesh{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_service_mesh_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServiceMesh) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceMesh) ProtoMessage() {}

func (x *ServiceMesh) ProtoReflect() protoreflect.Message {
	mi := &file_protos_service_mesh_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceMesh.ProtoReflect.Descriptor instead.
func (*ServiceMesh) Descriptor() ([]byte, []int) {
	return file_protos_service_mesh_proto_rawDescGZIP(), []int{0}
}

func (x *ServiceMesh) GetType() ServiceMesh_Type {
	if x != nil {
		return x.Type
	}
	return ServiceMesh_INVALID_MESH
}

func (x *ServiceMesh) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *ServiceMesh) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *ServiceMesh) GetAnnotations() map[string]string {
	if x != nil {
		return x.Annotations
	}
	return nil
}

var File_protos_service_mesh_proto protoreflect.FileDescriptor

var file_protos_service_mesh_proto_rawDesc = []byte{
	0x0a, 0x19, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x5f, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x73, 0x6d, 0x70,
	0x22, 0x8c, 0x04, 0x0a, 0x0b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4d, 0x65, 0x73, 0x68,
	0x12, 0x29, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15,
	0x2e, 0x73, 0x6d, 0x70, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4d, 0x65, 0x73, 0x68,
	0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x34, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x73, 0x6d, 0x70, 0x2e, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x4d, 0x65, 0x73, 0x68, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x43, 0x0a, 0x0b, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x21, 0x2e, 0x73, 0x6d, 0x70, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4d, 0x65,
	0x73, 0x68, 0x2e, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x0b, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x1a, 0x39, 0x0a, 0x0b, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x3e, 0x0a, 0x10, 0x41,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xc1, 0x01, 0x0a, 0x04,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x10, 0x0a, 0x0c, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f,
	0x4d, 0x45, 0x53, 0x48, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x41, 0x50, 0x50, 0x5f, 0x4d, 0x45,
	0x53, 0x48, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x43, 0x4f, 0x4e, 0x53, 0x55, 0x4c, 0x10, 0x02,
	0x12, 0x09, 0x0a, 0x05, 0x49, 0x53, 0x54, 0x49, 0x4f, 0x10, 0x03, 0x12, 0x08, 0x0a, 0x04, 0x4b,
	0x55, 0x4d, 0x41, 0x10, 0x04, 0x12, 0x0b, 0x0a, 0x07, 0x4c, 0x49, 0x4e, 0x4b, 0x45, 0x52, 0x44,
	0x10, 0x05, 0x12, 0x09, 0x0a, 0x05, 0x4d, 0x41, 0x45, 0x53, 0x48, 0x10, 0x06, 0x12, 0x0c, 0x0a,
	0x08, 0x4f, 0x43, 0x54, 0x41, 0x52, 0x49, 0x4e, 0x45, 0x10, 0x07, 0x12, 0x18, 0x0a, 0x14, 0x4e,
	0x45, 0x54, 0x57, 0x4f, 0x52, 0x4b, 0x5f, 0x53, 0x45, 0x52, 0x56, 0x49, 0x43, 0x45, 0x5f, 0x4d,
	0x45, 0x53, 0x48, 0x10, 0x08, 0x12, 0x09, 0x0a, 0x05, 0x54, 0x41, 0x4e, 0x5a, 0x55, 0x10, 0x09,
	0x12, 0x15, 0x0a, 0x11, 0x4f, 0x50, 0x45, 0x4e, 0x5f, 0x53, 0x45, 0x52, 0x56, 0x49, 0x43, 0x45,
	0x5f, 0x4d, 0x45, 0x53, 0x48, 0x10, 0x0a, 0x12, 0x16, 0x0a, 0x12, 0x4e, 0x47, 0x49, 0x4e, 0x58,
	0x5f, 0x53, 0x45, 0x52, 0x56, 0x49, 0x43, 0x45, 0x5f, 0x4d, 0x45, 0x53, 0x48, 0x10, 0x0b, 0x42,
	0x3f, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x6d, 0x70, 0x42, 0x10, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x4d, 0x65, 0x73, 0x68, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x08,
	0x73, 0x70, 0x65, 0x63, 0x3b, 0x73, 0x6d, 0x70, 0xa2, 0x02, 0x03, 0x73, 0x6d, 0x70, 0xaa, 0x02,
	0x03, 0x53, 0x6d, 0x70, 0xca, 0x02, 0x03, 0x53, 0x6d, 0x70, 0xea, 0x02, 0x03, 0x53, 0x6d, 0x70,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_service_mesh_proto_rawDescOnce sync.Once
	file_protos_service_mesh_proto_rawDescData = file_protos_service_mesh_proto_rawDesc
)

func file_protos_service_mesh_proto_rawDescGZIP() []byte {
	file_protos_service_mesh_proto_rawDescOnce.Do(func() {
		file_protos_service_mesh_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_service_mesh_proto_rawDescData)
	})
	return file_protos_service_mesh_proto_rawDescData
}

var file_protos_service_mesh_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_protos_service_mesh_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_protos_service_mesh_proto_goTypes = []interface{}{
	(ServiceMesh_Type)(0), // 0: smp.ServiceMesh.Type
	(*ServiceMesh)(nil),   // 1: smp.ServiceMesh
	nil,                   // 2: smp.ServiceMesh.LabelsEntry
	nil,                   // 3: smp.ServiceMesh.AnnotationsEntry
}
var file_protos_service_mesh_proto_depIdxs = []int32{
	0, // 0: smp.ServiceMesh.type:type_name -> smp.ServiceMesh.Type
	2, // 1: smp.ServiceMesh.labels:type_name -> smp.ServiceMesh.LabelsEntry
	3, // 2: smp.ServiceMesh.annotations:type_name -> smp.ServiceMesh.AnnotationsEntry
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_protos_service_mesh_proto_init() }
func file_protos_service_mesh_proto_init() {
	if File_protos_service_mesh_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_service_mesh_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServiceMesh); i {
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
			RawDescriptor: file_protos_service_mesh_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protos_service_mesh_proto_goTypes,
		DependencyIndexes: file_protos_service_mesh_proto_depIdxs,
		EnumInfos:         file_protos_service_mesh_proto_enumTypes,
		MessageInfos:      file_protos_service_mesh_proto_msgTypes,
	}.Build()
	File_protos_service_mesh_proto = out.File
	file_protos_service_mesh_proto_rawDesc = nil
	file_protos_service_mesh_proto_goTypes = nil
	file_protos_service_mesh_proto_depIdxs = nil
}
