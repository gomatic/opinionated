package login

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/go-kit/kit/endpoint"

	"golang.org/x/net/context"
)

// Errors

// ErrNoCredentials.
var ErrNoCredentials = errors.New("empty credentials")

// Model

// Login operations.
type Service interface {
	Login(string) (string, error)
}

// Private Login model.
type loginService struct{}

// Login requests.
type loginRequest struct {
	S string `json:"s"`
}

// Login response.
type loginResponse struct {
	V   string `json:"v"`
	Err string `json:"err,omitempty"` // errors don't define JSON marshaling
}

//
func (loginService) Login(s string) (string, error) {
	if s == "" {
		return "", ErrNoCredentials
	}
	return strings.ToUpper(s), nil
}

// Initialization

//
type loginEndpoint endpoint.Endpoint

//
func New() loginEndpoint {
	svc := loginService{}

	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(loginRequest)
		v, err := svc.Login(req.S)
		if err != nil {
			return loginResponse{v, err.Error()}, nil
		}
		return loginResponse{v, ""}, nil
	}
}

// Convenience to get a go-kit type back of login's private endpoint type.
func (e loginEndpoint) Endpoint() endpoint.Endpoint {
	return endpoint.Endpoint(e)
}

// Request/response encoding and decoding.

//
func (e loginEndpoint) Decoder(_ context.Context, r *http.Request) (interface{}, error) {
	var request loginRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

//
func (e loginEndpoint) Encoder(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
