// Code generated by go-swagger; DO NOT EDIT.

package domain_prereg_tasks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/hostwithquantum/internetx-autodns-go/models"
)

// DomainPreregListReader is a Reader for the DomainPreregList structure.
type DomainPreregListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DomainPreregListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDomainPreregListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDomainPreregListOK creates a DomainPreregListOK with default headers values
func NewDomainPreregListOK() *DomainPreregListOK {
	return &DomainPreregListOK{}
}

/*DomainPreregListOK handles this case with default header values.

successful operation
*/
type DomainPreregListOK struct {
	Payload *models.JSONResponseDataDomainPrereg
}

func (o *DomainPreregListOK) Error() string {
	return fmt.Sprintf("[POST /domainPrereg/_search][%d] domainPreregListOK  %+v", 200, o.Payload)
}

func (o *DomainPreregListOK) GetPayload() *models.JSONResponseDataDomainPrereg {
	return o.Payload
}

func (o *DomainPreregListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JSONResponseDataDomainPrereg)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
