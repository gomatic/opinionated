package starwars

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
var ErrNoQuery = errors.New("empty starwars")

// Model

// Query operations.
type Service interface {
	Query(string) (string, error)
}

// Private Query model.
type starwarsService struct{}

// Query requests.
type starwarsRequest struct {
	S string `json:"s"`
}

// Query response.
type starwarsResponse struct {
	V   string `json:"v"`
	Err string `json:"err,omitempty"` // errors don't define JSON marshaling
}

//
func (starwarsService) Query(s string) (string, error) {
	if s == "" {
		return "", ErrNoQuery
	}
	return strings.ToUpper(s), nil
}

// Initialization

//
type starwarsEndpoint endpoint.Endpoint

//
func New() starwarsEndpoint {
	svc := starwarsService{}

	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(starwarsRequest)
		v, err := svc.Query(req.S)
		if err != nil {
			return starwarsResponse{v, err.Error()}, nil
		}
		return starwarsResponse{v, ""}, nil
	}
}

// Convenience to get a go-kit type back of starwars's private endpoint type.
func (e starwarsEndpoint) Endpoint() endpoint.Endpoint {
	return endpoint.Endpoint(e)
}

// Request/response encoding and decoding.

//
func (e starwarsEndpoint) Decoder(_ context.Context, r *http.Request) (interface{}, error) {
	var request starwarsRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

//
func (e starwarsEndpoint) Encoder(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
