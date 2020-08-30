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

// DomainPreregDeleteReader is a Reader for the DomainPreregDelete structure.
type DomainPreregDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DomainPreregDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDomainPreregDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDomainPreregDeleteOK creates a DomainPreregDeleteOK with default headers values
func NewDomainPreregDeleteOK() *DomainPreregDeleteOK {
	return &DomainPreregDeleteOK{}
}

/*DomainPreregDeleteOK handles this case with default header values.

successful operation
*/
type DomainPreregDeleteOK struct {
	Payload *models.JSONResponseDataVoid
}

func (o *DomainPreregDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /domainPrereg/{reference}][%d] domainPreregDeleteOK  %+v", 200, o.Payload)
}

func (o *DomainPreregDeleteOK) GetPayload() *models.JSONResponseDataVoid {
	return o.Payload
}

func (o *DomainPreregDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JSONResponseDataVoid)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
