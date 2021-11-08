// Code generated by goa v3.5.2, DO NOT EDIT.
//
// HTTP request path constructors for the organisms service.
//
// Command:
// $ goa gen github.com/derbexuk/poieventservice/server/design

package client

import (
	"fmt"
)

// PostOrganismsPath returns the URL path to the organisms service post HTTP endpoint.
func PostOrganismsPath(path string) string {
	return fmt.Sprintf("/things/organisms/%v", path)
}

// ShowOrganismsPath returns the URL path to the organisms service show HTTP endpoint.
func ShowOrganismsPath(path string) string {
	return fmt.Sprintf("/things/organisms/path/%v", path)
}

// UpdateOrganismsPath returns the URL path to the organisms service update HTTP endpoint.
func UpdateOrganismsPath(path string) string {
	return fmt.Sprintf("/things/organisms/update/%v", path)
}

// DeleteOrganismsPath returns the URL path to the organisms service delete HTTP endpoint.
func DeleteOrganismsPath(path string) string {
	return fmt.Sprintf("/things/organisms/delete/%v", path)
}

// DeactivateOrganismsPath returns the URL path to the organisms service deactivate HTTP endpoint.
func DeactivateOrganismsPath(path string) string {
	return fmt.Sprintf("/things/organisms/deactivate/%v", path)
}

// ListByPathOrganismsPath returns the URL path to the organisms service ListByPath HTTP endpoint.
func ListByPathOrganismsPath(path string) string {
	return fmt.Sprintf("/things/organisms/list/%v", path)
}

// ListByReferenceOrganismsPath returns the URL path to the organisms service ListByReference HTTP endpoint.
func ListByReferenceOrganismsPath(path string) string {
	return fmt.Sprintf("/things/organisms/refs/%v", path)
}
