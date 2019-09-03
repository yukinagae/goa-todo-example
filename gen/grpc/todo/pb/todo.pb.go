// Code generated by protoc-gen-go. DO NOT EDIT.
// source: todo.proto

package todopb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type GetRequest struct {
	// Todo ID
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRequest) Reset()         { *m = GetRequest{} }
func (m *GetRequest) String() string { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()    {}
func (*GetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e4b95d0c4e09639, []int{0}
}

func (m *GetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRequest.Unmarshal(m, b)
}
func (m *GetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRequest.Marshal(b, m, deterministic)
}
func (m *GetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRequest.Merge(m, src)
}
func (m *GetRequest) XXX_Size() int {
	return xxx_messageInfo_GetRequest.Size(m)
}
func (m *GetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRequest proto.InternalMessageInfo

func (m *GetRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type GetResponse struct {
	// Todo ID
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Task                 string   `protobuf:"bytes,2,opt,name=task,proto3" json:"task,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetResponse) Reset()         { *m = GetResponse{} }
func (m *GetResponse) String() string { return proto.CompactTextString(m) }
func (*GetResponse) ProtoMessage()    {}
func (*GetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e4b95d0c4e09639, []int{1}
}

func (m *GetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetResponse.Unmarshal(m, b)
}
func (m *GetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetResponse.Marshal(b, m, deterministic)
}
func (m *GetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetResponse.Merge(m, src)
}
func (m *GetResponse) XXX_Size() int {
	return xxx_messageInfo_GetResponse.Size(m)
}
func (m *GetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetResponse proto.InternalMessageInfo

func (m *GetResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *GetResponse) GetTask() string {
	if m != nil {
		return m.Task
	}
	return ""
}

type ListRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListRequest) Reset()         { *m = ListRequest{} }
func (m *ListRequest) String() string { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()    {}
func (*ListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e4b95d0c4e09639, []int{2}
}

func (m *ListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListRequest.Unmarshal(m, b)
}
func (m *ListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListRequest.Marshal(b, m, deterministic)
}
func (m *ListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListRequest.Merge(m, src)
}
func (m *ListRequest) XXX_Size() int {
	return xxx_messageInfo_ListRequest.Size(m)
}
func (m *ListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListRequest proto.InternalMessageInfo

type TodoCollection struct {
	Field                []*Todo1 `protobuf:"bytes,1,rep,name=field,proto3" json:"field,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TodoCollection) Reset()         { *m = TodoCollection{} }
func (m *TodoCollection) String() string { return proto.CompactTextString(m) }
func (*TodoCollection) ProtoMessage()    {}
func (*TodoCollection) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e4b95d0c4e09639, []int{3}
}

func (m *TodoCollection) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TodoCollection.Unmarshal(m, b)
}
func (m *TodoCollection) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TodoCollection.Marshal(b, m, deterministic)
}
func (m *TodoCollection) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TodoCollection.Merge(m, src)
}
func (m *TodoCollection) XXX_Size() int {
	return xxx_messageInfo_TodoCollection.Size(m)
}
func (m *TodoCollection) XXX_DiscardUnknown() {
	xxx_messageInfo_TodoCollection.DiscardUnknown(m)
}

var xxx_messageInfo_TodoCollection proto.InternalMessageInfo

func (m *TodoCollection) GetField() []*Todo1 {
	if m != nil {
		return m.Field
	}
	return nil
}

type Todo1 struct {
	// Todo ID
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Task                 string   `protobuf:"bytes,2,opt,name=task,proto3" json:"task,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Todo1) Reset()         { *m = Todo1{} }
func (m *Todo1) String() string { return proto.CompactTextString(m) }
func (*Todo1) ProtoMessage()    {}
func (*Todo1) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e4b95d0c4e09639, []int{4}
}

func (m *Todo1) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Todo1.Unmarshal(m, b)
}
func (m *Todo1) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Todo1.Marshal(b, m, deterministic)
}
func (m *Todo1) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Todo1.Merge(m, src)
}
func (m *Todo1) XXX_Size() int {
	return xxx_messageInfo_Todo1.Size(m)
}
func (m *Todo1) XXX_DiscardUnknown() {
	xxx_messageInfo_Todo1.DiscardUnknown(m)
}

var xxx_messageInfo_Todo1 proto.InternalMessageInfo

func (m *Todo1) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Todo1) GetTask() string {
	if m != nil {
		return m.Task
	}
	return ""
}

type AddRequest struct {
	// Todo ID
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Todo task
	Task                 string   `protobuf:"bytes,2,opt,name=task,proto3" json:"task,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddRequest) Reset()         { *m = AddRequest{} }
func (m *AddRequest) String() string { return proto.CompactTextString(m) }
func (*AddRequest) ProtoMessage()    {}
func (*AddRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e4b95d0c4e09639, []int{5}
}

func (m *AddRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddRequest.Unmarshal(m, b)
}
func (m *AddRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddRequest.Marshal(b, m, deterministic)
}
func (m *AddRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddRequest.Merge(m, src)
}
func (m *AddRequest) XXX_Size() int {
	return xxx_messageInfo_AddRequest.Size(m)
}
func (m *AddRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddRequest proto.InternalMessageInfo

func (m *AddRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *AddRequest) GetTask() string {
	if m != nil {
		return m.Task
	}
	return ""
}

type AddResponse struct {
	Field                string   `protobuf:"bytes,1,opt,name=field,proto3" json:"field,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddResponse) Reset()         { *m = AddResponse{} }
func (m *AddResponse) String() string { return proto.CompactTextString(m) }
func (*AddResponse) ProtoMessage()    {}
func (*AddResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e4b95d0c4e09639, []int{6}
}

func (m *AddResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddResponse.Unmarshal(m, b)
}
func (m *AddResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddResponse.Marshal(b, m, deterministic)
}
func (m *AddResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddResponse.Merge(m, src)
}
func (m *AddResponse) XXX_Size() int {
	return xxx_messageInfo_AddResponse.Size(m)
}
func (m *AddResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AddResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AddResponse proto.InternalMessageInfo

func (m *AddResponse) GetField() string {
	if m != nil {
		return m.Field
	}
	return ""
}

type RemoveRequest struct {
	// Todo ID
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RemoveRequest) Reset()         { *m = RemoveRequest{} }
func (m *RemoveRequest) String() string { return proto.CompactTextString(m) }
func (*RemoveRequest) ProtoMessage()    {}
func (*RemoveRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e4b95d0c4e09639, []int{7}
}

func (m *RemoveRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoveRequest.Unmarshal(m, b)
}
func (m *RemoveRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoveRequest.Marshal(b, m, deterministic)
}
func (m *RemoveRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoveRequest.Merge(m, src)
}
func (m *RemoveRequest) XXX_Size() int {
	return xxx_messageInfo_RemoveRequest.Size(m)
}
func (m *RemoveRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoveRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RemoveRequest proto.InternalMessageInfo

func (m *RemoveRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type RemoveResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RemoveResponse) Reset()         { *m = RemoveResponse{} }
func (m *RemoveResponse) String() string { return proto.CompactTextString(m) }
func (*RemoveResponse) ProtoMessage()    {}
func (*RemoveResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e4b95d0c4e09639, []int{8}
}

func (m *RemoveResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoveResponse.Unmarshal(m, b)
}
func (m *RemoveResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoveResponse.Marshal(b, m, deterministic)
}
func (m *RemoveResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoveResponse.Merge(m, src)
}
func (m *RemoveResponse) XXX_Size() int {
	return xxx_messageInfo_RemoveResponse.Size(m)
}
func (m *RemoveResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoveResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RemoveResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*GetRequest)(nil), "todo.GetRequest")
	proto.RegisterType((*GetResponse)(nil), "todo.GetResponse")
	proto.RegisterType((*ListRequest)(nil), "todo.ListRequest")
	proto.RegisterType((*TodoCollection)(nil), "todo.TodoCollection")
	proto.RegisterType((*Todo1)(nil), "todo.Todo1")
	proto.RegisterType((*AddRequest)(nil), "todo.AddRequest")
	proto.RegisterType((*AddResponse)(nil), "todo.AddResponse")
	proto.RegisterType((*RemoveRequest)(nil), "todo.RemoveRequest")
	proto.RegisterType((*RemoveResponse)(nil), "todo.RemoveResponse")
}

func init() { proto.RegisterFile("todo.proto", fileDescriptor_0e4b95d0c4e09639) }

var fileDescriptor_0e4b95d0c4e09639 = []byte{
	// 283 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0xcd, 0x4e, 0xf3, 0x30,
	0x10, 0x54, 0x7e, 0x1a, 0x7d, 0xdf, 0x44, 0x8d, 0x8a, 0xe9, 0xa1, 0xaa, 0x90, 0x00, 0x73, 0xa9,
	0x8a, 0x54, 0x68, 0xf3, 0x04, 0x85, 0x43, 0x2f, 0x9c, 0x22, 0x4e, 0xdc, 0x28, 0x36, 0x92, 0x45,
	0xe8, 0x06, 0x6c, 0x78, 0x42, 0x1e, 0x0c, 0xd9, 0x8e, 0x95, 0x14, 0x51, 0x89, 0xdb, 0x7a, 0x67,
	0x77, 0x76, 0x66, 0x64, 0xc0, 0x90, 0xa0, 0x45, 0xf3, 0x4e, 0x86, 0x58, 0x6a, 0x6b, 0x7e, 0x02,
	0x6c, 0xa4, 0xa9, 0xe4, 0xdb, 0x87, 0xd4, 0x86, 0x15, 0x88, 0x95, 0x98, 0x44, 0x67, 0xd1, 0xec,
	0x7f, 0x15, 0x2b, 0xc1, 0x97, 0xc8, 0x1d, 0xaa, 0x1b, 0xda, 0x69, 0xf9, 0x13, 0x66, 0x0c, 0xa9,
	0x79, 0xd4, 0x2f, 0x93, 0xd8, 0x75, 0x5c, 0xcd, 0x87, 0xc8, 0xef, 0x94, 0x0e, 0x8c, 0xbc, 0x44,
	0x71, 0x4f, 0x82, 0x6e, 0xa9, 0xae, 0xe5, 0x93, 0x51, 0xb4, 0x63, 0xe7, 0x18, 0x3c, 0x2b, 0x59,
	0x5b, 0x9e, 0x64, 0x96, 0xaf, 0xf2, 0x85, 0xd3, 0x64, 0x87, 0x96, 0x95, 0x47, 0xf8, 0x25, 0x06,
	0xee, 0xfd, 0xa7, 0x83, 0xd7, 0xc0, 0x5a, 0x88, 0x03, 0x0e, 0x7e, 0xdd, 0xb8, 0x40, 0xee, 0x36,
	0x5a, 0x57, 0xe3, 0x4e, 0x90, 0x9d, 0x69, 0x35, 0x9c, 0x62, 0x58, 0xc9, 0x57, 0xfa, 0x94, 0x87,
	0xb2, 0x19, 0xa1, 0x08, 0x03, 0x9e, 0x68, 0xf5, 0x15, 0x21, 0xb5, 0xba, 0xd9, 0x1c, 0xc9, 0x46,
	0x1a, 0x36, 0xf2, 0xd6, 0xba, 0x7c, 0xa7, 0x47, 0xbd, 0x4e, 0x7b, 0xfd, 0x0a, 0xa9, 0xcd, 0x8b,
	0xb5, 0x50, 0x2f, 0xbb, 0xe9, 0xb8, 0x8b, 0xa6, 0x97, 0xdf, 0x1c, 0xc9, 0x5a, 0x88, 0x40, 0xde,
	0x59, 0x0f, 0xe4, 0x7d, 0x6b, 0x25, 0x32, 0xaf, 0x91, 0x1d, 0x7b, 0x70, 0xcf, 0x52, 0x38, 0xb0,
	0x6f, 0xe3, 0xe6, 0xdf, 0x43, 0x66, 0xdb, 0xcd, 0x76, 0x9b, 0xb9, 0x9f, 0x52, 0x7e, 0x07, 0x00,
	0x00, 0xff, 0xff, 0xe1, 0x0e, 0x74, 0xa1, 0x37, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TodoClient is the client API for Todo service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TodoClient interface {
	// Get implements get.
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	// List implements list.
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*TodoCollection, error)
	// Add implements add.
	Add(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*AddResponse, error)
	// Remove implements remove.
	Remove(ctx context.Context, in *RemoveRequest, opts ...grpc.CallOption) (*RemoveResponse, error)
}

type todoClient struct {
	cc *grpc.ClientConn
}

func NewTodoClient(cc *grpc.ClientConn) TodoClient {
	return &todoClient{cc}
}

func (c *todoClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/todo.Todo/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*TodoCollection, error) {
	out := new(TodoCollection)
	err := c.cc.Invoke(ctx, "/todo.Todo/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoClient) Add(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*AddResponse, error) {
	out := new(AddResponse)
	err := c.cc.Invoke(ctx, "/todo.Todo/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoClient) Remove(ctx context.Context, in *RemoveRequest, opts ...grpc.CallOption) (*RemoveResponse, error) {
	out := new(RemoveResponse)
	err := c.cc.Invoke(ctx, "/todo.Todo/Remove", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TodoServer is the server API for Todo service.
type TodoServer interface {
	// Get implements get.
	Get(context.Context, *GetRequest) (*GetResponse, error)
	// List implements list.
	List(context.Context, *ListRequest) (*TodoCollection, error)
	// Add implements add.
	Add(context.Context, *AddRequest) (*AddResponse, error)
	// Remove implements remove.
	Remove(context.Context, *RemoveRequest) (*RemoveResponse, error)
}

// UnimplementedTodoServer can be embedded to have forward compatible implementations.
type UnimplementedTodoServer struct {
}

func (*UnimplementedTodoServer) Get(ctx context.Context, req *GetRequest) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedTodoServer) List(ctx context.Context, req *ListRequest) (*TodoCollection, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (*UnimplementedTodoServer) Add(ctx context.Context, req *AddRequest) (*AddResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (*UnimplementedTodoServer) Remove(ctx context.Context, req *RemoveRequest) (*RemoveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Remove not implemented")
}

func RegisterTodoServer(s *grpc.Server, srv TodoServer) {
	s.RegisterService(&_Todo_serviceDesc, srv)
}

func _Todo_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todo.Todo/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Todo_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todo.Todo/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServer).List(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Todo_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todo.Todo/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServer).Add(ctx, req.(*AddRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Todo_Remove_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServer).Remove(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todo.Todo/Remove",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServer).Remove(ctx, req.(*RemoveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Todo_serviceDesc = grpc.ServiceDesc{
	ServiceName: "todo.Todo",
	HandlerType: (*TodoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _Todo_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _Todo_List_Handler,
		},
		{
			MethodName: "Add",
			Handler:    _Todo_Add_Handler,
		},
		{
			MethodName: "Remove",
			Handler:    _Todo_Remove_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "todo.proto",
}
