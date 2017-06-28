package http

import (
	"context"
	"encoding/json"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/will7200/mdar/mda/endpoints"
)

// NewHTTPHandler returns a handler that makes a set of endpoints available on
// predefined paths.
func NewHTTPHandler(endpoints endpoints.Endpoints) http.Handler {
	m := mux.NewRouter()
	opts := []httptransport.ServerOption{}
	m.Handle("add", httptransport.NewServer(
		endpoints.AddEndpoint,
		DecodeAddRequest,
		EncodeAddResponse,
		opts...,
	))
	m.Handle("start", httptransport.NewServer(
		endpoints.StartEndpoint,
		DecodeStartRequest,
		EncodeStartResponse,
		opts...,
	))
	m.Handle("remove", httptransport.NewServer(
		endpoints.RemoveEndpoint,
		DecodeRemoveRequest,
		EncodeRemoveResponse,
		opts...,
	))
	m.Handle("change", httptransport.NewServer(
		endpoints.ChangeEndpoint,
		DecodeChangeRequest,
		EncodeChangeResponse,
		opts...,
	))
	m.Handle("get", httptransport.NewServer(
		endpoints.GetEndpoint,
		DecodeGetRequest,
		EncodeGetResponse,
		opts...,
	))
	m.Handle("list", httptransport.NewServer(
		endpoints.ListEndpoint,
		DecodeListRequest,
		EncodeListResponse,
		opts...,
	))
	m.Handle("enable", httptransport.NewServer(
		endpoints.EnableEndpoint,
		DecodeEnableRequest,
		EncodeEnableResponse,
		opts...,
	))
	m.Handle("disable", httptransport.NewServer(
		endpoints.DisableEndpoint,
		DecodeDisableRequest,
		EncodeDisableResponse,
		opts...,
	))
	m.Handle("try", httptransport.NewServer(
		endpoints.TryEndpoint,
		DecodeTryRequest,
		EncodeTryResponse,
		opts...,
	)).Methods("GET", "PUT", "DELETE")
	return m
}

// DecodeAddRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body. Primarily useful in a server.
func DecodeAddRequest(_ context.Context, r *http.Request) (req interface{}, err error) {
	req = endpoints.AddRequest{}
	err = json.NewDecoder(r.Body).Decode(&r)
	return req, err
}

// EncodeAddResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer. Primarily useful in a server.
func EncodeAddResponse(_ context.Context, w http.ResponseWriter, response interface{}) (err error) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return err
}

// DecodeStartRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body. Primarily useful in a server.
func DecodeStartRequest(_ context.Context, r *http.Request) (req interface{}, err error) {
	req = endpoints.StartRequest{}
	err = json.NewDecoder(r.Body).Decode(&r)
	return req, err
}

// EncodeStartResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer. Primarily useful in a server.
func EncodeStartResponse(_ context.Context, w http.ResponseWriter, response interface{}) (err error) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return err
}

// DecodeRemoveRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body. Primarily useful in a server.
func DecodeRemoveRequest(_ context.Context, r *http.Request) (req interface{}, err error) {
	req = endpoints.RemoveRequest{}
	err = json.NewDecoder(r.Body).Decode(&r)
	return req, err
}

// EncodeRemoveResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer. Primarily useful in a server.
func EncodeRemoveResponse(_ context.Context, w http.ResponseWriter, response interface{}) (err error) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return err
}

// DecodeChangeRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body. Primarily useful in a server.
func DecodeChangeRequest(_ context.Context, r *http.Request) (req interface{}, err error) {
	req = endpoints.ChangeRequest{}
	err = json.NewDecoder(r.Body).Decode(&r)
	return req, err
}

// EncodeChangeResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer. Primarily useful in a server.
func EncodeChangeResponse(_ context.Context, w http.ResponseWriter, response interface{}) (err error) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return err
}

// DecodeGetRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body. Primarily useful in a server.
func DecodeGetRequest(_ context.Context, r *http.Request) (req interface{}, err error) {
	req = endpoints.GetRequest{}
	err = json.NewDecoder(r.Body).Decode(&r)
	return req, err
}

// EncodeGetResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer. Primarily useful in a server.
func EncodeGetResponse(_ context.Context, w http.ResponseWriter, response interface{}) (err error) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return err
}

// DecodeListRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body. Primarily useful in a server.
func DecodeListRequest(_ context.Context, r *http.Request) (req interface{}, err error) {
	req = endpoints.ListRequest{}
	err = json.NewDecoder(r.Body).Decode(&r)
	return req, err
}

// EncodeListResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer. Primarily useful in a server.
func EncodeListResponse(_ context.Context, w http.ResponseWriter, response interface{}) (err error) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return err
}

// DecodeEnableRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body. Primarily useful in a server.
func DecodeEnableRequest(_ context.Context, r *http.Request) (req interface{}, err error) {
	req = endpoints.EnableRequest{}
	err = json.NewDecoder(r.Body).Decode(&r)
	return req, err
}

// EncodeEnableResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer. Primarily useful in a server.
func EncodeEnableResponse(_ context.Context, w http.ResponseWriter, response interface{}) (err error) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return err
}

// DecodeDisableRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body. Primarily useful in a server.
func DecodeDisableRequest(_ context.Context, r *http.Request) (req interface{}, err error) {
	req = endpoints.DisableRequest{}
	err = json.NewDecoder(r.Body).Decode(&r)
	return req, err
}

// EncodeDisableResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer. Primarily useful in a server.
func EncodeDisableResponse(_ context.Context, w http.ResponseWriter, response interface{}) (err error) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return err
}

// DecodeTryRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body. Primarily useful in a server.
func DecodeTryRequest(_ context.Context, r *http.Request) (req interface{}, err error) {
	req = endpoints.TryRequest{}
	err = json.NewDecoder(r.Body).Decode(&r)
	return req, err
}

// EncodeTryResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer. Primarily useful in a server.
func EncodeTryResponse(_ context.Context, w http.ResponseWriter, response interface{}) (err error) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return err
}
