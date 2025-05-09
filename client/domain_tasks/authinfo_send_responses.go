// Code generated by go-swagger; DO NOT EDIT.

package domain_tasks

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

// AuthinfoSendReader is a Reader for the AuthinfoSend structure.
type AuthinfoSendReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AuthinfoSendReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAuthinfoSendOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PUT /domain/{name}/_sendAuthinfoToOwnerc] authinfoSend", response, response.Code())
	}
}

// NewAuthinfoSendOK creates a AuthinfoSendOK with default headers values
func NewAuthinfoSendOK() *AuthinfoSendOK {
	return &AuthinfoSendOK{}
}

/*
AuthinfoSendOK describes a response with status code 200, with default header values.

successful operation
*/
type AuthinfoSendOK struct {
	Payload *models.JSONResponseDataJSONNoData
}

// IsSuccess returns true when this authinfo send o k response has a 2xx status code
func (o *AuthinfoSendOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this authinfo send o k response has a 3xx status code
func (o *AuthinfoSendOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this authinfo send o k response has a 4xx status code
func (o *AuthinfoSendOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this authinfo send o k response has a 5xx status code
func (o *AuthinfoSendOK) IsServerError() bool {
	return false
}

// IsCode returns true when this authinfo send o k response a status code equal to that given
func (o *AuthinfoSendOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the authinfo send o k response
func (o *AuthinfoSendOK) Code() int {
	return 200
}

func (o *AuthinfoSendOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /domain/{name}/_sendAuthinfoToOwnerc][%d] authinfoSendOK %s", 200, payload)
}

func (o *AuthinfoSendOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /domain/{name}/_sendAuthinfoToOwnerc][%d] authinfoSendOK %s", 200, payload)
}

func (o *AuthinfoSendOK) GetPayload() *models.JSONResponseDataJSONNoData {
	return o.Payload
}

func (o *AuthinfoSendOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JSONResponseDataJSONNoData)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
