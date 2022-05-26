// Code generated by go-swagger; DO NOT EDIT.

package internal_storage_regions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/IBM-Cloud/power-go-client/power/models"
)

// InternalV1StorageRegionsThresholdsGetReader is a Reader for the InternalV1StorageRegionsThresholdsGet structure.
type InternalV1StorageRegionsThresholdsGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InternalV1StorageRegionsThresholdsGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewInternalV1StorageRegionsThresholdsGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewInternalV1StorageRegionsThresholdsGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewInternalV1StorageRegionsThresholdsGetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewInternalV1StorageRegionsThresholdsGetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewInternalV1StorageRegionsThresholdsGetOK creates a InternalV1StorageRegionsThresholdsGetOK with default headers values
func NewInternalV1StorageRegionsThresholdsGetOK() *InternalV1StorageRegionsThresholdsGetOK {
	return &InternalV1StorageRegionsThresholdsGetOK{}
}

/* InternalV1StorageRegionsThresholdsGetOK describes a response with status code 200, with default header values.

OK
*/
type InternalV1StorageRegionsThresholdsGetOK struct {
	Payload *models.Thresholds
}

func (o *InternalV1StorageRegionsThresholdsGetOK) Error() string {
	return fmt.Sprintf("[GET /internal/v1/storage/regions/{region_zone_id}/thresholds][%d] internalV1StorageRegionsThresholdsGetOK  %+v", 200, o.Payload)
}
func (o *InternalV1StorageRegionsThresholdsGetOK) GetPayload() *models.Thresholds {
	return o.Payload
}

func (o *InternalV1StorageRegionsThresholdsGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Thresholds)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInternalV1StorageRegionsThresholdsGetUnauthorized creates a InternalV1StorageRegionsThresholdsGetUnauthorized with default headers values
func NewInternalV1StorageRegionsThresholdsGetUnauthorized() *InternalV1StorageRegionsThresholdsGetUnauthorized {
	return &InternalV1StorageRegionsThresholdsGetUnauthorized{}
}

/* InternalV1StorageRegionsThresholdsGetUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type InternalV1StorageRegionsThresholdsGetUnauthorized struct {
	Payload *models.Error
}

func (o *InternalV1StorageRegionsThresholdsGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /internal/v1/storage/regions/{region_zone_id}/thresholds][%d] internalV1StorageRegionsThresholdsGetUnauthorized  %+v", 401, o.Payload)
}
func (o *InternalV1StorageRegionsThresholdsGetUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *InternalV1StorageRegionsThresholdsGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInternalV1StorageRegionsThresholdsGetNotFound creates a InternalV1StorageRegionsThresholdsGetNotFound with default headers values
func NewInternalV1StorageRegionsThresholdsGetNotFound() *InternalV1StorageRegionsThresholdsGetNotFound {
	return &InternalV1StorageRegionsThresholdsGetNotFound{}
}

/* InternalV1StorageRegionsThresholdsGetNotFound describes a response with status code 404, with default header values.

Not Found
*/
type InternalV1StorageRegionsThresholdsGetNotFound struct {
	Payload *models.Error
}

func (o *InternalV1StorageRegionsThresholdsGetNotFound) Error() string {
	return fmt.Sprintf("[GET /internal/v1/storage/regions/{region_zone_id}/thresholds][%d] internalV1StorageRegionsThresholdsGetNotFound  %+v", 404, o.Payload)
}
func (o *InternalV1StorageRegionsThresholdsGetNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *InternalV1StorageRegionsThresholdsGetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInternalV1StorageRegionsThresholdsGetInternalServerError creates a InternalV1StorageRegionsThresholdsGetInternalServerError with default headers values
func NewInternalV1StorageRegionsThresholdsGetInternalServerError() *InternalV1StorageRegionsThresholdsGetInternalServerError {
	return &InternalV1StorageRegionsThresholdsGetInternalServerError{}
}

/* InternalV1StorageRegionsThresholdsGetInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type InternalV1StorageRegionsThresholdsGetInternalServerError struct {
	Payload *models.Error
}

func (o *InternalV1StorageRegionsThresholdsGetInternalServerError) Error() string {
	return fmt.Sprintf("[GET /internal/v1/storage/regions/{region_zone_id}/thresholds][%d] internalV1StorageRegionsThresholdsGetInternalServerError  %+v", 500, o.Payload)
}
func (o *InternalV1StorageRegionsThresholdsGetInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *InternalV1StorageRegionsThresholdsGetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}