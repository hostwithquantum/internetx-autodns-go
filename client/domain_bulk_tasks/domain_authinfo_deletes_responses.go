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

// DomainAuthinfoDeletesReader is a Reader for the DomainAuthinfoDeletes structure.
type DomainAuthinfoDeletesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DomainAuthinfoDeletesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDomainAuthinfoDeletesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDomainAuthinfoDeletesOK creates a DomainAuthinfoDeletesOK with default headers values
func NewDomainAuthinfoDeletesOK() *DomainAuthinfoDeletesOK {
	return &DomainAuthinfoDeletesOK{}
}

/*DomainAuthinfoDeletesOK handles this case with default header values.

successful operation
*/
type DomainAuthinfoDeletesOK struct {
	Payload *models.JSONResponseDataListJSONResponseDataDomain
}

func (o *DomainAuthinfoDeletesOK) Error() string {
	return fmt.Sprintf("[DELETE /bulk/domain/_authinfo1][%d] domainAuthinfoDeletesOK  %+v", 200, o.Payload)
}

func (o *DomainAuthinfoDeletesOK) GetPayload() *models.JSONResponseDataListJSONResponseDataDomain {
	return o.Payload
}

func (o *DomainAuthinfoDeletesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JSONResponseDataListJSONResponseDataDomain)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
