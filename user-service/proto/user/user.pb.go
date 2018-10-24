// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/user/user.proto

package go_micro_srv_user

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "golang.org/x/net/context"
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

type User struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Company              string   `protobuf:"bytes,3,opt,name=company,proto3" json:"company,omitempty"`
	Email                string   `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	Password             string   `protobuf:"bytes,5,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{0}
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

func (m *User) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetCompany() string {
	if m != nil {
		return m.Company
	}
	return ""
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

type Request struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{1}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

type CreateResponse struct {
	Created              bool     `protobuf:"varint,1,opt,name=created,proto3" json:"created,omitempty"`
	User                 *User    `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateResponse) Reset()         { *m = CreateResponse{} }
func (m *CreateResponse) String() string { return proto.CompactTextString(m) }
func (*CreateResponse) ProtoMessage()    {}
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{2}
}

func (m *CreateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateResponse.Unmarshal(m, b)
}
func (m *CreateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateResponse.Marshal(b, m, deterministic)
}
func (m *CreateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateResponse.Merge(m, src)
}
func (m *CreateResponse) XXX_Size() int {
	return xxx_messageInfo_CreateResponse.Size(m)
}
func (m *CreateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateResponse proto.InternalMessageInfo

func (m *CreateResponse) GetCreated() bool {
	if m != nil {
		return m.Created
	}
	return false
}

func (m *CreateResponse) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type UserResponse struct {
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserResponse) Reset()         { *m = UserResponse{} }
func (m *UserResponse) String() string { return proto.CompactTextString(m) }
func (*UserResponse) ProtoMessage()    {}
func (*UserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{3}
}

func (m *UserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserResponse.Unmarshal(m, b)
}
func (m *UserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserResponse.Marshal(b, m, deterministic)
}
func (m *UserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserResponse.Merge(m, src)
}
func (m *UserResponse) XXX_Size() int {
	return xxx_messageInfo_UserResponse.Size(m)
}
func (m *UserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UserResponse proto.InternalMessageInfo

func (m *UserResponse) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type UsersResponse struct {
	Users                []*User  `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UsersResponse) Reset()         { *m = UsersResponse{} }
func (m *UsersResponse) String() string { return proto.CompactTextString(m) }
func (*UsersResponse) ProtoMessage()    {}
func (*UsersResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{4}
}

func (m *UsersResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UsersResponse.Unmarshal(m, b)
}
func (m *UsersResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UsersResponse.Marshal(b, m, deterministic)
}
func (m *UsersResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UsersResponse.Merge(m, src)
}
func (m *UsersResponse) XXX_Size() int {
	return xxx_messageInfo_UsersResponse.Size(m)
}
func (m *UsersResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UsersResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UsersResponse proto.InternalMessageInfo

func (m *UsersResponse) GetUsers() []*User {
	if m != nil {
		return m.Users
	}
	return nil
}

type Token struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Valid                bool     `protobuf:"varint,2,opt,name=valid,proto3" json:"valid,omitempty"`
	Errors               []*Error `protobuf:"bytes,3,rep,name=errors,proto3" json:"errors,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Token) Reset()         { *m = Token{} }
func (m *Token) String() string { return proto.CompactTextString(m) }
func (*Token) ProtoMessage()    {}
func (*Token) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{5}
}

func (m *Token) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Token.Unmarshal(m, b)
}
func (m *Token) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Token.Marshal(b, m, deterministic)
}
func (m *Token) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Token.Merge(m, src)
}
func (m *Token) XXX_Size() int {
	return xxx_messageInfo_Token.Size(m)
}
func (m *Token) XXX_DiscardUnknown() {
	xxx_messageInfo_Token.DiscardUnknown(m)
}

var xxx_messageInfo_Token proto.InternalMessageInfo

func (m *Token) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *Token) GetValid() bool {
	if m != nil {
		return m.Valid
	}
	return false
}

func (m *Token) GetErrors() []*Error {
	if m != nil {
		return m.Errors
	}
	return nil
}

type Error struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Description          string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Error) Reset()         { *m = Error{} }
func (m *Error) String() string { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()    {}
func (*Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{6}
}

func (m *Error) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Error.Unmarshal(m, b)
}
func (m *Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Error.Marshal(b, m, deterministic)
}
func (m *Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Error.Merge(m, src)
}
func (m *Error) XXX_Size() int {
	return xxx_messageInfo_Error.Size(m)
}
func (m *Error) XXX_DiscardUnknown() {
	xxx_messageInfo_Error.DiscardUnknown(m)
}

var xxx_messageInfo_Error proto.InternalMessageInfo

func (m *Error) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Error) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func init() {
	proto.RegisterType((*User)(nil), "go.micro.srv.user.User")
	proto.RegisterType((*Request)(nil), "go.micro.srv.user.Request")
	proto.RegisterType((*CreateResponse)(nil), "go.micro.srv.user.CreateResponse")
	proto.RegisterType((*UserResponse)(nil), "go.micro.srv.user.UserResponse")
	proto.RegisterType((*UsersResponse)(nil), "go.micro.srv.user.UsersResponse")
	proto.RegisterType((*Token)(nil), "go.micro.srv.user.Token")
	proto.RegisterType((*Error)(nil), "go.micro.srv.user.Error")
}

func init() { proto.RegisterFile("proto/user/user.proto", fileDescriptor_9b283a848145d6b7) }

var fileDescriptor_9b283a848145d6b7 = []byte{
	// 404 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0x4d, 0x6f, 0xd4, 0x30,
	0x10, 0xed, 0xe6, 0xab, 0xdb, 0x59, 0x5a, 0x09, 0x0b, 0x84, 0x95, 0x0b, 0xc1, 0x27, 0x24, 0x44,
	0x40, 0xe5, 0x58, 0x81, 0xb4, 0x42, 0xa5, 0x9c, 0xcd, 0xd7, 0x39, 0x24, 0x23, 0xb0, 0x48, 0xe2,
	0x60, 0x7b, 0x83, 0xf8, 0xc3, 0xfc, 0x0e, 0xe4, 0x71, 0x52, 0x8a, 0x48, 0x2a, 0x2e, 0x2b, 0xbf,
	0x37, 0x33, 0x6f, 0x3e, 0xde, 0x06, 0xee, 0x0f, 0x46, 0x3b, 0xfd, 0xec, 0x60, 0xd1, 0xd0, 0x4f,
	0x49, 0x98, 0xdd, 0xfd, 0xa2, 0xcb, 0x4e, 0xd5, 0x46, 0x97, 0xd6, 0x8c, 0xa5, 0x0f, 0x88, 0x11,
	0x92, 0x0f, 0x16, 0x0d, 0x3b, 0x83, 0x48, 0x35, 0x7c, 0x53, 0x6c, 0x1e, 0x9f, 0xc8, 0x48, 0x35,
	0x8c, 0x41, 0xd2, 0x57, 0x1d, 0xf2, 0x88, 0x18, 0x7a, 0x33, 0x0e, 0xc7, 0xb5, 0xee, 0x86, 0xaa,
	0xff, 0xc9, 0x63, 0xa2, 0x67, 0xc8, 0xee, 0x41, 0x8a, 0x5d, 0xa5, 0x5a, 0x9e, 0x10, 0x1f, 0x00,
	0xcb, 0x61, 0x3b, 0x54, 0xd6, 0xfe, 0xd0, 0xa6, 0xe1, 0x29, 0x05, 0xae, 0xb1, 0x38, 0x81, 0x63,
	0x89, 0xdf, 0x0f, 0x68, 0x9d, 0xf8, 0x04, 0x67, 0xaf, 0x0d, 0x56, 0x0e, 0x25, 0xda, 0x41, 0xf7,
	0x36, 0x34, 0x22, 0x26, 0x4c, 0xb4, 0x95, 0x33, 0x64, 0x4f, 0x20, 0xf1, 0x63, 0xd3, 0x58, 0xbb,
	0xf3, 0x07, 0xe5, 0x3f, 0x0b, 0x95, 0x7e, 0x1b, 0x49, 0x49, 0xe2, 0x02, 0xee, 0x10, 0x9a, 0x65,
	0xe7, 0xe2, 0xcd, 0xff, 0x14, 0xbf, 0x82, 0x53, 0x8f, 0xec, 0x75, 0xf5, 0x53, 0x48, 0x7d, 0xc0,
	0xf2, 0x4d, 0x11, 0xdf, 0x56, 0x1e, 0xb2, 0x04, 0x42, 0xfa, 0x5e, 0x7f, 0xc3, 0xde, 0xdf, 0xc6,
	0xf9, 0xc7, 0x74, 0xdc, 0x00, 0x3c, 0x3b, 0x56, 0xad, 0x6a, 0x68, 0x93, 0xad, 0x0c, 0x80, 0x3d,
	0x87, 0x0c, 0x8d, 0xd1, 0xc6, 0xf2, 0x98, 0x9a, 0xf0, 0x85, 0x26, 0x97, 0x3e, 0x41, 0x4e, 0x79,
	0xe2, 0x25, 0xa4, 0x44, 0x78, 0xc3, 0x6a, 0xdd, 0x20, 0x75, 0x49, 0x25, 0xbd, 0x59, 0x01, 0xbb,
	0x06, 0x6d, 0x6d, 0xd4, 0xe0, 0x94, 0xee, 0x27, 0x2f, 0x6f, 0x52, 0xe7, 0xbf, 0x22, 0xd8, 0xf9,
	0xa9, 0xdf, 0xa1, 0x19, 0x55, 0x8d, 0xec, 0x0d, 0x64, 0xc1, 0x0b, 0xb6, 0xb6, 0x5f, 0xfe, 0x68,
	0x21, 0xf0, 0xb7, 0x7f, 0xe2, 0x88, 0xed, 0x21, 0xbe, 0x42, 0xb7, 0x2e, 0xf2, 0x70, 0xed, 0x7a,
	0x7f, 0x24, 0xde, 0x42, 0x76, 0x85, 0x6e, 0xdf, 0xb6, 0x2c, 0x5f, 0x48, 0x9e, 0xfe, 0x3c, 0x79,
	0xb1, 0x22, 0x64, 0x6f, 0x28, 0x5d, 0x40, 0xb2, 0x3f, 0xb8, 0xaf, 0xeb, 0xd3, 0x2c, 0x9d, 0x99,
	0xcc, 0x13, 0x47, 0xec, 0x12, 0x4e, 0x3f, 0x7a, 0x6f, 0x2a, 0x87, 0xc1, 0xcf, 0xd5, 0xe4, 0xdb,
	0x64, 0x3e, 0x67, 0xf4, 0x05, 0xbe, 0xf8, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x95, 0xe3, 0x6d, 0xc1,
	0x9a, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for UserService service

type UserServiceClient interface {
	Create(ctx context.Context, in *User, opts ...client.CallOption) (*CreateResponse, error)
	Get(ctx context.Context, in *User, opts ...client.CallOption) (*UserResponse, error)
	GetAll(ctx context.Context, in *Request, opts ...client.CallOption) (*UsersResponse, error)
	Auth(ctx context.Context, in *User, opts ...client.CallOption) (*Token, error)
	ValidateToken(ctx context.Context, in *Token, opts ...client.CallOption) (*Token, error)
}

type userServiceClient struct {
	c           client.Client
	serviceName string
}

func NewUserServiceClient(serviceName string, c client.Client) UserServiceClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "go.micro.srv.user"
	}
	return &userServiceClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *userServiceClient) Create(ctx context.Context, in *User, opts ...client.CallOption) (*CreateResponse, error) {
	req := c.c.NewRequest(c.serviceName, "UserService.Create", in)
	out := new(CreateResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Get(ctx context.Context, in *User, opts ...client.CallOption) (*UserResponse, error) {
	req := c.c.NewRequest(c.serviceName, "UserService.Get", in)
	out := new(UserResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetAll(ctx context.Context, in *Request, opts ...client.CallOption) (*UsersResponse, error) {
	req := c.c.NewRequest(c.serviceName, "UserService.GetAll", in)
	out := new(UsersResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Auth(ctx context.Context, in *User, opts ...client.CallOption) (*Token, error) {
	req := c.c.NewRequest(c.serviceName, "UserService.Auth", in)
	out := new(Token)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) ValidateToken(ctx context.Context, in *Token, opts ...client.CallOption) (*Token, error) {
	req := c.c.NewRequest(c.serviceName, "UserService.ValidateToken", in)
	out := new(Token)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for UserService service

type UserServiceHandler interface {
	Create(context.Context, *User, *CreateResponse) error
	Get(context.Context, *User, *UserResponse) error
	GetAll(context.Context, *Request, *UsersResponse) error
	Auth(context.Context, *User, *Token) error
	ValidateToken(context.Context, *Token, *Token) error
}

func RegisterUserServiceHandler(s server.Server, hdlr UserServiceHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&UserService{hdlr}, opts...))
}

type UserService struct {
	UserServiceHandler
}

func (h *UserService) Create(ctx context.Context, in *User, out *CreateResponse) error {
	return h.UserServiceHandler.Create(ctx, in, out)
}

func (h *UserService) Get(ctx context.Context, in *User, out *UserResponse) error {
	return h.UserServiceHandler.Get(ctx, in, out)
}

func (h *UserService) GetAll(ctx context.Context, in *Request, out *UsersResponse) error {
	return h.UserServiceHandler.GetAll(ctx, in, out)
}

func (h *UserService) Auth(ctx context.Context, in *User, out *Token) error {
	return h.UserServiceHandler.Auth(ctx, in, out)
}

func (h *UserService) ValidateToken(ctx context.Context, in *Token, out *Token) error {
	return h.UserServiceHandler.ValidateToken(ctx, in, out)
}
