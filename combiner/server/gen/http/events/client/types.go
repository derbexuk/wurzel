// Code generated by goa v3.5.2, DO NOT EDIT.
//
// events HTTP client types
//
// Command:
// $ goa gen github.com/derbexuk/wurzel/combiner/server/design

package client

import (
	events "github.com/derbexuk/wurzel/combiner/server/gen/events"
	eventsviews "github.com/derbexuk/wurzel/combiner/server/gen/events/views"
	goa "goa.design/goa/v3/pkg"
)

// PostRequestBody is the type of the "events" service "post" endpoint HTTP
// request body.
type PostRequestBody struct {
	Events []*EventPayloadRequestBody `form:"events,omitempty" json:"events,omitempty" xml:"events,omitempty"`
}

// UpdateRequestBody is the type of the "events" service "update" endpoint HTTP
// request body.
type UpdateRequestBody struct {
	Event *EventPayloadRequestBody `form:"event,omitempty" json:"event,omitempty" xml:"event,omitempty"`
}

// ShowResponseBody is the type of the "events" service "show" endpoint HTTP
// response body.
type ShowResponseBody struct {
	// ID of Event
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Event title/name
	Title       *string `form:"title,omitempty" json:"title,omitempty" xml:"title,omitempty"`
	Description *string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	Deactivated *bool   `form:"deactivated,omitempty" json:"deactivated,omitempty" xml:"deactivated,omitempty"`
	Start       *string `form:"start,omitempty" json:"start,omitempty" xml:"start,omitempty"`
	End         *string `form:"end,omitempty" json:"end,omitempty" xml:"end,omitempty"`
	// Hash of application specific properties
	Properties map[string]string `form:"properties,omitempty" json:"properties,omitempty" xml:"properties,omitempty"`
	// refs
	Refs []string `form:"refs,omitempty" json:"refs,omitempty" xml:"refs,omitempty"`
	Path *string  `form:"path,omitempty" json:"path,omitempty" xml:"path,omitempty"`
}

// TimeSearchResponseBody is the type of the "events" service "TimeSearch"
// endpoint HTTP response body.
type TimeSearchResponseBody []*CspaceEventResponse

// ListByTimeAndPathResponseBody is the type of the "events" service
// "ListByTimeAndPath" endpoint HTTP response body.
type ListByTimeAndPathResponseBody []*CspaceEventResponse

// ListByPathResponseBody is the type of the "events" service "ListByPath"
// endpoint HTTP response body.
type ListByPathResponseBody []*CspaceEventResponse

// PostBadReqResponseBody is the type of the "events" service "post" endpoint
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

// ShowNotFoundResponseBody is the type of the "events" service "show" endpoint
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

// TimeSearchNotFoundResponseBody is the type of the "events" service
// "TimeSearch" endpoint HTTP response body for the "not_found" error.
type TimeSearchNotFoundResponseBody struct {
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

// TimeSearchBadReqResponseBody is the type of the "events" service
// "TimeSearch" endpoint HTTP response body for the "bad_req" error.
type TimeSearchBadReqResponseBody struct {
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

// ListByTimeAndPathNotFoundResponseBody is the type of the "events" service
// "ListByTimeAndPath" endpoint HTTP response body for the "not_found" error.
type ListByTimeAndPathNotFoundResponseBody struct {
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

// ListByPathNotFoundResponseBody is the type of the "events" service
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

// UpdateBadReqResponseBody is the type of the "events" service "update"
// endpoint HTTP response body for the "bad_req" error.
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

// DeactivateNotFoundResponseBody is the type of the "events" service
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

// DeleteNotFoundResponseBody is the type of the "events" service "delete"
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

// EventPayloadRequestBody is used to define fields on request body types.
type EventPayloadRequestBody struct {
	// ID of Event
	ID string `form:"id" json:"id" xml:"id"`
	// Event title/name
	Title       string  `form:"title" json:"title" xml:"title"`
	Description *string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	Deactivated *bool   `form:"deactivated,omitempty" json:"deactivated,omitempty" xml:"deactivated,omitempty"`
	Start       string  `form:"start" json:"start" xml:"start"`
	End         *string `form:"end,omitempty" json:"end,omitempty" xml:"end,omitempty"`
	// Hash of application specific properties
	Properties map[string]string `form:"properties,omitempty" json:"properties,omitempty" xml:"properties,omitempty"`
	// references
	Refs []string `form:"refs,omitempty" json:"refs,omitempty" xml:"refs,omitempty"`
	Path *string  `form:"path,omitempty" json:"path,omitempty" xml:"path,omitempty"`
}

// CspaceEventResponse is used to define fields on response body types.
type CspaceEventResponse struct {
	// ID of Event
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Event title/name
	Title       *string `form:"title,omitempty" json:"title,omitempty" xml:"title,omitempty"`
	Description *string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	Deactivated *bool   `form:"deactivated,omitempty" json:"deactivated,omitempty" xml:"deactivated,omitempty"`
	Start       *string `form:"start,omitempty" json:"start,omitempty" xml:"start,omitempty"`
	End         *string `form:"end,omitempty" json:"end,omitempty" xml:"end,omitempty"`
	// Hash of application specific properties
	Properties map[string]string `form:"properties,omitempty" json:"properties,omitempty" xml:"properties,omitempty"`
	// refs
	Refs []string `form:"refs,omitempty" json:"refs,omitempty" xml:"refs,omitempty"`
	Path *string  `form:"path,omitempty" json:"path,omitempty" xml:"path,omitempty"`
}

// NewPostRequestBody builds the HTTP request body from the payload of the
// "post" endpoint of the "events" service.
func NewPostRequestBody(p *events.PostPayload) *PostRequestBody {
	body := &PostRequestBody{}
	if p.Events != nil {
		body.Events = make([]*EventPayloadRequestBody, len(p.Events))
		for i, val := range p.Events {
			body.Events[i] = marshalEventsEventPayloadToEventPayloadRequestBody(val)
		}
	}
	return body
}

// NewUpdateRequestBody builds the HTTP request body from the payload of the
// "update" endpoint of the "events" service.
func NewUpdateRequestBody(p *events.UpdatePayload) *UpdateRequestBody {
	body := &UpdateRequestBody{}
	if p.Event != nil {
		body.Event = marshalEventsEventPayloadToEventPayloadRequestBody(p.Event)
	}
	return body
}

// NewPostBadReq builds a events service post endpoint bad_req error.
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

// NewShowCspaceEventOK builds a "events" service "show" endpoint result from a
// HTTP "OK" response.
func NewShowCspaceEventOK(body *ShowResponseBody) *eventsviews.CspaceEventView {
	v := &eventsviews.CspaceEventView{
		ID:          body.ID,
		Title:       body.Title,
		Description: body.Description,
		Deactivated: body.Deactivated,
		Start:       body.Start,
		End:         body.End,
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

// NewShowNotFound builds a events service show endpoint not_found error.
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

// NewTimeSearchCspaceEventCollectionOK builds a "events" service "TimeSearch"
// endpoint result from a HTTP "OK" response.
func NewTimeSearchCspaceEventCollectionOK(body TimeSearchResponseBody) eventsviews.CspaceEventCollectionView {
	v := make([]*eventsviews.CspaceEventView, len(body))
	for i, val := range body {
		v[i] = unmarshalCspaceEventResponseToEventsviewsCspaceEventView(val)
	}

	return v
}

// NewTimeSearchNotFound builds a events service TimeSearch endpoint not_found
// error.
func NewTimeSearchNotFound(body *TimeSearchNotFoundResponseBody) *goa.ServiceError {
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

// NewTimeSearchBadReq builds a events service TimeSearch endpoint bad_req
// error.
func NewTimeSearchBadReq(body *TimeSearchBadReqResponseBody) *goa.ServiceError {
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

// NewListByTimeAndPathCspaceEventCollectionOK builds a "events" service
// "ListByTimeAndPath" endpoint result from a HTTP "OK" response.
func NewListByTimeAndPathCspaceEventCollectionOK(body ListByTimeAndPathResponseBody) eventsviews.CspaceEventCollectionView {
	v := make([]*eventsviews.CspaceEventView, len(body))
	for i, val := range body {
		v[i] = unmarshalCspaceEventResponseToEventsviewsCspaceEventView(val)
	}

	return v
}

// NewListByTimeAndPathNotFound builds a events service ListByTimeAndPath
// endpoint not_found error.
func NewListByTimeAndPathNotFound(body *ListByTimeAndPathNotFoundResponseBody) *goa.ServiceError {
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

// NewListByPathCspaceEventCollectionOK builds a "events" service "ListByPath"
// endpoint result from a HTTP "OK" response.
func NewListByPathCspaceEventCollectionOK(body ListByPathResponseBody) eventsviews.CspaceEventCollectionView {
	v := make([]*eventsviews.CspaceEventView, len(body))
	for i, val := range body {
		v[i] = unmarshalCspaceEventResponseToEventsviewsCspaceEventView(val)
	}

	return v
}

// NewListByPathNotFound builds a events service ListByPath endpoint not_found
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

// NewUpdateBadReq builds a events service update endpoint bad_req error.
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

// NewDeactivateNotFound builds a events service deactivate endpoint not_found
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

// NewDeleteNotFound builds a events service delete endpoint not_found error.
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

// ValidateTimeSearchNotFoundResponseBody runs the validations defined on
// TimeSearch_not_found_Response_Body
func ValidateTimeSearchNotFoundResponseBody(body *TimeSearchNotFoundResponseBody) (err error) {
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

// ValidateTimeSearchBadReqResponseBody runs the validations defined on
// TimeSearch_bad_req_Response_Body
func ValidateTimeSearchBadReqResponseBody(body *TimeSearchBadReqResponseBody) (err error) {
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

// ValidateListByTimeAndPathNotFoundResponseBody runs the validations defined
// on ListByTimeAndPath_not_found_Response_Body
func ValidateListByTimeAndPathNotFoundResponseBody(body *ListByTimeAndPathNotFoundResponseBody) (err error) {
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

// ValidateEventPayloadRequestBody runs the validations defined on
// EventPayloadRequestBody
func ValidateEventPayloadRequestBody(body *EventPayloadRequestBody) (err error) {
	err = goa.MergeErrors(err, goa.ValidateFormat("body.start", body.Start, goa.FormatDateTime))

	if body.End != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.end", *body.End, goa.FormatDateTime))
	}
	return
}

// ValidateCspaceEventResponse runs the validations defined on
// CspaceEventResponse
func ValidateCspaceEventResponse(body *CspaceEventResponse) (err error) {
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Title == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("title", "body"))
	}
	if body.Start == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("start", "body"))
	}
	if body.Start != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.start", *body.Start, goa.FormatDateTime))
	}
	if body.End != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.end", *body.End, goa.FormatDateTime))
	}
	return
}
