// Code generated by go-swagger; DO NOT EDIT.

package subscription_tasks

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

// SubscriptionCreateReader is a Reader for the SubscriptionCreate structure.
type SubscriptionCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SubscriptionCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSubscriptionCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[POST /subscription] subscriptionCreate", response, response.Code())
	}
}

// NewSubscriptionCreateOK creates a SubscriptionCreateOK with default headers values
func NewSubscriptionCreateOK() *SubscriptionCreateOK {
	return &SubscriptionCreateOK{}
}

/*
SubscriptionCreateOK describes a response with status code 200, with default header values.

successful operation
*/
type SubscriptionCreateOK struct {
	Payload *models.JSONResponseDataPeriodicBilling
}

// IsSuccess returns true when this subscription create o k response has a 2xx status code
func (o *SubscriptionCreateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this subscription create o k response has a 3xx status code
func (o *SubscriptionCreateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this subscription create o k response has a 4xx status code
func (o *SubscriptionCreateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this subscription create o k response has a 5xx status code
func (o *SubscriptionCreateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this subscription create o k response a status code equal to that given
func (o *SubscriptionCreateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the subscription create o k response
func (o *SubscriptionCreateOK) Code() int {
	return 200
}

func (o *SubscriptionCreateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /subscription][%d] subscriptionCreateOK %s", 200, payload)
}

func (o *SubscriptionCreateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /subscription][%d] subscriptionCreateOK %s", 200, payload)
}

func (o *SubscriptionCreateOK) GetPayload() *models.JSONResponseDataPeriodicBilling {
	return o.Payload
}

func (o *SubscriptionCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JSONResponseDataPeriodicBilling)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
