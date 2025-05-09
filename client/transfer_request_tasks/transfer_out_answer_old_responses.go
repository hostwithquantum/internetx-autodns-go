// Code generated by go-swagger; DO NOT EDIT.

package transfer_request_tasks

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

// TransferOutAnswerOldReader is a Reader for the TransferOutAnswerOld structure.
type TransferOutAnswerOldReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TransferOutAnswerOldReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTransferOutAnswerOldOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PUT /transferout/{domain}/{type}] transferOutAnswerOld", response, response.Code())
	}
}

// NewTransferOutAnswerOldOK creates a TransferOutAnswerOldOK with default headers values
func NewTransferOutAnswerOldOK() *TransferOutAnswerOldOK {
	return &TransferOutAnswerOldOK{}
}

/*
TransferOutAnswerOldOK describes a response with status code 200, with default header values.

successful operation
*/
type TransferOutAnswerOldOK struct {
	Payload *models.JSONResponseDataTransferOut
}

// IsSuccess returns true when this transfer out answer old o k response has a 2xx status code
func (o *TransferOutAnswerOldOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this transfer out answer old o k response has a 3xx status code
func (o *TransferOutAnswerOldOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this transfer out answer old o k response has a 4xx status code
func (o *TransferOutAnswerOldOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this transfer out answer old o k response has a 5xx status code
func (o *TransferOutAnswerOldOK) IsServerError() bool {
	return false
}

// IsCode returns true when this transfer out answer old o k response a status code equal to that given
func (o *TransferOutAnswerOldOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the transfer out answer old o k response
func (o *TransferOutAnswerOldOK) Code() int {
	return 200
}

func (o *TransferOutAnswerOldOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /transferout/{domain}/{type}][%d] transferOutAnswerOldOK %s", 200, payload)
}

func (o *TransferOutAnswerOldOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /transferout/{domain}/{type}][%d] transferOutAnswerOldOK %s", 200, payload)
}

func (o *TransferOutAnswerOldOK) GetPayload() *models.JSONResponseDataTransferOut {
	return o.Payload
}

func (o *TransferOutAnswerOldOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JSONResponseDataTransferOut)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
