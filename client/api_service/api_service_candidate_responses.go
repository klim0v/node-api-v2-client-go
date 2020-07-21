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

// APIServiceCandidateReader is a Reader for the APIServiceCandidate structure.
type APIServiceCandidateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *APIServiceCandidateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAPIServiceCandidateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAPIServiceCandidateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAPIServiceCandidateOK creates a APIServiceCandidateOK with default headers values
func NewAPIServiceCandidateOK() *APIServiceCandidateOK {
	return &APIServiceCandidateOK{}
}

/*APIServiceCandidateOK handles this case with default header values.

A successful response.
*/
type APIServiceCandidateOK struct {
	Payload *models.APIPbCandidateResponse
}

func (o *APIServiceCandidateOK) Error() string {
	return fmt.Sprintf("[GET /candidate/{public_key}][%d] apiServiceCandidateOK  %+v", 200, o.Payload)
}

func (o *APIServiceCandidateOK) GetPayload() *models.APIPbCandidateResponse {
	return o.Payload
}

func (o *APIServiceCandidateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIPbCandidateResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAPIServiceCandidateDefault creates a APIServiceCandidateDefault with default headers values
func NewAPIServiceCandidateDefault(code int) *APIServiceCandidateDefault {
	return &APIServiceCandidateDefault{
		_statusCode: code,
	}
}

/*APIServiceCandidateDefault handles this case with default header values.

An unexpected error response
*/
type APIServiceCandidateDefault struct {
	_statusCode int

	Payload *models.RuntimeError
}

// Code gets the status code for the Api service candidate default response
func (o *APIServiceCandidateDefault) Code() int {
	return o._statusCode
}

func (o *APIServiceCandidateDefault) Error() string {
	return fmt.Sprintf("[GET /candidate/{public_key}][%d] ApiService_Candidate default  %+v", o._statusCode, o.Payload)
}

func (o *APIServiceCandidateDefault) GetPayload() *models.RuntimeError {
	return o.Payload
}

func (o *APIServiceCandidateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
