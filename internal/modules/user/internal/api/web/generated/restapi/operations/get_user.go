// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"github.com/Meat-Hook/back-template/internal/modules/user/internal/app"
)

// GetUserHandlerFunc turns a function with the right signature into a get user handler
type GetUserHandlerFunc func(GetUserParams, *app.Session) GetUserResponder

// Handle executing the request and returning a response
func (fn GetUserHandlerFunc) Handle(params GetUserParams, principal *app.Session) GetUserResponder {
	return fn(params, principal)
}

// GetUserHandler interface for that can handle valid get user params
type GetUserHandler interface {
	Handle(GetUserParams, *app.Session) GetUserResponder
}

// NewGetUser creates a new http.Handler for the get user operation
func NewGetUser(ctx *middleware.Context, handler GetUserHandler) *GetUser {
	return &GetUser{Context: ctx, Handler: handler}
}

/*GetUser swagger:route GET /user getUser

Open user profile by id. If id not set returns self info.

*/
type GetUser struct {
	Context *middleware.Context
	Handler GetUserHandler
}

func (o *GetUser) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetUserParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal *app.Session
	if uprinc != nil {
		principal = uprinc.(*app.Session) // this is really a app.Session, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}