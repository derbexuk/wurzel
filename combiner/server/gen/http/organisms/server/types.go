// Code generated by goa v3.5.2, DO NOT EDIT.
//
// organisms HTTP server types
//
// Command:
// $ goa gen github.com/derbexuk/poieventservice/server/design

package server

import (
	organisms "github.com/derbexuk/poieventservice/server/gen/organisms"
	organismsviews "github.com/derbexuk/poieventservice/server/gen/organisms/views"
	goa "goa.design/goa/v3/pkg"
)

// PostRequestBody is the type of the "organisms" service "post" endpoint HTTP
// request body.
type PostRequestBody struct {
	Organisms []*OrganismPayloadRequestBody `form:"organisms,omitempty" json:"organisms,omitempty" xml:"organisms,omitempty"`
}

// UpdateRequestBody is the type of the "organisms" service "update" endpoint
// HTTP request body.
type UpdateRequestBody struct {
	Organism *OrganismPayloadRequestBody `form:"organism,omitempty" json:"organism,omitempty" xml:"organism,omitempty"`
}

// ShowResponseBody is the type of the "organisms" service "show" endpoint HTTP
// response body.
type ShowResponseBody struct {
	// ID of Organism
	ID string `form:"id" json:"id" xml:"id"`
	// Organism title
	Title       string  `form:"title" json:"title" xml:"title"`
	Description *string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	Deactivated *bool   `form:"deactivated,omitempty" json:"deactivated,omitempty" xml:"deactivated,omitempty"`
	// Hash of application specific properties
	Properties map[string]string `form:"properties,omitempty" json:"properties,omitempty" xml:"properties,omitempty"`
	// refs
	Refs []string `form:"refs,omitempty" json:"refs,omitempty" xml:"refs,omitempty"`
	Path *string  `form:"path,omitempty" json:"path,omitempty" xml:"path,omitempty"`
}

// ShowResponseBodySummary is the type of the "organisms" service "show"
// endpoint HTTP response body.
type ShowResponseBodySummary struct {
	// ID of Organism
	ID string `form:"id" json:"id" xml:"id"`
	// Organism title
	Title string  `form:"title" json:"title" xml:"title"`
	Path  *string `form:"path,omitempty" json:"path,omitempty" xml:"path,omitempty"`
}

// CspaceOrganismResponseCollection is the type of the "organisms" service
// "ListByPath" endpoint HTTP response body.
type CspaceOrganismResponseCollection []*CspaceOrganismResponse

// CspaceOrganismResponseSummaryCollection is the type of the "organisms"
// service "ListByPath" endpoint HTTP response body.
type CspaceOrganismResponseSummaryCollection []*CspaceOrganismResponseSummary

// PostBadReqResponseBody is the type of the "organisms" service "post"
// endpoint HTTP response body for the "bad_req" error.
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

// ShowNotFoundResponseBody is the type of the "organisms" service "show"
// endpoint HTTP response body for the "not_found" error.
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

// UpdateBadReqResponseBody is the type of the "organisms" service "update"
// endpoint HTTP response body for the "bad_req" error.
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

// DeleteNotFoundResponseBody is the type of the "organisms" service "delete"
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

// DeactivateNotFoundResponseBody is the type of the "organisms" service
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

// ListByPathNotFoundResponseBody is the type of the "organisms" service
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

// ListByReferenceNotFoundResponseBody is the type of the "organisms" service
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

// CspaceOrganismResponse is used to define fields on response body types.
type CspaceOrganismResponse struct {
	// ID of Organism
	ID string `form:"id" json:"id" xml:"id"`
	// Organism title
	Title       string  `form:"title" json:"title" xml:"title"`
	Description *string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	Deactivated *bool   `form:"deactivated,omitempty" json:"deactivated,omitempty" xml:"deactivated,omitempty"`
	// Hash of application specific properties
	Properties map[string]string `form:"properties,omitempty" json:"properties,omitempty" xml:"properties,omitempty"`
	// refs
	Refs []string `form:"refs,omitempty" json:"refs,omitempty" xml:"refs,omitempty"`
	Path *string  `form:"path,omitempty" json:"path,omitempty" xml:"path,omitempty"`
}

// CspaceOrganismResponseSummary is used to define fields on response body
// types.
type CspaceOrganismResponseSummary struct {
	// ID of Organism
	ID string `form:"id" json:"id" xml:"id"`
	// Organism title
	Title string  `form:"title" json:"title" xml:"title"`
	Path  *string `form:"path,omitempty" json:"path,omitempty" xml:"path,omitempty"`
}

// OrganismPayloadRequestBody is used to define fields on request body types.
type OrganismPayloadRequestBody struct {
	// ID of Organism
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Organism title
	Title       *string `form:"title,omitempty" json:"title,omitempty" xml:"title,omitempty"`
	Description *string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	Deactivated *bool   `form:"deactivated,omitempty" json:"deactivated,omitempty" xml:"deactivated,omitempty"`
	// Hash of application specific properties
	Properties map[string]string `form:"properties,omitempty" json:"properties,omitempty" xml:"properties,omitempty"`
	// references
	Refs []string `form:"refs,omitempty" json:"refs,omitempty" xml:"refs,omitempty"`
	Path *string  `form:"path,omitempty" json:"path,omitempty" xml:"path,omitempty"`
}

// NewShowResponseBody builds the HTTP response body from the result of the
// "show" endpoint of the "organisms" service.
func NewShowResponseBody(res *organismsviews.CspaceOrganismView) *ShowResponseBody {
	body := &ShowResponseBody{
		ID:          *res.ID,
		Title:       *res.Title,
		Description: res.Description,
		Deactivated: res.Deactivated,
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
// the "show" endpoint of the "organisms" service.
func NewShowResponseBodySummary(res *organismsviews.CspaceOrganismView) *ShowResponseBodySummary {
	body := &ShowResponseBodySummary{
		ID:    *res.ID,
		Title: *res.Title,
		Path:  res.Path,
	}
	return body
}

// NewCspaceOrganismResponseCollection builds the HTTP response body from the
// result of the "ListByPath" endpoint of the "organisms" service.
func NewCspaceOrganismResponseCollection(res organismsviews.CspaceOrganismCollectionView) CspaceOrganismResponseCollection {
	body := make([]*CspaceOrganismResponse, len(res))
	for i, val := range res {
		body[i] = marshalOrganismsviewsCspaceOrganismViewToCspaceOrganismResponse(val)
	}
	return body
}

// NewCspaceOrganismResponseSummaryCollection builds the HTTP response body
// from the result of the "ListByPath" endpoint of the "organisms" service.
func NewCspaceOrganismResponseSummaryCollection(res organismsviews.CspaceOrganismCollectionView) CspaceOrganismResponseSummaryCollection {
	body := make([]*CspaceOrganismResponseSummary, len(res))
	for i, val := range res {
		body[i] = marshalOrganismsviewsCspaceOrganismViewToCspaceOrganismResponseSummary(val)
	}
	return body
}

// NewPostBadReqResponseBody builds the HTTP response body from the result of
// the "post" endpoint of the "organisms" service.
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
// the "show" endpoint of the "organisms" service.
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

// NewUpdateBadReqResponseBody builds the HTTP response body from the result of
// the "update" endpoint of the "organisms" service.
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

// NewDeleteNotFoundResponseBody builds the HTTP response body from the result
// of the "delete" endpoint of the "organisms" service.
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

// NewDeactivateNotFoundResponseBody builds the HTTP response body from the
// result of the "deactivate" endpoint of the "organisms" service.
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

// NewListByPathNotFoundResponseBody builds the HTTP response body from the
// result of the "ListByPath" endpoint of the "organisms" service.
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
// the result of the "ListByReference" endpoint of the "organisms" service.
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

// NewPostPayload builds a organisms service post endpoint payload.
func NewPostPayload(body *PostRequestBody, path string) *organisms.PostPayload {
	v := &organisms.PostPayload{}
	if body.Organisms != nil {
		v.Organisms = make([]*organisms.OrganismPayload, len(body.Organisms))
		for i, val := range body.Organisms {
			v.Organisms[i] = unmarshalOrganismPayloadRequestBodyToOrganismsOrganismPayload(val)
		}
	}
	v.Path = path

	return v
}

// NewShowPayload builds a organisms service show endpoint payload.
func NewShowPayload(path string) *organisms.ShowPayload {
	v := &organisms.ShowPayload{}
	v.Path = path

	return v
}

// NewUpdatePayload builds a organisms service update endpoint payload.
func NewUpdatePayload(body *UpdateRequestBody, path string) *organisms.UpdatePayload {
	v := &organisms.UpdatePayload{}
	if body.Organism != nil {
		v.Organism = unmarshalOrganismPayloadRequestBodyToOrganismsOrganismPayload(body.Organism)
	}
	v.Path = path

	return v
}

// NewDeletePayload builds a organisms service delete endpoint payload.
func NewDeletePayload(path string) *organisms.DeletePayload {
	v := &organisms.DeletePayload{}
	v.Path = path

	return v
}

// NewDeactivatePayload builds a organisms service deactivate endpoint payload.
func NewDeactivatePayload(path string) *organisms.DeactivatePayload {
	v := &organisms.DeactivatePayload{}
	v.Path = path

	return v
}

// NewListByPathPayload builds a organisms service ListByPath endpoint payload.
func NewListByPathPayload(path string) *organisms.ListByPathPayload {
	v := &organisms.ListByPathPayload{}
	v.Path = path

	return v
}

// NewListByReferencePayload builds a organisms service ListByReference
// endpoint payload.
func NewListByReferencePayload(path string) *organisms.ListByReferencePayload {
	v := &organisms.ListByReferencePayload{}
	v.Path = path

	return v
}

// ValidatePostRequestBody runs the validations defined on PostRequestBody
func ValidatePostRequestBody(body *PostRequestBody) (err error) {
	for _, e := range body.Organisms {
		if e != nil {
			if err2 := ValidateOrganismPayloadRequestBody(e); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// ValidateUpdateRequestBody runs the validations defined on UpdateRequestBody
func ValidateUpdateRequestBody(body *UpdateRequestBody) (err error) {
	if body.Organism != nil {
		if err2 := ValidateOrganismPayloadRequestBody(body.Organism); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateOrganismPayloadRequestBody runs the validations defined on
// OrganismPayloadRequestBody
func ValidateOrganismPayloadRequestBody(body *OrganismPayloadRequestBody) (err error) {
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Title == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("title", "body"))
	}
	return
}
