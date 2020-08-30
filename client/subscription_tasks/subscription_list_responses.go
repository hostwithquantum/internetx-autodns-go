// Code generated by go-swagger; DO NOT EDIT.

package subscription_tasks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
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
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSubscriptionListOK creates a SubscriptionListOK with default headers values
func NewSubscriptionListOK() *SubscriptionListOK {
	return &SubscriptionListOK{}
}

/*SubscriptionListOK handles this case with default header values.

successful operation
*/
type SubscriptionListOK struct {
	Payload *models.JSONResponseDataPeriodicBilling
}

func (o *SubscriptionListOK) Error() string {
	return fmt.Sprintf("[POST /subscription/_search][%d] subscriptionListOK  %+v", 200, o.Payload)
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