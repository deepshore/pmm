// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: managementpb/proxysql.proto

package managementpb

import (
	context "context"
	_ "github.com/mwitkow/go-proto-validators"
	inventorypb "github.com/percona/pmm/api/inventorypb"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type AddProxySQLRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Node identifier on which a service is been running.
	// Exactly one of these parameters should be present: node_id, node_name, add_node.
	NodeId string `protobuf:"bytes,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	// Node name on which a service is been running.
	// Exactly one of these parameters should be present: node_id, node_name, add_node.
	NodeName string `protobuf:"bytes,2,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	// Create a new Node with those parameters.
	// Exactly one of these parameters should be present: node_id, node_name, add_node.
	AddNode *AddNodeParams `protobuf:"bytes,3,opt,name=add_node,json=addNode,proto3" json:"add_node,omitempty"`
	// Unique across all Services user-defined name. Required.
	ServiceName string `protobuf:"bytes,4,opt,name=service_name,json=serviceName,proto3" json:"service_name,omitempty"`
	// Node and Service access address (DNS name or IP).
	// Address (and port) or socket is required.
	Address string `protobuf:"bytes,5,opt,name=address,proto3" json:"address,omitempty"`
	// Service Access port.
	// Port is required when the address present.
	Port uint32 `protobuf:"varint,6,opt,name=port,proto3" json:"port,omitempty"`
	// Service Access socket.
	// Address (and port) or socket is required.
	Socket string `protobuf:"bytes,17,opt,name=socket,proto3" json:"socket,omitempty"`
	// The "pmm-agent" identifier which should run agents. Required.
	PmmAgentId string `protobuf:"bytes,7,opt,name=pmm_agent_id,json=pmmAgentId,proto3" json:"pmm_agent_id,omitempty"`
	// Environment name.
	Environment string `protobuf:"bytes,8,opt,name=environment,proto3" json:"environment,omitempty"`
	// Cluster name.
	Cluster string `protobuf:"bytes,9,opt,name=cluster,proto3" json:"cluster,omitempty"`
	// Replication set name.
	ReplicationSet string `protobuf:"bytes,10,opt,name=replication_set,json=replicationSet,proto3" json:"replication_set,omitempty"`
	// ProxySQL username for scraping metrics.
	Username string `protobuf:"bytes,11,opt,name=username,proto3" json:"username,omitempty"`
	// ProxySQL password for scraping metrics.
	Password string `protobuf:"bytes,12,opt,name=password,proto3" json:"password,omitempty"`
	// Custom user-assigned labels for Service.
	CustomLabels map[string]string `protobuf:"bytes,13,rep,name=custom_labels,json=customLabels,proto3" json:"custom_labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Skip connection check.
	SkipConnectionCheck bool `protobuf:"varint,14,opt,name=skip_connection_check,json=skipConnectionCheck,proto3" json:"skip_connection_check,omitempty"`
	// Use TLS for database connections.
	Tls bool `protobuf:"varint,15,opt,name=tls,proto3" json:"tls,omitempty"`
	// Skip TLS certificate and hostname validation.
	TlsSkipVerify bool `protobuf:"varint,16,opt,name=tls_skip_verify,json=tlsSkipVerify,proto3" json:"tls_skip_verify,omitempty"`
}

func (x *AddProxySQLRequest) Reset() {
	*x = AddProxySQLRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_managementpb_proxysql_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddProxySQLRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddProxySQLRequest) ProtoMessage() {}

func (x *AddProxySQLRequest) ProtoReflect() protoreflect.Message {
	mi := &file_managementpb_proxysql_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddProxySQLRequest.ProtoReflect.Descriptor instead.
func (*AddProxySQLRequest) Descriptor() ([]byte, []int) {
	return file_managementpb_proxysql_proto_rawDescGZIP(), []int{0}
}

func (x *AddProxySQLRequest) GetNodeId() string {
	if x != nil {
		return x.NodeId
	}
	return ""
}

func (x *AddProxySQLRequest) GetNodeName() string {
	if x != nil {
		return x.NodeName
	}
	return ""
}

func (x *AddProxySQLRequest) GetAddNode() *AddNodeParams {
	if x != nil {
		return x.AddNode
	}
	return nil
}

func (x *AddProxySQLRequest) GetServiceName() string {
	if x != nil {
		return x.ServiceName
	}
	return ""
}

func (x *AddProxySQLRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *AddProxySQLRequest) GetPort() uint32 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *AddProxySQLRequest) GetSocket() string {
	if x != nil {
		return x.Socket
	}
	return ""
}

func (x *AddProxySQLRequest) GetPmmAgentId() string {
	if x != nil {
		return x.PmmAgentId
	}
	return ""
}

func (x *AddProxySQLRequest) GetEnvironment() string {
	if x != nil {
		return x.Environment
	}
	return ""
}

func (x *AddProxySQLRequest) GetCluster() string {
	if x != nil {
		return x.Cluster
	}
	return ""
}

func (x *AddProxySQLRequest) GetReplicationSet() string {
	if x != nil {
		return x.ReplicationSet
	}
	return ""
}

func (x *AddProxySQLRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *AddProxySQLRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *AddProxySQLRequest) GetCustomLabels() map[string]string {
	if x != nil {
		return x.CustomLabels
	}
	return nil
}

func (x *AddProxySQLRequest) GetSkipConnectionCheck() bool {
	if x != nil {
		return x.SkipConnectionCheck
	}
	return false
}

func (x *AddProxySQLRequest) GetTls() bool {
	if x != nil {
		return x.Tls
	}
	return false
}

func (x *AddProxySQLRequest) GetTlsSkipVerify() bool {
	if x != nil {
		return x.TlsSkipVerify
	}
	return false
}

type AddProxySQLResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Service          *inventorypb.ProxySQLService  `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	ProxysqlExporter *inventorypb.ProxySQLExporter `protobuf:"bytes,2,opt,name=proxysql_exporter,json=proxysqlExporter,proto3" json:"proxysql_exporter,omitempty"`
}

func (x *AddProxySQLResponse) Reset() {
	*x = AddProxySQLResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_managementpb_proxysql_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddProxySQLResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddProxySQLResponse) ProtoMessage() {}

func (x *AddProxySQLResponse) ProtoReflect() protoreflect.Message {
	mi := &file_managementpb_proxysql_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddProxySQLResponse.ProtoReflect.Descriptor instead.
func (*AddProxySQLResponse) Descriptor() ([]byte, []int) {
	return file_managementpb_proxysql_proto_rawDescGZIP(), []int{1}
}

func (x *AddProxySQLResponse) GetService() *inventorypb.ProxySQLService {
	if x != nil {
		return x.Service
	}
	return nil
}

func (x *AddProxySQLResponse) GetProxysqlExporter() *inventorypb.ProxySQLExporter {
	if x != nil {
		return x.ProxysqlExporter
	}
	return nil
}

var File_managementpb_proxysql_proto protoreflect.FileDescriptor

var file_managementpb_proxysql_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x2f, 0x70,
	0x72, 0x6f, 0x78, 0x79, 0x73, 0x71, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x1a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x77, 0x69, 0x74, 0x6b, 0x6f, 0x77, 0x2f, 0x67, 0x6f,
	0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72,
	0x73, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x18, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x70, 0x62, 0x2f, 0x61, 0x67, 0x65,
	0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x69, 0x6e, 0x76, 0x65, 0x6e,
	0x74, 0x6f, 0x72, 0x79, 0x70, 0x62, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x70, 0x62, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xc6, 0x05, 0x0a, 0x12, 0x41, 0x64, 0x64, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x53, 0x51,
	0x4c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x6e, 0x6f, 0x64, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x6f, 0x64, 0x65, 0x49,
	0x64, 0x12, 0x1b, 0x0a, 0x09, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x6f, 0x64, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x34,
	0x0a, 0x08, 0x61, 0x64, 0x64, 0x5f, 0x6e, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x19, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x41, 0x64,
	0x64, 0x4e, 0x6f, 0x64, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x52, 0x07, 0x61, 0x64, 0x64,
	0x4e, 0x6f, 0x64, 0x65, 0x12, 0x29, 0x0a, 0x0c, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xe2, 0xdf, 0x1f, 0x02,
	0x58, 0x01, 0x52, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72,
	0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73,
	0x6f, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x28, 0x0a, 0x0c, 0x70, 0x6d, 0x6d, 0x5f, 0x61, 0x67, 0x65,
	0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xe2, 0xdf, 0x1f,
	0x02, 0x58, 0x01, 0x52, 0x0a, 0x70, 0x6d, 0x6d, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12,
	0x20, 0x0a, 0x0b, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e,
	0x74, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x12, 0x27, 0x0a, 0x0f, 0x72,
	0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x74, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x53, 0x65, 0x74, 0x12, 0x22, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xe2, 0xdf, 0x1f, 0x02, 0x58, 0x01, 0x52, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x12, 0x55, 0x0a, 0x0d, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x6c,
	0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x0d, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x6d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x41, 0x64, 0x64, 0x50, 0x72, 0x6f, 0x78,
	0x79, 0x53, 0x51, 0x4c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x43, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0c, 0x63,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x32, 0x0a, 0x15, 0x73,
	0x6b, 0x69, 0x70, 0x5f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63,
	0x68, 0x65, 0x63, 0x6b, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x08, 0x52, 0x13, 0x73, 0x6b, 0x69, 0x70,
	0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x12,
	0x10, 0x0a, 0x03, 0x74, 0x6c, 0x73, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x08, 0x52, 0x03, 0x74, 0x6c,
	0x73, 0x12, 0x26, 0x0a, 0x0f, 0x74, 0x6c, 0x73, 0x5f, 0x73, 0x6b, 0x69, 0x70, 0x5f, 0x76, 0x65,
	0x72, 0x69, 0x66, 0x79, 0x18, 0x10, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x74, 0x6c, 0x73, 0x53,
	0x6b, 0x69, 0x70, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x1a, 0x3f, 0x0a, 0x11, 0x43, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x95, 0x01, 0x0a, 0x13, 0x41,
	0x64, 0x64, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x53, 0x51, 0x4c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x34, 0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e,
	0x50, 0x72, 0x6f, 0x78, 0x79, 0x53, 0x51, 0x4c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52,
	0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x48, 0x0a, 0x11, 0x70, 0x72, 0x6f, 0x78,
	0x79, 0x73, 0x71, 0x6c, 0x5f, 0x65, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e,
	0x50, 0x72, 0x6f, 0x78, 0x79, 0x53, 0x51, 0x4c, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72,
	0x52, 0x10, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x73, 0x71, 0x6c, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74,
	0x65, 0x72, 0x32, 0x82, 0x01, 0x0a, 0x08, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x53, 0x51, 0x4c, 0x12,
	0x76, 0x0a, 0x0b, 0x41, 0x64, 0x64, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x53, 0x51, 0x4c, 0x12, 0x1e,
	0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x41, 0x64, 0x64, 0x50,
	0x72, 0x6f, 0x78, 0x79, 0x53, 0x51, 0x4c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f,
	0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x41, 0x64, 0x64, 0x50,
	0x72, 0x6f, 0x78, 0x79, 0x53, 0x51, 0x4c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x26, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x20, 0x22, 0x1b, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x53, 0x51, 0x4c,
	0x2f, 0x41, 0x64, 0x64, 0x3a, 0x01, 0x2a, 0x42, 0x1f, 0x5a, 0x1d, 0x61, 0x70, 0x69, 0x2f, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x3b, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_managementpb_proxysql_proto_rawDescOnce sync.Once
	file_managementpb_proxysql_proto_rawDescData = file_managementpb_proxysql_proto_rawDesc
)

func file_managementpb_proxysql_proto_rawDescGZIP() []byte {
	file_managementpb_proxysql_proto_rawDescOnce.Do(func() {
		file_managementpb_proxysql_proto_rawDescData = protoimpl.X.CompressGZIP(file_managementpb_proxysql_proto_rawDescData)
	})
	return file_managementpb_proxysql_proto_rawDescData
}

var file_managementpb_proxysql_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_managementpb_proxysql_proto_goTypes = []interface{}{
	(*AddProxySQLRequest)(nil),           // 0: management.AddProxySQLRequest
	(*AddProxySQLResponse)(nil),          // 1: management.AddProxySQLResponse
	nil,                                  // 2: management.AddProxySQLRequest.CustomLabelsEntry
	(*AddNodeParams)(nil),                // 3: management.AddNodeParams
	(*inventorypb.ProxySQLService)(nil),  // 4: inventory.ProxySQLService
	(*inventorypb.ProxySQLExporter)(nil), // 5: inventory.ProxySQLExporter
}
var file_managementpb_proxysql_proto_depIdxs = []int32{
	3, // 0: management.AddProxySQLRequest.add_node:type_name -> management.AddNodeParams
	2, // 1: management.AddProxySQLRequest.custom_labels:type_name -> management.AddProxySQLRequest.CustomLabelsEntry
	4, // 2: management.AddProxySQLResponse.service:type_name -> inventory.ProxySQLService
	5, // 3: management.AddProxySQLResponse.proxysql_exporter:type_name -> inventory.ProxySQLExporter
	0, // 4: management.ProxySQL.AddProxySQL:input_type -> management.AddProxySQLRequest
	1, // 5: management.ProxySQL.AddProxySQL:output_type -> management.AddProxySQLResponse
	5, // [5:6] is the sub-list for method output_type
	4, // [4:5] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_managementpb_proxysql_proto_init() }
func file_managementpb_proxysql_proto_init() {
	if File_managementpb_proxysql_proto != nil {
		return
	}
	file_managementpb_service_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_managementpb_proxysql_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddProxySQLRequest); i {
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
		file_managementpb_proxysql_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddProxySQLResponse); i {
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
			RawDescriptor: file_managementpb_proxysql_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_managementpb_proxysql_proto_goTypes,
		DependencyIndexes: file_managementpb_proxysql_proto_depIdxs,
		MessageInfos:      file_managementpb_proxysql_proto_msgTypes,
	}.Build()
	File_managementpb_proxysql_proto = out.File
	file_managementpb_proxysql_proto_rawDesc = nil
	file_managementpb_proxysql_proto_goTypes = nil
	file_managementpb_proxysql_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ProxySQLClient is the client API for ProxySQL service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ProxySQLClient interface {
	// AddProxySQL adds ProxySQL Service and starts several Agents.
	// It automatically adds a service to inventory, which is running on provided "node_id",
	// then adds "proxysql_exporter" with provided "pmm_agent_id" and other parameters.
	AddProxySQL(ctx context.Context, in *AddProxySQLRequest, opts ...grpc.CallOption) (*AddProxySQLResponse, error)
}

type proxySQLClient struct {
	cc grpc.ClientConnInterface
}

func NewProxySQLClient(cc grpc.ClientConnInterface) ProxySQLClient {
	return &proxySQLClient{cc}
}

func (c *proxySQLClient) AddProxySQL(ctx context.Context, in *AddProxySQLRequest, opts ...grpc.CallOption) (*AddProxySQLResponse, error) {
	out := new(AddProxySQLResponse)
	err := c.cc.Invoke(ctx, "/management.ProxySQL/AddProxySQL", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProxySQLServer is the server API for ProxySQL service.
type ProxySQLServer interface {
	// AddProxySQL adds ProxySQL Service and starts several Agents.
	// It automatically adds a service to inventory, which is running on provided "node_id",
	// then adds "proxysql_exporter" with provided "pmm_agent_id" and other parameters.
	AddProxySQL(context.Context, *AddProxySQLRequest) (*AddProxySQLResponse, error)
}

// UnimplementedProxySQLServer can be embedded to have forward compatible implementations.
type UnimplementedProxySQLServer struct {
}

func (*UnimplementedProxySQLServer) AddProxySQL(context.Context, *AddProxySQLRequest) (*AddProxySQLResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddProxySQL not implemented")
}

func RegisterProxySQLServer(s *grpc.Server, srv ProxySQLServer) {
	s.RegisterService(&_ProxySQL_serviceDesc, srv)
}

func _ProxySQL_AddProxySQL_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddProxySQLRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProxySQLServer).AddProxySQL(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/management.ProxySQL/AddProxySQL",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProxySQLServer).AddProxySQL(ctx, req.(*AddProxySQLRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ProxySQL_serviceDesc = grpc.ServiceDesc{
	ServiceName: "management.ProxySQL",
	HandlerType: (*ProxySQLServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddProxySQL",
			Handler:    _ProxySQL_AddProxySQL_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "managementpb/proxysql.proto",
}
