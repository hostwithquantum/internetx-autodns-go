// Code generated by go-swagger; DO NOT EDIT.

package certificate_tasks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/hostwithquantum/internetx-autodns-go/models"
)

// CertificateRevokeReader is a Reader for the CertificateRevoke structure.
type CertificateRevokeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CertificateRevokeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCertificateRevokeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCertificateRevokeOK creates a CertificateRevokeOK with default headers values
func NewCertificateRevokeOK() *CertificateRevokeOK {
	return &CertificateRevokeOK{}
}

/*CertificateRevokeOK handles this case with default header values.

successful operation
*/
type CertificateRevokeOK struct {
	Payload *models.JSONResponseDataJob
}

func (o *CertificateRevokeOK) Error() string {
	return fmt.Sprintf("[DELETE /certificate/{id}/_revoke][%d] certificateRevokeOK  %+v", 200, o.Payload)
}

func (o *CertificateRevokeOK) GetPayload() *models.JSONResponseDataJob {
	return o.Payload
}

func (o *CertificateRevokeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JSONResponseDataJob)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
