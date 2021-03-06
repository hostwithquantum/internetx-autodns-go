// Code generated by go-swagger; DO NOT EDIT.

package job_tasks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/hostwithquantum/internetx-autodns-go/models"
)

// JobRestartReader is a Reader for the JobRestart structure.
type JobRestartReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *JobRestartReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewJobRestartOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewJobRestartOK creates a JobRestartOK with default headers values
func NewJobRestartOK() *JobRestartOK {
	return &JobRestartOK{}
}

/*JobRestartOK handles this case with default header values.

successful operation
*/
type JobRestartOK struct {
	Payload *models.JSONResponseDataJSONNoData
}

func (o *JobRestartOK) Error() string {
	return fmt.Sprintf("[PUT /job/{id}/_restart][%d] jobRestartOK  %+v", 200, o.Payload)
}

func (o *JobRestartOK) GetPayload() *models.JSONResponseDataJSONNoData {
	return o.Payload
}

func (o *JobRestartOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JSONResponseDataJSONNoData)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
