// Code generated by go-swagger; DO NOT EDIT.

package user_tasks

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

// UserListReader is a Reader for the UserList structure.
type UserListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UserListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUserListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[POST /user/_search] userList", response, response.Code())
	}
}

// NewUserListOK creates a UserListOK with default headers values
func NewUserListOK() *UserListOK {
	return &UserListOK{}
}

/*
UserListOK describes a response with status code 200, with default header values.

successful operation
*/
type UserListOK struct {
	Payload *models.JSONResponseDataBasicUser
}

// IsSuccess returns true when this user list o k response has a 2xx status code
func (o *UserListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this user list o k response has a 3xx status code
func (o *UserListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user list o k response has a 4xx status code
func (o *UserListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this user list o k response has a 5xx status code
func (o *UserListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this user list o k response a status code equal to that given
func (o *UserListOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the user list o k response
func (o *UserListOK) Code() int {
	return 200
}

func (o *UserListOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /user/_search][%d] userListOK %s", 200, payload)
}

func (o *UserListOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /user/_search][%d] userListOK %s", 200, payload)
}

func (o *UserListOK) GetPayload() *models.JSONResponseDataBasicUser {
	return o.Payload
}

func (o *UserListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JSONResponseDataBasicUser)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
