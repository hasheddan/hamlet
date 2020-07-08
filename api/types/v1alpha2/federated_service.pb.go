// Code generated by protoc-gen-go. DO NOT EDIT.
// source: federated_service.proto

package v1alpha2

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// FederatedService represents a federated service which can be discovered by
// a federated service mesh consumer.
type FederatedService struct {
	// REQUIRED. A human readable name for the federated service.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The description of the federated service.
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	// Informative values for filtering purposes.
	Tags []string `protobuf:"bytes,3,rep,name=tags,proto3" json:"tags,omitempty"`
	// Informative KV pairs for filtering purposes.
	Labels map[string]string `protobuf:"bytes,4,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// REQUIRED. The FQDN (Fully Qualified Domain Name) uniquely identifies the
	// federated service. The FQDN is used by consumer platforms as a DNS name to
	// access the imported FederatedService. The consumer platform DNS can be
	// programmed to reference the Ingress of the imported FederatedService.
	Fqdn string `protobuf:"bytes,5,opt,name=fqdn,proto3" json:"fqdn,omitempty"`
	// REQUIRED. The unique name that identifies the service within a federated
	// service mesh owner. The value of this field will be set as the SNI header
	// by the federated service mesh consumer. Each vendor may possibly have its
	// own SNI format, so this specification doesn't define a particular format
	// to use for this field.
	//
	// Following are some sample values for this field.
	// Example: foo.acme.com
	// Example: outbound_.8080_.v1_.foo.acme.com
	// Example: foo.acme.64a95d14-92f8-11e9-bc42-526af7764f64.com
	Id string `protobuf:"bytes,6,opt,name=id,proto3" json:"id,omitempty"`
	// DEPRECATED. Usage of sans is recommended.
	//
	// The SAN (Subject Alternative Name) to enable end-to-end security for the
	// federated service.
	San string `protobuf:"bytes,7,opt,name=san,proto3" json:"san,omitempty"`
	// REQUIRED. A list of SANs (Subject Alternative Names) to enable end-to-end
	// security for the federated service.
	Sans string `protobuf:"bytes,10,opt,name=sans,proto3" json:"sans,omitempty"`
	// REQUIRED. The protocols supported by the federated service.
	Protocols []string `protobuf:"bytes,8,rep,name=protocols,proto3" json:"protocols,omitempty"`
	// REQUIRED. The service mesh ingress endpoints that provide access to the
	// federated service.
	Endpoints []*FederatedService_Endpoint `protobuf:"bytes,9,rep,name=endpoints,proto3" json:"endpoints,omitempty"`
	// REQUIRED. The service instances corresponding to the federated service.
	Instances            []*FederatedService_Instance `protobuf:"bytes,11,rep,name=instances,proto3" json:"instances,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *FederatedService) Reset()         { *m = FederatedService{} }
func (m *FederatedService) String() string { return proto.CompactTextString(m) }
func (*FederatedService) ProtoMessage()    {}
func (*FederatedService) Descriptor() ([]byte, []int) {
	return fileDescriptor_0ca57f12e11e216c, []int{0}
}

func (m *FederatedService) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FederatedService.Unmarshal(m, b)
}
func (m *FederatedService) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FederatedService.Marshal(b, m, deterministic)
}
func (m *FederatedService) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FederatedService.Merge(m, src)
}
func (m *FederatedService) XXX_Size() int {
	return xxx_messageInfo_FederatedService.Size(m)
}
func (m *FederatedService) XXX_DiscardUnknown() {
	xxx_messageInfo_FederatedService.DiscardUnknown(m)
}

var xxx_messageInfo_FederatedService proto.InternalMessageInfo

func (m *FederatedService) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *FederatedService) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *FederatedService) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *FederatedService) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *FederatedService) GetFqdn() string {
	if m != nil {
		return m.Fqdn
	}
	return ""
}

func (m *FederatedService) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *FederatedService) GetSan() string {
	if m != nil {
		return m.San
	}
	return ""
}

func (m *FederatedService) GetSans() string {
	if m != nil {
		return m.Sans
	}
	return ""
}

func (m *FederatedService) GetProtocols() []string {
	if m != nil {
		return m.Protocols
	}
	return nil
}

func (m *FederatedService) GetEndpoints() []*FederatedService_Endpoint {
	if m != nil {
		return m.Endpoints
	}
	return nil
}

func (m *FederatedService) GetInstances() []*FederatedService_Instance {
	if m != nil {
		return m.Instances
	}
	return nil
}

// Endpoint represents a service mesh ingress endpoint that provides access to
// a federated service over mTLS (mutual TLS).
type FederatedService_Endpoint struct {
	// REQUIRED. The service mesh ingress endpoint address (i.e. IP, CIDR, or a
	// domain name).
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	// REQUIRED. The port exposed by the service mesh ingress endpoint.
	Port uint32 `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
	// REQUIRED. The endpoint labels decouple IP from routing information. The
	// endpoint_selector field in Instance matches with at least one endpoint
	// label here.
	Labels               []string `protobuf:"bytes,3,rep,name=labels,proto3" json:"labels,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FederatedService_Endpoint) Reset()         { *m = FederatedService_Endpoint{} }
func (m *FederatedService_Endpoint) String() string { return proto.CompactTextString(m) }
func (*FederatedService_Endpoint) ProtoMessage()    {}
func (*FederatedService_Endpoint) Descriptor() ([]byte, []int) {
	return fileDescriptor_0ca57f12e11e216c, []int{0, 1}
}

func (m *FederatedService_Endpoint) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FederatedService_Endpoint.Unmarshal(m, b)
}
func (m *FederatedService_Endpoint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FederatedService_Endpoint.Marshal(b, m, deterministic)
}
func (m *FederatedService_Endpoint) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FederatedService_Endpoint.Merge(m, src)
}
func (m *FederatedService_Endpoint) XXX_Size() int {
	return xxx_messageInfo_FederatedService_Endpoint.Size(m)
}
func (m *FederatedService_Endpoint) XXX_DiscardUnknown() {
	xxx_messageInfo_FederatedService_Endpoint.DiscardUnknown(m)
}

var xxx_messageInfo_FederatedService_Endpoint proto.InternalMessageInfo

func (m *FederatedService_Endpoint) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *FederatedService_Endpoint) GetPort() uint32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *FederatedService_Endpoint) GetLabels() []string {
	if m != nil {
		return m.Labels
	}
	return nil
}

// Instance represents an instance of the federated service.
type FederatedService_Instance struct {
	// REQUIRED. The unique identifier for the federated service instance within
	// the group of instances for a FederatedService.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// REQUIRED. The protocol corresponding to the federated service instance.
	//
	// E.g. In the case where protocol is TLS, the connection will be routed
	// based on the SNI header to the destination.
	Protocol string `protobuf:"bytes,2,opt,name=protocol,proto3" json:"protocol,omitempty"`
	// Additional metadata for the purpose of establishing connectivity.
	//
	// E.g. In the case where protocol is TLS, the connection will be routed
	// based on the SNI header to the destination. The SNI header can be made
	// available in the metadata using a key named 'sni'.
	Metadata map[string]string `protobuf:"bytes,3,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// REQUIRED. A list of endpoint labels to match over Endpoint instances in a
	// FederatedService. At least one matching label must be provided.
	EndpointSelector     []string `protobuf:"bytes,4,rep,name=endpoint_selector,json=endpointSelector,proto3" json:"endpoint_selector,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FederatedService_Instance) Reset()         { *m = FederatedService_Instance{} }
func (m *FederatedService_Instance) String() string { return proto.CompactTextString(m) }
func (*FederatedService_Instance) ProtoMessage()    {}
func (*FederatedService_Instance) Descriptor() ([]byte, []int) {
	return fileDescriptor_0ca57f12e11e216c, []int{0, 2}
}

func (m *FederatedService_Instance) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FederatedService_Instance.Unmarshal(m, b)
}
func (m *FederatedService_Instance) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FederatedService_Instance.Marshal(b, m, deterministic)
}
func (m *FederatedService_Instance) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FederatedService_Instance.Merge(m, src)
}
func (m *FederatedService_Instance) XXX_Size() int {
	return xxx_messageInfo_FederatedService_Instance.Size(m)
}
func (m *FederatedService_Instance) XXX_DiscardUnknown() {
	xxx_messageInfo_FederatedService_Instance.DiscardUnknown(m)
}

var xxx_messageInfo_FederatedService_Instance proto.InternalMessageInfo

func (m *FederatedService_Instance) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *FederatedService_Instance) GetProtocol() string {
	if m != nil {
		return m.Protocol
	}
	return ""
}

func (m *FederatedService_Instance) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *FederatedService_Instance) GetEndpointSelector() []string {
	if m != nil {
		return m.EndpointSelector
	}
	return nil
}

func init() {
	proto.RegisterType((*FederatedService)(nil), "federation.types.v1alpha2.FederatedService")
	proto.RegisterMapType((map[string]string)(nil), "federation.types.v1alpha2.FederatedService.LabelsEntry")
	proto.RegisterType((*FederatedService_Endpoint)(nil), "federation.types.v1alpha2.FederatedService.Endpoint")
	proto.RegisterType((*FederatedService_Instance)(nil), "federation.types.v1alpha2.FederatedService.Instance")
	proto.RegisterMapType((map[string]string)(nil), "federation.types.v1alpha2.FederatedService.Instance.MetadataEntry")
}

func init() { proto.RegisterFile("federated_service.proto", fileDescriptor_0ca57f12e11e216c) }

var fileDescriptor_0ca57f12e11e216c = []byte{
	// 421 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x52, 0x4d, 0xab, 0xd3, 0x40,
	0x14, 0x25, 0xc9, 0x7b, 0x79, 0xc9, 0x0d, 0xef, 0x11, 0x07, 0xd1, 0x31, 0xb8, 0x08, 0xae, 0x0a,
	0x42, 0xc4, 0x2a, 0xf8, 0xb5, 0x2b, 0x54, 0x10, 0x14, 0x25, 0xdd, 0xb9, 0xb0, 0x4c, 0x33, 0xb7,
	0x1a, 0x4c, 0x27, 0x71, 0x66, 0x2c, 0x74, 0xe7, 0x1f, 0xf4, 0x3f, 0xc9, 0x4c, 0x66, 0xd2, 0x2a,
	0xb8, 0xe8, 0xdb, 0x9d, 0x7b, 0x2f, 0xe7, 0xe4, 0xe4, 0x9c, 0x81, 0xfb, 0x5b, 0xe4, 0x28, 0x99,
	0x46, 0xbe, 0x56, 0x28, 0xf7, 0x6d, 0x83, 0xd5, 0x20, 0x7b, 0xdd, 0x93, 0x07, 0xee, 0xd0, 0xf6,
	0xa2, 0xd2, 0x87, 0x01, 0x55, 0xb5, 0x7f, 0xca, 0xba, 0xe1, 0x1b, 0x9b, 0x3f, 0xfa, 0x1d, 0x43,
	0xfe, 0xd6, 0xd3, 0x56, 0x23, 0x8b, 0x10, 0xb8, 0x10, 0x6c, 0x87, 0x34, 0x28, 0x83, 0x59, 0x5a,
	0x5b, 0x4c, 0x4a, 0xc8, 0x38, 0xaa, 0x46, 0xb6, 0x83, 0x91, 0xa1, 0xa1, 0x3d, 0x9d, 0xae, 0x0c,
	0x4b, 0xb3, 0xaf, 0x8a, 0x46, 0x65, 0x64, 0x58, 0x06, 0x93, 0x8f, 0x10, 0x77, 0x6c, 0x83, 0x9d,
	0xa2, 0x17, 0x65, 0x34, 0xcb, 0xe6, 0x2f, 0xaa, 0xff, 0x5a, 0xa9, 0xfe, 0xb5, 0x51, 0xbd, 0xb7,
	0xcc, 0xa5, 0xd0, 0xf2, 0x50, 0x3b, 0x19, 0xf3, 0x91, 0xed, 0x0f, 0x2e, 0xe8, 0xe5, 0x68, 0xcd,
	0x60, 0x72, 0x03, 0x61, 0xcb, 0x69, 0x6c, 0x37, 0x61, 0xcb, 0x49, 0x0e, 0x91, 0x62, 0x82, 0x5e,
	0xd9, 0x85, 0x81, 0x86, 0xa5, 0x98, 0x50, 0x14, 0x46, 0x96, 0xc1, 0xe4, 0x21, 0xa4, 0x36, 0x9d,
	0xa6, 0xef, 0x14, 0x4d, 0xac, 0xe7, 0xe3, 0x82, 0xd4, 0x90, 0xa2, 0xe0, 0x43, 0xdf, 0x0a, 0xad,
	0x68, 0x6a, 0xbd, 0x3f, 0x3f, 0xc7, 0xfb, 0xd2, 0x91, 0xeb, 0xa3, 0x8c, 0xd1, 0x6c, 0x85, 0xd2,
	0x4c, 0x34, 0xa8, 0x68, 0x76, 0xbe, 0xe6, 0x3b, 0x47, 0xae, 0x8f, 0x32, 0xc5, 0x2b, 0xc8, 0x4e,
	0x62, 0x32, 0xbf, 0xfe, 0x1d, 0x0f, 0xae, 0x38, 0x03, 0xc9, 0x5d, 0xb8, 0xdc, 0xb3, 0xee, 0x27,
	0xba, 0xc6, 0xc6, 0xe1, 0x75, 0xf8, 0x32, 0x28, 0x3e, 0x41, 0xe2, 0x5d, 0x12, 0x0a, 0x57, 0x8c,
	0x73, 0x89, 0x4a, 0x39, 0xae, 0x1f, 0x4d, 0x74, 0x43, 0x2f, 0xb5, 0xa5, 0x5f, 0xd7, 0x16, 0x93,
	0x7b, 0x53, 0xab, 0x63, 0xd7, 0x6e, 0x2a, 0x7e, 0x85, 0x90, 0x78, 0x93, 0xae, 0x95, 0x60, 0x6a,
	0xa5, 0x80, 0xc4, 0xc7, 0xeb, 0xbc, 0x4c, 0x33, 0xf9, 0x02, 0xc9, 0x0e, 0x35, 0xe3, 0x4c, 0x33,
	0x2b, 0x99, 0xcd, 0x17, 0xb7, 0x09, 0xa6, 0xfa, 0xe0, 0x44, 0xc6, 0x37, 0x33, 0x69, 0x92, 0xc7,
	0x70, 0xc7, 0xd7, 0xb0, 0x56, 0xd8, 0x61, 0xa3, 0x7b, 0x69, 0x5f, 0x64, 0x5a, 0xe7, 0xfe, 0xb0,
	0x72, 0xfb, 0xe2, 0x0d, 0x5c, 0xff, 0xa5, 0x73, 0x4e, 0xa8, 0x8b, 0xfc, 0xf3, 0x8d, 0x75, 0xfb,
	0xc4, 0xbb, 0xdd, 0xc4, 0xf6, 0x2f, 0x9f, 0xfd, 0x09, 0x00, 0x00, 0xff, 0xff, 0xc1, 0x1a, 0xe9,
	0x95, 0x9e, 0x03, 0x00, 0x00,
}