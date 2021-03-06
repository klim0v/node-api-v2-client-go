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

// APIServiceEstimateTxCommission2Reader is a Reader for the APIServiceEstimateTxCommission2 structure.
type APIServiceEstimateTxCommission2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *APIServiceEstimateTxCommission2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAPIServiceEstimateTxCommission2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAPIServiceEstimateTxCommission2Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAPIServiceEstimateTxCommission2OK creates a APIServiceEstimateTxCommission2OK with default headers values
func NewAPIServiceEstimateTxCommission2OK() *APIServiceEstimateTxCommission2OK {
	return &APIServiceEstimateTxCommission2OK{}
}

/*APIServiceEstimateTxCommission2OK handles this case with default header values.

A successful response.
*/
type APIServiceEstimateTxCommission2OK struct {
	Payload *models.APIPbEstimateTxCommissionResponse
}

func (o *APIServiceEstimateTxCommission2OK) Error() string {
	return fmt.Sprintf("[GET /estimate_tx_commission][%d] apiServiceEstimateTxCommission2OK  %+v", 200, o.Payload)
}

func (o *APIServiceEstimateTxCommission2OK) GetPayload() *models.APIPbEstimateTxCommissionResponse {
	return o.Payload
}

func (o *APIServiceEstimateTxCommission2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIPbEstimateTxCommissionResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAPIServiceEstimateTxCommission2Default creates a APIServiceEstimateTxCommission2Default with default headers values
func NewAPIServiceEstimateTxCommission2Default(code int) *APIServiceEstimateTxCommission2Default {
	return &APIServiceEstimateTxCommission2Default{
		_statusCode: code,
	}
}

/*APIServiceEstimateTxCommission2Default handles this case with default header values.

An unexpected error response
*/
type APIServiceEstimateTxCommission2Default struct {
	_statusCode int

	Payload *models.RuntimeError
}

// Code gets the status code for the Api service estimate tx commission2 default response
func (o *APIServiceEstimateTxCommission2Default) Code() int {
	return o._statusCode
}

func (o *APIServiceEstimateTxCommission2Default) Error() string {
	return fmt.Sprintf("[GET /estimate_tx_commission][%d] ApiService_EstimateTxCommission2 default  %+v", o._statusCode, o.Payload)
}

func (o *APIServiceEstimateTxCommission2Default) GetPayload() *models.RuntimeError {
	return o.Payload
}

func (o *APIServiceEstimateTxCommission2Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
