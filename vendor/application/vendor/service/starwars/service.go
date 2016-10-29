package starwars

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	"github.com/gorilla/mux"
	"github.com/graphql-go/graphql"
	"golang.org/x/net/context"
)

// Errors

// ErrNoQuery.
var ErrNoQuery = errors.New("empty starwars")

// Model

// Query operations.
type Service interface {
	Query(string) (*graphql.Result, error)
}

// Private Query model.
type starwarsService struct{}

//
func (starwarsService) Query(query string) (*graphql.Result, error) {
	params := graphql.Params{Schema: StarWarsSchema, RequestString: query}
	r := graphql.Do(params)
	if len(r.Errors) > 0 {
		return nil, fmt.Errorf("failed to execute graphql operation, errors: %+v", r.Errors)
	}
	return r, nil
}

// Initialization

//
type starwarsEndpoint endpoint.Endpoint

//
func New() (endpoint starwarsEndpoint, err error) {
	svc := starwarsService{}

	endpoint = func(ctx context.Context, request interface{}) (interface{}, error) {
		return svc.Query(request.(string))
	}

	return endpoint, nil
}

// Convenience to get a go-kit type back of starwars's private endpoint type.
func (e starwarsEndpoint) Endpoint() endpoint.Endpoint {
	return endpoint.Endpoint(e)
}

// Request/response encoding and decoding.

//
func (e starwarsEndpoint) Decoder(ctx context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	if q, exists := vars["query"]; exists && q != "" {
		return q, nil
	}
	query := r.URL.Query()
	if q := query.Get("query"); q != "" {
		return string(q), nil
	}
	if q, err := ioutil.ReadAll(r.Body); err != nil {
		return nil, err
	} else if len(q) != 0 {
		return string(q), nil
	}
	return nil, fmt.Errorf("Requires a query")
}

//
func (e starwarsEndpoint) Encoder(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
