package query

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/go-kit/kit/endpoint"

	"golang.org/x/net/context"
)

// Errors

// ErrNoQuery.
var ErrNoQuery = errors.New("empty query")

// Model

// Query operations.
type Service interface {
	Query(string) (string, error)
}

// Private Query model.
type queryService struct{}

// Query requests.
type queryRequest struct {
	S string `json:"s"`
}

// Query response.
type queryResponse struct {
	V   string `json:"v"`
	Err string `json:"err,omitempty"` // errors don't define JSON marshaling
}

//
func (queryService) Query(s string) (string, error) {
	if s == "" {
		return "", ErrNoQuery
	}
	return strings.ToUpper(s), nil
}

// Initialization

//
type queryEndpoint endpoint.Endpoint

//
func New() queryEndpoint {
	svc := queryService{}

	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(queryRequest)
		v, err := svc.Query(req.S)
		if err != nil {
			return queryResponse{v, err.Error()}, nil
		}
		return queryResponse{v, ""}, nil
	}
}

// Convenience to get a go-kit type back of query's private endpoint type.
func (e queryEndpoint) Endpoint() endpoint.Endpoint {
	return endpoint.Endpoint(e)
}

// Request/response encoding and decoding.

//
func (e queryEndpoint) Decoder(_ context.Context, r *http.Request) (interface{}, error) {
	var request queryRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

//
func (e queryEndpoint) Encoder(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
