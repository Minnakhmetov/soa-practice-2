// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.22.3
// source: mafia.proto

package mafia

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// MafiaClient is the client API for Mafia service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MafiaClient interface {
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (Mafia_LoginClient, error)
	EndTurn(ctx context.Context, in *EndTurnRequest, opts ...grpc.CallOption) (*EndTurnResponse, error)
	VoteAgainst(ctx context.Context, in *VoteAgainstRequest, opts ...grpc.CallOption) (*VoteAgainstResponse, error)
	Shoot(ctx context.Context, in *ShootRequest, opts ...grpc.CallOption) (*ShootResponse, error)
	Check(ctx context.Context, in *CheckRequest, opts ...grpc.CallOption) (*CheckResponse, error)
	PublishCheckResult(ctx context.Context, in *PublishCheckResultRequest, opts ...grpc.CallOption) (*PublishCheckResultResponse, error)
	GetAlivePlayers(ctx context.Context, in *GetAlivePlayersRequest, opts ...grpc.CallOption) (*GetAlivePlayersResponse, error)
}

type mafiaClient struct {
	cc grpc.ClientConnInterface
}

func NewMafiaClient(cc grpc.ClientConnInterface) MafiaClient {
	return &mafiaClient{cc}
}

func (c *mafiaClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (Mafia_LoginClient, error) {
	stream, err := c.cc.NewStream(ctx, &Mafia_ServiceDesc.Streams[0], "/mafia.Mafia/Login", opts...)
	if err != nil {
		return nil, err
	}
	x := &mafiaLoginClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Mafia_LoginClient interface {
	Recv() (*LoginResponse, error)
	grpc.ClientStream
}

type mafiaLoginClient struct {
	grpc.ClientStream
}

func (x *mafiaLoginClient) Recv() (*LoginResponse, error) {
	m := new(LoginResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *mafiaClient) EndTurn(ctx context.Context, in *EndTurnRequest, opts ...grpc.CallOption) (*EndTurnResponse, error) {
	out := new(EndTurnResponse)
	err := c.cc.Invoke(ctx, "/mafia.Mafia/EndTurn", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mafiaClient) VoteAgainst(ctx context.Context, in *VoteAgainstRequest, opts ...grpc.CallOption) (*VoteAgainstResponse, error) {
	out := new(VoteAgainstResponse)
	err := c.cc.Invoke(ctx, "/mafia.Mafia/VoteAgainst", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mafiaClient) Shoot(ctx context.Context, in *ShootRequest, opts ...grpc.CallOption) (*ShootResponse, error) {
	out := new(ShootResponse)
	err := c.cc.Invoke(ctx, "/mafia.Mafia/Shoot", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mafiaClient) Check(ctx context.Context, in *CheckRequest, opts ...grpc.CallOption) (*CheckResponse, error) {
	out := new(CheckResponse)
	err := c.cc.Invoke(ctx, "/mafia.Mafia/Check", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mafiaClient) PublishCheckResult(ctx context.Context, in *PublishCheckResultRequest, opts ...grpc.CallOption) (*PublishCheckResultResponse, error) {
	out := new(PublishCheckResultResponse)
	err := c.cc.Invoke(ctx, "/mafia.Mafia/PublishCheckResult", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mafiaClient) GetAlivePlayers(ctx context.Context, in *GetAlivePlayersRequest, opts ...grpc.CallOption) (*GetAlivePlayersResponse, error) {
	out := new(GetAlivePlayersResponse)
	err := c.cc.Invoke(ctx, "/mafia.Mafia/GetAlivePlayers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MafiaServer is the server API for Mafia service.
// All implementations must embed UnimplementedMafiaServer
// for forward compatibility
type MafiaServer interface {
	Login(*LoginRequest, Mafia_LoginServer) error
	EndTurn(context.Context, *EndTurnRequest) (*EndTurnResponse, error)
	VoteAgainst(context.Context, *VoteAgainstRequest) (*VoteAgainstResponse, error)
	Shoot(context.Context, *ShootRequest) (*ShootResponse, error)
	Check(context.Context, *CheckRequest) (*CheckResponse, error)
	PublishCheckResult(context.Context, *PublishCheckResultRequest) (*PublishCheckResultResponse, error)
	GetAlivePlayers(context.Context, *GetAlivePlayersRequest) (*GetAlivePlayersResponse, error)
	mustEmbedUnimplementedMafiaServer()
}

// UnimplementedMafiaServer must be embedded to have forward compatible implementations.
type UnimplementedMafiaServer struct {
}

func (UnimplementedMafiaServer) Login(*LoginRequest, Mafia_LoginServer) error {
	return status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedMafiaServer) EndTurn(context.Context, *EndTurnRequest) (*EndTurnResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EndTurn not implemented")
}
func (UnimplementedMafiaServer) VoteAgainst(context.Context, *VoteAgainstRequest) (*VoteAgainstResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VoteAgainst not implemented")
}
func (UnimplementedMafiaServer) Shoot(context.Context, *ShootRequest) (*ShootResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Shoot not implemented")
}
func (UnimplementedMafiaServer) Check(context.Context, *CheckRequest) (*CheckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Check not implemented")
}
func (UnimplementedMafiaServer) PublishCheckResult(context.Context, *PublishCheckResultRequest) (*PublishCheckResultResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PublishCheckResult not implemented")
}
func (UnimplementedMafiaServer) GetAlivePlayers(context.Context, *GetAlivePlayersRequest) (*GetAlivePlayersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAlivePlayers not implemented")
}
func (UnimplementedMafiaServer) mustEmbedUnimplementedMafiaServer() {}

// UnsafeMafiaServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MafiaServer will
// result in compilation errors.
type UnsafeMafiaServer interface {
	mustEmbedUnimplementedMafiaServer()
}

func RegisterMafiaServer(s grpc.ServiceRegistrar, srv MafiaServer) {
	s.RegisterService(&Mafia_ServiceDesc, srv)
}

func _Mafia_Login_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(LoginRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MafiaServer).Login(m, &mafiaLoginServer{stream})
}

type Mafia_LoginServer interface {
	Send(*LoginResponse) error
	grpc.ServerStream
}

type mafiaLoginServer struct {
	grpc.ServerStream
}

func (x *mafiaLoginServer) Send(m *LoginResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Mafia_EndTurn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EndTurnRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MafiaServer).EndTurn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mafia.Mafia/EndTurn",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MafiaServer).EndTurn(ctx, req.(*EndTurnRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mafia_VoteAgainst_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VoteAgainstRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MafiaServer).VoteAgainst(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mafia.Mafia/VoteAgainst",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MafiaServer).VoteAgainst(ctx, req.(*VoteAgainstRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mafia_Shoot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShootRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MafiaServer).Shoot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mafia.Mafia/Shoot",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MafiaServer).Shoot(ctx, req.(*ShootRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mafia_Check_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MafiaServer).Check(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mafia.Mafia/Check",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MafiaServer).Check(ctx, req.(*CheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mafia_PublishCheckResult_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PublishCheckResultRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MafiaServer).PublishCheckResult(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mafia.Mafia/PublishCheckResult",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MafiaServer).PublishCheckResult(ctx, req.(*PublishCheckResultRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mafia_GetAlivePlayers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAlivePlayersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MafiaServer).GetAlivePlayers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mafia.Mafia/GetAlivePlayers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MafiaServer).GetAlivePlayers(ctx, req.(*GetAlivePlayersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Mafia_ServiceDesc is the grpc.ServiceDesc for Mafia service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Mafia_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "mafia.Mafia",
	HandlerType: (*MafiaServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "EndTurn",
			Handler:    _Mafia_EndTurn_Handler,
		},
		{
			MethodName: "VoteAgainst",
			Handler:    _Mafia_VoteAgainst_Handler,
		},
		{
			MethodName: "Shoot",
			Handler:    _Mafia_Shoot_Handler,
		},
		{
			MethodName: "Check",
			Handler:    _Mafia_Check_Handler,
		},
		{
			MethodName: "PublishCheckResult",
			Handler:    _Mafia_PublishCheckResult_Handler,
		},
		{
			MethodName: "GetAlivePlayers",
			Handler:    _Mafia_GetAlivePlayers_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Login",
			Handler:       _Mafia_Login_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "mafia.proto",
}
