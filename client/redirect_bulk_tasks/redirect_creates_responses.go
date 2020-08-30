// Code generated by go-swagger; DO NOT EDIT.

package redirect_bulk_tasks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/hostwithquantum/internetx-autodns-go/models"
)

// RedirectCreatesReader is a Reader for the RedirectCreates structure.
type RedirectCreatesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RedirectCreatesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRedirectCreatesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewRedirectCreatesOK creates a RedirectCreatesOK with default headers values
func NewRedirectCreatesOK() *RedirectCreatesOK {
	return &RedirectCreatesOK{}
}

/*RedirectCreatesOK handles this case with default header values.

successful operation
*/
type RedirectCreatesOK struct {
	Payload *models.JSONResponseDataListJSONResponseDataRedirect
}

func (o *RedirectCreatesOK) Error() string {
	return fmt.Sprintf("[POST /bulk/redirect][%d] redirectCreatesOK  %+v", 200, o.Payload)
}

func (o *RedirectCreatesOK) GetPayload() *models.JSONResponseDataListJSONResponseDataRedirect {
	return o.Payload
}

func (o *RedirectCreatesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JSONResponseDataListJSONResponseDataRedirect)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}