// Code generated by go-swagger; DO NOT EDIT.

package zone_bulk_tasks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/hostwithquantum/internetx-autodns-go/models"
)

// ZoneBulkPatchesReader is a Reader for the ZoneBulkPatches structure.
type ZoneBulkPatchesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ZoneBulkPatchesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewZoneBulkPatchesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewZoneBulkPatchesOK creates a ZoneBulkPatchesOK with default headers values
func NewZoneBulkPatchesOK() *ZoneBulkPatchesOK {
	return &ZoneBulkPatchesOK{}
}

/*ZoneBulkPatchesOK handles this case with default header values.

successful operation
*/
type ZoneBulkPatchesOK struct {
	Payload *models.JSONResponseDataListJSONResponseDataZone
}

func (o *ZoneBulkPatchesOK) Error() string {
	return fmt.Sprintf("[PATCH /bulk/zone][%d] zoneBulkPatchesOK  %+v", 200, o.Payload)
}

func (o *ZoneBulkPatchesOK) GetPayload() *models.JSONResponseDataListJSONResponseDataZone {
	return o.Payload
}

func (o *ZoneBulkPatchesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JSONResponseDataListJSONResponseDataZone)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
