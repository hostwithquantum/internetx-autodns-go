// Code generated by go-swagger; DO NOT EDIT.

package user_2fa_tasks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/hostwithquantum/internetx-autodns-go/models"
)

// User2faAktivateReader is a Reader for the User2faAktivate structure.
type User2faAktivateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *User2faAktivateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUser2faAktivateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUser2faAktivateOK creates a User2faAktivateOK with default headers values
func NewUser2faAktivateOK() *User2faAktivateOK {
	return &User2faAktivateOK{}
}

/*User2faAktivateOK handles this case with default header values.

successful operation
*/
type User2faAktivateOK struct {
	Payload *models.JSONResponseDataOTPAuth
}

func (o *User2faAktivateOK) Error() string {
	return fmt.Sprintf("[PUT /user/_2fa][%d] user2faAktivateOK  %+v", 200, o.Payload)
}

func (o *User2faAktivateOK) GetPayload() *models.JSONResponseDataOTPAuth {
	return o.Payload
}

func (o *User2faAktivateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JSONResponseDataOTPAuth)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
