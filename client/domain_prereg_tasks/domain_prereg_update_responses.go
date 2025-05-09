// Code generated by go-swagger; DO NOT EDIT.

package domain_prereg_tasks

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

// DomainPreregUpdateReader is a Reader for the DomainPreregUpdate structure.
type DomainPreregUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DomainPreregUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDomainPreregUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PUT /domainPrereg/{reference}] domainPreregUpdate", response, response.Code())
	}
}

// NewDomainPreregUpdateOK creates a DomainPreregUpdateOK with default headers values
func NewDomainPreregUpdateOK() *DomainPreregUpdateOK {
	return &DomainPreregUpdateOK{}
}

/*
DomainPreregUpdateOK describes a response with status code 200, with default header values.

successful operation
*/
type DomainPreregUpdateOK struct {
	Payload *models.JSONResponseDataVoid
}

// IsSuccess returns true when this domain prereg update o k response has a 2xx status code
func (o *DomainPreregUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this domain prereg update o k response has a 3xx status code
func (o *DomainPreregUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this domain prereg update o k response has a 4xx status code
func (o *DomainPreregUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this domain prereg update o k response has a 5xx status code
func (o *DomainPreregUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this domain prereg update o k response a status code equal to that given
func (o *DomainPreregUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the domain prereg update o k response
func (o *DomainPreregUpdateOK) Code() int {
	return 200
}

func (o *DomainPreregUpdateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /domainPrereg/{reference}][%d] domainPreregUpdateOK %s", 200, payload)
}

func (o *DomainPreregUpdateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /domainPrereg/{reference}][%d] domainPreregUpdateOK %s", 200, payload)
}

func (o *DomainPreregUpdateOK) GetPayload() *models.JSONResponseDataVoid {
	return o.Payload
}

func (o *DomainPreregUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JSONResponseDataVoid)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
