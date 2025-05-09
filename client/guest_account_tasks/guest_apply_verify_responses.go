// Code generated by go-swagger; DO NOT EDIT.

package guest_account_tasks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/hostwithquantum/internetx-autodns-go/models"
)

// GuestApplyVerifyReader is a Reader for the GuestApplyVerify structure.
type GuestApplyVerifyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GuestApplyVerifyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGuestApplyVerifyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /user/_guestverify/{token}] guestApplyVerify", response, response.Code())
	}
}

// NewGuestApplyVerifyOK creates a GuestApplyVerifyOK with default headers values
func NewGuestApplyVerifyOK() *GuestApplyVerifyOK {
	return &GuestApplyVerifyOK{}
}

/*
GuestApplyVerifyOK describes a response with status code 200, with default header values.

successful operation
*/
type GuestApplyVerifyOK struct {
	Payload *models.JSONResponseDataUser
}

// IsSuccess returns true when this guest apply verify o k response has a 2xx status code
func (o *GuestApplyVerifyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this guest apply verify o k response has a 3xx status code
func (o *GuestApplyVerifyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this guest apply verify o k response has a 4xx status code
func (o *GuestApplyVerifyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this guest apply verify o k response has a 5xx status code
func (o *GuestApplyVerifyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this guest apply verify o k response a status code equal to that given
func (o *GuestApplyVerifyOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the guest apply verify o k response
func (o *GuestApplyVerifyOK) Code() int {
	return 200
}

func (o *GuestApplyVerifyOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /user/_guestverify/{token}][%d] guestApplyVerifyOK %s", 200, payload)
}

func (o *GuestApplyVerifyOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /user/_guestverify/{token}][%d] guestApplyVerifyOK %s", 200, payload)
}

func (o *GuestApplyVerifyOK) GetPayload() *models.JSONResponseDataUser {
	return o.Payload
}

func (o *GuestApplyVerifyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JSONResponseDataUser)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
