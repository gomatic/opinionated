package http

import httptransport "github.com/go-kit/kit/transport/http"

//
type ChainResponseEncoders func(next httptransport.EncodeResponseFunc) httptransport.EncodeResponseFunc

//
type ChainRequestDecoders func(next httptransport.DecodeRequestFunc) httptransport.DecodeRequestFunc
