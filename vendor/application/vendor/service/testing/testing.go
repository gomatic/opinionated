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

// Convenience to get a go-kit type back of test's private endpoint type.
func (e testEndpoint) Endpoint() endpoint.Endpoint {
	return endpoint.Endpoint(e)
}

// Request/response encoding and decoding.

//
func (e testEndpoint) Decoder(_ context.Context, r *http.Request) (interface{}, error) {
	var request testRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

//
func (e testEndpoint) Encoder(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
