// Code generated by go-swagger; DO NOT EDIT.

package accounting_tasks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/hostwithquantum/internetx-autodns-go/models"
)

// AccountingDocumentCreateReader is a Reader for the AccountingDocumentCreate structure.
type AccountingDocumentCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AccountingDocumentCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAccountingDocumentCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAccountingDocumentCreateOK creates a AccountingDocumentCreateOK with default headers values
func NewAccountingDocumentCreateOK() *AccountingDocumentCreateOK {
	return &AccountingDocumentCreateOK{}
}

/*AccountingDocumentCreateOK handles this case with default header values.

successful operation
*/
type AccountingDocumentCreateOK struct {
	Payload *models.JSONResponseDataAccountingDocument
}

func (o *AccountingDocumentCreateOK) Error() string {
	return fmt.Sprintf("[POST /accounting][%d] accountingDocumentCreateOK  %+v", 200, o.Payload)
}

func (o *AccountingDocumentCreateOK) GetPayload() *models.JSONResponseDataAccountingDocument {
	return o.Payload
}

func (o *AccountingDocumentCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JSONResponseDataAccountingDocument)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
