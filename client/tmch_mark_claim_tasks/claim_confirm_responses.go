// Code generated by go-swagger; DO NOT EDIT.

package tmch_mark_claim_tasks

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

// ClaimConfirmReader is a Reader for the ClaimConfirm structure.
type ClaimConfirmReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ClaimConfirmReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewClaimConfirmOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[POST /tmchMark/claim/{reference}/_confirm] claimConfirm", response, response.Code())
	}
}

// NewClaimConfirmOK creates a ClaimConfirmOK with default headers values
func NewClaimConfirmOK() *ClaimConfirmOK {
	return &ClaimConfirmOK{}
}

/*
ClaimConfirmOK describes a response with status code 200, with default header values.

successful operation
*/
type ClaimConfirmOK struct {
	Payload *models.JSONResponseDataTmchClaimsNotice
}

// IsSuccess returns true when this claim confirm o k response has a 2xx status code
func (o *ClaimConfirmOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this claim confirm o k response has a 3xx status code
func (o *ClaimConfirmOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this claim confirm o k response has a 4xx status code
func (o *ClaimConfirmOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this claim confirm o k response has a 5xx status code
func (o *ClaimConfirmOK) IsServerError() bool {
	return false
}

// IsCode returns true when this claim confirm o k response a status code equal to that given
func (o *ClaimConfirmOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the claim confirm o k response
func (o *ClaimConfirmOK) Code() int {
	return 200
}

func (o *ClaimConfirmOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /tmchMark/claim/{reference}/_confirm][%d] claimConfirmOK %s", 200, payload)
}

func (o *ClaimConfirmOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /tmchMark/claim/{reference}/_confirm][%d] claimConfirmOK %s", 200, payload)
}

func (o *ClaimConfirmOK) GetPayload() *models.JSONResponseDataTmchClaimsNotice {
	return o.Payload
}

func (o *ClaimConfirmOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JSONResponseDataTmchClaimsNotice)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
