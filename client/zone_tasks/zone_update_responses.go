// Code generated by go-swagger; DO NOT EDIT.

package zone_tasks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/hostwithquantum/internetx-autodns-go/models"
)

// ZoneUpdateReader is a Reader for the ZoneUpdate structure.
type ZoneUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ZoneUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewZoneUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewZoneUpdateOK creates a ZoneUpdateOK with default headers values
func NewZoneUpdateOK() *ZoneUpdateOK {
	return &ZoneUpdateOK{}
}

/*ZoneUpdateOK handles this case with default header values.

successful operation
*/
type ZoneUpdateOK struct {
	Payload *models.JSONResponseDataZone
}

func (o *ZoneUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /zone/{name}/{systemNameServer}][%d] zoneUpdateOK  %+v", 200, o.Payload)
}

func (o *ZoneUpdateOK) GetPayload() *models.JSONResponseDataZone {
	return o.Payload
}

func (o *ZoneUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JSONResponseDataZone)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
