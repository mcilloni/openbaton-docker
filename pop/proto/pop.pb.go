// Code generated by protoc-gen-go.
// source: pop.proto
// DO NOT EDIT!

/*
Package pop is a generated protocol buffer package.

It is generated from these files:
	pop.proto

It has these top-level messages:
	Container
	ContainerList
	Credentials
	Endpoint
	Filter
	Image
	ImageList
	Infos
	Ip
	Network
	NetworkList
	Subnet
	Token
*/
package pop

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/empty"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Container struct {
	Id             string               `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Names          []string             `protobuf:"bytes,2,rep,name=names" json:"names,omitempty"`
	ImageId        string               `protobuf:"bytes,3,opt,name=image_id,json=imageId" json:"image_id,omitempty"`
	Command        string               `protobuf:"bytes,4,opt,name=command" json:"command,omitempty"`
	Created        int64                `protobuf:"varint,5,opt,name=created" json:"created,omitempty"`
	Status         string               `protobuf:"bytes,6,opt,name=status" json:"status,omitempty"`
	ExtendedStatus string               `protobuf:"bytes,7,opt,name=extended_status,json=extendedStatus" json:"extended_status,omitempty"`
	Endpoints      map[string]*Endpoint `protobuf:"bytes,8,rep,name=endpoints" json:"endpoints,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *Container) Reset()                    { *m = Container{} }
func (m *Container) String() string            { return proto.CompactTextString(m) }
func (*Container) ProtoMessage()               {}
func (*Container) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Container) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Container) GetNames() []string {
	if m != nil {
		return m.Names
	}
	return nil
}

func (m *Container) GetImageId() string {
	if m != nil {
		return m.ImageId
	}
	return ""
}

func (m *Container) GetCommand() string {
	if m != nil {
		return m.Command
	}
	return ""
}

func (m *Container) GetCreated() int64 {
	if m != nil {
		return m.Created
	}
	return 0
}

func (m *Container) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *Container) GetExtendedStatus() string {
	if m != nil {
		return m.ExtendedStatus
	}
	return ""
}

func (m *Container) GetEndpoints() map[string]*Endpoint {
	if m != nil {
		return m.Endpoints
	}
	return nil
}

type ContainerList struct {
	List []*Container `protobuf:"bytes,1,rep,name=list" json:"list,omitempty"`
}

func (m *ContainerList) Reset()                    { *m = ContainerList{} }
func (m *ContainerList) String() string            { return proto.CompactTextString(m) }
func (*ContainerList) ProtoMessage()               {}
func (*ContainerList) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ContainerList) GetList() []*Container {
	if m != nil {
		return m.List
	}
	return nil
}

// Credentials represents the login credentials for a given user.
type Credentials struct {
	Username string `protobuf:"bytes,1,opt,name=username" json:"username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password" json:"password,omitempty"`
}

func (m *Credentials) Reset()                    { *m = Credentials{} }
func (m *Credentials) String() string            { return proto.CompactTextString(m) }
func (*Credentials) ProtoMessage()               {}
func (*Credentials) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Credentials) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *Credentials) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type Endpoint struct {
	NetId      string `protobuf:"bytes,1,opt,name=net_id,json=netId" json:"net_id,omitempty"`
	EndpointId string `protobuf:"bytes,2,opt,name=endpoint_id,json=endpointId" json:"endpoint_id,omitempty"`
	Ipv4       *Ip    `protobuf:"bytes,3,opt,name=ipv4" json:"ipv4,omitempty"`
	Ipv6       *Ip    `protobuf:"bytes,4,opt,name=ipv6" json:"ipv6,omitempty"`
	Mac        string `protobuf:"bytes,5,opt,name=mac" json:"mac,omitempty"`
}

func (m *Endpoint) Reset()                    { *m = Endpoint{} }
func (m *Endpoint) String() string            { return proto.CompactTextString(m) }
func (*Endpoint) ProtoMessage()               {}
func (*Endpoint) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Endpoint) GetNetId() string {
	if m != nil {
		return m.NetId
	}
	return ""
}

func (m *Endpoint) GetEndpointId() string {
	if m != nil {
		return m.EndpointId
	}
	return ""
}

func (m *Endpoint) GetIpv4() *Ip {
	if m != nil {
		return m.Ipv4
	}
	return nil
}

func (m *Endpoint) GetIpv6() *Ip {
	if m != nil {
		return m.Ipv6
	}
	return nil
}

func (m *Endpoint) GetMac() string {
	if m != nil {
		return m.Mac
	}
	return ""
}

type Filter struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *Filter) Reset()                    { *m = Filter{} }
func (m *Filter) String() string            { return proto.CompactTextString(m) }
func (*Filter) ProtoMessage()               {}
func (*Filter) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Filter) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type Image struct {
	Id      string   `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Names   []string `protobuf:"bytes,2,rep,name=names" json:"names,omitempty"`
	Created int64    `protobuf:"varint,3,opt,name=created" json:"created,omitempty"`
}

func (m *Image) Reset()                    { *m = Image{} }
func (m *Image) String() string            { return proto.CompactTextString(m) }
func (*Image) ProtoMessage()               {}
func (*Image) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *Image) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Image) GetNames() []string {
	if m != nil {
		return m.Names
	}
	return nil
}

func (m *Image) GetCreated() int64 {
	if m != nil {
		return m.Created
	}
	return 0
}

type ImageList struct {
	List []*Image `protobuf:"bytes,1,rep,name=list" json:"list,omitempty"`
}

func (m *ImageList) Reset()                    { *m = ImageList{} }
func (m *ImageList) String() string            { return proto.CompactTextString(m) }
func (*ImageList) ProtoMessage()               {}
func (*ImageList) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *ImageList) GetList() []*Image {
	if m != nil {
		return m.List
	}
	return nil
}

type Infos struct {
	Type      string `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Timestamp int64  `protobuf:"varint,3,opt,name=timestamp" json:"timestamp,omitempty"`
}

func (m *Infos) Reset()                    { *m = Infos{} }
func (m *Infos) String() string            { return proto.CompactTextString(m) }
func (*Infos) ProtoMessage()               {}
func (*Infos) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *Infos) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Infos) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Infos) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

type Ip struct {
	Address string  `protobuf:"bytes,1,opt,name=address" json:"address,omitempty"`
	Subnet  *Subnet `protobuf:"bytes,2,opt,name=subnet" json:"subnet,omitempty"`
}

func (m *Ip) Reset()                    { *m = Ip{} }
func (m *Ip) String() string            { return proto.CompactTextString(m) }
func (*Ip) ProtoMessage()               {}
func (*Ip) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *Ip) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Ip) GetSubnet() *Subnet {
	if m != nil {
		return m.Subnet
	}
	return nil
}

type Network struct {
	Id       string    `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Name     string    `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	External bool      `protobuf:"varint,3,opt,name=external" json:"external,omitempty"`
	Subnets  []*Subnet `protobuf:"bytes,4,rep,name=subnets" json:"subnets,omitempty"`
}

func (m *Network) Reset()                    { *m = Network{} }
func (m *Network) String() string            { return proto.CompactTextString(m) }
func (*Network) ProtoMessage()               {}
func (*Network) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *Network) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Network) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Network) GetExternal() bool {
	if m != nil {
		return m.External
	}
	return false
}

func (m *Network) GetSubnets() []*Subnet {
	if m != nil {
		return m.Subnets
	}
	return nil
}

type NetworkList struct {
	List []*Network `protobuf:"bytes,1,rep,name=list" json:"list,omitempty"`
}

func (m *NetworkList) Reset()                    { *m = NetworkList{} }
func (m *NetworkList) String() string            { return proto.CompactTextString(m) }
func (*NetworkList) ProtoMessage()               {}
func (*NetworkList) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *NetworkList) GetList() []*Network {
	if m != nil {
		return m.List
	}
	return nil
}

type Subnet struct {
	Cidr    string `protobuf:"bytes,1,opt,name=cidr" json:"cidr,omitempty"`
	Gateway string `protobuf:"bytes,2,opt,name=gateway" json:"gateway,omitempty"`
}

func (m *Subnet) Reset()                    { *m = Subnet{} }
func (m *Subnet) String() string            { return proto.CompactTextString(m) }
func (*Subnet) ProtoMessage()               {}
func (*Subnet) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *Subnet) GetCidr() string {
	if m != nil {
		return m.Cidr
	}
	return ""
}

func (m *Subnet) GetGateway() string {
	if m != nil {
		return m.Gateway
	}
	return ""
}

// Token is a token generated by the server after a successful login.
// This token should be set as metadata, to authenticate every other
type Token struct {
	Value string `protobuf:"bytes,1,opt,name=value" json:"value,omitempty"`
}

func (m *Token) Reset()                    { *m = Token{} }
func (m *Token) String() string            { return proto.CompactTextString(m) }
func (*Token) ProtoMessage()               {}
func (*Token) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *Token) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func init() {
	proto.RegisterType((*Container)(nil), "pop.Container")
	proto.RegisterType((*ContainerList)(nil), "pop.ContainerList")
	proto.RegisterType((*Credentials)(nil), "pop.Credentials")
	proto.RegisterType((*Endpoint)(nil), "pop.Endpoint")
	proto.RegisterType((*Filter)(nil), "pop.Filter")
	proto.RegisterType((*Image)(nil), "pop.Image")
	proto.RegisterType((*ImageList)(nil), "pop.ImageList")
	proto.RegisterType((*Infos)(nil), "pop.Infos")
	proto.RegisterType((*Ip)(nil), "pop.Ip")
	proto.RegisterType((*Network)(nil), "pop.Network")
	proto.RegisterType((*NetworkList)(nil), "pop.NetworkList")
	proto.RegisterType((*Subnet)(nil), "pop.Subnet")
	proto.RegisterType((*Token)(nil), "pop.Token")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Pop service

type PopClient interface {
	// Containers returns the containers available in the PoP, either
	// created or running.
	Containers(ctx context.Context, in *Filter, opts ...grpc.CallOption) (*ContainerList, error)
	// Images returns the images available in the PoP.
	Images(ctx context.Context, in *Filter, opts ...grpc.CallOption) (*ImageList, error)
	// Info can be used to check if the Pop is alive and if your credentials to this service are valid.
	// It also returns informations about this server.
	Info(ctx context.Context, in *google_protobuf.Empty, opts ...grpc.CallOption) (*Infos, error)
	// Networks returns the available retworks in the PoP.
	Networks(ctx context.Context, in *Filter, opts ...grpc.CallOption) (*NetworkList, error)
	// Login logs an user in and sets up a session.
	// The returned token should be set into the metadata
	// of the gRPC session with key "token" to authenticate your client.
	Login(ctx context.Context, in *Credentials, opts ...grpc.CallOption) (*Token, error)
	// Logout invalids the current token.
	Logout(ctx context.Context, in *google_protobuf.Empty, opts ...grpc.CallOption) (*google_protobuf.Empty, error)
}

type popClient struct {
	cc *grpc.ClientConn
}

func NewPopClient(cc *grpc.ClientConn) PopClient {
	return &popClient{cc}
}

func (c *popClient) Containers(ctx context.Context, in *Filter, opts ...grpc.CallOption) (*ContainerList, error) {
	out := new(ContainerList)
	err := grpc.Invoke(ctx, "/pop.Pop/Containers", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *popClient) Images(ctx context.Context, in *Filter, opts ...grpc.CallOption) (*ImageList, error) {
	out := new(ImageList)
	err := grpc.Invoke(ctx, "/pop.Pop/Images", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *popClient) Info(ctx context.Context, in *google_protobuf.Empty, opts ...grpc.CallOption) (*Infos, error) {
	out := new(Infos)
	err := grpc.Invoke(ctx, "/pop.Pop/Info", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *popClient) Networks(ctx context.Context, in *Filter, opts ...grpc.CallOption) (*NetworkList, error) {
	out := new(NetworkList)
	err := grpc.Invoke(ctx, "/pop.Pop/Networks", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *popClient) Login(ctx context.Context, in *Credentials, opts ...grpc.CallOption) (*Token, error) {
	out := new(Token)
	err := grpc.Invoke(ctx, "/pop.Pop/Login", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *popClient) Logout(ctx context.Context, in *google_protobuf.Empty, opts ...grpc.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := grpc.Invoke(ctx, "/pop.Pop/Logout", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Pop service

type PopServer interface {
	// Containers returns the containers available in the PoP, either
	// created or running.
	Containers(context.Context, *Filter) (*ContainerList, error)
	// Images returns the images available in the PoP.
	Images(context.Context, *Filter) (*ImageList, error)
	// Info can be used to check if the Pop is alive and if your credentials to this service are valid.
	// It also returns informations about this server.
	Info(context.Context, *google_protobuf.Empty) (*Infos, error)
	// Networks returns the available retworks in the PoP.
	Networks(context.Context, *Filter) (*NetworkList, error)
	// Login logs an user in and sets up a session.
	// The returned token should be set into the metadata
	// of the gRPC session with key "token" to authenticate your client.
	Login(context.Context, *Credentials) (*Token, error)
	// Logout invalids the current token.
	Logout(context.Context, *google_protobuf.Empty) (*google_protobuf.Empty, error)
}

func RegisterPopServer(s *grpc.Server, srv PopServer) {
	s.RegisterService(&_Pop_serviceDesc, srv)
}

func _Pop_Containers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Filter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PopServer).Containers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pop.Pop/Containers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PopServer).Containers(ctx, req.(*Filter))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pop_Images_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Filter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PopServer).Images(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pop.Pop/Images",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PopServer).Images(ctx, req.(*Filter))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pop_Info_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(google_protobuf.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PopServer).Info(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pop.Pop/Info",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PopServer).Info(ctx, req.(*google_protobuf.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pop_Networks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Filter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PopServer).Networks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pop.Pop/Networks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PopServer).Networks(ctx, req.(*Filter))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pop_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Credentials)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PopServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pop.Pop/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PopServer).Login(ctx, req.(*Credentials))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pop_Logout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(google_protobuf.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PopServer).Logout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pop.Pop/Logout",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PopServer).Logout(ctx, req.(*google_protobuf.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _Pop_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pop.Pop",
	HandlerType: (*PopServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Containers",
			Handler:    _Pop_Containers_Handler,
		},
		{
			MethodName: "Images",
			Handler:    _Pop_Images_Handler,
		},
		{
			MethodName: "Info",
			Handler:    _Pop_Info_Handler,
		},
		{
			MethodName: "Networks",
			Handler:    _Pop_Networks_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _Pop_Login_Handler,
		},
		{
			MethodName: "Logout",
			Handler:    _Pop_Logout_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pop.proto",
}

func init() { proto.RegisterFile("pop.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 701 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x54, 0x51, 0x4f, 0xdb, 0x3a,
	0x14, 0xa6, 0x49, 0x93, 0x36, 0x27, 0x97, 0x5e, 0x64, 0xdd, 0x8b, 0x72, 0x7b, 0xc7, 0x56, 0x05,
	0x21, 0xba, 0x4d, 0xeb, 0xa4, 0x32, 0xa1, 0x89, 0x3d, 0xa2, 0x6e, 0xaa, 0xc6, 0xa6, 0x29, 0xec,
	0x1d, 0x19, 0x6c, 0xaa, 0x88, 0xc6, 0xb6, 0x62, 0x17, 0xd6, 0x3f, 0xb1, 0xd7, 0xfd, 0xa8, 0xfd,
	0xa9, 0xc9, 0xc7, 0x4e, 0xa1, 0x30, 0xa4, 0xbd, 0xf9, 0x9c, 0xef, 0xe4, 0xf8, 0x3b, 0xdf, 0xf9,
	0x1c, 0x48, 0x94, 0x54, 0x23, 0x55, 0x4b, 0x23, 0x49, 0xa8, 0xa4, 0xea, 0xa7, 0xbc, 0x52, 0x66,
	0xe9, 0x32, 0xf9, 0xcf, 0x00, 0x92, 0x63, 0x29, 0x0c, 0x2d, 0x05, 0xaf, 0x49, 0x0f, 0x82, 0x92,
	0x65, 0xad, 0x41, 0x6b, 0x98, 0x14, 0x41, 0xc9, 0xc8, 0x3f, 0x10, 0x09, 0x5a, 0x71, 0x9d, 0x05,
	0x83, 0x70, 0x98, 0x14, 0x2e, 0x20, 0xff, 0x41, 0xb7, 0xac, 0xe8, 0x8c, 0x9f, 0x95, 0x2c, 0x0b,
	0xb1, 0xb6, 0x83, 0xf1, 0x94, 0x91, 0x0c, 0x3a, 0x17, 0xb2, 0xaa, 0xa8, 0x60, 0x59, 0xdb, 0x21,
	0x3e, 0x44, 0xa4, 0xe6, 0xd4, 0x70, 0x96, 0x45, 0x83, 0xd6, 0x30, 0x2c, 0x9a, 0x90, 0x6c, 0x43,
	0xac, 0x0d, 0x35, 0x0b, 0x9d, 0xc5, 0xf8, 0x89, 0x8f, 0xc8, 0x3e, 0xfc, 0xcd, 0xbf, 0x19, 0x2e,
	0x18, 0x67, 0x67, 0xbe, 0xa0, 0x83, 0x05, 0xbd, 0x26, 0x7d, 0xea, 0x0a, 0xdf, 0x41, 0xc2, 0x05,
	0x53, 0xb2, 0x14, 0x46, 0x67, 0xdd, 0x41, 0x38, 0x4c, 0xc7, 0x3b, 0x23, 0x3b, 0xf4, 0x6a, 0xb0,
	0xd1, 0xa4, 0xc1, 0x27, 0xc2, 0xd4, 0xcb, 0xe2, 0xb6, 0xbe, 0xff, 0x11, 0x7a, 0xeb, 0x20, 0xd9,
	0x82, 0xf0, 0x8a, 0x2f, 0xbd, 0x0a, 0xf6, 0x48, 0x76, 0x21, 0xba, 0xa6, 0xf3, 0x05, 0xcf, 0x82,
	0x41, 0x6b, 0x98, 0x8e, 0x37, 0xb1, 0x79, 0xf3, 0x55, 0xe1, 0xb0, 0xa3, 0xe0, 0x6d, 0x2b, 0x3f,
	0x80, 0xcd, 0xd5, 0x9d, 0x27, 0xa5, 0x36, 0x24, 0x87, 0xf6, 0xbc, 0xd4, 0x26, 0x6b, 0x21, 0xab,
	0xde, 0x3a, 0xab, 0x02, 0xb1, 0x7c, 0x02, 0xe9, 0x71, 0xcd, 0x19, 0x17, 0xa6, 0xa4, 0x73, 0x4d,
	0xfa, 0xd0, 0x5d, 0x68, 0x5e, 0x5b, 0xa9, 0x3d, 0x87, 0x55, 0x6c, 0x31, 0x45, 0xb5, 0xbe, 0x91,
	0x35, 0x43, 0x2e, 0x49, 0xb1, 0x8a, 0xf3, 0xef, 0x2d, 0xe8, 0x36, 0x9c, 0xc8, 0xbf, 0x10, 0x0b,
	0x6e, 0xce, 0x56, 0xcb, 0x8c, 0x04, 0x37, 0x53, 0x46, 0x9e, 0x41, 0xda, 0x4c, 0x6e, 0x31, 0xd7,
	0x02, 0x9a, 0xd4, 0x94, 0x91, 0xff, 0xa1, 0x5d, 0xaa, 0xeb, 0x37, 0xb8, 0xd6, 0x74, 0xdc, 0x41,
	0xbe, 0x53, 0x55, 0x60, 0xd2, 0x83, 0x87, 0xb8, 0xd9, 0x7b, 0xe0, 0xa1, 0x55, 0xad, 0xa2, 0x17,
	0xb8, 0xdb, 0xa4, 0xb0, 0xc7, 0x3c, 0x83, 0xf8, 0x7d, 0x39, 0x37, 0x0f, 0x6d, 0x95, 0x7f, 0x80,
	0x68, 0x6a, 0x0d, 0xf3, 0x87, 0x7e, 0xbb, 0x63, 0x9d, 0x70, 0xcd, 0x3a, 0xf9, 0x4b, 0x48, 0xb0,
	0x11, 0x6a, 0xfd, 0x74, 0x4d, 0x6b, 0x70, 0xf4, 0x2c, 0xea, 0x75, 0xfe, 0x04, 0xd1, 0x54, 0x5c,
	0x4a, 0x4d, 0x08, 0xb4, 0xcd, 0x52, 0x35, 0xea, 0xe2, 0xd9, 0xe6, 0x50, 0x71, 0x27, 0x09, 0x9e,
	0xc9, 0x13, 0x48, 0x4c, 0x59, 0x71, 0x6d, 0x68, 0xa5, 0xfc, 0xcd, 0xb7, 0x89, 0xfc, 0x18, 0x82,
	0xa9, 0xb2, 0xdc, 0x28, 0x63, 0x35, 0xd7, 0xda, 0xb7, 0x6b, 0x42, 0xb2, 0x0b, 0xb1, 0x5e, 0x9c,
	0x0b, 0x6e, 0xbc, 0x6b, 0x52, 0x24, 0x74, 0x8a, 0xa9, 0xc2, 0x43, 0xb9, 0x82, 0xce, 0x67, 0x6e,
	0x6e, 0x64, 0x7d, 0xf5, 0x40, 0x8b, 0xdf, 0x31, 0xea, 0x43, 0xd7, 0x7a, 0xbf, 0x16, 0x74, 0x8e,
	0x84, 0xba, 0xc5, 0x2a, 0x26, 0x7b, 0xd0, 0x71, 0x4d, 0x75, 0xd6, 0x46, 0x05, 0xd6, 0x2e, 0x6c,
	0xb0, 0xfc, 0x35, 0xa4, 0xfe, 0x46, 0x14, 0x6d, 0xb0, 0x26, 0xda, 0x5f, 0xf8, 0x89, 0xc7, 0xbd,
	0x6c, 0x87, 0x10, 0xbb, 0x1e, 0x96, 0xd1, 0x45, 0xc9, 0xea, 0x46, 0x37, 0x7b, 0xb6, 0xf3, 0xcf,
	0xa8, 0xe1, 0x37, 0x74, 0xe9, 0x89, 0x36, 0x61, 0xbe, 0x03, 0xd1, 0x57, 0x79, 0xc5, 0x85, 0x5d,
	0xaa, 0x7b, 0x3d, 0xde, 0x8a, 0x18, 0x8c, 0x7f, 0x04, 0x10, 0x7e, 0x91, 0x8a, 0xbc, 0x02, 0x58,
	0x3d, 0x08, 0x4d, 0x1c, 0x67, 0x67, 0x9b, 0x3e, 0x59, 0x7f, 0x2e, 0xc8, 0x77, 0x0f, 0x62, 0xdc,
	0xe9, 0xbd, 0xd2, 0xde, 0xed, 0xb6, 0xb1, 0xec, 0x05, 0xb4, 0xed, 0xae, 0xc9, 0xf6, 0x68, 0x26,
	0xe5, 0x6c, 0xce, 0xdd, 0xdf, 0xee, 0x7c, 0x71, 0x39, 0x9a, 0xd8, 0x9f, 0x5f, 0xdf, 0xbb, 0x03,
	0xed, 0xf0, 0x1c, 0xba, 0x7e, 0xe2, 0x7b, 0x4d, 0xb7, 0xee, 0xaa, 0x81, 0x6d, 0xf7, 0x21, 0x3a,
	0x91, 0xb3, 0x52, 0x10, 0x07, 0xdd, 0x79, 0xb6, 0xbe, 0x23, 0x4e, 0x9c, 0x6f, 0x90, 0x23, 0x88,
	0x4f, 0xe4, 0x4c, 0x2e, 0xcc, 0xa3, 0x0c, 0x1e, 0xc9, 0xe7, 0x1b, 0xe7, 0x31, 0x66, 0x0e, 0x7e,
	0x05, 0x00, 0x00, 0xff, 0xff, 0xca, 0x3e, 0x0a, 0x7c, 0xb8, 0x05, 0x00, 0x00,
}