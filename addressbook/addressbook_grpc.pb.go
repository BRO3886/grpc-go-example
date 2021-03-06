// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package addressbook

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// AddressBookServiceClient is the client API for AddressBookService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AddressBookServiceClient interface {
	CreateAddressBook(ctx context.Context, in *AddressBook, opts ...grpc.CallOption) (*AddressBook, error)
	AddPerson(ctx context.Context, opts ...grpc.CallOption) (AddressBookService_AddPersonClient, error)
	GetAddressBook(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*AddressBook, error)
}

type addressBookServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAddressBookServiceClient(cc grpc.ClientConnInterface) AddressBookServiceClient {
	return &addressBookServiceClient{cc}
}

func (c *addressBookServiceClient) CreateAddressBook(ctx context.Context, in *AddressBook, opts ...grpc.CallOption) (*AddressBook, error) {
	out := new(AddressBook)
	err := c.cc.Invoke(ctx, "/addressbook.AddressBookService/CreateAddressBook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *addressBookServiceClient) AddPerson(ctx context.Context, opts ...grpc.CallOption) (AddressBookService_AddPersonClient, error) {
	stream, err := c.cc.NewStream(ctx, &AddressBookService_ServiceDesc.Streams[0], "/addressbook.AddressBookService/AddPerson", opts...)
	if err != nil {
		return nil, err
	}
	x := &addressBookServiceAddPersonClient{stream}
	return x, nil
}

type AddressBookService_AddPersonClient interface {
	Send(*Person) error
	Recv() (*AddressBook, error)
	grpc.ClientStream
}

type addressBookServiceAddPersonClient struct {
	grpc.ClientStream
}

func (x *addressBookServiceAddPersonClient) Send(m *Person) error {
	return x.ClientStream.SendMsg(m)
}

func (x *addressBookServiceAddPersonClient) Recv() (*AddressBook, error) {
	m := new(AddressBook)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *addressBookServiceClient) GetAddressBook(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*AddressBook, error) {
	out := new(AddressBook)
	err := c.cc.Invoke(ctx, "/addressbook.AddressBookService/GetAddressBook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AddressBookServiceServer is the server API for AddressBookService service.
// All implementations must embed UnimplementedAddressBookServiceServer
// for forward compatibility
type AddressBookServiceServer interface {
	CreateAddressBook(context.Context, *AddressBook) (*AddressBook, error)
	AddPerson(AddressBookService_AddPersonServer) error
	GetAddressBook(context.Context, *emptypb.Empty) (*AddressBook, error)
	mustEmbedUnimplementedAddressBookServiceServer()
}

// UnimplementedAddressBookServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAddressBookServiceServer struct {
}

func (UnimplementedAddressBookServiceServer) CreateAddressBook(context.Context, *AddressBook) (*AddressBook, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAddressBook not implemented")
}
func (UnimplementedAddressBookServiceServer) AddPerson(AddressBookService_AddPersonServer) error {
	return status.Errorf(codes.Unimplemented, "method AddPerson not implemented")
}
func (UnimplementedAddressBookServiceServer) GetAddressBook(context.Context, *emptypb.Empty) (*AddressBook, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAddressBook not implemented")
}
func (UnimplementedAddressBookServiceServer) mustEmbedUnimplementedAddressBookServiceServer() {}

// UnsafeAddressBookServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AddressBookServiceServer will
// result in compilation errors.
type UnsafeAddressBookServiceServer interface {
	mustEmbedUnimplementedAddressBookServiceServer()
}

func RegisterAddressBookServiceServer(s grpc.ServiceRegistrar, srv AddressBookServiceServer) {
	s.RegisterService(&AddressBookService_ServiceDesc, srv)
}

func _AddressBookService_CreateAddressBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddressBook)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AddressBookServiceServer).CreateAddressBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/addressbook.AddressBookService/CreateAddressBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AddressBookServiceServer).CreateAddressBook(ctx, req.(*AddressBook))
	}
	return interceptor(ctx, in, info, handler)
}

func _AddressBookService_AddPerson_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(AddressBookServiceServer).AddPerson(&addressBookServiceAddPersonServer{stream})
}

type AddressBookService_AddPersonServer interface {
	Send(*AddressBook) error
	Recv() (*Person, error)
	grpc.ServerStream
}

type addressBookServiceAddPersonServer struct {
	grpc.ServerStream
}

func (x *addressBookServiceAddPersonServer) Send(m *AddressBook) error {
	return x.ServerStream.SendMsg(m)
}

func (x *addressBookServiceAddPersonServer) Recv() (*Person, error) {
	m := new(Person)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _AddressBookService_GetAddressBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AddressBookServiceServer).GetAddressBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/addressbook.AddressBookService/GetAddressBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AddressBookServiceServer).GetAddressBook(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// AddressBookService_ServiceDesc is the grpc.ServiceDesc for AddressBookService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AddressBookService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "addressbook.AddressBookService",
	HandlerType: (*AddressBookServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateAddressBook",
			Handler:    _AddressBookService_CreateAddressBook_Handler,
		},
		{
			MethodName: "GetAddressBook",
			Handler:    _AddressBookService_GetAddressBook_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "AddPerson",
			Handler:       _AddressBookService_AddPerson_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "addressbook/addressbook.proto",
}
