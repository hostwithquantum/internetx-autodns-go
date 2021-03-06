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

// VerificationInfoWithReferenceReader is a Reader for the VerificationInfoWithReference structure.
type VerificationInfoWithReferenceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VerificationInfoWithReferenceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewVerificationInfoWithReferenceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewVerificationInfoWithReferenceOK creates a VerificationInfoWithReferenceOK with default headers values
func NewVerificationInfoWithReferenceOK() *VerificationInfoWithReferenceOK {
	return &VerificationInfoWithReferenceOK{}
}

/*VerificationInfoWithReferenceOK handles this case with default header values.

successful operation
*/
type VerificationInfoWithReferenceOK struct {
	Payload *models.JSONResponseDataContactVerification
}

func (o *VerificationInfoWithReferenceOK) Error() string {
	return fmt.Sprintf("[GET /contact/verification][%d] verificationInfoWithReferenceOK  %+v", 200, o.Payload)
}

func (o *VerificationInfoWithReferenceOK) GetPayload() *models.JSONResponseDataContactVerification {
	return o.Payload
}

func (o *VerificationInfoWithReferenceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JSONResponseDataContactVerification)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
