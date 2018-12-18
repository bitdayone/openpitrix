// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// OpenpitrixAppVersion openpitrix app version
// swagger:model openpitrixAppVersion
type OpenpitrixAppVersion struct {

	// active
	Active bool `json:"active,omitempty"`

	// app id
	AppID string `json:"app_id,omitempty"`

	// create time
	CreateTime strfmt.DateTime `json:"create_time,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// home
	Home string `json:"home,omitempty"`

	// icon
	Icon string `json:"icon,omitempty"`

	// keywords
	Keywords string `json:"keywords,omitempty"`

	// maintainers
	Maintainers string `json:"maintainers,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// owner
	Owner string `json:"owner,omitempty"`

	// package name
	PackageName string `json:"package_name,omitempty"`

	// readme
	Readme string `json:"readme,omitempty"`

	// review id
	ReviewID string `json:"review_id,omitempty"`

	// screenshots
	Screenshots string `json:"screenshots,omitempty"`

	// sequence
	Sequence int64 `json:"sequence,omitempty"`

	// sources
	Sources string `json:"sources,omitempty"`

	// status
	Status string `json:"status,omitempty"`

	// status time
	StatusTime strfmt.DateTime `json:"status_time,omitempty"`

	// type
	Type string `json:"type,omitempty"`

	// update time
	UpdateTime strfmt.DateTime `json:"update_time,omitempty"`

	// version id
	VersionID string `json:"version_id,omitempty"`
}

// Validate validates this openpitrix app version
func (m *OpenpitrixAppVersion) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *OpenpitrixAppVersion) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenpitrixAppVersion) UnmarshalBinary(b []byte) error {
	var res OpenpitrixAppVersion
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
