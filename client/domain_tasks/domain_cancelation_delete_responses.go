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

// DomainCancelationDeleteReader is a Reader for the DomainCancelationDelete structure.
type DomainCancelationDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DomainCancelationDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDomainCancelationDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDomainCancelationDeleteOK creates a DomainCancelationDeleteOK with default headers values
func NewDomainCancelationDeleteOK() *DomainCancelationDeleteOK {
	return &DomainCancelationDeleteOK{}
}

/*DomainCancelationDeleteOK handles this case with default header values.

successful operation
*/
type DomainCancelationDeleteOK struct {
	Payload *models.JSONResponseDataJSONNoData
}

func (o *DomainCancelationDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /domain/{name}/cancelation][%d] domainCancelationDeleteOK  %+v", 200, o.Payload)
}

func (o *DomainCancelationDeleteOK) GetPayload() *models.JSONResponseDataJSONNoData {
	return o.Payload
}

func (o *DomainCancelationDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JSONResponseDataJSONNoData)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
