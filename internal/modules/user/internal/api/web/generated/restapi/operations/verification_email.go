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

// VerificationEmailHandlerFunc turns a function with the right signature into a verification email handler
type VerificationEmailHandlerFunc func(VerificationEmailParams) VerificationEmailResponder

// Handle executing the request and returning a response
func (fn VerificationEmailHandlerFunc) Handle(params VerificationEmailParams) VerificationEmailResponder {
	return fn(params)
}

// VerificationEmailHandler interface for that can handle valid verification email params
type VerificationEmailHandler interface {
	Handle(VerificationEmailParams) VerificationEmailResponder
}

// NewVerificationEmail creates a new http.Handler for the verification email operation
func NewVerificationEmail(ctx *middleware.Context, handler VerificationEmailHandler) *VerificationEmail {
	return &VerificationEmail{Context: ctx, Handler: handler}
}

/*VerificationEmail swagger:route POST /email/verification verificationEmail

VerificationEmail verification email API

*/
type VerificationEmail struct {
	Context *middleware.Context
	Handler VerificationEmailHandler
}

func (o *VerificationEmail) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewVerificationEmailParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// VerificationEmailBody verification email body
//
// swagger:model VerificationEmailBody
type VerificationEmailBody struct {

	// email
	// Required: true
	// Format: email
	Email models.Email `json:"email"`
}

// Validate validates this verification email body
func (o *VerificationEmailBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateEmail(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *VerificationEmailBody) validateEmail(formats strfmt.Registry) error {

	if err := o.Email.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("args" + "." + "email")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *VerificationEmailBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *VerificationEmailBody) UnmarshalBinary(b []byte) error {
	var res VerificationEmailBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
