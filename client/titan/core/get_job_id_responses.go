package core

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/httpkit"

	strfmt "github.com/go-swagger/go-swagger/strfmt"

	"github.com/iron-io/titan/runner/client/models"
)

// GetJobIDReader is a Reader for the GetJobID structure.
type GetJobIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *GetJobIDReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetJobIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewGetJobIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewGetJobIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewGetJobIDOK creates a GetJobIDOK with default headers values
func NewGetJobIDOK() *GetJobIDOK {
	return &GetJobIDOK{}
}

/*GetJobIDOK handles this case with default header values.

Job information
*/
type GetJobIDOK struct {
	Payload *models.JobWrapper
}

func (o *GetJobIDOK) Error() string {
	return fmt.Sprintf("[GET /job/{id}][%d] getJobIdOK  %+v", 200, o.Payload)
}

func (o *GetJobIDOK) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JobWrapper)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetJobIDNotFound creates a GetJobIDNotFound with default headers values
func NewGetJobIDNotFound() *GetJobIDNotFound {
	return &GetJobIDNotFound{}
}

/*GetJobIDNotFound handles this case with default header values.

Job does not exist.
*/
type GetJobIDNotFound struct {
	Payload *models.Error
}

func (o *GetJobIDNotFound) Error() string {
	return fmt.Sprintf("[GET /job/{id}][%d] getJobIdNotFound  %+v", 404, o.Payload)
}

func (o *GetJobIDNotFound) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetJobIDDefault creates a GetJobIDDefault with default headers values
func NewGetJobIDDefault(code int) *GetJobIDDefault {
	return &GetJobIDDefault{
		_statusCode: code,
	}
}

/*GetJobIDDefault handles this case with default header values.

Unexpected error
*/
type GetJobIDDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get job ID default response
func (o *GetJobIDDefault) Code() int {
	return o._statusCode
}

func (o *GetJobIDDefault) Error() string {
	return fmt.Sprintf("[GET /job/{id}][%d] GetJobID default  %+v", o._statusCode, o.Payload)
}

func (o *GetJobIDDefault) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
