package grpc

import (
	"context"
	"errors"

	grpctransport "github.com/go-kit/kit/transport/grpc"
	"github.com/will7200/mda/mda/endpoints"
	"github.com/will7200/mda/mda/grpc/pb"
	oldcontext "golang.org/x/net/context"
)

type grpcServer struct {
	add     grpctransport.Handler
	start   grpctransport.Handler
	remove  grpctransport.Handler
	change  grpctransport.Handler
	get     grpctransport.Handler
	list    grpctransport.Handler
	enable  grpctransport.Handler
	disable grpctransport.Handler
}

// MakeGRPCServer makes a set of endpoints available as a gRPC server.
func MakeGRPCServer(endpoints endpoints.Endpoints) (req pb.MdaServer) {
	req = &grpcServer{
		add: grpctransport.NewServer(
			endpoints.AddEndpoint,
			DecodeGRPCAddRequest,
			EncodeGRPCAddResponse,
		),

		start: grpctransport.NewServer(
			endpoints.StartEndpoint,
			DecodeGRPCStartRequest,
			EncodeGRPCStartResponse,
		),

		remove: grpctransport.NewServer(
			endpoints.RemoveEndpoint,
			DecodeGRPCRemoveRequest,
			EncodeGRPCRemoveResponse,
		),

		change: grpctransport.NewServer(
			endpoints.ChangeEndpoint,
			DecodeGRPCChangeRequest,
			EncodeGRPCChangeResponse,
		),

		get: grpctransport.NewServer(
			endpoints.GetEndpoint,
			DecodeGRPCGetRequest,
			EncodeGRPCGetResponse,
		),

		list: grpctransport.NewServer(
			endpoints.ListEndpoint,
			DecodeGRPCListRequest,
			EncodeGRPCListResponse,
		),

		enable: grpctransport.NewServer(
			endpoints.EnableEndpoint,
			DecodeGRPCEnableRequest,
			EncodeGRPCEnableResponse,
		),

		disable: grpctransport.NewServer(
			endpoints.DisableEndpoint,
			DecodeGRPCDisableRequest,
			EncodeGRPCDisableResponse,
		),
	}
	return req
}

// DecodeGRPCAddRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain request. Primarily useful in a server.
// TODO: Do not forget to implement the decoder, you can find an example here :
// https://github.com/go-kit/kit/blob/master/examples/addsvc/transport_grpc.go#L62-L65
func DecodeGRPCAddRequest(_ context.Context, grpcReq interface{}) (req interface{}, err error) {
	err = errors.New("'Add' Decoder is not impelement")
	return req, err
}

// EncodeGRPCAddResponse is a transport/grpc.EncodeResponseFunc that converts a
// user-domain response to a gRPC reply. Primarily useful in a server.
// TODO: Do not forget to implement the encoder, you can find an example here :
// https://github.com/go-kit/kit/blob/master/examples/addsvc/transport_grpc.go#L62-L65
func EncodeGRPCAddResponse(_ context.Context, grpcReply interface{}) (res interface{}, err error) {
	err = errors.New("'Add' Encoder is not impelement")
	return res, err
}

func (s *grpcServer) Add(ctx oldcontext.Context, req *pb.AddRequest) (rep *pb.AddReply, err error) {
	_, rp, err := s.add.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	rep = rp.(*pb.AddReply)
	return rep, err
}

// DecodeGRPCStartRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain request. Primarily useful in a server.
// TODO: Do not forget to implement the decoder, you can find an example here :
// https://github.com/go-kit/kit/blob/master/examples/addsvc/transport_grpc.go#L62-L65
func DecodeGRPCStartRequest(_ context.Context, grpcReq interface{}) (req interface{}, err error) {
	err = errors.New("'Start' Decoder is not impelement")
	return req, err
}

// EncodeGRPCStartResponse is a transport/grpc.EncodeResponseFunc that converts a
// user-domain response to a gRPC reply. Primarily useful in a server.
// TODO: Do not forget to implement the encoder, you can find an example here :
// https://github.com/go-kit/kit/blob/master/examples/addsvc/transport_grpc.go#L62-L65
func EncodeGRPCStartResponse(_ context.Context, grpcReply interface{}) (res interface{}, err error) {
	err = errors.New("'Start' Encoder is not impelement")
	return res, err
}

func (s *grpcServer) Start(ctx oldcontext.Context, req *pb.StartRequest) (rep *pb.StartReply, err error) {
	_, rp, err := s.start.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	rep = rp.(*pb.StartReply)
	return rep, err
}

// DecodeGRPCRemoveRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain request. Primarily useful in a server.
// TODO: Do not forget to implement the decoder, you can find an example here :
// https://github.com/go-kit/kit/blob/master/examples/addsvc/transport_grpc.go#L62-L65
func DecodeGRPCRemoveRequest(_ context.Context, grpcReq interface{}) (req interface{}, err error) {
	err = errors.New("'Remove' Decoder is not impelement")
	return req, err
}

// EncodeGRPCRemoveResponse is a transport/grpc.EncodeResponseFunc that converts a
// user-domain response to a gRPC reply. Primarily useful in a server.
// TODO: Do not forget to implement the encoder, you can find an example here :
// https://github.com/go-kit/kit/blob/master/examples/addsvc/transport_grpc.go#L62-L65
func EncodeGRPCRemoveResponse(_ context.Context, grpcReply interface{}) (res interface{}, err error) {
	err = errors.New("'Remove' Encoder is not impelement")
	return res, err
}

func (s *grpcServer) Remove(ctx oldcontext.Context, req *pb.RemoveRequest) (rep *pb.RemoveReply, err error) {
	_, rp, err := s.remove.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	rep = rp.(*pb.RemoveReply)
	return rep, err
}

// DecodeGRPCChangeRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain request. Primarily useful in a server.
// TODO: Do not forget to implement the decoder, you can find an example here :
// https://github.com/go-kit/kit/blob/master/examples/addsvc/transport_grpc.go#L62-L65
func DecodeGRPCChangeRequest(_ context.Context, grpcReq interface{}) (req interface{}, err error) {
	err = errors.New("'Change' Decoder is not impelement")
	return req, err
}

// EncodeGRPCChangeResponse is a transport/grpc.EncodeResponseFunc that converts a
// user-domain response to a gRPC reply. Primarily useful in a server.
// TODO: Do not forget to implement the encoder, you can find an example here :
// https://github.com/go-kit/kit/blob/master/examples/addsvc/transport_grpc.go#L62-L65
func EncodeGRPCChangeResponse(_ context.Context, grpcReply interface{}) (res interface{}, err error) {
	err = errors.New("'Change' Encoder is not impelement")
	return res, err
}

func (s *grpcServer) Change(ctx oldcontext.Context, req *pb.ChangeRequest) (rep *pb.ChangeReply, err error) {
	_, rp, err := s.change.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	rep = rp.(*pb.ChangeReply)
	return rep, err
}

// DecodeGRPCGetRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain request. Primarily useful in a server.
// TODO: Do not forget to implement the decoder, you can find an example here :
// https://github.com/go-kit/kit/blob/master/examples/addsvc/transport_grpc.go#L62-L65
func DecodeGRPCGetRequest(_ context.Context, grpcReq interface{}) (req interface{}, err error) {
	err = errors.New("'Get' Decoder is not impelement")
	return req, err
}

// EncodeGRPCGetResponse is a transport/grpc.EncodeResponseFunc that converts a
// user-domain response to a gRPC reply. Primarily useful in a server.
// TODO: Do not forget to implement the encoder, you can find an example here :
// https://github.com/go-kit/kit/blob/master/examples/addsvc/transport_grpc.go#L62-L65
func EncodeGRPCGetResponse(_ context.Context, grpcReply interface{}) (res interface{}, err error) {
	err = errors.New("'Get' Encoder is not impelement")
	return res, err
}

func (s *grpcServer) Get(ctx oldcontext.Context, req *pb.GetRequest) (rep *pb.GetReply, err error) {
	_, rp, err := s.get.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	rep = rp.(*pb.GetReply)
	return rep, err
}

// DecodeGRPCListRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain request. Primarily useful in a server.
// TODO: Do not forget to implement the decoder, you can find an example here :
// https://github.com/go-kit/kit/blob/master/examples/addsvc/transport_grpc.go#L62-L65
func DecodeGRPCListRequest(_ context.Context, grpcReq interface{}) (req interface{}, err error) {
	err = errors.New("'List' Decoder is not impelement")
	return req, err
}

// EncodeGRPCListResponse is a transport/grpc.EncodeResponseFunc that converts a
// user-domain response to a gRPC reply. Primarily useful in a server.
// TODO: Do not forget to implement the encoder, you can find an example here :
// https://github.com/go-kit/kit/blob/master/examples/addsvc/transport_grpc.go#L62-L65
func EncodeGRPCListResponse(_ context.Context, grpcReply interface{}) (res interface{}, err error) {
	err = errors.New("'List' Encoder is not impelement")
	return res, err
}

func (s *grpcServer) List(ctx oldcontext.Context, req *pb.ListRequest) (rep *pb.ListReply, err error) {
	_, rp, err := s.list.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	rep = rp.(*pb.ListReply)
	return rep, err
}

// DecodeGRPCEnableRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain request. Primarily useful in a server.
// TODO: Do not forget to implement the decoder, you can find an example here :
// https://github.com/go-kit/kit/blob/master/examples/addsvc/transport_grpc.go#L62-L65
func DecodeGRPCEnableRequest(_ context.Context, grpcReq interface{}) (req interface{}, err error) {
	err = errors.New("'Enable' Decoder is not impelement")
	return req, err
}

// EncodeGRPCEnableResponse is a transport/grpc.EncodeResponseFunc that converts a
// user-domain response to a gRPC reply. Primarily useful in a server.
// TODO: Do not forget to implement the encoder, you can find an example here :
// https://github.com/go-kit/kit/blob/master/examples/addsvc/transport_grpc.go#L62-L65
func EncodeGRPCEnableResponse(_ context.Context, grpcReply interface{}) (res interface{}, err error) {
	err = errors.New("'Enable' Encoder is not impelement")
	return res, err
}

func (s *grpcServer) Enable(ctx oldcontext.Context, req *pb.EnableRequest) (rep *pb.EnableReply, err error) {
	_, rp, err := s.enable.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	rep = rp.(*pb.EnableReply)
	return rep, err
}

// DecodeGRPCDisableRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain request. Primarily useful in a server.
// TODO: Do not forget to implement the decoder, you can find an example here :
// https://github.com/go-kit/kit/blob/master/examples/addsvc/transport_grpc.go#L62-L65
func DecodeGRPCDisableRequest(_ context.Context, grpcReq interface{}) (req interface{}, err error) {
	err = errors.New("'Disable' Decoder is not impelement")
	return req, err
}

// EncodeGRPCDisableResponse is a transport/grpc.EncodeResponseFunc that converts a
// user-domain response to a gRPC reply. Primarily useful in a server.
// TODO: Do not forget to implement the encoder, you can find an example here :
// https://github.com/go-kit/kit/blob/master/examples/addsvc/transport_grpc.go#L62-L65
func EncodeGRPCDisableResponse(_ context.Context, grpcReply interface{}) (res interface{}, err error) {
	err = errors.New("'Disable' Encoder is not impelement")
	return res, err
}

func (s *grpcServer) Disable(ctx oldcontext.Context, req *pb.DisableRequest) (rep *pb.DisableReply, err error) {
	_, rp, err := s.disable.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	rep = rp.(*pb.DisableReply)
	return rep, err
}
