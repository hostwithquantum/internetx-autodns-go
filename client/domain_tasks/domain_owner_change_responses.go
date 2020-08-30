// Code generated by go-swagger; DO NOT EDIT.

package domain_tasks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/hostwithquantum/internetx-autodns-go/models"
)

// DomainOwnerChangeReader is a Reader for the DomainOwnerChange structure.
type DomainOwnerChangeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DomainOwnerChangeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDomainOwnerChangeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDomainOwnerChangeOK creates a DomainOwnerChangeOK with default headers values
func NewDomainOwnerChangeOK() *DomainOwnerChangeOK {
	return &DomainOwnerChangeOK{}
}

/*DomainOwnerChangeOK handles this case with default header values.

successful operation
*/
type DomainOwnerChangeOK struct {
	Payload *models.JSONResponseDataJob
}

func (o *DomainOwnerChangeOK) Error() string {
	return fmt.Sprintf("[POST /domain/_ownerChange][%d] domainOwnerChangeOK  %+v", 200, o.Payload)
}

func (o *DomainOwnerChangeOK) GetPayload() *models.JSONResponseDataJob {
	return o.Payload
}

func (o *DomainOwnerChangeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JSONResponseDataJob)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
