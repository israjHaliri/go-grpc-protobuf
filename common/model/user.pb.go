// Code generated by protoc-gen-go. DO NOT EDIT.
// source: common/model/user.proto

package model

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type User struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Email                string   `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Password             string   `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	Fullname             string   `protobuf:"bytes,4,opt,name=fullname,proto3" json:"fullname,omitempty"`
	Token                string   `protobuf:"bytes,5,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_d6e08ea5475f97cc, []int{0}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *User) GetFullname() string {
	if m != nil {
		return m.Fullname
	}
	return ""
}

func (m *User) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type UserList struct {
	List                 []*User  `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserList) Reset()         { *m = UserList{} }
func (m *UserList) String() string { return proto.CompactTextString(m) }
func (*UserList) ProtoMessage()    {}
func (*UserList) Descriptor() ([]byte, []int) {
	return fileDescriptor_d6e08ea5475f97cc, []int{1}
}

func (m *UserList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserList.Unmarshal(m, b)
}
func (m *UserList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserList.Marshal(b, m, deterministic)
}
func (m *UserList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserList.Merge(m, src)
}
func (m *UserList) XXX_Size() int {
	return xxx_messageInfo_UserList.Size(m)
}
func (m *UserList) XXX_DiscardUnknown() {
	xxx_messageInfo_UserList.DiscardUnknown(m)
}

var xxx_messageInfo_UserList proto.InternalMessageInfo

func (m *UserList) GetList() []*User {
	if m != nil {
		return m.List
	}
	return nil
}

type EmailAndPassword struct {
	Email                string   `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EmailAndPassword) Reset()         { *m = EmailAndPassword{} }
func (m *EmailAndPassword) String() string { return proto.CompactTextString(m) }
func (*EmailAndPassword) ProtoMessage()    {}
func (*EmailAndPassword) Descriptor() ([]byte, []int) {
	return fileDescriptor_d6e08ea5475f97cc, []int{2}
}

func (m *EmailAndPassword) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EmailAndPassword.Unmarshal(m, b)
}
func (m *EmailAndPassword) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EmailAndPassword.Marshal(b, m, deterministic)
}
func (m *EmailAndPassword) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EmailAndPassword.Merge(m, src)
}
func (m *EmailAndPassword) XXX_Size() int {
	return xxx_messageInfo_EmailAndPassword.Size(m)
}
func (m *EmailAndPassword) XXX_DiscardUnknown() {
	xxx_messageInfo_EmailAndPassword.DiscardUnknown(m)
}

var xxx_messageInfo_EmailAndPassword proto.InternalMessageInfo

func (m *EmailAndPassword) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *EmailAndPassword) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func init() {
	proto.RegisterType((*User)(nil), "model.User")
	proto.RegisterType((*UserList)(nil), "model.UserList")
	proto.RegisterType((*EmailAndPassword)(nil), "model.EmailAndPassword")
}

func init() { proto.RegisterFile("common/model/user.proto", fileDescriptor_d6e08ea5475f97cc) }

var fileDescriptor_d6e08ea5475f97cc = []byte{
	// 283 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0x51, 0x4b, 0xc3, 0x30,
	0x14, 0x85, 0x97, 0xae, 0x95, 0x79, 0x27, 0x2a, 0x41, 0x5c, 0xa8, 0x0f, 0x96, 0x3e, 0x15, 0x84,
	0x14, 0xe6, 0x93, 0x8f, 0x95, 0xcd, 0x27, 0x1f, 0xa4, 0xe0, 0x0f, 0xe8, 0x4c, 0x36, 0x82, 0x49,
	0x53, 0x9a, 0x16, 0x99, 0xbf, 0xc0, 0x9f, 0x2d, 0x49, 0xd6, 0x32, 0x85, 0x3d, 0x9e, 0x7b, 0x72,
	0xef, 0xf9, 0x72, 0x60, 0xf1, 0xa1, 0x95, 0xd2, 0x75, 0xae, 0x34, 0xe3, 0x32, 0xef, 0x0d, 0x6f,
	0x69, 0xd3, 0xea, 0x4e, 0xe3, 0xc8, 0x4d, 0xe2, 0xbb, 0x9d, 0xd6, 0x3b, 0xc9, 0x73, 0x37, 0xdc,
	0xf4, 0xdb, 0x7c, 0xad, 0x9a, 0x6e, 0xef, 0xdf, 0xa4, 0xdf, 0x10, 0xbe, 0x1b, 0xde, 0xe2, 0x4b,
	0x08, 0x04, 0x23, 0x28, 0x41, 0x59, 0x54, 0x06, 0x82, 0xe1, 0x1b, 0x88, 0xb8, 0xaa, 0x84, 0x24,
	0x41, 0x82, 0xb2, 0xf3, 0xd2, 0x0b, 0x1c, 0xc3, 0xac, 0xa9, 0x8c, 0xf9, 0xd2, 0x2d, 0x23, 0x53,
	0x67, 0x8c, 0xda, 0x7a, 0xdb, 0x5e, 0xca, 0xba, 0x52, 0x9c, 0x84, 0xde, 0x1b, 0xb4, 0xbd, 0xd6,
	0xe9, 0x4f, 0x5e, 0x93, 0xc8, 0x5f, 0x73, 0x22, 0x7d, 0x80, 0x99, 0xcd, 0x7e, 0x15, 0xa6, 0xc3,
	0xf7, 0x10, 0x4a, 0x61, 0x3a, 0x82, 0x92, 0x69, 0x36, 0x5f, 0xce, 0xa9, 0x43, 0xa7, 0xd6, 0x2e,
	0x9d, 0x91, 0xae, 0xe0, 0x7a, 0x6d, 0x19, 0x8a, 0x9a, 0xbd, 0x0d, 0x91, 0x23, 0x24, 0x3a, 0x05,
	0x19, 0xfc, 0x85, 0x5c, 0xfe, 0x20, 0x88, 0xec, 0x51, 0x83, 0x9f, 0xe0, 0xe2, 0x45, 0xd4, 0xac,
	0x90, 0xd2, 0xeb, 0x5b, 0xea, 0x6b, 0xa2, 0x43, 0x4d, 0xd4, 0xd5, 0x14, 0x5f, 0x1d, 0xa1, 0x58,
	0xd2, 0x74, 0x82, 0x57, 0x10, 0xdb, 0x55, 0x3b, 0x79, 0xde, 0x8f, 0x50, 0xc5, 0x01, 0x6a, 0x71,
	0x58, 0xf8, 0x4f, 0x1b, 0x1f, 0x7f, 0x2a, 0x9d, 0x6c, 0xce, 0x5c, 0xd0, 0xe3, 0x6f, 0x00, 0x00,
	0x00, 0xff, 0xff, 0xb2, 0xe9, 0x17, 0x97, 0xbf, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UsersClient is the client API for Users service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UsersClient interface {
	FindAllUsers(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*UserList, error)
	FindUserByEmailAndPAssword(ctx context.Context, in *EmailAndPassword, opts ...grpc.CallOption) (*User, error)
}

type usersClient struct {
	cc *grpc.ClientConn
}

func NewUsersClient(cc *grpc.ClientConn) UsersClient {
	return &usersClient{cc}
}

func (c *usersClient) FindAllUsers(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*UserList, error) {
	out := new(UserList)
	err := c.cc.Invoke(ctx, "/model.Users/FindAllUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) FindUserByEmailAndPAssword(ctx context.Context, in *EmailAndPassword, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/model.Users/FindUserByEmailAndPAssword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UsersServer is the server API for Users service.
type UsersServer interface {
	FindAllUsers(context.Context, *empty.Empty) (*UserList, error)
	FindUserByEmailAndPAssword(context.Context, *EmailAndPassword) (*User, error)
}

// UnimplementedUsersServer can be embedded to have forward compatible implementations.
type UnimplementedUsersServer struct {
}

func (*UnimplementedUsersServer) FindAllUsers(ctx context.Context, req *empty.Empty) (*UserList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindAllUsers not implemented")
}
func (*UnimplementedUsersServer) FindUserByEmailAndPAssword(ctx context.Context, req *EmailAndPassword) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindUserByEmailAndPAssword not implemented")
}

func RegisterUsersServer(s *grpc.Server, srv UsersServer) {
	s.RegisterService(&_Users_serviceDesc, srv)
}

func _Users_FindAllUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).FindAllUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/model.Users/FindAllUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).FindAllUsers(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_FindUserByEmailAndPAssword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmailAndPassword)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).FindUserByEmailAndPAssword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/model.Users/FindUserByEmailAndPAssword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).FindUserByEmailAndPAssword(ctx, req.(*EmailAndPassword))
	}
	return interceptor(ctx, in, info, handler)
}

var _Users_serviceDesc = grpc.ServiceDesc{
	ServiceName: "model.Users",
	HandlerType: (*UsersServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FindAllUsers",
			Handler:    _Users_FindAllUsers_Handler,
		},
		{
			MethodName: "FindUserByEmailAndPAssword",
			Handler:    _Users_FindUserByEmailAndPAssword_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "common/model/user.proto",
}
