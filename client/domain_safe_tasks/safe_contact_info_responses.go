// Code generated by go-swagger; DO NOT EDIT.

package domain_safe_tasks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/hostwithquantum/internetx-autodns-go/models"
)

// SafeContactInfoReader is a Reader for the SafeContactInfo structure.
type SafeContactInfoReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SafeContactInfoReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSafeContactInfoOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSafeContactInfoOK creates a SafeContactInfoOK with default headers values
func NewSafeContactInfoOK() *SafeContactInfoOK {
	return &SafeContactInfoOK{}
}

/*SafeContactInfoOK handles this case with default header values.

successful operation
*/
type SafeContactInfoOK struct {
	Payload *models.JSONResponseDataDomainSafeAccount
}

func (o *SafeContactInfoOK) Error() string {
	return fmt.Sprintf("[GET /domainSafeContact/{id}][%d] safeContactInfoOK  %+v", 200, o.Payload)
}

func (o *SafeContactInfoOK) GetPayload() *models.JSONResponseDataDomainSafeAccount {
	return o.Payload
}

func (o *SafeContactInfoOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JSONResponseDataDomainSafeAccount)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}