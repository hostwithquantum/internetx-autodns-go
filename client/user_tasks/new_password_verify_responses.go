// Code generated by go-swagger; DO NOT EDIT.

package user_tasks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/hostwithquantum/internetx-autodns-go/models"
)

// NewPasswordVerifyReader is a Reader for the NewPasswordVerify structure.
type NewPasswordVerifyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NewPasswordVerifyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNewPasswordVerifyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewNewPasswordVerifyOK creates a NewPasswordVerifyOK with default headers values
func NewNewPasswordVerifyOK() *NewPasswordVerifyOK {
	return &NewPasswordVerifyOK{}
}

/*NewPasswordVerifyOK handles this case with default header values.

successful operation
*/
type NewPasswordVerifyOK struct {
	Payload *models.JSONResponseDataVoid
}

func (o *NewPasswordVerifyOK) Error() string {
	return fmt.Sprintf("[GET /user/_newPasswordVerify?token={token}][%d] newPasswordVerifyOK  %+v", 200, o.Payload)
}

func (o *NewPasswordVerifyOK) GetPayload() *models.JSONResponseDataVoid {
	return o.Payload
}

func (o *NewPasswordVerifyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JSONResponseDataVoid)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
