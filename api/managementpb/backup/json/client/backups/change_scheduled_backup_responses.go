// Code generated by go-swagger; DO NOT EDIT.

package backups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ChangeScheduledBackupReader is a Reader for the ChangeScheduledBackup structure.
type ChangeScheduledBackupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ChangeScheduledBackupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewChangeScheduledBackupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewChangeScheduledBackupDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewChangeScheduledBackupOK creates a ChangeScheduledBackupOK with default headers values
func NewChangeScheduledBackupOK() *ChangeScheduledBackupOK {
	return &ChangeScheduledBackupOK{}
}

/* ChangeScheduledBackupOK describes a response with status code 200, with default header values.

A successful response.
*/
type ChangeScheduledBackupOK struct {
	Payload interface{}
}

func (o *ChangeScheduledBackupOK) Error() string {
	return fmt.Sprintf("[POST /v1/management/backup/Backups/ChangeScheduled][%d] changeScheduledBackupOk  %+v", 200, o.Payload)
}

func (o *ChangeScheduledBackupOK) GetPayload() interface{} {
	return o.Payload
}

func (o *ChangeScheduledBackupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewChangeScheduledBackupDefault creates a ChangeScheduledBackupDefault with default headers values
func NewChangeScheduledBackupDefault(code int) *ChangeScheduledBackupDefault {
	return &ChangeScheduledBackupDefault{
		_statusCode: code,
	}
}

/* ChangeScheduledBackupDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ChangeScheduledBackupDefault struct {
	_statusCode int

	Payload *ChangeScheduledBackupDefaultBody
}

// Code gets the status code for the change scheduled backup default response
func (o *ChangeScheduledBackupDefault) Code() int {
	return o._statusCode
}

func (o *ChangeScheduledBackupDefault) Error() string {
	return fmt.Sprintf("[POST /v1/management/backup/Backups/ChangeScheduled][%d] ChangeScheduledBackup default  %+v", o._statusCode, o.Payload)
}

func (o *ChangeScheduledBackupDefault) GetPayload() *ChangeScheduledBackupDefaultBody {
	return o.Payload
}

func (o *ChangeScheduledBackupDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(ChangeScheduledBackupDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*ChangeScheduledBackupBody change scheduled backup body
swagger:model ChangeScheduledBackupBody
*/
type ChangeScheduledBackupBody struct {
	// scheduled backup id
	ScheduledBackupID string `json:"scheduled_backup_id,omitempty"`

	// enabled
	Enabled bool `json:"enabled,omitempty"`

	// How often backup should be run in cron format.
	CronExpression string `json:"cron_expression,omitempty"`

	// First backup wouldn't happen before this time.
	// Format: date-time
	StartTime strfmt.DateTime `json:"start_time,omitempty"`

	// Name of backup.
	Name string `json:"name,omitempty"`

	// Human-readable description.
	Description string `json:"description,omitempty"`

	// Delay between each retry. Should have a suffix in JSON: 1s, 1m, 1h.
	RetryInterval string `json:"retry_interval,omitempty"`

	// How many times to retry a failed backup before giving up.
	Retries int64 `json:"retries,omitempty"`

	// How many artifacts keep. 0 - unlimited.
	Retention int64 `json:"retention,omitempty"`
}

// Validate validates this change scheduled backup body
func (o *ChangeScheduledBackupBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateStartTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ChangeScheduledBackupBody) validateStartTime(formats strfmt.Registry) error {
	if swag.IsZero(o.StartTime) { // not required
		return nil
	}

	if err := validate.FormatOf("body"+"."+"start_time", "body", "date-time", o.StartTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this change scheduled backup body based on context it is used
func (o *ChangeScheduledBackupBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ChangeScheduledBackupBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangeScheduledBackupBody) UnmarshalBinary(b []byte) error {
	var res ChangeScheduledBackupBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ChangeScheduledBackupDefaultBody change scheduled backup default body
swagger:model ChangeScheduledBackupDefaultBody
*/
type ChangeScheduledBackupDefaultBody struct {
	// error
	Error string `json:"error,omitempty"`

	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*ChangeScheduledBackupDefaultBodyDetailsItems0 `json:"details"`
}

// Validate validates this change scheduled backup default body
func (o *ChangeScheduledBackupDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ChangeScheduledBackupDefaultBody) validateDetails(formats strfmt.Registry) error {
	if swag.IsZero(o.Details) { // not required
		return nil
	}

	for i := 0; i < len(o.Details); i++ {
		if swag.IsZero(o.Details[i]) { // not required
			continue
		}

		if o.Details[i] != nil {
			if err := o.Details[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ChangeScheduledBackup default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ChangeScheduledBackup default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this change scheduled backup default body based on the context it is used
func (o *ChangeScheduledBackupDefaultBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ChangeScheduledBackupDefaultBody) contextValidateDetails(ctx context.Context, formats strfmt.Registry) error {
	for i := 0; i < len(o.Details); i++ {
		if o.Details[i] != nil {
			if err := o.Details[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ChangeScheduledBackup default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ChangeScheduledBackup default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ChangeScheduledBackupDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangeScheduledBackupDefaultBody) UnmarshalBinary(b []byte) error {
	var res ChangeScheduledBackupDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ChangeScheduledBackupDefaultBodyDetailsItems0 change scheduled backup default body details items0
swagger:model ChangeScheduledBackupDefaultBodyDetailsItems0
*/
type ChangeScheduledBackupDefaultBodyDetailsItems0 struct {
	// type url
	TypeURL string `json:"type_url,omitempty"`

	// value
	// Format: byte
	Value strfmt.Base64 `json:"value,omitempty"`
}

// Validate validates this change scheduled backup default body details items0
func (o *ChangeScheduledBackupDefaultBodyDetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this change scheduled backup default body details items0 based on context it is used
func (o *ChangeScheduledBackupDefaultBodyDetailsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ChangeScheduledBackupDefaultBodyDetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangeScheduledBackupDefaultBodyDetailsItems0) UnmarshalBinary(b []byte) error {
	var res ChangeScheduledBackupDefaultBodyDetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
