// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/Meat-Hook/point-bank/internal/modules/session/internal/api/web/generated/models"
)

// LogoutReader is a Reader for the Logout structure.
type LogoutReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LogoutReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewLogoutNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewLogoutDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewLogoutNoContent creates a LogoutNoContent with default headers values
func NewLogoutNoContent() *LogoutNoContent {
	return &LogoutNoContent{}
}

/*LogoutNoContent handles this case with default header values.

The server successfully processed the request and is not returning any content.
*/
type LogoutNoContent struct {
}

func (o *LogoutNoContent) Error() string {
	return fmt.Sprintf("[POST /logout][%d] logoutNoContent ", 204)
}

func (o *LogoutNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewLogoutDefault creates a LogoutDefault with default headers values
func NewLogoutDefault(code int) *LogoutDefault {
	return &LogoutDefault{
		_statusCode: code,
	}
}

/*LogoutDefault handles this case with default header values.

Generic error response.
*/
type LogoutDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the logout default response
func (o *LogoutDefault) Code() int {
	return o._statusCode
}

func (o *LogoutDefault) Error() string {
	return fmt.Sprintf("[POST /logout][%d] logout default  %+v", o._statusCode, o.Payload)
}

func (o *LogoutDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *LogoutDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
