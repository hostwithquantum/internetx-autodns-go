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

// VerificationHistoryInfoReader is a Reader for the VerificationHistoryInfo structure.
type VerificationHistoryInfoReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VerificationHistoryInfoReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewVerificationHistoryInfoOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewVerificationHistoryInfoOK creates a VerificationHistoryInfoOK with default headers values
func NewVerificationHistoryInfoOK() *VerificationHistoryInfoOK {
	return &VerificationHistoryInfoOK{}
}

/*VerificationHistoryInfoOK handles this case with default header values.

successful operation
*/
type VerificationHistoryInfoOK struct {
	Payload *models.JSONResponseDataContactVerification
}

func (o *VerificationHistoryInfoOK) Error() string {
	return fmt.Sprintf("[GET /contact/verification/history][%d] verificationHistoryInfoOK  %+v", 200, o.Payload)
}

func (o *VerificationHistoryInfoOK) GetPayload() *models.JSONResponseDataContactVerification {
	return o.Payload
}

func (o *VerificationHistoryInfoOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JSONResponseDataContactVerification)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
