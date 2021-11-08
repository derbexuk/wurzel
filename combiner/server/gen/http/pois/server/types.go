// Code generated by goa v3.5.2, DO NOT EDIT.
//
// pois HTTP server types
//
// Command:
// $ goa gen github.com/derbexuk/wurzel/combiner/server/design

package server

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
	ID string `form:"id" json:"id" xml:"id"`
	// Poi title/name
	Title       string  `form:"title" json:"title" xml:"title"`
	Description *string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	Deactivated *bool   `form:"deactivated,omitempty" json:"deactivated,omitempty" xml:"deactivated,omitempty"`
	// name of location
	Location *string `form:"location,omitempty" json:"location,omitempty" xml:"location,omitempty"`
	Geojson  string  `form:"geojson" json:"geojson" xml:"geojson"`
	// Hash of application specific properties
	Properties map[string]string `form:"properties,omitempty" json:"properties,omitempty" xml:"properties,omitempty"`
	// refs
	Refs []string `form:"refs,omitempty" json:"refs,omitempty" xml:"refs,omitempty"`
	Path *string  `form:"path,omitempty" json:"path,omitempty" xml:"path,omitempty"`
}

// ShowResponseBodySummary is the type of the "pois" service "show" endpoint
// HTTP response body.
type ShowResponseBodySummary struct {
	// ID of poi
	ID string `form:"id" json:"id" xml:"id"`
	// Poi title/name
	Title   string  `form:"title" json:"title" xml:"title"`
	Geojson string  `form:"geojson" json:"geojson" xml:"geojson"`
	Path    *string `form:"path,omitempty" json:"path,omitempty" xml:"path,omitempty"`
}

// CspacePoiResponseCollection is the type of the "pois" service "ListByPath"
// endpoint HTTP response body.
type CspacePoiResponseCollection []*CspacePoiResponse

// CspacePoiResponseSummaryCollection is the type of the "pois" service
// "ListByPath" endpoint HTTP response body.
type CspacePoiResponseSummaryCollection []*CspacePoiResponseSummary

// PostUnAuthResponseBody is the type of the "pois" service "post" endpoint
// HTTP response body for the "un_auth" error.
type PostUnAuthResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// PostBadReqResponseBody is the type of the "pois" service "post" endpoint
// HTTP response body for the "bad_req" error.
type PostBadReqResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// ShowNotFoundResponseBody is the type of the "pois" service "show" endpoint
// HTTP response body for the "not_found" error.
type ShowNotFoundResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// ListByPathNotFoundResponseBody is the type of the "pois" service
// "ListByPath" endpoint HTTP response body for the "not_found" error.
type ListByPathNotFoundResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// ListByReferenceNotFoundResponseBody is the type of the "pois" service
// "ListByReference" endpoint HTTP response body for the "not_found" error.
type ListByReferenceNotFoundResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// UpdateBadReqResponseBody is the type of the "pois" service "update" endpoint
// HTTP response body for the "bad_req" error.
type UpdateBadReqResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// DeactivateNotFoundResponseBody is the type of the "pois" service
// "deactivate" endpoint HTTP response body for the "not_found" error.
type DeactivateNotFoundResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// DeleteNotFoundResponseBody is the type of the "pois" service "delete"
// endpoint HTTP response body for the "not_found" error.
type DeleteNotFoundResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// CspacePoiResponse is used to define fields on response body types.
type CspacePoiResponse struct {
	// ID of poi
	ID string `form:"id" json:"id" xml:"id"`
	// Poi title/name
	Title       string  `form:"title" json:"title" xml:"title"`
	Description *string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	Deactivated *bool   `form:"deactivated,omitempty" json:"deactivated,omitempty" xml:"deactivated,omitempty"`
	// name of location
	Location *string `form:"location,omitempty" json:"location,omitempty" xml:"location,omitempty"`
	Geojson  string  `form:"geojson" json:"geojson" xml:"geojson"`
	// Hash of application specific properties
	Properties map[string]string `form:"properties,omitempty" json:"properties,omitempty" xml:"properties,omitempty"`
	// refs
	Refs []string `form:"refs,omitempty" json:"refs,omitempty" xml:"refs,omitempty"`
	Path *string  `form:"path,omitempty" json:"path,omitempty" xml:"path,omitempty"`
}

// CspacePoiResponseSummary is used to define fields on response body types.
type CspacePoiResponseSummary struct {
	// ID of poi
	ID string `form:"id" json:"id" xml:"id"`
	// Poi title/name
	Title   string  `form:"title" json:"title" xml:"title"`
	Geojson string  `form:"geojson" json:"geojson" xml:"geojson"`
	Path    *string `form:"path,omitempty" json:"path,omitempty" xml:"path,omitempty"`
}

// PoiPayloadRequestBody is used to define fields on request body types.
type PoiPayloadRequestBody struct {
	// ID of poi
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Poi title/name
	Title       *string `form:"title,omitempty" json:"title,omitempty" xml:"title,omitempty"`
	Description *string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	Deactivated *bool   `form:"deactivated,omitempty" json:"deactivated,omitempty" xml:"deactivated,omitempty"`
	// name of location
	Location *string             `form:"location,omitempty" json:"location,omitempty" xml:"location,omitempty"`
	Geopath  *GeoPathRequestBody `form:"geopath,omitempty" json:"geopath,omitempty" xml:"geopath,omitempty"`
	Geojson  *string             `form:"geojson,omitempty" json:"geojson,omitempty" xml:"geojson,omitempty"`
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

// NewShowResponseBody builds the HTTP response body from the result of the
// "show" endpoint of the "pois" service.
func NewShowResponseBody(res *poisviews.CspacePoiView) *ShowResponseBody {
	body := &ShowResponseBody{
		ID:          *res.ID,
		Title:       *res.Title,
		Description: res.Description,
		Deactivated: res.Deactivated,
		Location:    res.Location,
		Geojson:     *res.Geojson,
		Path:        res.Path,
	}
	if res.Properties != nil {
		body.Properties = make(map[string]string, len(res.Properties))
		for key, val := range res.Properties {
			tk := key
			tv := val
			body.Properties[tk] = tv
		}
	}
	if res.Refs != nil {
		body.Refs = make([]string, len(res.Refs))
		for i, val := range res.Refs {
			body.Refs[i] = val
		}
	}
	return body
}

// NewShowResponseBodySummary builds the HTTP response body from the result of
// the "show" endpoint of the "pois" service.
func NewShowResponseBodySummary(res *poisviews.CspacePoiView) *ShowResponseBodySummary {
	body := &ShowResponseBodySummary{
		ID:      *res.ID,
		Title:   *res.Title,
		Geojson: *res.Geojson,
		Path:    res.Path,
	}
	return body
}

// NewCspacePoiResponseCollection builds the HTTP response body from the result
// of the "ListByPath" endpoint of the "pois" service.
func NewCspacePoiResponseCollection(res poisviews.CspacePoiCollectionView) CspacePoiResponseCollection {
	body := make([]*CspacePoiResponse, len(res))
	for i, val := range res {
		body[i] = marshalPoisviewsCspacePoiViewToCspacePoiResponse(val)
	}
	return body
}

// NewCspacePoiResponseSummaryCollection builds the HTTP response body from the
// result of the "ListByPath" endpoint of the "pois" service.
func NewCspacePoiResponseSummaryCollection(res poisviews.CspacePoiCollectionView) CspacePoiResponseSummaryCollection {
	body := make([]*CspacePoiResponseSummary, len(res))
	for i, val := range res {
		body[i] = marshalPoisviewsCspacePoiViewToCspacePoiResponseSummary(val)
	}
	return body
}

// NewPostUnAuthResponseBody builds the HTTP response body from the result of
// the "post" endpoint of the "pois" service.
func NewPostUnAuthResponseBody(res *goa.ServiceError) *PostUnAuthResponseBody {
	body := &PostUnAuthResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewPostBadReqResponseBody builds the HTTP response body from the result of
// the "post" endpoint of the "pois" service.
func NewPostBadReqResponseBody(res *goa.ServiceError) *PostBadReqResponseBody {
	body := &PostBadReqResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewShowNotFoundResponseBody builds the HTTP response body from the result of
// the "show" endpoint of the "pois" service.
func NewShowNotFoundResponseBody(res *goa.ServiceError) *ShowNotFoundResponseBody {
	body := &ShowNotFoundResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewListByPathNotFoundResponseBody builds the HTTP response body from the
// result of the "ListByPath" endpoint of the "pois" service.
func NewListByPathNotFoundResponseBody(res *goa.ServiceError) *ListByPathNotFoundResponseBody {
	body := &ListByPathNotFoundResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewListByReferenceNotFoundResponseBody builds the HTTP response body from
// the result of the "ListByReference" endpoint of the "pois" service.
func NewListByReferenceNotFoundResponseBody(res *goa.ServiceError) *ListByReferenceNotFoundResponseBody {
	body := &ListByReferenceNotFoundResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewUpdateBadReqResponseBody builds the HTTP response body from the result of
// the "update" endpoint of the "pois" service.
func NewUpdateBadReqResponseBody(res *goa.ServiceError) *UpdateBadReqResponseBody {
	body := &UpdateBadReqResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewDeactivateNotFoundResponseBody builds the HTTP response body from the
// result of the "deactivate" endpoint of the "pois" service.
func NewDeactivateNotFoundResponseBody(res *goa.ServiceError) *DeactivateNotFoundResponseBody {
	body := &DeactivateNotFoundResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewDeleteNotFoundResponseBody builds the HTTP response body from the result
// of the "delete" endpoint of the "pois" service.
func NewDeleteNotFoundResponseBody(res *goa.ServiceError) *DeleteNotFoundResponseBody {
	body := &DeleteNotFoundResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewPostPayload builds a pois service post endpoint payload.
func NewPostPayload(body *PostRequestBody, path string) *pois.PostPayload {
	v := &pois.PostPayload{}
	if body.Pois != nil {
		v.Pois = make([]*pois.PoiPayload, len(body.Pois))
		for i, val := range body.Pois {
			v.Pois[i] = unmarshalPoiPayloadRequestBodyToPoisPoiPayload(val)
		}
	}
	v.Path = path

	return v
}

// NewShowPayload builds a pois service show endpoint payload.
func NewShowPayload(path string) *pois.ShowPayload {
	v := &pois.ShowPayload{}
	v.Path = path

	return v
}

// NewListByPathPayload builds a pois service ListByPath endpoint payload.
func NewListByPathPayload(path string) *pois.ListByPathPayload {
	v := &pois.ListByPathPayload{}
	v.Path = path

	return v
}

// NewListByReferencePayload builds a pois service ListByReference endpoint
// payload.
func NewListByReferencePayload(path string) *pois.ListByReferencePayload {
	v := &pois.ListByReferencePayload{}
	v.Path = path

	return v
}

// NewUpdatePayload builds a pois service update endpoint payload.
func NewUpdatePayload(body *UpdateRequestBody, path string) *pois.UpdatePayload {
	v := &pois.UpdatePayload{}
	if body.Poi != nil {
		v.Poi = unmarshalPoiPayloadRequestBodyToPoisPoiPayload(body.Poi)
	}
	v.Path = path

	return v
}

// NewDeactivatePayload builds a pois service deactivate endpoint payload.
func NewDeactivatePayload(path string) *pois.DeactivatePayload {
	v := &pois.DeactivatePayload{}
	v.Path = path

	return v
}

// NewDeletePayload builds a pois service delete endpoint payload.
func NewDeletePayload(path string) *pois.DeletePayload {
	v := &pois.DeletePayload{}
	v.Path = path

	return v
}

// ValidatePostRequestBody runs the validations defined on PostRequestBody
func ValidatePostRequestBody(body *PostRequestBody) (err error) {
	for _, e := range body.Pois {
		if e != nil {
			if err2 := ValidatePoiPayloadRequestBody(e); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// ValidateUpdateRequestBody runs the validations defined on UpdateRequestBody
func ValidateUpdateRequestBody(body *UpdateRequestBody) (err error) {
	if body.Poi != nil {
		if err2 := ValidatePoiPayloadRequestBody(body.Poi); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidatePoiPayloadRequestBody runs the validations defined on
// PoiPayloadRequestBody
func ValidatePoiPayloadRequestBody(body *PoiPayloadRequestBody) (err error) {
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
