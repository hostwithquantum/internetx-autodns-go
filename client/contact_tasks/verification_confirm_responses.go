// Code generated by go-swagger; DO NOT EDIT.

package contact_tasks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/hostwithquantum/internetx-autodns-go/models"
)

// VerificationConfirmReader is a Reader for the VerificationConfirm structure.
type VerificationConfirmReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VerificationConfirmReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewVerificationConfirmOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewVerificationConfirmOK creates a VerificationConfirmOK with default headers values
func NewVerificationConfirmOK() *VerificationConfirmOK {
	return &VerificationConfirmOK{}
}

/*VerificationConfirmOK handles this case with default header values.

successful operation
*/
type VerificationConfirmOK struct {
	Payload *models.JSONResponseDataContactVerification
}

func (o *VerificationConfirmOK) Error() string {
	return fmt.Sprintf("[PUT /contact/verification/_confirm][%d] verificationConfirmOK  %+v", 200, o.Payload)
}

func (o *VerificationConfirmOK) GetPayload() *models.JSONResponseDataContactVerification {
	return o.Payload
}

func (o *VerificationConfirmOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JSONResponseDataContactVerification)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
