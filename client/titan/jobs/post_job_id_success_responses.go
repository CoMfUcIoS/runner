package jobs

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

// PostJobIDSuccessReader is a Reader for the PostJobIDSuccess structure.
type PostJobIDSuccessReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *PostJobIDSuccessReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostJobIDSuccessOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewPostJobIDSuccessNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewPostJobIDSuccessConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewPostJobIDSuccessDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewPostJobIDSuccessOK creates a PostJobIDSuccessOK with default headers values
func NewPostJobIDSuccessOK() *PostJobIDSuccessOK {
	return &PostJobIDSuccessOK{}
}

/*PostJobIDSuccessOK handles this case with default header values.

Job updated.
*/
type PostJobIDSuccessOK struct {
	Payload *models.JobWrapper
}

func (o *PostJobIDSuccessOK) Error() string {
	return fmt.Sprintf("[POST /job/{id}/success][%d] postJobIdSuccessOK  %+v", 200, o.Payload)
}

func (o *PostJobIDSuccessOK) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JobWrapper)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostJobIDSuccessNotFound creates a PostJobIDSuccessNotFound with default headers values
func NewPostJobIDSuccessNotFound() *PostJobIDSuccessNotFound {
	return &PostJobIDSuccessNotFound{}
}

/*PostJobIDSuccessNotFound handles this case with default header values.

Job does not exist.
*/
type PostJobIDSuccessNotFound struct {
	Payload *models.Error
}

func (o *PostJobIDSuccessNotFound) Error() string {
	return fmt.Sprintf("[POST /job/{id}/success][%d] postJobIdSuccessNotFound  %+v", 404, o.Payload)
}

func (o *PostJobIDSuccessNotFound) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostJobIDSuccessConflict creates a PostJobIDSuccessConflict with default headers values
func NewPostJobIDSuccessConflict() *PostJobIDSuccessConflict {
	return &PostJobIDSuccessConflict{}
}

/*PostJobIDSuccessConflict handles this case with default header values.

Job was not in running state.
*/
type PostJobIDSuccessConflict struct {
	Payload *models.IDStatus
}

func (o *PostJobIDSuccessConflict) Error() string {
	return fmt.Sprintf("[POST /job/{id}/success][%d] postJobIdSuccessConflict  %+v", 409, o.Payload)
}

func (o *PostJobIDSuccessConflict) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IDStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostJobIDSuccessDefault creates a PostJobIDSuccessDefault with default headers values
func NewPostJobIDSuccessDefault(code int) *PostJobIDSuccessDefault {
	return &PostJobIDSuccessDefault{
		_statusCode: code,
	}
}

/*PostJobIDSuccessDefault handles this case with default header values.

Unexpected error
*/
type PostJobIDSuccessDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the post job ID success default response
func (o *PostJobIDSuccessDefault) Code() int {
	return o._statusCode
}

func (o *PostJobIDSuccessDefault) Error() string {
	return fmt.Sprintf("[POST /job/{id}/success][%d] PostJobIDSuccess default  %+v", o._statusCode, o.Payload)
}

func (o *PostJobIDSuccessDefault) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
