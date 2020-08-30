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

// UserSSOVerificationReader is a Reader for the UserSSOVerification structure.
type UserSSOVerificationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UserSSOVerificationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUserSSOVerificationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUserSSOVerificationOK creates a UserSSOVerificationOK with default headers values
func NewUserSSOVerificationOK() *UserSSOVerificationOK {
	return &UserSSOVerificationOK{}
}

/*UserSSOVerificationOK handles this case with default header values.

successful operation
*/
type UserSSOVerificationOK struct {
	Payload *models.JSONResponseDataUser
}

func (o *UserSSOVerificationOK) Error() string {
	return fmt.Sprintf("[GET /user/_sso/{token}][%d] userSSOVerificationOK  %+v", 200, o.Payload)
}

func (o *UserSSOVerificationOK) GetPayload() *models.JSONResponseDataUser {
	return o.Payload
}

func (o *UserSSOVerificationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JSONResponseDataUser)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}