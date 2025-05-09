// Code generated by go-swagger; DO NOT EDIT.

package domain_tasks

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

// DomainPatchesReader is a Reader for the DomainPatches structure.
type DomainPatchesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DomainPatchesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDomainPatchesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PATCH /bulk/domain] domainPatches", response, response.Code())
	}
}

// NewDomainPatchesOK creates a DomainPatchesOK with default headers values
func NewDomainPatchesOK() *DomainPatchesOK {
	return &DomainPatchesOK{}
}

/*
DomainPatchesOK describes a response with status code 200, with default header values.

successful operation
*/
type DomainPatchesOK struct {
	Payload *models.JSONResponseDataListJSONResponseDataDomain
}

// IsSuccess returns true when this domain patches o k response has a 2xx status code
func (o *DomainPatchesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this domain patches o k response has a 3xx status code
func (o *DomainPatchesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this domain patches o k response has a 4xx status code
func (o *DomainPatchesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this domain patches o k response has a 5xx status code
func (o *DomainPatchesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this domain patches o k response a status code equal to that given
func (o *DomainPatchesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the domain patches o k response
func (o *DomainPatchesOK) Code() int {
	return 200
}

func (o *DomainPatchesOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /bulk/domain][%d] domainPatchesOK %s", 200, payload)
}

func (o *DomainPatchesOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /bulk/domain][%d] domainPatchesOK %s", 200, payload)
}

func (o *DomainPatchesOK) GetPayload() *models.JSONResponseDataListJSONResponseDataDomain {
	return o.Payload
}

func (o *DomainPatchesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JSONResponseDataListJSONResponseDataDomain)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
