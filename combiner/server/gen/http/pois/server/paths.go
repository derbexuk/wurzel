// Code generated by goa v3.5.2, DO NOT EDIT.
//
// HTTP request path constructors for the pois service.
//
// Command:
// $ goa gen github.com/derbexuk/wurzel/combiner/server/design

package server

import (
	"fmt"
)

// PostPoisPath returns the URL path to the pois service post HTTP endpoint.
func PostPoisPath(path string) string {
	return fmt.Sprintf("/things/pois/%v", path)
}

// ShowPoisPath returns the URL path to the pois service show HTTP endpoint.
func ShowPoisPath(path string) string {
	return fmt.Sprintf("/things/pois/path/%v", path)
}

// ListByPathPoisPath returns the URL path to the pois service ListByPath HTTP endpoint.
func ListByPathPoisPath(path string) string {
	return fmt.Sprintf("/things/pois/paths/%v", path)
}

// ListByReferencePoisPath returns the URL path to the pois service ListByReference HTTP endpoint.
func ListByReferencePoisPath(path string) string {
	return fmt.Sprintf("/things/pois/refs/%v", path)
}

// UpdatePoisPath returns the URL path to the pois service update HTTP endpoint.
func UpdatePoisPath(path string) string {
	return fmt.Sprintf("/things/pois/update/%v", path)
}

// DeactivatePoisPath returns the URL path to the pois service deactivate HTTP endpoint.
func DeactivatePoisPath(path string) string {
	return fmt.Sprintf("/things/pois/deactivate/%v", path)
}

// DeletePoisPath returns the URL path to the pois service delete HTTP endpoint.
func DeletePoisPath(path string) string {
	return fmt.Sprintf("/things/pois/delete/%v", path)
}
