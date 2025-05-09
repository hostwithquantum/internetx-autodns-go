// Code generated by go-swagger; DO NOT EDIT.

package customer_account_tasks

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

// ContextHostCreateReader is a Reader for the ContextHostCreate structure.
type ContextHostCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ContextHostCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewContextHostCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[POST /account] contextHostCreate", response, response.Code())
	}
}

// NewContextHostCreateOK creates a ContextHostCreateOK with default headers values
func NewContextHostCreateOK() *ContextHostCreateOK {
	return &ContextHostCreateOK{}
}

/*
ContextHostCreateOK describes a response with status code 200, with default header values.

successful operation
*/
type ContextHostCreateOK struct {
	Payload *models.JSONResponseDataContextHost
}

// IsSuccess returns true when this context host create o k response has a 2xx status code
func (o *ContextHostCreateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this context host create o k response has a 3xx status code
func (o *ContextHostCreateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this context host create o k response has a 4xx status code
func (o *ContextHostCreateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this context host create o k response has a 5xx status code
func (o *ContextHostCreateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this context host create o k response a status code equal to that given
func (o *ContextHostCreateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the context host create o k response
func (o *ContextHostCreateOK) Code() int {
	return 200
}

func (o *ContextHostCreateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /account][%d] contextHostCreateOK %s", 200, payload)
}

func (o *ContextHostCreateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /account][%d] contextHostCreateOK %s", 200, payload)
}

func (o *ContextHostCreateOK) GetPayload() *models.JSONResponseDataContextHost {
	return o.Payload
}

func (o *ContextHostCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JSONResponseDataContextHost)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
