// Code generated by go-swagger; DO NOT EDIT.

package ssl_contact_tasks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/hostwithquantum/internetx-autodns-go/models"
)

// SslContactUpdateReader is a Reader for the SslContactUpdate structure.
type SslContactUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SslContactUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSslContactUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSslContactUpdateOK creates a SslContactUpdateOK with default headers values
func NewSslContactUpdateOK() *SslContactUpdateOK {
	return &SslContactUpdateOK{}
}

/*SslContactUpdateOK handles this case with default header values.

successful operation
*/
type SslContactUpdateOK struct {
	Payload *models.JSONResponseDataSslContact
}

func (o *SslContactUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /sslcontact/{id}][%d] sslContactUpdateOK  %+v", 200, o.Payload)
}

func (o *SslContactUpdateOK) GetPayload() *models.JSONResponseDataSslContact {
	return o.Payload
}

func (o *SslContactUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JSONResponseDataSslContact)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
