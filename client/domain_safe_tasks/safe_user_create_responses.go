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

// SafeUserCreateReader is a Reader for the SafeUserCreate structure.
type SafeUserCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SafeUserCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSafeUserCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSafeUserCreateOK creates a SafeUserCreateOK with default headers values
func NewSafeUserCreateOK() *SafeUserCreateOK {
	return &SafeUserCreateOK{}
}

/*SafeUserCreateOK handles this case with default header values.

successful operation
*/
type SafeUserCreateOK struct {
	Payload *models.JSONResponseDataDomainSafeUser
}

func (o *SafeUserCreateOK) Error() string {
	return fmt.Sprintf("[POST /domainSafeContact/user][%d] safeUserCreateOK  %+v", 200, o.Payload)
}

func (o *SafeUserCreateOK) GetPayload() *models.JSONResponseDataDomainSafeUser {
	return o.Payload
}

func (o *SafeUserCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JSONResponseDataDomainSafeUser)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}