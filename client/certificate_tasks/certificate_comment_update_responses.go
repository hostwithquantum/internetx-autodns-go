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

// CertificateCommentUpdateReader is a Reader for the CertificateCommentUpdate structure.
type CertificateCommentUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CertificateCommentUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCertificateCommentUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCertificateCommentUpdateOK creates a CertificateCommentUpdateOK with default headers values
func NewCertificateCommentUpdateOK() *CertificateCommentUpdateOK {
	return &CertificateCommentUpdateOK{}
}

/*CertificateCommentUpdateOK handles this case with default header values.

successful operation
*/
type CertificateCommentUpdateOK struct {
	Payload *models.JSONResponseDataVoid
}

func (o *CertificateCommentUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /certificate/{id}/_comment][%d] certificateCommentUpdateOK  %+v", 200, o.Payload)
}

func (o *CertificateCommentUpdateOK) GetPayload() *models.JSONResponseDataVoid {
	return o.Payload
}

func (o *CertificateCommentUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JSONResponseDataVoid)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
