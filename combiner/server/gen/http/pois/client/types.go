// Code generated by goa v3.5.2, DO NOT EDIT.
//
// pois HTTP client types
//
// Command:
// $ goa gen github.com/derbexuk/wurzel/combiner/server/design

package client

import (
	pois "github.com/derbexuk/wurzel/combiner/server/gen/pois"
	poisviews "github.com/derbexuk/wurzel/combiner/server/gen/pois/views"
	goa "goa.design/goa/v3/pkg"
)

// PostRequestBody is the type of the "pois" service "post" endpoint HTTP
// request body.
type PostRequestBody struct {
	Pois []*PoiPayloadRequestBody `form:"pois,omitempty" json:"pois,omitempty" xml:"pois,omitempty"`
}

// UpdateRequestBody is the type of the "pois" service "update" endpoint HTTP
// request body.
type UpdateRequestBody struct {
	Poi *PoiPayloadRequestBody `form:"poi,omitempty" json:"poi,omitempty" xml:"poi,omitempty"`
}

// ShowResponseBody is the type of the "pois" service "show" endpoint HTTP
// response body.
type ShowResponseBody struct {
	// ID of poi
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Poi title/name
	Title       *string `form:"title,omitempty" json:"title,omitempty" xml:"title,omitempty"`
	Description *string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	Deactivated *bool   `form:"deactivated,omitempty" json:"deactivated,omitempty" xml:"deactivated,omitempty"`
	// name of location
	Location *string `form:"location,omitempty" json:"location,omitempty" xml:"location,omitempty"`
	Geojson  *string `form:"geojson,omitempty" json:"geojson,omitempty" xml:"geojson,omitempty"`
	// Hash of application specific properties
	Properties map[string]string `form:"properties,omitempty" json:"properties,omitempty" xml:"properties,omitempty"`
	// refs
	Refs []string `form:"refs,omitempty" json:"refs,omitempty" xml:"refs,omitempty"`
	Path *string  `form:"path,omitempty" json:"path,omitempty" xml:"path,omitempty"`
}

// ListByPathResponseBody is the type of the "pois" service "ListByPath"
// endpoint HTTP response body.
type ListByPathResponseBody []*CspacePoiResponse

// ListByReferenceResponseBody is the type of the "pois" service
// "ListByReference" endpoint HTTP response body.
type ListByReferenceResponseBody []*CspacePoiResponse

// PostUnAuthResponseBody is the type of the "pois" service "post" endpoint
// HTTP response body for the "un_auth" error.
type PostUnAuthResponseBody struct {
	// Name is the name of this class of errors.
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// Is the error temporary?
	Temporary *bool `form:"temporary,omitempty" json:"temporary,omitempty" xml:"temporary,omitempty"`
	// Is the error a timeout?
	Timeout *bool `form:"timeout,omitempty" json:"timeout,omitempty" xml:"timeout,omitempty"`
	// Is the error a server-side fault?
	Fault *bool `form:"fault,omitempty" json:"fault,omitempty" xml:"fault,omitempty"`
}

// PostBadReqResponseBody is the type of the "pois" service "post" endpoint
// HTTP response body for the "bad_req" error.
type PostBadReqResponseBody struct {
	// Name is the name of this class of errors.
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// Is the error temporary?
	Temporary *bool `form:"temporary,omitempty" json:"temporary,omitempty" xml:"temporary,omitempty"`
	// Is the error a timeout?
	Timeout *bool `form:"timeout,omitempty" json:"timeout,omitempty" xml:"timeout,omitempty"`
	// Is the error a server-side fault?
	Fault *bool `form:"fault,omitempty" json:"fault,omitempty" xml:"fault,omitempty"`
}

// ShowNotFoundResponseBody is the type of the "pois" service "show" endpoint
// HTTP response body for the "not_found" error.
type ShowNotFoundResponseBody struct {
	// Name is the name of this class of errors.
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// Is the error temporary?
	Temporary *bool `form:"temporary,omitempty" json:"temporary,omitempty" xml:"temporary,omitempty"`
	// Is the error a timeout?
	Timeout *bool `form:"timeout,omitempty" json:"timeout,omitempty" xml:"timeout,omitempty"`
	// Is the error a server-side fault?
	Fault *bool `form:"fault,omitempty" json:"fault,omitempty" xml:"fault,omitempty"`
}

// ListByPathNotFoundResponseBody is the type of the "pois" service
// "ListByPath" endpoint HTTP response body for the "not_found" error.
type ListByPathNotFoundResponseBody struct {
	// Name is the name of this class of errors.
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// Is the error temporary?
	Temporary *bool `form:"temporary,omitempty" json:"temporary,omitempty" xml:"temporary,omitempty"`
	// Is the error a timeout?
	Timeout *bool `form:"timeout,omitempty" json:"timeout,omitempty" xml:"timeout,omitempty"`
	// Is the error a server-side fault?
	Fault *bool `form:"fault,omitempty" json:"fault,omitempty" xml:"fault,omitempty"`
}

// ListByReferenceNotFoundResponseBody is the type of the "pois" service
// "ListByReference" endpoint HTTP response body for the "not_found" error.
type ListByReferenceNotFoundResponseBody struct {
	// Name is the name of this class of errors.
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// Is the error temporary?
	Temporary *bool `form:"temporary,omitempty" json:"temporary,omitempty" xml:"temporary,omitempty"`
	// Is the error a timeout?
	Timeout *bool `form:"timeout,omitempty" json:"timeout,omitempty" xml:"timeout,omitempty"`
	// Is the error a server-side fault?
	Fault *bool `form:"fault,omitempty" json:"fault,omitempty" xml:"fault,omitempty"`
}

// UpdateBadReqResponseBody is the type of the "pois" service "update" endpoint
// HTTP response body for the "bad_req" error.
type UpdateBadReqResponseBody struct {
	// Name is the name of this class of errors.
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// Is the error temporary?
	Temporary *bool `form:"temporary,omitempty" json:"temporary,omitempty" xml:"temporary,omitempty"`
	// Is the error a timeout?
	Timeout *bool `form:"timeout,omitempty" json:"timeout,omitempty" xml:"timeout,omitempty"`
	// Is the error a server-side fault?
	Fault *bool `form:"fault,omitempty" json:"fault,omitempty" xml:"fault,omitempty"`
}

// DeactivateNotFoundResponseBody is the type of the "pois" service
// "deactivate" endpoint HTTP response body for the "not_found" error.
type DeactivateNotFoundResponseBody struct {
	// Name is the name of this class of errors.
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// Is the error temporary?
	Temporary *bool `form:"temporary,omitempty" json:"temporary,omitempty" xml:"temporary,omitempty"`
	// Is the error a timeout?
	Timeout *bool `form:"timeout,omitempty" json:"timeout,omitempty" xml:"timeout,omitempty"`
	// Is the error a server-side fault?
	Fault *bool `form:"fault,omitempty" json:"fault,omitempty" xml:"fault,omitempty"`
}

// DeleteNotFoundResponseBody is the type of the "pois" service "delete"
// endpoint HTTP response body for the "not_found" error.
type DeleteNotFoundResponseBody struct {
	// Name is the name of this class of errors.
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// Is the error temporary?
	Temporary *bool `form:"temporary,omitempty" json:"temporary,omitempty" xml:"temporary,omitempty"`
	// Is the error a timeout?
	Timeout *bool `form:"timeout,omitempty" json:"timeout,omitempty" xml:"timeout,omitempty"`
	// Is the error a server-side fault?
	Fault *bool `form:"fault,omitempty" json:"fault,omitempty" xml:"fault,omitempty"`
}

// PoiPayloadRequestBody is used to define fields on request body types.
type PoiPayloadRequestBody struct {
	// ID of poi
	ID string `form:"id" json:"id" xml:"id"`
	// Poi title/name
	Title       string  `form:"title" json:"title" xml:"title"`
	Description *string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	Deactivated *bool   `form:"deactivated,omitempty" json:"deactivated,omitempty" xml:"deactivated,omitempty"`
	// name of location
	Location *string             `form:"location,omitempty" json:"location,omitempty" xml:"location,omitempty"`
	Geopath  *GeoPathRequestBody `form:"geopath,omitempty" json:"geopath,omitempty" xml:"geopath,omitempty"`
	Geojson  string              `form:"geojson" json:"geojson" xml:"geojson"`
	// Hash of application specific properties
	Properties map[string]string `form:"properties,omitempty" json:"properties,omitempty" xml:"properties,omitempty"`
	// references
	Refs []string `form:"refs,omitempty" json:"refs,omitempty" xml:"refs,omitempty"`
	Path *string  `form:"path,omitempty" json:"path,omitempty" xml:"path,omitempty"`
}

// GeoPathRequestBody is used to define fields on request body types.
type GeoPathRequestBody struct {
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Path title/name
	Title *string `form:"title,omitempty" json:"title,omitempty" xml:"title,omitempty"`
	// must be route | point | area
	Type *string `form:"type,omitempty" json:"type,omitempty" xml:"type,omitempty"`
	// path
	Geopoints [][]float32 `form:"geopoints,omitempty" json:"geopoints,omitempty" xml:"geopoints,omitempty"`
}

// CspacePoiResponse is used to define fields on response body types.
type CspacePoiResponse struct {
	// ID of poi
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Poi title/name
	Title       *string `form:"title,omitempty" json:"title,omitempty" xml:"title,omitempty"`
	Description *string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	Deactivated *bool   `form:"deactivated,omitempty" json:"deactivated,omitempty" xml:"deactivated,omitempty"`
	// name of location
	Location *string `form:"location,omitempty" json:"location,omitempty" xml:"location,omitempty"`
	Geojson  *string `form:"geojson,omitempty" json:"geojson,omitempty" xml:"geojson,omitempty"`
	// Hash of application specific properties
	Properties map[string]string `form:"properties,omitempty" json:"properties,omitempty" xml:"properties,omitempty"`
	// refs
	Refs []string `form:"refs,omitempty" json:"refs,omitempty" xml:"refs,omitempty"`
	Path *string  `form:"path,omitempty" json:"path,omitempty" xml:"path,omitempty"`
}

// NewPostRequestBody builds the HTTP request body from the payload of the
// "post" endpoint of the "pois" service.
func NewPostRequestBody(p *pois.PostPayload) *PostRequestBody {
	body := &PostRequestBody{}
	if p.Pois != nil {
		body.Pois = make([]*PoiPayloadRequestBody, len(p.Pois))
		for i, val := range p.Pois {
			body.Pois[i] = marshalPoisPoiPayloadToPoiPayloadRequestBody(val)
		}
	}
	return body
}

// NewUpdateRequestBody builds the HTTP request body from the payload of the
// "update" endpoint of the "pois" service.
func NewUpdateRequestBody(p *pois.UpdatePayload) *UpdateRequestBody {
	body := &UpdateRequestBody{}
	if p.Poi != nil {
		body.Poi = marshalPoisPoiPayloadToPoiPayloadRequestBody(p.Poi)
	}
	return body
}

// NewPostUnAuth builds a pois service post endpoint un_auth error.
func NewPostUnAuth(body *PostUnAuthResponseBody) *goa.ServiceError {
	v := &goa.ServiceError{
		Name:      *body.Name,
		ID:        *body.ID,
		Message:   *body.Message,
		Temporary: *body.Temporary,
		Timeout:   *body.Timeout,
		Fault:     *body.Fault,
	}

	return v
}

// NewPostBadReq builds a pois service post endpoint bad_req error.
func NewPostBadReq(body *PostBadReqResponseBody) *goa.ServiceError {
	v := &goa.ServiceError{
		Name:      *body.Name,
		ID:        *body.ID,
		Message:   *body.Message,
		Temporary: *body.Temporary,
		Timeout:   *body.Timeout,
		Fault:     *body.Fault,
	}

	return v
}

// NewShowCspacePoiOK builds a "pois" service "show" endpoint result from a
// HTTP "OK" response.
func NewShowCspacePoiOK(body *ShowResponseBody) *poisviews.CspacePoiView {
	v := &poisviews.CspacePoiView{
		ID:          body.ID,
		Title:       body.Title,
		Description: body.Description,
		Deactivated: body.Deactivated,
		Location:    body.Location,
		Geojson:     body.Geojson,
		Path:        body.Path,
	}
	if body.Properties != nil {
		v.Properties = make(map[string]string, len(body.Properties))
		for key, val := range body.Properties {
			tk := key
			tv := val
			v.Properties[tk] = tv
		}
	}
	if body.Refs != nil {
		v.Refs = make([]string, len(body.Refs))
		for i, val := range body.Refs {
			v.Refs[i] = val
		}
	}

	return v
}

// NewShowNotFound builds a pois service show endpoint not_found error.
func NewShowNotFound(body *ShowNotFoundResponseBody) *goa.ServiceError {
	v := &goa.ServiceError{
		Name:      *body.Name,
		ID:        *body.ID,
		Message:   *body.Message,
		Temporary: *body.Temporary,
		Timeout:   *body.Timeout,
		Fault:     *body.Fault,
	}

	return v
}

// NewListByPathCspacePoiCollectionOK builds a "pois" service "ListByPath"
// endpoint result from a HTTP "OK" response.
func NewListByPathCspacePoiCollectionOK(body ListByPathResponseBody) poisviews.CspacePoiCollectionView {
	v := make([]*poisviews.CspacePoiView, len(body))
	for i, val := range body {
		v[i] = unmarshalCspacePoiResponseToPoisviewsCspacePoiView(val)
	}

	return v
}

// NewListByPathNotFound builds a pois service ListByPath endpoint not_found
// error.
func NewListByPathNotFound(body *ListByPathNotFoundResponseBody) *goa.ServiceError {
	v := &goa.ServiceError{
		Name:      *body.Name,
		ID:        *body.ID,
		Message:   *body.Message,
		Temporary: *body.Temporary,
		Timeout:   *body.Timeout,
		Fault:     *body.Fault,
	}

	return v
}

// NewListByReferenceCspacePoiCollectionOK builds a "pois" service
// "ListByReference" endpoint result from a HTTP "OK" response.
func NewListByReferenceCspacePoiCollectionOK(body ListByReferenceResponseBody) poisviews.CspacePoiCollectionView {
	v := make([]*poisviews.CspacePoiView, len(body))
	for i, val := range body {
		v[i] = unmarshalCspacePoiResponseToPoisviewsCspacePoiView(val)
	}

	return v
}

// NewListByReferenceNotFound builds a pois service ListByReference endpoint
// not_found error.
func NewListByReferenceNotFound(body *ListByReferenceNotFoundResponseBody) *goa.ServiceError {
	v := &goa.ServiceError{
		Name:      *body.Name,
		ID:        *body.ID,
		Message:   *body.Message,
		Temporary: *body.Temporary,
		Timeout:   *body.Timeout,
		Fault:     *body.Fault,
	}

	return v
}

// NewUpdateBadReq builds a pois service update endpoint bad_req error.
func NewUpdateBadReq(body *UpdateBadReqResponseBody) *goa.ServiceError {
	v := &goa.ServiceError{
		Name:      *body.Name,
		ID:        *body.ID,
		Message:   *body.Message,
		Temporary: *body.Temporary,
		Timeout:   *body.Timeout,
		Fault:     *body.Fault,
	}

	return v
}

// NewDeactivateNotFound builds a pois service deactivate endpoint not_found
// error.
func NewDeactivateNotFound(body *DeactivateNotFoundResponseBody) *goa.ServiceError {
	v := &goa.ServiceError{
		Name:      *body.Name,
		ID:        *body.ID,
		Message:   *body.Message,
		Temporary: *body.Temporary,
		Timeout:   *body.Timeout,
		Fault:     *body.Fault,
	}

	return v
}

// NewDeleteNotFound builds a pois service delete endpoint not_found error.
func NewDeleteNotFound(body *DeleteNotFoundResponseBody) *goa.ServiceError {
	v := &goa.ServiceError{
		Name:      *body.Name,
		ID:        *body.ID,
		Message:   *body.Message,
		Temporary: *body.Temporary,
		Timeout:   *body.Timeout,
		Fault:     *body.Fault,
	}

	return v
}

// ValidatePostUnAuthResponseBody runs the validations defined on
// post_un_auth_response_body
func ValidatePostUnAuthResponseBody(body *PostUnAuthResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.Temporary == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("temporary", "body"))
	}
	if body.Timeout == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("timeout", "body"))
	}
	if body.Fault == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("fault", "body"))
	}
	return
}

// ValidatePostBadReqResponseBody runs the validations defined on
// post_bad_req_response_body
func ValidatePostBadReqResponseBody(body *PostBadReqResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.Temporary == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("temporary", "body"))
	}
	if body.Timeout == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("timeout", "body"))
	}
	if body.Fault == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("fault", "body"))
	}
	return
}

// ValidateShowNotFoundResponseBody runs the validations defined on
// show_not_found_response_body
func ValidateShowNotFoundResponseBody(body *ShowNotFoundResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.Temporary == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("temporary", "body"))
	}
	if body.Timeout == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("timeout", "body"))
	}
	if body.Fault == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("fault", "body"))
	}
	return
}

// ValidateListByPathNotFoundResponseBody runs the validations defined on
// ListByPath_not_found_Response_Body
func ValidateListByPathNotFoundResponseBody(body *ListByPathNotFoundResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.Temporary == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("temporary", "body"))
	}
	if body.Timeout == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("timeout", "body"))
	}
	if body.Fault == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("fault", "body"))
	}
	return
}

// ValidateListByReferenceNotFoundResponseBody runs the validations defined on
// ListByReference_not_found_Response_Body
func ValidateListByReferenceNotFoundResponseBody(body *ListByReferenceNotFoundResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.Temporary == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("temporary", "body"))
	}
	if body.Timeout == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("timeout", "body"))
	}
	if body.Fault == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("fault", "body"))
	}
	return
}

// ValidateUpdateBadReqResponseBody runs the validations defined on
// update_bad_req_response_body
func ValidateUpdateBadReqResponseBody(body *UpdateBadReqResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.Temporary == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("temporary", "body"))
	}
	if body.Timeout == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("timeout", "body"))
	}
	if body.Fault == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("fault", "body"))
	}
	return
}

// ValidateDeactivateNotFoundResponseBody runs the validations defined on
// deactivate_not_found_response_body
func ValidateDeactivateNotFoundResponseBody(body *DeactivateNotFoundResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.Temporary == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("temporary", "body"))
	}
	if body.Timeout == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("timeout", "body"))
	}
	if body.Fault == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("fault", "body"))
	}
	return
}

// ValidateDeleteNotFoundResponseBody runs the validations defined on
// delete_not_found_response_body
func ValidateDeleteNotFoundResponseBody(body *DeleteNotFoundResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.Temporary == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("temporary", "body"))
	}
	if body.Timeout == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("timeout", "body"))
	}
	if body.Fault == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("fault", "body"))
	}
	return
}

// ValidateCspacePoiResponse runs the validations defined on CspacePoiResponse
func ValidateCspacePoiResponse(body *CspacePoiResponse) (err error) {
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Title == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("title", "body"))
	}
	if body.Geojson == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("geojson", "body"))
	}
	return
}
