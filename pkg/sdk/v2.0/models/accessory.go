// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Accessory The accessory of the artifact
//
// swagger:model Accessory
type Accessory struct {

	// The artifact id of the accessory
	ArtifactID int64 `json:"artifact_id"`

	// The creation time of the accessory
	// Format: date-time
	CreationTime strfmt.DateTime `json:"creation_time,omitempty"`

	// The artifact digest of the accessory
	Digest string `json:"digest"`

	// The icon of the accessory
	Icon string `json:"icon"`

	// The ID of the accessory
	ID int64 `json:"id,omitempty"`

	// The artifact size of the accessory
	Size int64 `json:"size"`

	// The subject artifact id of the accessory
	SubjectArtifactID int64 `json:"subject_artifact_id"`

	// The artifact size of the accessory
	Type string `json:"type"`
}

// Validate validates this accessory
func (m *Accessory) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreationTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Accessory) validateCreationTime(formats strfmt.Registry) error {

	if swag.IsZero(m.CreationTime) { // not required
		return nil
	}

	if err := validate.FormatOf("creation_time", "body", "date-time", m.CreationTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Accessory) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Accessory) UnmarshalBinary(b []byte) error {
	var res Accessory
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
