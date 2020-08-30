// Code generated by go-swagger; DO NOT EDIT.

package domain_bulk_tasks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
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
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDomainPatchesOK creates a DomainPatchesOK with default headers values
func NewDomainPatchesOK() *DomainPatchesOK {
	return &DomainPatchesOK{}
}

/*DomainPatchesOK handles this case with default header values.

successful operation
*/
type DomainPatchesOK struct {
	Payload *models.JSONResponseDataListJSONResponseDataDomain
}

func (o *DomainPatchesOK) Error() string {
	return fmt.Sprintf("[PATCH /bulk/domain][%d] domainPatchesOK  %+v", 200, o.Payload)
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