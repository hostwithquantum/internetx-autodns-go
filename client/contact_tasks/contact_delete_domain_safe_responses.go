// Code generated by go-swagger; DO NOT EDIT.

package contact_tasks

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

// ContactDeleteDomainSafeReader is a Reader for the ContactDeleteDomainSafe structure.
type ContactDeleteDomainSafeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ContactDeleteDomainSafeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewContactDeleteDomainSafeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[DELETE /contact/{id}/_domainSafe] contactDeleteDomainSafe", response, response.Code())
	}
}

// NewContactDeleteDomainSafeOK creates a ContactDeleteDomainSafeOK with default headers values
func NewContactDeleteDomainSafeOK() *ContactDeleteDomainSafeOK {
	return &ContactDeleteDomainSafeOK{}
}

/*
ContactDeleteDomainSafeOK describes a response with status code 200, with default header values.

successful operation
*/
type ContactDeleteDomainSafeOK struct {
	Payload *models.JSONResponseDataJSONNoData
}

// IsSuccess returns true when this contact delete domain safe o k response has a 2xx status code
func (o *ContactDeleteDomainSafeOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this contact delete domain safe o k response has a 3xx status code
func (o *ContactDeleteDomainSafeOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this contact delete domain safe o k response has a 4xx status code
func (o *ContactDeleteDomainSafeOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this contact delete domain safe o k response has a 5xx status code
func (o *ContactDeleteDomainSafeOK) IsServerError() bool {
	return false
}

// IsCode returns true when this contact delete domain safe o k response a status code equal to that given
func (o *ContactDeleteDomainSafeOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the contact delete domain safe o k response
func (o *ContactDeleteDomainSafeOK) Code() int {
	return 200
}

func (o *ContactDeleteDomainSafeOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /contact/{id}/_domainSafe][%d] contactDeleteDomainSafeOK %s", 200, payload)
}

func (o *ContactDeleteDomainSafeOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /contact/{id}/_domainSafe][%d] contactDeleteDomainSafeOK %s", 200, payload)
}

func (o *ContactDeleteDomainSafeOK) GetPayload() *models.JSONResponseDataJSONNoData {
	return o.Payload
}

func (o *ContactDeleteDomainSafeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JSONResponseDataJSONNoData)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
