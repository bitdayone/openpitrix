// Code generated by go-swagger; DO NOT EDIT.

package apps

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"openpitrix.io/openpitrix/pkg/swagger/models"
)

// GetAppsReader is a Reader for the GetApps structure.
type GetAppsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAppsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetAppsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 500:
		result := NewGetAppsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetAppsOK creates a GetAppsOK with default headers values
func NewGetAppsOK() *GetAppsOK {
	return &GetAppsOK{}
}

/*GetAppsOK handles this case with default header values.

A list of apps
*/
type GetAppsOK struct {
	Payload *models.GetAppsOKBody
}

func (o *GetAppsOK) Error() string {
	return fmt.Sprintf("[GET /apps][%d] getAppsOK  %+v", 200, o.Payload)
}

func (o *GetAppsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetAppsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAppsInternalServerError creates a GetAppsInternalServerError with default headers values
func NewGetAppsInternalServerError() *GetAppsInternalServerError {
	return &GetAppsInternalServerError{}
}

/*GetAppsInternalServerError handles this case with default header values.

An unexpected error occured.
*/
type GetAppsInternalServerError struct {
	Payload *models.Error
}

func (o *GetAppsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /apps][%d] getAppsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetAppsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}