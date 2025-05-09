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

// SubscriptionListReader is a Reader for the SubscriptionList structure.
type SubscriptionListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SubscriptionListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSubscriptionListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[POST /subscription/_search] subscriptionList", response, response.Code())
	}
}

// NewSubscriptionListOK creates a SubscriptionListOK with default headers values
func NewSubscriptionListOK() *SubscriptionListOK {
	return &SubscriptionListOK{}
}

/*
SubscriptionListOK describes a response with status code 200, with default header values.

successful operation
*/
type SubscriptionListOK struct {
	Payload *models.JSONResponseDataPeriodicBilling
}

// IsSuccess returns true when this subscription list o k response has a 2xx status code
func (o *SubscriptionListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this subscription list o k response has a 3xx status code
func (o *SubscriptionListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this subscription list o k response has a 4xx status code
func (o *SubscriptionListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this subscription list o k response has a 5xx status code
func (o *SubscriptionListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this subscription list o k response a status code equal to that given
func (o *SubscriptionListOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the subscription list o k response
func (o *SubscriptionListOK) Code() int {
	return 200
}

func (o *SubscriptionListOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /subscription/_search][%d] subscriptionListOK %s", 200, payload)
}

func (o *SubscriptionListOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /subscription/_search][%d] subscriptionListOK %s", 200, payload)
}

func (o *SubscriptionListOK) GetPayload() *models.JSONResponseDataPeriodicBilling {
	return o.Payload
}

func (o *SubscriptionListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JSONResponseDataPeriodicBilling)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
