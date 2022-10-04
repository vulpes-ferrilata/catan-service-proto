// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1
// source: protobuf/catan.proto

package pb

import (
	context "context"
	requests "github.com/vulpes-ferrilata/catan-service-proto/pb/requests"
	responses "github.com/vulpes-ferrilata/catan-service-proto/pb/responses"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// CatanClient is the client API for Catan service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CatanClient interface {
	FindGamesByUserID(ctx context.Context, in *requests.FindGamesByUserID, opts ...grpc.CallOption) (*responses.GameList, error)
	GetGameByIDByUserID(ctx context.Context, in *requests.GetGameByIDByUserID, opts ...grpc.CallOption) (*responses.Game, error)
	CreateGame(ctx context.Context, in *requests.CreateGame, opts ...grpc.CallOption) (*emptypb.Empty, error)
	JoinGame(ctx context.Context, in *requests.JoinGame, opts ...grpc.CallOption) (*emptypb.Empty, error)
	StartGame(ctx context.Context, in *requests.StartGame, opts ...grpc.CallOption) (*emptypb.Empty, error)
	BuildSettlementAndRoad(ctx context.Context, in *requests.BuildSettlementAndRoad, opts ...grpc.CallOption) (*emptypb.Empty, error)
	RollDices(ctx context.Context, in *requests.RollDices, opts ...grpc.CallOption) (*emptypb.Empty, error)
	MoveRobber(ctx context.Context, in *requests.MoveRobber, opts ...grpc.CallOption) (*emptypb.Empty, error)
	EndTurn(ctx context.Context, in *requests.EndTurn, opts ...grpc.CallOption) (*emptypb.Empty, error)
	BuildSettlement(ctx context.Context, in *requests.BuildSettlement, opts ...grpc.CallOption) (*emptypb.Empty, error)
	BuildRoad(ctx context.Context, in *requests.BuildRoad, opts ...grpc.CallOption) (*emptypb.Empty, error)
	UpgradeCity(ctx context.Context, in *requests.UpgradeCity, opts ...grpc.CallOption) (*emptypb.Empty, error)
	BuyDevelopmentCard(ctx context.Context, in *requests.BuyDevelopmentCard, opts ...grpc.CallOption) (*emptypb.Empty, error)
	ToggleResourceCards(ctx context.Context, in *requests.ToggleResourceCards, opts ...grpc.CallOption) (*emptypb.Empty, error)
	MaritimeTrade(ctx context.Context, in *requests.MaritimeTrade, opts ...grpc.CallOption) (*emptypb.Empty, error)
	SendTradeOffer(ctx context.Context, in *requests.SendTradeOffer, opts ...grpc.CallOption) (*emptypb.Empty, error)
	ConfirmTradeOffer(ctx context.Context, in *requests.ConfirmTradeOffer, opts ...grpc.CallOption) (*emptypb.Empty, error)
	CancelTradeOffer(ctx context.Context, in *requests.CancelTradeOffer, opts ...grpc.CallOption) (*emptypb.Empty, error)
	PlayKnightCard(ctx context.Context, in *requests.PlayKnightCard, opts ...grpc.CallOption) (*emptypb.Empty, error)
	PlayRoadBuildingCard(ctx context.Context, in *requests.PlayRoadBuildingCard, opts ...grpc.CallOption) (*emptypb.Empty, error)
	PlayYearOfPlentyCard(ctx context.Context, in *requests.PlayYearOfPlentyCard, opts ...grpc.CallOption) (*emptypb.Empty, error)
	PlayMonopolyCard(ctx context.Context, in *requests.PlayMonopolyCard, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type catanClient struct {
	cc grpc.ClientConnInterface
}

func NewCatanClient(cc grpc.ClientConnInterface) CatanClient {
	return &catanClient{cc}
}

func (c *catanClient) FindGamesByUserID(ctx context.Context, in *requests.FindGamesByUserID, opts ...grpc.CallOption) (*responses.GameList, error) {
	out := new(responses.GameList)
	err := c.cc.Invoke(ctx, "/pb.Catan/FindGamesByUserID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catanClient) GetGameByIDByUserID(ctx context.Context, in *requests.GetGameByIDByUserID, opts ...grpc.CallOption) (*responses.Game, error) {
	out := new(responses.Game)
	err := c.cc.Invoke(ctx, "/pb.Catan/GetGameByIDByUserID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catanClient) CreateGame(ctx context.Context, in *requests.CreateGame, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/pb.Catan/CreateGame", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catanClient) JoinGame(ctx context.Context, in *requests.JoinGame, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/pb.Catan/JoinGame", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catanClient) StartGame(ctx context.Context, in *requests.StartGame, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/pb.Catan/StartGame", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catanClient) BuildSettlementAndRoad(ctx context.Context, in *requests.BuildSettlementAndRoad, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/pb.Catan/BuildSettlementAndRoad", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catanClient) RollDices(ctx context.Context, in *requests.RollDices, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/pb.Catan/RollDices", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catanClient) MoveRobber(ctx context.Context, in *requests.MoveRobber, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/pb.Catan/MoveRobber", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catanClient) EndTurn(ctx context.Context, in *requests.EndTurn, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/pb.Catan/EndTurn", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catanClient) BuildSettlement(ctx context.Context, in *requests.BuildSettlement, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/pb.Catan/BuildSettlement", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catanClient) BuildRoad(ctx context.Context, in *requests.BuildRoad, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/pb.Catan/BuildRoad", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catanClient) UpgradeCity(ctx context.Context, in *requests.UpgradeCity, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/pb.Catan/UpgradeCity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catanClient) BuyDevelopmentCard(ctx context.Context, in *requests.BuyDevelopmentCard, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/pb.Catan/BuyDevelopmentCard", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catanClient) ToggleResourceCards(ctx context.Context, in *requests.ToggleResourceCards, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/pb.Catan/ToggleResourceCards", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catanClient) MaritimeTrade(ctx context.Context, in *requests.MaritimeTrade, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/pb.Catan/MaritimeTrade", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catanClient) SendTradeOffer(ctx context.Context, in *requests.SendTradeOffer, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/pb.Catan/SendTradeOffer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catanClient) ConfirmTradeOffer(ctx context.Context, in *requests.ConfirmTradeOffer, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/pb.Catan/ConfirmTradeOffer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catanClient) CancelTradeOffer(ctx context.Context, in *requests.CancelTradeOffer, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/pb.Catan/CancelTradeOffer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catanClient) PlayKnightCard(ctx context.Context, in *requests.PlayKnightCard, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/pb.Catan/PlayKnightCard", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catanClient) PlayRoadBuildingCard(ctx context.Context, in *requests.PlayRoadBuildingCard, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/pb.Catan/PlayRoadBuildingCard", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catanClient) PlayYearOfPlentyCard(ctx context.Context, in *requests.PlayYearOfPlentyCard, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/pb.Catan/PlayYearOfPlentyCard", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catanClient) PlayMonopolyCard(ctx context.Context, in *requests.PlayMonopolyCard, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/pb.Catan/PlayMonopolyCard", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CatanServer is the server API for Catan service.
// All implementations must embed UnimplementedCatanServer
// for forward compatibility
type CatanServer interface {
	FindGamesByUserID(context.Context, *requests.FindGamesByUserID) (*responses.GameList, error)
	GetGameByIDByUserID(context.Context, *requests.GetGameByIDByUserID) (*responses.Game, error)
	CreateGame(context.Context, *requests.CreateGame) (*emptypb.Empty, error)
	JoinGame(context.Context, *requests.JoinGame) (*emptypb.Empty, error)
	StartGame(context.Context, *requests.StartGame) (*emptypb.Empty, error)
	BuildSettlementAndRoad(context.Context, *requests.BuildSettlementAndRoad) (*emptypb.Empty, error)
	RollDices(context.Context, *requests.RollDices) (*emptypb.Empty, error)
	MoveRobber(context.Context, *requests.MoveRobber) (*emptypb.Empty, error)
	EndTurn(context.Context, *requests.EndTurn) (*emptypb.Empty, error)
	BuildSettlement(context.Context, *requests.BuildSettlement) (*emptypb.Empty, error)
	BuildRoad(context.Context, *requests.BuildRoad) (*emptypb.Empty, error)
	UpgradeCity(context.Context, *requests.UpgradeCity) (*emptypb.Empty, error)
	BuyDevelopmentCard(context.Context, *requests.BuyDevelopmentCard) (*emptypb.Empty, error)
	ToggleResourceCards(context.Context, *requests.ToggleResourceCards) (*emptypb.Empty, error)
	MaritimeTrade(context.Context, *requests.MaritimeTrade) (*emptypb.Empty, error)
	SendTradeOffer(context.Context, *requests.SendTradeOffer) (*emptypb.Empty, error)
	ConfirmTradeOffer(context.Context, *requests.ConfirmTradeOffer) (*emptypb.Empty, error)
	CancelTradeOffer(context.Context, *requests.CancelTradeOffer) (*emptypb.Empty, error)
	PlayKnightCard(context.Context, *requests.PlayKnightCard) (*emptypb.Empty, error)
	PlayRoadBuildingCard(context.Context, *requests.PlayRoadBuildingCard) (*emptypb.Empty, error)
	PlayYearOfPlentyCard(context.Context, *requests.PlayYearOfPlentyCard) (*emptypb.Empty, error)
	PlayMonopolyCard(context.Context, *requests.PlayMonopolyCard) (*emptypb.Empty, error)
	mustEmbedUnimplementedCatanServer()
}

// UnimplementedCatanServer must be embedded to have forward compatible implementations.
type UnimplementedCatanServer struct {
}

func (UnimplementedCatanServer) FindGamesByUserID(context.Context, *requests.FindGamesByUserID) (*responses.GameList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindGamesByUserID not implemented")
}
func (UnimplementedCatanServer) GetGameByIDByUserID(context.Context, *requests.GetGameByIDByUserID) (*responses.Game, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGameByIDByUserID not implemented")
}
func (UnimplementedCatanServer) CreateGame(context.Context, *requests.CreateGame) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateGame not implemented")
}
func (UnimplementedCatanServer) JoinGame(context.Context, *requests.JoinGame) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JoinGame not implemented")
}
func (UnimplementedCatanServer) StartGame(context.Context, *requests.StartGame) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartGame not implemented")
}
func (UnimplementedCatanServer) BuildSettlementAndRoad(context.Context, *requests.BuildSettlementAndRoad) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BuildSettlementAndRoad not implemented")
}
func (UnimplementedCatanServer) RollDices(context.Context, *requests.RollDices) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RollDices not implemented")
}
func (UnimplementedCatanServer) MoveRobber(context.Context, *requests.MoveRobber) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MoveRobber not implemented")
}
func (UnimplementedCatanServer) EndTurn(context.Context, *requests.EndTurn) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EndTurn not implemented")
}
func (UnimplementedCatanServer) BuildSettlement(context.Context, *requests.BuildSettlement) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BuildSettlement not implemented")
}
func (UnimplementedCatanServer) BuildRoad(context.Context, *requests.BuildRoad) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BuildRoad not implemented")
}
func (UnimplementedCatanServer) UpgradeCity(context.Context, *requests.UpgradeCity) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpgradeCity not implemented")
}
func (UnimplementedCatanServer) BuyDevelopmentCard(context.Context, *requests.BuyDevelopmentCard) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BuyDevelopmentCard not implemented")
}
func (UnimplementedCatanServer) ToggleResourceCards(context.Context, *requests.ToggleResourceCards) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ToggleResourceCards not implemented")
}
func (UnimplementedCatanServer) MaritimeTrade(context.Context, *requests.MaritimeTrade) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MaritimeTrade not implemented")
}
func (UnimplementedCatanServer) SendTradeOffer(context.Context, *requests.SendTradeOffer) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendTradeOffer not implemented")
}
func (UnimplementedCatanServer) ConfirmTradeOffer(context.Context, *requests.ConfirmTradeOffer) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConfirmTradeOffer not implemented")
}
func (UnimplementedCatanServer) CancelTradeOffer(context.Context, *requests.CancelTradeOffer) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelTradeOffer not implemented")
}
func (UnimplementedCatanServer) PlayKnightCard(context.Context, *requests.PlayKnightCard) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PlayKnightCard not implemented")
}
func (UnimplementedCatanServer) PlayRoadBuildingCard(context.Context, *requests.PlayRoadBuildingCard) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PlayRoadBuildingCard not implemented")
}
func (UnimplementedCatanServer) PlayYearOfPlentyCard(context.Context, *requests.PlayYearOfPlentyCard) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PlayYearOfPlentyCard not implemented")
}
func (UnimplementedCatanServer) PlayMonopolyCard(context.Context, *requests.PlayMonopolyCard) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PlayMonopolyCard not implemented")
}
func (UnimplementedCatanServer) mustEmbedUnimplementedCatanServer() {}

// UnsafeCatanServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CatanServer will
// result in compilation errors.
type UnsafeCatanServer interface {
	mustEmbedUnimplementedCatanServer()
}

func RegisterCatanServer(s grpc.ServiceRegistrar, srv CatanServer) {
	s.RegisterService(&Catan_ServiceDesc, srv)
}

func _Catan_FindGamesByUserID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(requests.FindGamesByUserID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatanServer).FindGamesByUserID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Catan/FindGamesByUserID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatanServer).FindGamesByUserID(ctx, req.(*requests.FindGamesByUserID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Catan_GetGameByIDByUserID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(requests.GetGameByIDByUserID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatanServer).GetGameByIDByUserID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Catan/GetGameByIDByUserID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatanServer).GetGameByIDByUserID(ctx, req.(*requests.GetGameByIDByUserID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Catan_CreateGame_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(requests.CreateGame)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatanServer).CreateGame(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Catan/CreateGame",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatanServer).CreateGame(ctx, req.(*requests.CreateGame))
	}
	return interceptor(ctx, in, info, handler)
}

func _Catan_JoinGame_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(requests.JoinGame)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatanServer).JoinGame(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Catan/JoinGame",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatanServer).JoinGame(ctx, req.(*requests.JoinGame))
	}
	return interceptor(ctx, in, info, handler)
}

func _Catan_StartGame_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(requests.StartGame)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatanServer).StartGame(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Catan/StartGame",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatanServer).StartGame(ctx, req.(*requests.StartGame))
	}
	return interceptor(ctx, in, info, handler)
}

func _Catan_BuildSettlementAndRoad_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(requests.BuildSettlementAndRoad)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatanServer).BuildSettlementAndRoad(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Catan/BuildSettlementAndRoad",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatanServer).BuildSettlementAndRoad(ctx, req.(*requests.BuildSettlementAndRoad))
	}
	return interceptor(ctx, in, info, handler)
}

func _Catan_RollDices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(requests.RollDices)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatanServer).RollDices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Catan/RollDices",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatanServer).RollDices(ctx, req.(*requests.RollDices))
	}
	return interceptor(ctx, in, info, handler)
}

func _Catan_MoveRobber_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(requests.MoveRobber)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatanServer).MoveRobber(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Catan/MoveRobber",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatanServer).MoveRobber(ctx, req.(*requests.MoveRobber))
	}
	return interceptor(ctx, in, info, handler)
}

func _Catan_EndTurn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(requests.EndTurn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatanServer).EndTurn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Catan/EndTurn",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatanServer).EndTurn(ctx, req.(*requests.EndTurn))
	}
	return interceptor(ctx, in, info, handler)
}

func _Catan_BuildSettlement_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(requests.BuildSettlement)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatanServer).BuildSettlement(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Catan/BuildSettlement",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatanServer).BuildSettlement(ctx, req.(*requests.BuildSettlement))
	}
	return interceptor(ctx, in, info, handler)
}

func _Catan_BuildRoad_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(requests.BuildRoad)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatanServer).BuildRoad(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Catan/BuildRoad",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatanServer).BuildRoad(ctx, req.(*requests.BuildRoad))
	}
	return interceptor(ctx, in, info, handler)
}

func _Catan_UpgradeCity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(requests.UpgradeCity)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatanServer).UpgradeCity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Catan/UpgradeCity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatanServer).UpgradeCity(ctx, req.(*requests.UpgradeCity))
	}
	return interceptor(ctx, in, info, handler)
}

func _Catan_BuyDevelopmentCard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(requests.BuyDevelopmentCard)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatanServer).BuyDevelopmentCard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Catan/BuyDevelopmentCard",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatanServer).BuyDevelopmentCard(ctx, req.(*requests.BuyDevelopmentCard))
	}
	return interceptor(ctx, in, info, handler)
}

func _Catan_ToggleResourceCards_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(requests.ToggleResourceCards)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatanServer).ToggleResourceCards(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Catan/ToggleResourceCards",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatanServer).ToggleResourceCards(ctx, req.(*requests.ToggleResourceCards))
	}
	return interceptor(ctx, in, info, handler)
}

func _Catan_MaritimeTrade_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(requests.MaritimeTrade)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatanServer).MaritimeTrade(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Catan/MaritimeTrade",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatanServer).MaritimeTrade(ctx, req.(*requests.MaritimeTrade))
	}
	return interceptor(ctx, in, info, handler)
}

func _Catan_SendTradeOffer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(requests.SendTradeOffer)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatanServer).SendTradeOffer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Catan/SendTradeOffer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatanServer).SendTradeOffer(ctx, req.(*requests.SendTradeOffer))
	}
	return interceptor(ctx, in, info, handler)
}

func _Catan_ConfirmTradeOffer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(requests.ConfirmTradeOffer)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatanServer).ConfirmTradeOffer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Catan/ConfirmTradeOffer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatanServer).ConfirmTradeOffer(ctx, req.(*requests.ConfirmTradeOffer))
	}
	return interceptor(ctx, in, info, handler)
}

func _Catan_CancelTradeOffer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(requests.CancelTradeOffer)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatanServer).CancelTradeOffer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Catan/CancelTradeOffer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatanServer).CancelTradeOffer(ctx, req.(*requests.CancelTradeOffer))
	}
	return interceptor(ctx, in, info, handler)
}

func _Catan_PlayKnightCard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(requests.PlayKnightCard)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatanServer).PlayKnightCard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Catan/PlayKnightCard",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatanServer).PlayKnightCard(ctx, req.(*requests.PlayKnightCard))
	}
	return interceptor(ctx, in, info, handler)
}

func _Catan_PlayRoadBuildingCard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(requests.PlayRoadBuildingCard)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatanServer).PlayRoadBuildingCard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Catan/PlayRoadBuildingCard",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatanServer).PlayRoadBuildingCard(ctx, req.(*requests.PlayRoadBuildingCard))
	}
	return interceptor(ctx, in, info, handler)
}

func _Catan_PlayYearOfPlentyCard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(requests.PlayYearOfPlentyCard)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatanServer).PlayYearOfPlentyCard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Catan/PlayYearOfPlentyCard",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatanServer).PlayYearOfPlentyCard(ctx, req.(*requests.PlayYearOfPlentyCard))
	}
	return interceptor(ctx, in, info, handler)
}

func _Catan_PlayMonopolyCard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(requests.PlayMonopolyCard)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatanServer).PlayMonopolyCard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Catan/PlayMonopolyCard",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatanServer).PlayMonopolyCard(ctx, req.(*requests.PlayMonopolyCard))
	}
	return interceptor(ctx, in, info, handler)
}

// Catan_ServiceDesc is the grpc.ServiceDesc for Catan service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Catan_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Catan",
	HandlerType: (*CatanServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FindGamesByUserID",
			Handler:    _Catan_FindGamesByUserID_Handler,
		},
		{
			MethodName: "GetGameByIDByUserID",
			Handler:    _Catan_GetGameByIDByUserID_Handler,
		},
		{
			MethodName: "CreateGame",
			Handler:    _Catan_CreateGame_Handler,
		},
		{
			MethodName: "JoinGame",
			Handler:    _Catan_JoinGame_Handler,
		},
		{
			MethodName: "StartGame",
			Handler:    _Catan_StartGame_Handler,
		},
		{
			MethodName: "BuildSettlementAndRoad",
			Handler:    _Catan_BuildSettlementAndRoad_Handler,
		},
		{
			MethodName: "RollDices",
			Handler:    _Catan_RollDices_Handler,
		},
		{
			MethodName: "MoveRobber",
			Handler:    _Catan_MoveRobber_Handler,
		},
		{
			MethodName: "EndTurn",
			Handler:    _Catan_EndTurn_Handler,
		},
		{
			MethodName: "BuildSettlement",
			Handler:    _Catan_BuildSettlement_Handler,
		},
		{
			MethodName: "BuildRoad",
			Handler:    _Catan_BuildRoad_Handler,
		},
		{
			MethodName: "UpgradeCity",
			Handler:    _Catan_UpgradeCity_Handler,
		},
		{
			MethodName: "BuyDevelopmentCard",
			Handler:    _Catan_BuyDevelopmentCard_Handler,
		},
		{
			MethodName: "ToggleResourceCards",
			Handler:    _Catan_ToggleResourceCards_Handler,
		},
		{
			MethodName: "MaritimeTrade",
			Handler:    _Catan_MaritimeTrade_Handler,
		},
		{
			MethodName: "SendTradeOffer",
			Handler:    _Catan_SendTradeOffer_Handler,
		},
		{
			MethodName: "ConfirmTradeOffer",
			Handler:    _Catan_ConfirmTradeOffer_Handler,
		},
		{
			MethodName: "CancelTradeOffer",
			Handler:    _Catan_CancelTradeOffer_Handler,
		},
		{
			MethodName: "PlayKnightCard",
			Handler:    _Catan_PlayKnightCard_Handler,
		},
		{
			MethodName: "PlayRoadBuildingCard",
			Handler:    _Catan_PlayRoadBuildingCard_Handler,
		},
		{
			MethodName: "PlayYearOfPlentyCard",
			Handler:    _Catan_PlayYearOfPlentyCard_Handler,
		},
		{
			MethodName: "PlayMonopolyCard",
			Handler:    _Catan_PlayMonopolyCard_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protobuf/catan.proto",
}
