package application

import "net/url"

//
var Settings struct {
	Debugger bool
	Verbose  bool
	Mode     string

	Addr     string
	Port     int
	Insecure bool

	Server  string
	Powered string
	Version string

	Docs *url.URL
}
