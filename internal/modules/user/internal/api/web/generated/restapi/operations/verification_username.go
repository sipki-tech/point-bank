// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/Meat-Hook/point-bank/internal/modules/user/internal/api/web/generated/models"
)

// VerificationUsernameHandlerFunc turns a function with the right signature into a verification username handler
type VerificationUsernameHandlerFunc func(VerificationUsernameParams) VerificationUsernameResponder

// Handle executing the request and returning a response
func (fn VerificationUsernameHandlerFunc) Handle(params VerificationUsernameParams) VerificationUsernameResponder {
	return fn(params)
}

// VerificationUsernameHandler interface for that can handle valid verification username params
type VerificationUsernameHandler interface {
	Handle(VerificationUsernameParams) VerificationUsernameResponder
}

// NewVerificationUsername creates a new http.Handler for the verification username operation
func NewVerificationUsername(ctx *middleware.Context, handler VerificationUsernameHandler) *VerificationUsername {
	return &VerificationUsername{Context: ctx, Handler: handler}
}

/*VerificationUsername swagger:route POST /username/verification verificationUsername

VerificationUsername verification username API

*/
type VerificationUsername struct {
	Context *middleware.Context
	Handler VerificationUsernameHandler
}

func (o *VerificationUsername) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewVerificationUsernameParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// VerificationUsernameBody verification username body
//
// swagger:model VerificationUsernameBody
type VerificationUsernameBody struct {

	// username
	// Required: true
	Username models.Username `json:"username"`
}

// Validate validates this verification username body
func (o *VerificationUsernameBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateUsername(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *VerificationUsernameBody) validateUsername(formats strfmt.Registry) error {

	if err := o.Username.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("args" + "." + "username")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *VerificationUsernameBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *VerificationUsernameBody) UnmarshalBinary(b []byte) error {
	var res VerificationUsernameBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
