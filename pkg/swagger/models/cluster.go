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

// Cluster cluster
// swagger:model Cluster

type Cluster struct {

	// app id
	AppID string `json:"app_id,omitempty"`

	// app version
	AppVersion string `json:"app_version,omitempty"`

	// created
	// Read Only: true
	Created strfmt.DateTime `json:"created,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// id
	// Required: true
	// Max Length: 12
	// Min Length: 12
	// Pattern: cl-[a-zA-Z0-9]{9}
	ID *string `json:"id"`

	// last modified
	// Read Only: true
	LastModified strfmt.DateTime `json:"last_modified,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// status
	Status string `json:"status,omitempty"`

	// transition status
	TransitionStatus string `json:"transition_status,omitempty"`
}

/* polymorph Cluster app_id false */

/* polymorph Cluster app_version false */

/* polymorph Cluster created false */

/* polymorph Cluster description false */

/* polymorph Cluster id false */

/* polymorph Cluster last_modified false */

/* polymorph Cluster name false */

/* polymorph Cluster status false */

/* polymorph Cluster transition_status false */

// Validate validates this cluster
func (m *Cluster) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Cluster) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	if err := validate.MinLength("id", "body", string(*m.ID), 12); err != nil {
		return err
	}

	if err := validate.MaxLength("id", "body", string(*m.ID), 12); err != nil {
		return err
	}

	if err := validate.Pattern("id", "body", string(*m.ID), `cl-[a-zA-Z0-9]{9}`); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Cluster) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Cluster) UnmarshalBinary(b []byte) error {
	var res Cluster
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}