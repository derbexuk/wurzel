// Code generated by goa v3.5.2, DO NOT EDIT.
//
// pois endpoints
//
// Command:
// $ goa gen github.com/derbexuk/wurzel/combiner/server/design

package pois

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Endpoints wraps the "pois" service endpoints.
type Endpoints struct {
	Post            goa.Endpoint
	Show            goa.Endpoint
	ListByPath      goa.Endpoint
	ListByReference goa.Endpoint
	Update          goa.Endpoint
	Deactivate      goa.Endpoint
	Delete          goa.Endpoint
}

// NewEndpoints wraps the methods of the "pois" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	return &Endpoints{
		Post:            NewPostEndpoint(s),
		Show:            NewShowEndpoint(s),
		ListByPath:      NewListByPathEndpoint(s),
		ListByReference: NewListByReferenceEndpoint(s),
		Update:          NewUpdateEndpoint(s),
		Deactivate:      NewDeactivateEndpoint(s),
		Delete:          NewDeleteEndpoint(s),
	}
}

// Use applies the given middleware to all the "pois" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.Post = m(e.Post)
	e.Show = m(e.Show)
	e.ListByPath = m(e.ListByPath)
	e.ListByReference = m(e.ListByReference)
	e.Update = m(e.Update)
	e.Deactivate = m(e.Deactivate)
	e.Delete = m(e.Delete)
}

// NewPostEndpoint returns an endpoint function that calls the method "post" of
// service "pois".
func NewPostEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*PostPayload)
		return nil, s.Post(ctx, p)
	}
}

// NewShowEndpoint returns an endpoint function that calls the method "show" of
// service "pois".
func NewShowEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*ShowPayload)
		res, view, err := s.Show(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedCspacePoi(res, view)
		return vres, nil
	}
}

// NewListByPathEndpoint returns an endpoint function that calls the method
// "ListByPath" of service "pois".
func NewListByPathEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*ListByPathPayload)
		res, view, err := s.ListByPath(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedCspacePoiCollection(res, view)
		return vres, nil
	}
}

// NewListByReferenceEndpoint returns an endpoint function that calls the
// method "ListByReference" of service "pois".
func NewListByReferenceEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*ListByReferencePayload)
		res, view, err := s.ListByReference(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedCspacePoiCollection(res, view)
		return vres, nil
	}
}

// NewUpdateEndpoint returns an endpoint function that calls the method
// "update" of service "pois".
func NewUpdateEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*UpdatePayload)
		return nil, s.Update(ctx, p)
	}
}

// NewDeactivateEndpoint returns an endpoint function that calls the method
// "deactivate" of service "pois".
func NewDeactivateEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*DeactivatePayload)
		return nil, s.Deactivate(ctx, p)
	}
}

// NewDeleteEndpoint returns an endpoint function that calls the method
// "delete" of service "pois".
func NewDeleteEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*DeletePayload)
		return nil, s.Delete(ctx, p)
	}
}
