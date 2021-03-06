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

// APIServiceEstimateCoinSellAllReader is a Reader for the APIServiceEstimateCoinSellAll structure.
type APIServiceEstimateCoinSellAllReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *APIServiceEstimateCoinSellAllReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAPIServiceEstimateCoinSellAllOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAPIServiceEstimateCoinSellAllDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAPIServiceEstimateCoinSellAllOK creates a APIServiceEstimateCoinSellAllOK with default headers values
func NewAPIServiceEstimateCoinSellAllOK() *APIServiceEstimateCoinSellAllOK {
	return &APIServiceEstimateCoinSellAllOK{}
}

/*APIServiceEstimateCoinSellAllOK handles this case with default header values.

A successful response.
*/
type APIServiceEstimateCoinSellAllOK struct {
	Payload *models.APIPbEstimateCoinSellAllResponse
}

func (o *APIServiceEstimateCoinSellAllOK) Error() string {
	return fmt.Sprintf("[GET /estimate_coin_sell_all][%d] apiServiceEstimateCoinSellAllOK  %+v", 200, o.Payload)
}

func (o *APIServiceEstimateCoinSellAllOK) GetPayload() *models.APIPbEstimateCoinSellAllResponse {
	return o.Payload
}

func (o *APIServiceEstimateCoinSellAllOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIPbEstimateCoinSellAllResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAPIServiceEstimateCoinSellAllDefault creates a APIServiceEstimateCoinSellAllDefault with default headers values
func NewAPIServiceEstimateCoinSellAllDefault(code int) *APIServiceEstimateCoinSellAllDefault {
	return &APIServiceEstimateCoinSellAllDefault{
		_statusCode: code,
	}
}

/*APIServiceEstimateCoinSellAllDefault handles this case with default header values.

An unexpected error response
*/
type APIServiceEstimateCoinSellAllDefault struct {
	_statusCode int

	Payload *models.RuntimeError
}

// Code gets the status code for the Api service estimate coin sell all default response
func (o *APIServiceEstimateCoinSellAllDefault) Code() int {
	return o._statusCode
}

func (o *APIServiceEstimateCoinSellAllDefault) Error() string {
	return fmt.Sprintf("[GET /estimate_coin_sell_all][%d] ApiService_EstimateCoinSellAll default  %+v", o._statusCode, o.Payload)
}

func (o *APIServiceEstimateCoinSellAllDefault) GetPayload() *models.RuntimeError {
	return o.Payload
}

func (o *APIServiceEstimateCoinSellAllDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
