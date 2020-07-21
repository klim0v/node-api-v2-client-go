// Code generated by go-swagger; DO NOT EDIT.

package api_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/klim0v/node-api-v2-client-go/models"
)

// APIServiceStatusReader is a Reader for the APIServiceStatus structure.
type APIServiceStatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *APIServiceStatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAPIServiceStatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAPIServiceStatusDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAPIServiceStatusOK creates a APIServiceStatusOK with default headers values
func NewAPIServiceStatusOK() *APIServiceStatusOK {
	return &APIServiceStatusOK{}
}

/*APIServiceStatusOK handles this case with default header values.

A successful response.
*/
type APIServiceStatusOK struct {
	Payload *models.APIPbStatusResponse
}

func (o *APIServiceStatusOK) Error() string {
	return fmt.Sprintf("[GET /status][%d] apiServiceStatusOK  %+v", 200, o.Payload)
}

func (o *APIServiceStatusOK) GetPayload() *models.APIPbStatusResponse {
	return o.Payload
}

func (o *APIServiceStatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIPbStatusResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAPIServiceStatusDefault creates a APIServiceStatusDefault with default headers values
func NewAPIServiceStatusDefault(code int) *APIServiceStatusDefault {
	return &APIServiceStatusDefault{
		_statusCode: code,
	}
}

/*APIServiceStatusDefault handles this case with default header values.

An unexpected error response
*/
type APIServiceStatusDefault struct {
	_statusCode int

	Payload *models.RuntimeError
}

// Code gets the status code for the Api service status default response
func (o *APIServiceStatusDefault) Code() int {
	return o._statusCode
}

func (o *APIServiceStatusDefault) Error() string {
	return fmt.Sprintf("[GET /status][%d] ApiService_Status default  %+v", o._statusCode, o.Payload)
}

func (o *APIServiceStatusDefault) GetPayload() *models.RuntimeError {
	return o.Payload
}

func (o *APIServiceStatusDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
