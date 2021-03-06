// Code generated by go-swagger; DO NOT EDIT.

package api_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/klim0v/node-api-v2-client-go/models"
)

// APIServiceSubscribeReader is a Reader for the APIServiceSubscribe structure.
type APIServiceSubscribeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *APIServiceSubscribeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAPIServiceSubscribeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAPIServiceSubscribeDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAPIServiceSubscribeOK creates a APIServiceSubscribeOK with default headers values
func NewAPIServiceSubscribeOK() *APIServiceSubscribeOK {
	return &APIServiceSubscribeOK{}
}

/*APIServiceSubscribeOK handles this case with default header values.

A successful response.(streaming responses)
*/
type APIServiceSubscribeOK struct {
	Payload *APIServiceSubscribeOKBody
}

func (o *APIServiceSubscribeOK) Error() string {
	return fmt.Sprintf("[GET /subscribe][%d] apiServiceSubscribeOK  %+v", 200, o.Payload)
}

func (o *APIServiceSubscribeOK) GetPayload() *APIServiceSubscribeOKBody {
	return o.Payload
}

func (o *APIServiceSubscribeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(APIServiceSubscribeOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAPIServiceSubscribeDefault creates a APIServiceSubscribeDefault with default headers values
func NewAPIServiceSubscribeDefault(code int) *APIServiceSubscribeDefault {
	return &APIServiceSubscribeDefault{
		_statusCode: code,
	}
}

/*APIServiceSubscribeDefault handles this case with default header values.

An unexpected error response
*/
type APIServiceSubscribeDefault struct {
	_statusCode int

	Payload *models.RuntimeError
}

// Code gets the status code for the Api service subscribe default response
func (o *APIServiceSubscribeDefault) Code() int {
	return o._statusCode
}

func (o *APIServiceSubscribeDefault) Error() string {
	return fmt.Sprintf("[GET /subscribe][%d] ApiService_Subscribe default  %+v", o._statusCode, o.Payload)
}

func (o *APIServiceSubscribeDefault) GetPayload() *models.RuntimeError {
	return o.Payload
}

func (o *APIServiceSubscribeDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*APIServiceSubscribeOKBody Stream result of api_pbSubscribeResponse
swagger:model APIServiceSubscribeOKBody
*/
type APIServiceSubscribeOKBody struct {

	// error
	Error *models.RuntimeStreamError `json:"error,omitempty"`

	// result
	Result *models.APIPbSubscribeResponse `json:"result,omitempty"`
}

// Validate validates this API service subscribe o k body
func (o *APIServiceSubscribeOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateError(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateResult(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *APIServiceSubscribeOKBody) validateError(formats strfmt.Registry) error {

	if swag.IsZero(o.Error) { // not required
		return nil
	}

	if o.Error != nil {
		if err := o.Error.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("apiServiceSubscribeOK" + "." + "error")
			}
			return err
		}
	}

	return nil
}

func (o *APIServiceSubscribeOKBody) validateResult(formats strfmt.Registry) error {

	if swag.IsZero(o.Result) { // not required
		return nil
	}

	if o.Result != nil {
		if err := o.Result.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("apiServiceSubscribeOK" + "." + "result")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *APIServiceSubscribeOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *APIServiceSubscribeOKBody) UnmarshalBinary(b []byte) error {
	var res APIServiceSubscribeOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
