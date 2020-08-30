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

// DomainInfoReader is a Reader for the DomainInfo structure.
type DomainInfoReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DomainInfoReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDomainInfoOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDomainInfoOK creates a DomainInfoOK with default headers values
func NewDomainInfoOK() *DomainInfoOK {
	return &DomainInfoOK{}
}

/*DomainInfoOK handles this case with default header values.

successful operation
*/
type DomainInfoOK struct {
	Payload *models.JSONResponseDataDomain
}

func (o *DomainInfoOK) Error() string {
	return fmt.Sprintf("[GET /domain/{name}][%d] domainInfoOK  %+v", 200, o.Payload)
}

func (o *DomainInfoOK) GetPayload() *models.JSONResponseDataDomain {
	return o.Payload
}

func (o *DomainInfoOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JSONResponseDataDomain)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
