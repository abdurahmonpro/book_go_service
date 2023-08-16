// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: book_service.proto

package book_service

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	BookService_Create_FullMethodName         = "/book_service.BookService/Create"
	BookService_GetByID_FullMethodName        = "/book_service.BookService/GetByID"
	BookService_GetList_FullMethodName        = "/book_service.BookService/GetList"
	BookService_Update_FullMethodName         = "/book_service.BookService/Update"
	BookService_UpdatePatch_FullMethodName    = "/book_service.BookService/UpdatePatch"
	BookService_Delete_FullMethodName         = "/book_service.BookService/Delete"
	BookService_GetBookByTitle_FullMethodName = "/book_service.BookService/GetBookByTitle"
)

// BookServiceClient is the client API for BookService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BookServiceClient interface {
	Create(ctx context.Context, in *CreateBook, opts ...grpc.CallOption) (*Book, error)
	GetByID(ctx context.Context, in *BookPK, opts ...grpc.CallOption) (*Book, error)
	GetList(ctx context.Context, in *BookListRequest, opts ...grpc.CallOption) (*BookListResponse, error)
	Update(ctx context.Context, in *UpdateBook, opts ...grpc.CallOption) (*Book, error)
	UpdatePatch(ctx context.Context, in *UpdatePatchBook, opts ...grpc.CallOption) (*Book, error)
	Delete(ctx context.Context, in *BookPK, opts ...grpc.CallOption) (*empty.Empty, error)
	GetBookByTitle(ctx context.Context, in *BookByTitle, opts ...grpc.CallOption) (*Book, error)
}

type bookServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBookServiceClient(cc grpc.ClientConnInterface) BookServiceClient {
	return &bookServiceClient{cc}
}

func (c *bookServiceClient) Create(ctx context.Context, in *CreateBook, opts ...grpc.CallOption) (*Book, error) {
	out := new(Book)
	err := c.cc.Invoke(ctx, BookService_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookServiceClient) GetByID(ctx context.Context, in *BookPK, opts ...grpc.CallOption) (*Book, error) {
	out := new(Book)
	err := c.cc.Invoke(ctx, BookService_GetByID_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookServiceClient) GetList(ctx context.Context, in *BookListRequest, opts ...grpc.CallOption) (*BookListResponse, error) {
	out := new(BookListResponse)
	err := c.cc.Invoke(ctx, BookService_GetList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookServiceClient) Update(ctx context.Context, in *UpdateBook, opts ...grpc.CallOption) (*Book, error) {
	out := new(Book)
	err := c.cc.Invoke(ctx, BookService_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookServiceClient) UpdatePatch(ctx context.Context, in *UpdatePatchBook, opts ...grpc.CallOption) (*Book, error) {
	out := new(Book)
	err := c.cc.Invoke(ctx, BookService_UpdatePatch_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookServiceClient) Delete(ctx context.Context, in *BookPK, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, BookService_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookServiceClient) GetBookByTitle(ctx context.Context, in *BookByTitle, opts ...grpc.CallOption) (*Book, error) {
	out := new(Book)
	err := c.cc.Invoke(ctx, BookService_GetBookByTitle_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BookServiceServer is the server API for BookService service.
// All implementations must embed UnimplementedBookServiceServer
// for forward compatibility
type BookServiceServer interface {
	Create(context.Context, *CreateBook) (*Book, error)
	GetByID(context.Context, *BookPK) (*Book, error)
	GetList(context.Context, *BookListRequest) (*BookListResponse, error)
	Update(context.Context, *UpdateBook) (*Book, error)
	UpdatePatch(context.Context, *UpdatePatchBook) (*Book, error)
	Delete(context.Context, *BookPK) (*empty.Empty, error)
	GetBookByTitle(context.Context, *BookByTitle) (*Book, error)
	mustEmbedUnimplementedBookServiceServer()
}

// UnimplementedBookServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBookServiceServer struct {
}

func (UnimplementedBookServiceServer) Create(context.Context, *CreateBook) (*Book, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedBookServiceServer) GetByID(context.Context, *BookPK) (*Book, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByID not implemented")
}
func (UnimplementedBookServiceServer) GetList(context.Context, *BookListRequest) (*BookListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetList not implemented")
}
func (UnimplementedBookServiceServer) Update(context.Context, *UpdateBook) (*Book, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedBookServiceServer) UpdatePatch(context.Context, *UpdatePatchBook) (*Book, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePatch not implemented")
}
func (UnimplementedBookServiceServer) Delete(context.Context, *BookPK) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedBookServiceServer) GetBookByTitle(context.Context, *BookByTitle) (*Book, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBookByTitle not implemented")
}
func (UnimplementedBookServiceServer) mustEmbedUnimplementedBookServiceServer() {}

// UnsafeBookServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BookServiceServer will
// result in compilation errors.
type UnsafeBookServiceServer interface {
	mustEmbedUnimplementedBookServiceServer()
}

func RegisterBookServiceServer(s grpc.ServiceRegistrar, srv BookServiceServer) {
	s.RegisterService(&BookService_ServiceDesc, srv)
}

func _BookService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBook)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BookService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServiceServer).Create(ctx, req.(*CreateBook))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookService_GetByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BookPK)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServiceServer).GetByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BookService_GetByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServiceServer).GetByID(ctx, req.(*BookPK))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookService_GetList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BookListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServiceServer).GetList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BookService_GetList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServiceServer).GetList(ctx, req.(*BookListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateBook)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BookService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServiceServer).Update(ctx, req.(*UpdateBook))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookService_UpdatePatch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePatchBook)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServiceServer).UpdatePatch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BookService_UpdatePatch_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServiceServer).UpdatePatch(ctx, req.(*UpdatePatchBook))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BookPK)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BookService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServiceServer).Delete(ctx, req.(*BookPK))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookService_GetBookByTitle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BookByTitle)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServiceServer).GetBookByTitle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BookService_GetBookByTitle_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServiceServer).GetBookByTitle(ctx, req.(*BookByTitle))
	}
	return interceptor(ctx, in, info, handler)
}

// BookService_ServiceDesc is the grpc.ServiceDesc for BookService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BookService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "book_service.BookService",
	HandlerType: (*BookServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _BookService_Create_Handler,
		},
		{
			MethodName: "GetByID",
			Handler:    _BookService_GetByID_Handler,
		},
		{
			MethodName: "GetList",
			Handler:    _BookService_GetList_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _BookService_Update_Handler,
		},
		{
			MethodName: "UpdatePatch",
			Handler:    _BookService_UpdatePatch_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _BookService_Delete_Handler,
		},
		{
			MethodName: "GetBookByTitle",
			Handler:    _BookService_GetBookByTitle_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "book_service.proto",
}