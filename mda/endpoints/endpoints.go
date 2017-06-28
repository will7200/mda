package endpoints

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/will7200/mda/da"
	"github.com/will7200/mdar/mda/service"
)

// Endpoints collects all of the endpoints that compose an add service. It's
// meant to be used as a helper struct, to collect all of the endpoints into a
// single parameter.

type Endpoints struct {
	AddEndpoint     endpoint.Endpoint
	StartEndpoint   endpoint.Endpoint
	RemoveEndpoint  endpoint.Endpoint
	ChangeEndpoint  endpoint.Endpoint
	GetEndpoint     endpoint.Endpoint
	ListEndpoint    endpoint.Endpoint
	EnableEndpoint  endpoint.Endpoint
	DisableEndpoint endpoint.Endpoint
	TryEndpoint     endpoint.Endpoint
}
type AddRequest struct {
	Mdd da.DA
}
type AddResponse struct {
	S0 string
	E1 error
}
type StartRequest struct {
	Id string
}
type StartResponse struct {
	E0 error
}
type RemoveRequest struct {
	Id string
}
type RemoveResponse struct {
	E0 error
}
type ChangeRequest struct {
	Id  string
	Mdd da.DA
}
type ChangeResponse struct {
	E0 error
}
type GetRequest struct {
	Id string
}
type GetResponse struct {
	D   *da.DA
	Err error
}
type ListRequest struct{}
type ListResponse struct {
	Result *[]da.DA
	E      error
}
type EnableRequest struct {
	Id string
}
type EnableResponse struct {
	E0 error
}
type DisableRequest struct {
	Id string
}
type DisableResponse struct {
	E0 error
}
type TryRequest struct {
	Id string
}
type TryResponse struct {
	E0 error
}

func New(svc service.MdaService) (ep Endpoints) {
	ep.AddEndpoint = MakeAddEndpoint(svc)
	ep.StartEndpoint = MakeStartEndpoint(svc)
	ep.RemoveEndpoint = MakeRemoveEndpoint(svc)
	ep.ChangeEndpoint = MakeChangeEndpoint(svc)
	ep.GetEndpoint = MakeGetEndpoint(svc)
	ep.ListEndpoint = MakeListEndpoint(svc)
	ep.EnableEndpoint = MakeEnableEndpoint(svc)
	ep.DisableEndpoint = MakeDisableEndpoint(svc)
	ep.TryEndpoint = MakeTryEndpoint(svc)
	return ep
}

// MakeAddEndpoint returns an endpoint that invokes Add on the service.
// Primarily useful in a server.
func MakeAddEndpoint(svc service.MdaService) (ep endpoint.Endpoint) {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(AddRequest)
		s0, e1 := svc.Add(ctx, req.Mdd)
		return AddResponse{S0: s0, E1: e1}, nil
	}
}

// MakeStartEndpoint returns an endpoint that invokes Start on the service.
// Primarily useful in a server.
func MakeStartEndpoint(svc service.MdaService) (ep endpoint.Endpoint) {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(StartRequest)
		e0 := svc.Start(ctx, req.Id)
		return StartResponse{E0: e0}, nil
	}
}

// MakeRemoveEndpoint returns an endpoint that invokes Remove on the service.
// Primarily useful in a server.
func MakeRemoveEndpoint(svc service.MdaService) (ep endpoint.Endpoint) {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(RemoveRequest)
		e0 := svc.Remove(ctx, req.Id)
		return RemoveResponse{E0: e0}, nil
	}
}

// MakeChangeEndpoint returns an endpoint that invokes Change on the service.
// Primarily useful in a server.
func MakeChangeEndpoint(svc service.MdaService) (ep endpoint.Endpoint) {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(ChangeRequest)
		e0 := svc.Change(ctx, req.Id, req.Mdd)
		return ChangeResponse{E0: e0}, nil
	}
}

// MakeGetEndpoint returns an endpoint that invokes Get on the service.
// Primarily useful in a server.
func MakeGetEndpoint(svc service.MdaService) (ep endpoint.Endpoint) {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetRequest)
		d, err := svc.Get(ctx, req.Id)
		return GetResponse{D: d, Err: err}, nil
	}
}

// MakeListEndpoint returns an endpoint that invokes List on the service.
// Primarily useful in a server.
func MakeListEndpoint(svc service.MdaService) (ep endpoint.Endpoint) {
	return func(ctx context.Context, _ interface{}) (interface{}, error) {
		result, e := svc.List(ctx)
		return ListResponse{Result: result, E: e}, nil
	}
}

// MakeEnableEndpoint returns an endpoint that invokes Enable on the service.
// Primarily useful in a server.
func MakeEnableEndpoint(svc service.MdaService) (ep endpoint.Endpoint) {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(EnableRequest)
		e0 := svc.Enable(ctx, req.Id)
		return EnableResponse{E0: e0}, nil
	}
}

// MakeDisableEndpoint returns an endpoint that invokes Disable on the service.
// Primarily useful in a server.
func MakeDisableEndpoint(svc service.MdaService) (ep endpoint.Endpoint) {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DisableRequest)
		e0 := svc.Disable(ctx, req.Id)
		return DisableResponse{E0: e0}, nil
	}
}

// MakeTryEndpoint returns an endpoint that invokes Try on the service.
// Primarily useful in a server.
func MakeTryEndpoint(svc service.MdaService) (ep endpoint.Endpoint) {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(TryRequest)
		e0 := svc.Try(ctx, req.Id)
		return TryResponse{E0: e0}, nil
	}
}
