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

// DomainPreregCreateReader is a Reader for the DomainPreregCreate structure.
type DomainPreregCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DomainPreregCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDomainPreregCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDomainPreregCreateOK creates a DomainPreregCreateOK with default headers values
func NewDomainPreregCreateOK() *DomainPreregCreateOK {
	return &DomainPreregCreateOK{}
}

/*DomainPreregCreateOK handles this case with default header values.

successful operation
*/
type DomainPreregCreateOK struct {
	Payload *models.JSONResponseDataVoid
}

func (o *DomainPreregCreateOK) Error() string {
	return fmt.Sprintf("[POST /domainPrereg][%d] domainPreregCreateOK  %+v", 200, o.Payload)
}

func (o *DomainPreregCreateOK) GetPayload() *models.JSONResponseDataVoid {
	return o.Payload
}

func (o *DomainPreregCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JSONResponseDataVoid)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
