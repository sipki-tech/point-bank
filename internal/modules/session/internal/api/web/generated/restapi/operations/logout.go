// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"github.com/Meat-Hook/back-template/internal/modules/session/internal/app"
)

// LogoutHandlerFunc turns a function with the right signature into a logout handler
type LogoutHandlerFunc func(LogoutParams, *app.Session) LogoutResponder

// Handle executing the request and returning a response
func (fn LogoutHandlerFunc) Handle(params LogoutParams, principal *app.Session) LogoutResponder {
	return fn(params, principal)
}

// LogoutHandler interface for that can handle valid logout params
type LogoutHandler interface {
	Handle(LogoutParams, *app.Session) LogoutResponder
}

// NewLogout creates a new http.Handler for the logout operation
func NewLogout(ctx *middleware.Context, handler LogoutHandler) *Logout {
	return &Logout{Context: ctx, Handler: handler}
}

/*Logout swagger:route POST /logout logout

Logout for user.

*/
type Logout struct {
	Context *middleware.Context
	Handler LogoutHandler
}

func (o *Logout) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewLogoutParams()

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