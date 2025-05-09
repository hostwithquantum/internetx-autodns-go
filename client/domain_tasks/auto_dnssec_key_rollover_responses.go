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

// AutoDnssecKeyRolloverReader is a Reader for the AutoDnssecKeyRollover structure.
type AutoDnssecKeyRolloverReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AutoDnssecKeyRolloverReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAutoDnssecKeyRolloverOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PUT /domain/{name}/_autoDnssecKeyRollover] autoDnssecKeyRollover", response, response.Code())
	}
}

// NewAutoDnssecKeyRolloverOK creates a AutoDnssecKeyRolloverOK with default headers values
func NewAutoDnssecKeyRolloverOK() *AutoDnssecKeyRolloverOK {
	return &AutoDnssecKeyRolloverOK{}
}

/*
AutoDnssecKeyRolloverOK describes a response with status code 200, with default header values.

successful operation
*/
type AutoDnssecKeyRolloverOK struct {
	Payload *models.JSONResponseDataVoid
}

// IsSuccess returns true when this auto dnssec key rollover o k response has a 2xx status code
func (o *AutoDnssecKeyRolloverOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this auto dnssec key rollover o k response has a 3xx status code
func (o *AutoDnssecKeyRolloverOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this auto dnssec key rollover o k response has a 4xx status code
func (o *AutoDnssecKeyRolloverOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this auto dnssec key rollover o k response has a 5xx status code
func (o *AutoDnssecKeyRolloverOK) IsServerError() bool {
	return false
}

// IsCode returns true when this auto dnssec key rollover o k response a status code equal to that given
func (o *AutoDnssecKeyRolloverOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the auto dnssec key rollover o k response
func (o *AutoDnssecKeyRolloverOK) Code() int {
	return 200
}

func (o *AutoDnssecKeyRolloverOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /domain/{name}/_autoDnssecKeyRollover][%d] autoDnssecKeyRolloverOK %s", 200, payload)
}

func (o *AutoDnssecKeyRolloverOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /domain/{name}/_autoDnssecKeyRollover][%d] autoDnssecKeyRolloverOK %s", 200, payload)
}

func (o *AutoDnssecKeyRolloverOK) GetPayload() *models.JSONResponseDataVoid {
	return o.Payload
}

func (o *AutoDnssecKeyRolloverOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JSONResponseDataVoid)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
