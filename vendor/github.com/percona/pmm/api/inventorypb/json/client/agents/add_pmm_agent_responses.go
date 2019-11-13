// Code generated by go-swagger; DO NOT EDIT.

package agents

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// AddPMMAgentReader is a Reader for the AddPMMAgent structure.
type AddPMMAgentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddPMMAgentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAddPMMAgentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAddPMMAgentDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAddPMMAgentOK creates a AddPMMAgentOK with default headers values
func NewAddPMMAgentOK() *AddPMMAgentOK {
	return &AddPMMAgentOK{}
}

/*AddPMMAgentOK handles this case with default header values.

A successful response.
*/
type AddPMMAgentOK struct {
	Payload *AddPMMAgentOKBody
}

func (o *AddPMMAgentOK) Error() string {
	return fmt.Sprintf("[POST /v1/inventory/Agents/AddPMMAgent][%d] addPmmAgentOk  %+v", 200, o.Payload)
}

func (o *AddPMMAgentOK) GetPayload() *AddPMMAgentOKBody {
	return o.Payload
}

func (o *AddPMMAgentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(AddPMMAgentOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddPMMAgentDefault creates a AddPMMAgentDefault with default headers values
func NewAddPMMAgentDefault(code int) *AddPMMAgentDefault {
	return &AddPMMAgentDefault{
		_statusCode: code,
	}
}

/*AddPMMAgentDefault handles this case with default header values.

An error response.
*/
type AddPMMAgentDefault struct {
	_statusCode int

	Payload *AddPMMAgentDefaultBody
}

// Code gets the status code for the add PMM agent default response
func (o *AddPMMAgentDefault) Code() int {
	return o._statusCode
}

func (o *AddPMMAgentDefault) Error() string {
	return fmt.Sprintf("[POST /v1/inventory/Agents/AddPMMAgent][%d] AddPMMAgent default  %+v", o._statusCode, o.Payload)
}

func (o *AddPMMAgentDefault) GetPayload() *AddPMMAgentDefaultBody {
	return o.Payload
}

func (o *AddPMMAgentDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(AddPMMAgentDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*AddPMMAgentBody add PMM agent body
swagger:model AddPMMAgentBody
*/
type AddPMMAgentBody struct {

	// Node identifier where this instance runs.
	RunsOnNodeID string `json:"runs_on_node_id,omitempty"`

	// Custom user-assigned labels.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`
}

// Validate validates this add PMM agent body
func (o *AddPMMAgentBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AddPMMAgentBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddPMMAgentBody) UnmarshalBinary(b []byte) error {
	var res AddPMMAgentBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*AddPMMAgentDefaultBody ErrorResponse is a message returned on HTTP error.
swagger:model AddPMMAgentDefaultBody
*/
type AddPMMAgentDefaultBody struct {

	// code
	Code int32 `json:"code,omitempty"`

	// error
	Error string `json:"error,omitempty"`

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this add PMM agent default body
func (o *AddPMMAgentDefaultBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AddPMMAgentDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddPMMAgentDefaultBody) UnmarshalBinary(b []byte) error {
	var res AddPMMAgentDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*AddPMMAgentOKBody add PMM agent OK body
swagger:model AddPMMAgentOKBody
*/
type AddPMMAgentOKBody struct {

	// pmm agent
	PMMAgent *AddPMMAgentOKBodyPMMAgent `json:"pmm_agent,omitempty"`
}

// Validate validates this add PMM agent OK body
func (o *AddPMMAgentOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validatePMMAgent(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AddPMMAgentOKBody) validatePMMAgent(formats strfmt.Registry) error {

	if swag.IsZero(o.PMMAgent) { // not required
		return nil
	}

	if o.PMMAgent != nil {
		if err := o.PMMAgent.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("addPmmAgentOk" + "." + "pmm_agent")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *AddPMMAgentOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddPMMAgentOKBody) UnmarshalBinary(b []byte) error {
	var res AddPMMAgentOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*AddPMMAgentOKBodyPMMAgent PMMAgent runs on Generic or Container Node.
swagger:model AddPMMAgentOKBodyPMMAgent
*/
type AddPMMAgentOKBodyPMMAgent struct {

	// Unique randomly generated instance identifier.
	AgentID string `json:"agent_id,omitempty"`

	// Node identifier where this instance runs.
	RunsOnNodeID string `json:"runs_on_node_id,omitempty"`

	// Custom user-assigned labels.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`

	// True if Agent is running and connected to pmm-managed.
	Connected bool `json:"connected,omitempty"`
}

// Validate validates this add PMM agent OK body PMM agent
func (o *AddPMMAgentOKBodyPMMAgent) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AddPMMAgentOKBodyPMMAgent) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddPMMAgentOKBodyPMMAgent) UnmarshalBinary(b []byte) error {
	var res AddPMMAgentOKBodyPMMAgent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
