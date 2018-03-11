// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// SearchResultResultsItems search result results items
// swagger:model searchResultResultsItems
type SearchResultResultsItems struct {

	// item type
	ItemType string `json:"item_type,omitempty"`

	// package
	Package *Package `json:"package,omitempty"`

	// repository
	Repository *Repository `json:"repository,omitempty"`

	// url
	URL strfmt.URI `json:"url,omitempty"`

	// weight
	Weight int64 `json:"weight,omitempty"`
}

// Validate validates this search result results items
func (m *SearchResultResultsItems) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateItemType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePackage(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateRepository(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateURL(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var searchResultResultsItemsTypeItemTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["repository","package"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		searchResultResultsItemsTypeItemTypePropEnum = append(searchResultResultsItemsTypeItemTypePropEnum, v)
	}
}

const (

	// SearchResultResultsItemsItemTypeRepository captures enum value "repository"
	SearchResultResultsItemsItemTypeRepository string = "repository"

	// SearchResultResultsItemsItemTypePackage captures enum value "package"
	SearchResultResultsItemsItemTypePackage string = "package"
)

// prop value enum
func (m *SearchResultResultsItems) validateItemTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, searchResultResultsItemsTypeItemTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *SearchResultResultsItems) validateItemType(formats strfmt.Registry) error {

	if swag.IsZero(m.ItemType) { // not required
		return nil
	}

	// value enum
	if err := m.validateItemTypeEnum("item_type", "body", m.ItemType); err != nil {
		return err
	}

	return nil
}

func (m *SearchResultResultsItems) validatePackage(formats strfmt.Registry) error {

	if swag.IsZero(m.Package) { // not required
		return nil
	}

	if m.Package != nil {

		if err := m.Package.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("package")
			}
			return err
		}

	}

	return nil
}

func (m *SearchResultResultsItems) validateRepository(formats strfmt.Registry) error {

	if swag.IsZero(m.Repository) { // not required
		return nil
	}

	if m.Repository != nil {

		if err := m.Repository.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("repository")
			}
			return err
		}

	}

	return nil
}

func (m *SearchResultResultsItems) validateURL(formats strfmt.Registry) error {

	if swag.IsZero(m.URL) { // not required
		return nil
	}

	if err := validate.FormatOf("url", "body", "uri", m.URL.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SearchResultResultsItems) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SearchResultResultsItems) UnmarshalBinary(b []byte) error {
	var res SearchResultResultsItems
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}