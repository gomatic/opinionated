package main

import (
	"github.com/goadesign/goa"
	"github.com/gomatic/opinionated/app"
)

// UserController implements the user resource.
type UserController struct {
	*goa.Controller
}

// NewUserController creates a user controller.
func NewUserController(service *goa.Service) *UserController {
	return &UserController{Controller: service.NewController("UserController")}
}

// Login runs the login action.
func (c *UserController) Login(ctx *app.LoginUserContext) error {
	// UserController_Login: start_implement

	// Put your logic here

	// UserController_Login: end_implement
	res := &app.JSON{}
	return ctx.OK(res)
}
