// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ListOryAccessControlPolicies list ory access control policies
// swagger:model listOryAccessControlPolicies
type ListOryAccessControlPolicies struct {

	// The ORY Access Control Policy flavor. Can be "regex", "glob", and "exact"
	//
	// in: path
	// Required: true
	Flavor *string `json:"flavor"`

	// The maximum amount of policies returned.
	//
	// in: query
	Limit int64 `json:"limit,omitempty"`

	// The offset from where to start looking.
	//
	// in: query
	Offset int64 `json:"offset,omitempty"`
}

// Validate validates this list ory access control policies
func (m *ListOryAccessControlPolicies) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFlavor(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ListOryAccessControlPolicies) validateFlavor(formats strfmt.Registry) error {

	if err := validate.Required("flavor", "body", m.Flavor); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ListOryAccessControlPolicies) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ListOryAccessControlPolicies) UnmarshalBinary(b []byte) error {
	var res ListOryAccessControlPolicies
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
