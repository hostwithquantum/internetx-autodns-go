// Code generated by go-swagger; DO NOT EDIT.

package domain_prereg_bulk_tasks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/hostwithquantum/internetx-autodns-go/models"
)

// DomainPreregCreatesReader is a Reader for the DomainPreregCreates structure.
type DomainPreregCreatesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DomainPreregCreatesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDomainPreregCreatesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDomainPreregCreatesOK creates a DomainPreregCreatesOK with default headers values
func NewDomainPreregCreatesOK() *DomainPreregCreatesOK {
	return &DomainPreregCreatesOK{}
}

/*DomainPreregCreatesOK handles this case with default header values.

successful operation
*/
type DomainPreregCreatesOK struct {
	Payload *models.JSONResponseDataListJSONResponseDataDomainPrereg
}

func (o *DomainPreregCreatesOK) Error() string {
	return fmt.Sprintf("[POST /bulk/domainPrereg][%d] domainPreregCreatesOK  %+v", 200, o.Payload)
}

func (o *DomainPreregCreatesOK) GetPayload() *models.JSONResponseDataListJSONResponseDataDomainPrereg {
	return o.Payload
}

func (o *DomainPreregCreatesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JSONResponseDataListJSONResponseDataDomainPrereg)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}