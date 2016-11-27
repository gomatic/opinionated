//go:generate goagen bootstrap -d github.com/gomatic/opinionated/design

package main

import (
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/gomatic/opinionated/app"
)

//
func service() *goa.Service {
	// Create service
	service := goa.New("opinionated")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Mount "user" controller
	c := NewUserController(service)
	app.MountUserController(service, c)

	return service
}
