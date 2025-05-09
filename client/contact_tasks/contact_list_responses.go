// Code generated by go-swagger; DO NOT EDIT.

package contact_tasks

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

// ContactListReader is a Reader for the ContactList structure.
type ContactListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ContactListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewContactListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[POST /contact/_search] contactList", response, response.Code())
	}
}

// NewContactListOK creates a ContactListOK with default headers values
func NewContactListOK() *ContactListOK {
	return &ContactListOK{}
}

/*
ContactListOK describes a response with status code 200, with default header values.

successful operation
*/
type ContactListOK struct {
	Payload *models.JSONResponseDataContact
}

// IsSuccess returns true when this contact list o k response has a 2xx status code
func (o *ContactListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this contact list o k response has a 3xx status code
func (o *ContactListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this contact list o k response has a 4xx status code
func (o *ContactListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this contact list o k response has a 5xx status code
func (o *ContactListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this contact list o k response a status code equal to that given
func (o *ContactListOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the contact list o k response
func (o *ContactListOK) Code() int {
	return 200
}

func (o *ContactListOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /contact/_search][%d] contactListOK %s", 200, payload)
}

func (o *ContactListOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /contact/_search][%d] contactListOK %s", 200, payload)
}

func (o *ContactListOK) GetPayload() *models.JSONResponseDataContact {
	return o.Payload
}

func (o *ContactListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JSONResponseDataContact)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
