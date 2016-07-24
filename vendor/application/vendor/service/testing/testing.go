package testing

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/go-kit/kit/endpoint"

	"golang.org/x/net/context"
)

// Errors

// ErrEmpty.
var ErrEmpty = errors.New("empty")

// Model

// Operations.
type Service interface {
	Test(string) (string, error)
}

//
type testService struct{}

// Test requests.
type testRequest struct {
	S string `json:"s"`
}

// Test response.
type testResponse struct {
	V   string `json:"v"`
	Err string `json:"err,omitempty"` // errors don't define JSON marshaling
}

//
func (testService) Test(s string) (string, error) {
	if s == "" {
		return "", ErrEmpty
	}
	return strings.ToUpper(s), nil
}

// Initialization

//
type testEndpoint endpoint.Endpoint

//
func New() testEndpoint {
	svc := testService{}

	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(testRequest)
		v, err := svc.Test(req.S)
		if err != nil {
			return testResponse{v, err.Error()}, nil
		}
		return testResponse{v, ""}, nil
	}
}

//
func (e testEndpoint) Post() endpoint.Endpoint {
	return endpoint.Endpoint(e)
}

//
func (e testEndpoint) Get() endpoint.Endpoint {
	return endpoint.Endpoint(e)
}

//
func (e testEndpoint) Put() endpoint.Endpoint {
	return endpoint.Endpoint(e)
}

//
func (e testEndpoint) Patch() endpoint.Endpoint {
	return endpoint.Endpoint(e)
}

//
func (e testEndpoint) Delete() endpoint.Endpoint {
	return endpoint.Endpoint(e)
}

// Request/response encoding and decoding.

// POST
func (e testEndpoint) DecodePost(_ context.Context, r *http.Request) (interface{}, error) {
	var request testRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

//
func (e testEndpoint) EncodePost(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

// GET
func (e testEndpoint) DecodeGet(_ context.Context, r *http.Request) (interface{}, error) {
	var request testRequest
	return request, nil
}

//
func (e testEndpoint) EncodeGet(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

// PUT
func (e testEndpoint) DecodePut(_ context.Context, r *http.Request) (interface{}, error) {
	var request testRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

//
func (e testEndpoint) EncodePut(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

// PATCH
func (e testEndpoint) DecodePatch(_ context.Context, r *http.Request) (interface{}, error) {
	var request testRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

//
func (e testEndpoint) EncodePatch(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

// DELETE
func (e testEndpoint) DecodeDelete(_ context.Context, r *http.Request) (interface{}, error) {
	var request testRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

//
func (e testEndpoint) EncodeDelete(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
