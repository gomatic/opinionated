package http

import (
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

//
type ChainResponseEncoders func(next httptransport.EncodeResponseFunc) httptransport.EncodeResponseFunc

//
type ChainRequestDecoders func(next httptransport.DecodeRequestFunc) httptransport.DecodeRequestFunc

//
type ChainRouter func(next *mux.Router) *mux.Router
