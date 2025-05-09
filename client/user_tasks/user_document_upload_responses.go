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

// UserDocumentUploadReader is a Reader for the UserDocumentUpload structure.
type UserDocumentUploadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UserDocumentUploadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUserDocumentUploadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[POST /user/{name}/{context}/document/{key}] userDocumentUpload", response, response.Code())
	}
}

// NewUserDocumentUploadOK creates a UserDocumentUploadOK with default headers values
func NewUserDocumentUploadOK() *UserDocumentUploadOK {
	return &UserDocumentUploadOK{}
}

/*
UserDocumentUploadOK describes a response with status code 200, with default header values.

successful operation
*/
type UserDocumentUploadOK struct {
	Payload *models.JSONResponseDataVoid
}

// IsSuccess returns true when this user document upload o k response has a 2xx status code
func (o *UserDocumentUploadOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this user document upload o k response has a 3xx status code
func (o *UserDocumentUploadOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user document upload o k response has a 4xx status code
func (o *UserDocumentUploadOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this user document upload o k response has a 5xx status code
func (o *UserDocumentUploadOK) IsServerError() bool {
	return false
}

// IsCode returns true when this user document upload o k response a status code equal to that given
func (o *UserDocumentUploadOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the user document upload o k response
func (o *UserDocumentUploadOK) Code() int {
	return 200
}

func (o *UserDocumentUploadOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /user/{name}/{context}/document/{key}][%d] userDocumentUploadOK %s", 200, payload)
}

func (o *UserDocumentUploadOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /user/{name}/{context}/document/{key}][%d] userDocumentUploadOK %s", 200, payload)
}

func (o *UserDocumentUploadOK) GetPayload() *models.JSONResponseDataVoid {
	return o.Payload
}

func (o *UserDocumentUploadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JSONResponseDataVoid)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
