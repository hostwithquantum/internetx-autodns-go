// Code generated by go-swagger; DO NOT EDIT.

package zone_tasks

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

// ZoneRestoreReader is a Reader for the ZoneRestore structure.
type ZoneRestoreReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ZoneRestoreReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewZoneRestoreOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[POST /zone/{name}/{systemNameServer}/_restore] zoneRestore", response, response.Code())
	}
}

// NewZoneRestoreOK creates a ZoneRestoreOK with default headers values
func NewZoneRestoreOK() *ZoneRestoreOK {
	return &ZoneRestoreOK{}
}

/*
ZoneRestoreOK describes a response with status code 200, with default header values.

successful operation
*/
type ZoneRestoreOK struct {
	Payload *models.JSONResponseDataZone
}

// IsSuccess returns true when this zone restore o k response has a 2xx status code
func (o *ZoneRestoreOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this zone restore o k response has a 3xx status code
func (o *ZoneRestoreOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this zone restore o k response has a 4xx status code
func (o *ZoneRestoreOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this zone restore o k response has a 5xx status code
func (o *ZoneRestoreOK) IsServerError() bool {
	return false
}

// IsCode returns true when this zone restore o k response a status code equal to that given
func (o *ZoneRestoreOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the zone restore o k response
func (o *ZoneRestoreOK) Code() int {
	return 200
}

func (o *ZoneRestoreOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /zone/{name}/{systemNameServer}/_restore][%d] zoneRestoreOK %s", 200, payload)
}

func (o *ZoneRestoreOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /zone/{name}/{systemNameServer}/_restore][%d] zoneRestoreOK %s", 200, payload)
}

func (o *ZoneRestoreOK) GetPayload() *models.JSONResponseDataZone {
	return o.Payload
}

func (o *ZoneRestoreOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JSONResponseDataZone)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
