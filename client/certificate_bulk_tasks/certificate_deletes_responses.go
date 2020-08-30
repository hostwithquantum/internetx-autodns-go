// Code generated by go-swagger; DO NOT EDIT.

package certificate_bulk_tasks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/hostwithquantum/internetx-autodns-go/models"
)

// CertificateDeletesReader is a Reader for the CertificateDeletes structure.
type CertificateDeletesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CertificateDeletesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCertificateDeletesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCertificateDeletesOK creates a CertificateDeletesOK with default headers values
func NewCertificateDeletesOK() *CertificateDeletesOK {
	return &CertificateDeletesOK{}
}

/*CertificateDeletesOK handles this case with default header values.

successful operation
*/
type CertificateDeletesOK struct {
	Payload *models.JSONResponseDataListJSONResponseDataCertificate
}

func (o *CertificateDeletesOK) Error() string {
	return fmt.Sprintf("[DELETE /bulk/certificate][%d] certificateDeletesOK  %+v", 200, o.Payload)
}

func (o *CertificateDeletesOK) GetPayload() *models.JSONResponseDataListJSONResponseDataCertificate {
	return o.Payload
}

func (o *CertificateDeletesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JSONResponseDataListJSONResponseDataCertificate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
