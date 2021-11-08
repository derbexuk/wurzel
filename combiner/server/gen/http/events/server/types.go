// Code generated by goa v3.5.2, DO NOT EDIT.
//
// events HTTP server types
//
// Command:
// $ goa gen github.com/derbexuk/wurzel/combiner/server/design

package server

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
	ID string `form:"id" json:"id" xml:"id"`
	// Event title/name
	Title       string  `form:"title" json:"title" xml:"title"`
	Description *string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	Deactivated *bool   `form:"deactivated,omitempty" json:"deactivated,omitempty" xml:"deactivated,omitempty"`
	Start       string  `form:"start" json:"start" xml:"start"`
	End         *string `form:"end,omitempty" json:"end,omitempty" xml:"end,omitempty"`
	// Hash of application specific properties
	Properties map[string]string `form:"properties,omitempty" json:"properties,omitempty" xml:"properties,omitempty"`
	// refs
	Refs []string `form:"refs,omitempty" json:"refs,omitempty" xml:"refs,omitempty"`
	Path *string  `form:"path,omitempty" json:"path,omitempty" xml:"path,omitempty"`
}

// ShowResponseBodySummary is the type of the "events" service "show" endpoint
// HTTP response body.
type ShowResponseBodySummary struct {
	// ID of Event
	ID string `form:"id" json:"id" xml:"id"`
	// Event title/name
	Title string  `form:"title" json:"title" xml:"title"`
	Start string  `form:"start" json:"start" xml:"start"`
	End   *string `form:"end,omitempty" json:"end,omitempty" xml:"end,omitempty"`
	Path  *string `form:"path,omitempty" json:"path,omitempty" xml:"path,omitempty"`
}

// CspaceEventResponseCollection is the type of the "events" service
// "TimeSearch" endpoint HTTP response body.
type CspaceEventResponseCollection []*CspaceEventResponse

// CspaceEventResponseSummaryCollection is the type of the "events" service
// "TimeSearch" endpoint HTTP response body.
type CspaceEventResponseSummaryCollection []*CspaceEventResponseSummary

// PostBadReqResponseBody is the type of the "events" service "post" endpoint
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

// ShowNotFoundResponseBody is the type of the "events" service "show" endpoint
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

// TimeSearchNotFoundResponseBody is the type of the "events" service
// "TimeSearch" endpoint HTTP response body for the "not_found" error.
type TimeSearchNotFoundResponseBody struct {
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

// TimeSearchBadReqResponseBody is the type of the "events" service
// "TimeSearch" endpoint HTTP response body for the "bad_req" error.
type TimeSearchBadReqResponseBody struct {
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

// ListByTimeAndPathNotFoundResponseBody is the type of the "events" service
// "ListByTimeAndPath" endpoint HTTP response body for the "not_found" error.
type ListByTimeAndPathNotFoundResponseBody struct {
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

// ListByPathNotFoundResponseBody is the type of the "events" service
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

// UpdateBadReqResponseBody is the type of the "events" service "update"
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

// DeactivateNotFoundResponseBody is the type of the "events" service
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

// DeleteNotFoundResponseBody is the type of the "events" service "delete"
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

// CspaceEventResponse is used to define fields on response body types.
type CspaceEventResponse struct {
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
	// refs
	Refs []string `form:"refs,omitempty" json:"refs,omitempty" xml:"refs,omitempty"`
	Path *string  `form:"path,omitempty" json:"path,omitempty" xml:"path,omitempty"`
}

// CspaceEventResponseSummary is used to define fields on response body types.
type CspaceEventResponseSummary struct {
	// ID of Event
	ID string `form:"id" json:"id" xml:"id"`
	// Event title/name
	Title string  `form:"title" json:"title" xml:"title"`
	Start string  `form:"start" json:"start" xml:"start"`
	End   *string `form:"end,omitempty" json:"end,omitempty" xml:"end,omitempty"`
	Path  *string `form:"path,omitempty" json:"path,omitempty" xml:"path,omitempty"`
}

// EventPayloadRequestBody is used to define fields on request body types.
type EventPayloadRequestBody struct {
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
	// references
	Refs []string `form:"refs,omitempty" json:"refs,omitempty" xml:"refs,omitempty"`
	Path *string  `form:"path,omitempty" json:"path,omitempty" xml:"path,omitempty"`
}

// NewShowResponseBody builds the HTTP response body from the result of the
// "show" endpoint of the "events" service.
func NewShowResponseBody(res *eventsviews.CspaceEventView) *ShowResponseBody {
	body := &ShowResponseBody{
		ID:          *res.ID,
		Title:       *res.Title,
		Description: res.Description,
		Deactivated: res.Deactivated,
		Start:       *res.Start,
		End:         res.End,
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
// the "show" endpoint of the "events" service.
func NewShowResponseBodySummary(res *eventsviews.CspaceEventView) *ShowResponseBodySummary {
	body := &ShowResponseBodySummary{
		ID:    *res.ID,
		Title: *res.Title,
		Start: *res.Start,
		End:   res.End,
		Path:  res.Path,
	}
	return body
}

// NewCspaceEventResponseCollection builds the HTTP response body from the
// result of the "TimeSearch" endpoint of the "events" service.
func NewCspaceEventResponseCollection(res eventsviews.CspaceEventCollectionView) CspaceEventResponseCollection {
	body := make([]*CspaceEventResponse, len(res))
	for i, val := range res {
		body[i] = marshalEventsviewsCspaceEventViewToCspaceEventResponse(val)
	}
	return body
}

// NewCspaceEventResponseSummaryCollection builds the HTTP response body from
// the result of the "TimeSearch" endpoint of the "events" service.
func NewCspaceEventResponseSummaryCollection(res eventsviews.CspaceEventCollectionView) CspaceEventResponseSummaryCollection {
	body := make([]*CspaceEventResponseSummary, len(res))
	for i, val := range res {
		body[i] = marshalEventsviewsCspaceEventViewToCspaceEventResponseSummary(val)
	}
	return body
}

// NewPostBadReqResponseBody builds the HTTP response body from the result of
// the "post" endpoint of the "events" service.
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
// the "show" endpoint of the "events" service.
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

// NewTimeSearchNotFoundResponseBody builds the HTTP response body from the
// result of the "TimeSearch" endpoint of the "events" service.
func NewTimeSearchNotFoundResponseBody(res *goa.ServiceError) *TimeSearchNotFoundResponseBody {
	body := &TimeSearchNotFoundResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewTimeSearchBadReqResponseBody builds the HTTP response body from the
// result of the "TimeSearch" endpoint of the "events" service.
func NewTimeSearchBadReqResponseBody(res *goa.ServiceError) *TimeSearchBadReqResponseBody {
	body := &TimeSearchBadReqResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewListByTimeAndPathNotFoundResponseBody builds the HTTP response body from
// the result of the "ListByTimeAndPath" endpoint of the "events" service.
func NewListByTimeAndPathNotFoundResponseBody(res *goa.ServiceError) *ListByTimeAndPathNotFoundResponseBody {
	body := &ListByTimeAndPathNotFoundResponseBody{
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
// result of the "ListByPath" endpoint of the "events" service.
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

// NewUpdateBadReqResponseBody builds the HTTP response body from the result of
// the "update" endpoint of the "events" service.
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
// result of the "deactivate" endpoint of the "events" service.
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
// of the "delete" endpoint of the "events" service.
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

// NewPostPayload builds a events service post endpoint payload.
func NewPostPayload(body *PostRequestBody, path string) *events.PostPayload {
	v := &events.PostPayload{}
	if body.Events != nil {
		v.Events = make([]*events.EventPayload, len(body.Events))
		for i, val := range body.Events {
			v.Events[i] = unmarshalEventPayloadRequestBodyToEventsEventPayload(val)
		}
	}
	v.Path = path

	return v
}

// NewShowPayload builds a events service show endpoint payload.
func NewShowPayload(path string) *events.ShowPayload {
	v := &events.ShowPayload{}
	v.Path = path

	return v
}

// NewTimeSearchPayload builds a events service TimeSearch endpoint payload.
func NewTimeSearchPayload(start string, end string, order string, limit int, path string) *events.TimeSearchPayload {
	v := &events.TimeSearchPayload{}
	v.Start = start
	v.End = &end
	v.Order = &order
	v.Limit = &limit
	v.Path = path

	return v
}

// NewListByTimeAndPathPayload builds a events service ListByTimeAndPath
// endpoint payload.
func NewListByTimeAndPathPayload(start string, end string, path string) *events.ListByTimeAndPathPayload {
	v := &events.ListByTimeAndPathPayload{}
	v.Start = start
	v.End = &end
	v.Path = path

	return v
}

// NewListByPathPayload builds a events service ListByPath endpoint payload.
func NewListByPathPayload(path string) *events.ListByPathPayload {
	v := &events.ListByPathPayload{}
	v.Path = path

	return v
}

// NewUpdatePayload builds a events service update endpoint payload.
func NewUpdatePayload(body *UpdateRequestBody, path string) *events.UpdatePayload {
	v := &events.UpdatePayload{}
	if body.Event != nil {
		v.Event = unmarshalEventPayloadRequestBodyToEventsEventPayload(body.Event)
	}
	v.Path = path

	return v
}

// NewDeactivatePayload builds a events service deactivate endpoint payload.
func NewDeactivatePayload(path string) *events.DeactivatePayload {
	v := &events.DeactivatePayload{}
	v.Path = path

	return v
}

// NewDeletePayload builds a events service delete endpoint payload.
func NewDeletePayload(path string) *events.DeletePayload {
	v := &events.DeletePayload{}
	v.Path = path

	return v
}

// ValidatePostRequestBody runs the validations defined on PostRequestBody
func ValidatePostRequestBody(body *PostRequestBody) (err error) {
	for _, e := range body.Events {
		if e != nil {
			if err2 := ValidateEventPayloadRequestBody(e); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// ValidateUpdateRequestBody runs the validations defined on UpdateRequestBody
func ValidateUpdateRequestBody(body *UpdateRequestBody) (err error) {
	if body.Event != nil {
		if err2 := ValidateEventPayloadRequestBody(body.Event); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateEventPayloadRequestBody runs the validations defined on
// EventPayloadRequestBody
func ValidateEventPayloadRequestBody(body *EventPayloadRequestBody) (err error) {
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
