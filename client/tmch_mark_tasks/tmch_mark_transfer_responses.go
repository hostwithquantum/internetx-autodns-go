// Code generated by go-swagger; DO NOT EDIT.

package tmch_mark_tasks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/hostwithquantum/internetx-autodns-go/models"
)

// TmchMarkTransferReader is a Reader for the TmchMarkTransfer structure.
type TmchMarkTransferReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TmchMarkTransferReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTmchMarkTransferOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PUT /tmchMark/{reference}/_transfer] tmchMarkTransfer", response, response.Code())
	}
}

// NewTmchMarkTransferOK creates a TmchMarkTransferOK with default headers values
func NewTmchMarkTransferOK() *TmchMarkTransferOK {
	return &TmchMarkTransferOK{}
}

/*
TmchMarkTransferOK describes a response with status code 200, with default header values.

successful operation
*/
type TmchMarkTransferOK struct {
	Payload *models.JSONResponseDataVoid
}

// IsSuccess returns true when this tmch mark transfer o k response has a 2xx status code
func (o *TmchMarkTransferOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this tmch mark transfer o k response has a 3xx status code
func (o *TmchMarkTransferOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this tmch mark transfer o k response has a 4xx status code
func (o *TmchMarkTransferOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this tmch mark transfer o k response has a 5xx status code
func (o *TmchMarkTransferOK) IsServerError() bool {
	return false
}

// IsCode returns true when this tmch mark transfer o k response a status code equal to that given
func (o *TmchMarkTransferOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the tmch mark transfer o k response
func (o *TmchMarkTransferOK) Code() int {
	return 200
}

func (o *TmchMarkTransferOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /tmchMark/{reference}/_transfer][%d] tmchMarkTransferOK %s", 200, payload)
}

func (o *TmchMarkTransferOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /tmchMark/{reference}/_transfer][%d] tmchMarkTransferOK %s", 200, payload)
}

func (o *TmchMarkTransferOK) GetPayload() *models.JSONResponseDataVoid {
	return o.Payload
}

func (o *TmchMarkTransferOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JSONResponseDataVoid)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
