// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CreatesDomainOwnedByGivenOrganizationDeprecatedRequest creates domain owned by given organization deprecated request
//
// swagger:model createsDomainOwnedByGivenOrganizationDeprecatedRequest
type CreatesDomainOwnedByGivenOrganizationDeprecatedRequest struct {

	// The guid of the domain.
	GUID string `json:"guid,omitempty"`

	// The name of the domain.
	Name string `json:"name,omitempty"`

	// The organization that owns the domain. If not specified the domain is shared.
	OwningOrganizationGUID string `json:"owning_organization_guid,omitempty"`

	// Allow routes with non-empty hosts
	Wildcard bool `json:"wildcard,omitempty"`
}

// Validate validates this creates domain owned by given organization deprecated request
func (m *CreatesDomainOwnedByGivenOrganizationDeprecatedRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CreatesDomainOwnedByGivenOrganizationDeprecatedRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreatesDomainOwnedByGivenOrganizationDeprecatedRequest) UnmarshalBinary(b []byte) error {
	var res CreatesDomainOwnedByGivenOrganizationDeprecatedRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
