//
// Copyright (c) 2024, NVIDIA CORPORATION.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.0
// source: topology.proto

package protos

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

type TopologyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Provider    string   `protobuf:"bytes,1,opt,name=provider,proto3" json:"provider,omitempty"`
	Region      string   `protobuf:"bytes,2,opt,name=region,proto3" json:"region,omitempty"`
	InstanceIds []string `protobuf:"bytes,3,rep,name=instance_ids,json=instanceIds,proto3" json:"instance_ids,omitempty"`
}

func (x *TopologyRequest) Reset() {
	*x = TopologyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_topology_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TopologyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TopologyRequest) ProtoMessage() {}

func (x *TopologyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_topology_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TopologyRequest.ProtoReflect.Descriptor instead.
func (*TopologyRequest) Descriptor() ([]byte, []int) {
	return file_topology_proto_rawDescGZIP(), []int{0}
}

func (x *TopologyRequest) GetProvider() string {
	if x != nil {
		return x.Provider
	}
	return ""
}

func (x *TopologyRequest) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *TopologyRequest) GetInstanceIds() []string {
	if x != nil {
		return x.InstanceIds
	}
	return nil
}

type TopologyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Instances []*Instance `protobuf:"bytes,1,rep,name=instances,proto3" json:"instances,omitempty"`
}

func (x *TopologyResponse) Reset() {
	*x = TopologyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_topology_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TopologyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TopologyResponse) ProtoMessage() {}

func (x *TopologyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_topology_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TopologyResponse.ProtoReflect.Descriptor instead.
func (*TopologyResponse) Descriptor() ([]byte, []int) {
	return file_topology_proto_rawDescGZIP(), []int{1}
}

func (x *TopologyResponse) GetInstances() []*Instance {
	if x != nil {
		return x.Instances
	}
	return nil
}

type Instance struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	InstanceType  string   `protobuf:"bytes,2,opt,name=instance_type,json=instanceType,proto3" json:"instance_type,omitempty"`
	Provider      string   `protobuf:"bytes,3,opt,name=provider,proto3" json:"provider,omitempty"`
	Region        string   `protobuf:"bytes,4,opt,name=region,proto3" json:"region,omitempty"`
	DataCenter    string   `protobuf:"bytes,5,opt,name=data_center,json=dataCenter,proto3" json:"data_center,omitempty"`
	NetworkLayers []string `protobuf:"bytes,6,rep,name=network_layers,json=networkLayers,proto3" json:"network_layers,omitempty"`
	NvlinkDomain  string   `protobuf:"bytes,7,opt,name=nvlink_domain,json=nvlinkDomain,proto3" json:"nvlink_domain,omitempty"`
}

func (x *Instance) Reset() {
	*x = Instance{}
	if protoimpl.UnsafeEnabled {
		mi := &file_topology_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Instance) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Instance) ProtoMessage() {}

func (x *Instance) ProtoReflect() protoreflect.Message {
	mi := &file_topology_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Instance.ProtoReflect.Descriptor instead.
func (*Instance) Descriptor() ([]byte, []int) {
	return file_topology_proto_rawDescGZIP(), []int{2}
}

func (x *Instance) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Instance) GetInstanceType() string {
	if x != nil {
		return x.InstanceType
	}
	return ""
}

func (x *Instance) GetProvider() string {
	if x != nil {
		return x.Provider
	}
	return ""
}

func (x *Instance) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *Instance) GetDataCenter() string {
	if x != nil {
		return x.DataCenter
	}
	return ""
}

func (x *Instance) GetNetworkLayers() []string {
	if x != nil {
		return x.NetworkLayers
	}
	return nil
}

func (x *Instance) GetNvlinkDomain() string {
	if x != nil {
		return x.NvlinkDomain
	}
	return ""
}

var File_topology_proto protoreflect.FileDescriptor

var file_topology_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x74, 0x6f, 0x70, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x74, 0x6f, 0x70, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x22, 0x68, 0x0a, 0x0f, 0x54, 0x6f,
	0x70, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a,
	0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x67,
	0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f,
	0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x69, 0x64,
	0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63,
	0x65, 0x49, 0x64, 0x73, 0x22, 0x44, 0x0a, 0x10, 0x54, 0x6f, 0x70, 0x6f, 0x6c, 0x6f, 0x67, 0x79,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x09, 0x69, 0x6e, 0x73, 0x74,
	0x61, 0x6e, 0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x74, 0x6f,
	0x70, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x2e, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x52,
	0x09, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x22, 0xe0, 0x01, 0x0a, 0x08, 0x49,
	0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x69, 0x6e, 0x73, 0x74, 0x61,
	0x6e, 0x63, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x67, 0x69,
	0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e,
	0x12, 0x1f, 0x0a, 0x0b, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x61, 0x74, 0x61, 0x43, 0x65, 0x6e, 0x74, 0x65,
	0x72, 0x12, 0x25, 0x0a, 0x0e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0d, 0x6e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x4c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x6e, 0x76, 0x6c, 0x69,
	0x6e, 0x6b, 0x5f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x6e, 0x76, 0x6c, 0x69, 0x6e, 0x6b, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x32, 0x5e, 0x0a,
	0x0f, 0x54, 0x6f, 0x70, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x4b, 0x0a, 0x10, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x54, 0x6f, 0x70, 0x6f,
	0x6c, 0x6f, 0x67, 0x79, 0x12, 0x19, 0x2e, 0x74, 0x6f, 0x70, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x2e,
	0x54, 0x6f, 0x70, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1a, 0x2e, 0x74, 0x6f, 0x70, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x2e, 0x54, 0x6f, 0x70, 0x6f, 0x6c,
	0x6f, 0x67, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0b, 0x5a,
	0x09, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_topology_proto_rawDescOnce sync.Once
	file_topology_proto_rawDescData = file_topology_proto_rawDesc
)

func file_topology_proto_rawDescGZIP() []byte {
	file_topology_proto_rawDescOnce.Do(func() {
		file_topology_proto_rawDescData = protoimpl.X.CompressGZIP(file_topology_proto_rawDescData)
	})
	return file_topology_proto_rawDescData
}

var file_topology_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_topology_proto_goTypes = []any{
	(*TopologyRequest)(nil),  // 0: topology.TopologyRequest
	(*TopologyResponse)(nil), // 1: topology.TopologyResponse
	(*Instance)(nil),         // 2: topology.Instance
}
var file_topology_proto_depIdxs = []int32{
	2, // 0: topology.TopologyResponse.instances:type_name -> topology.Instance
	0, // 1: topology.TopologyService.DescribeTopology:input_type -> topology.TopologyRequest
	1, // 2: topology.TopologyService.DescribeTopology:output_type -> topology.TopologyResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_topology_proto_init() }
func file_topology_proto_init() {
	if File_topology_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_topology_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*TopologyRequest); i {
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
		file_topology_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*TopologyResponse); i {
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
		file_topology_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*Instance); i {
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
			RawDescriptor: file_topology_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_topology_proto_goTypes,
		DependencyIndexes: file_topology_proto_depIdxs,
		MessageInfos:      file_topology_proto_msgTypes,
	}.Build()
	File_topology_proto = out.File
	file_topology_proto_rawDesc = nil
	file_topology_proto_goTypes = nil
	file_topology_proto_depIdxs = nil
}