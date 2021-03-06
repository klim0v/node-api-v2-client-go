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

// APIServiceEstimateCoinBuyReader is a Reader for the APIServiceEstimateCoinBuy structure.
type APIServiceEstimateCoinBuyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *APIServiceEstimateCoinBuyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAPIServiceEstimateCoinBuyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAPIServiceEstimateCoinBuyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAPIServiceEstimateCoinBuyOK creates a APIServiceEstimateCoinBuyOK with default headers values
func NewAPIServiceEstimateCoinBuyOK() *APIServiceEstimateCoinBuyOK {
	return &APIServiceEstimateCoinBuyOK{}
}

/*APIServiceEstimateCoinBuyOK handles this case with default header values.

A successful response.
*/
type APIServiceEstimateCoinBuyOK struct {
	Payload *models.APIPbEstimateCoinBuyResponse
}

func (o *APIServiceEstimateCoinBuyOK) Error() string {
	return fmt.Sprintf("[GET /estimate_coin_buy][%d] apiServiceEstimateCoinBuyOK  %+v", 200, o.Payload)
}

func (o *APIServiceEstimateCoinBuyOK) GetPayload() *models.APIPbEstimateCoinBuyResponse {
	return o.Payload
}

func (o *APIServiceEstimateCoinBuyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIPbEstimateCoinBuyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAPIServiceEstimateCoinBuyDefault creates a APIServiceEstimateCoinBuyDefault with default headers values
func NewAPIServiceEstimateCoinBuyDefault(code int) *APIServiceEstimateCoinBuyDefault {
	return &APIServiceEstimateCoinBuyDefault{
		_statusCode: code,
	}
}

/*APIServiceEstimateCoinBuyDefault handles this case with default header values.

An unexpected error response
*/
type APIServiceEstimateCoinBuyDefault struct {
	_statusCode int

	Payload *models.RuntimeError
}

// Code gets the status code for the Api service estimate coin buy default response
func (o *APIServiceEstimateCoinBuyDefault) Code() int {
	return o._statusCode
}

func (o *APIServiceEstimateCoinBuyDefault) Error() string {
	return fmt.Sprintf("[GET /estimate_coin_buy][%d] ApiService_EstimateCoinBuy default  %+v", o._statusCode, o.Payload)
}

func (o *APIServiceEstimateCoinBuyDefault) GetPayload() *models.RuntimeError {
	return o.Payload
}

func (o *APIServiceEstimateCoinBuyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
