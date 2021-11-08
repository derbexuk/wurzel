// Code generated by goa v3.5.2, DO NOT EDIT.
//
// organisms endpoints
//
// Command:
// $ goa gen github.com/derbexuk/poieventservice/server/design

package organisms

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Endpoints wraps the "organisms" service endpoints.
type Endpoints struct {
	Post            goa.Endpoint
	Show            goa.Endpoint
	Update          goa.Endpoint
	Delete          goa.Endpoint
	Deactivate      goa.Endpoint
	ListByPath      goa.Endpoint
	ListByReference goa.Endpoint
}

// NewEndpoints wraps the methods of the "organisms" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	return &Endpoints{
		Post:            NewPostEndpoint(s),
		Show:            NewShowEndpoint(s),
		Update:          NewUpdateEndpoint(s),
		Delete:          NewDeleteEndpoint(s),
		Deactivate:      NewDeactivateEndpoint(s),
		ListByPath:      NewListByPathEndpoint(s),
		ListByReference: NewListByReferenceEndpoint(s),
	}
}

// Use applies the given middleware to all the "organisms" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.Post = m(e.Post)
	e.Show = m(e.Show)
	e.Update = m(e.Update)
	e.Delete = m(e.Delete)
	e.Deactivate = m(e.Deactivate)
	e.ListByPath = m(e.ListByPath)
	e.ListByReference = m(e.ListByReference)
}

// NewPostEndpoint returns an endpoint function that calls the method "post" of
// service "organisms".
func NewPostEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*PostPayload)
		return nil, s.Post(ctx, p)
	}
}

// NewShowEndpoint returns an endpoint function that calls the method "show" of
// service "organisms".
func NewShowEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*ShowPayload)
		res, view, err := s.Show(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedCspaceOrganism(res, view)
		return vres, nil
	}
}

// NewUpdateEndpoint returns an endpoint function that calls the method
// "update" of service "organisms".
func NewUpdateEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*UpdatePayload)
		return nil, s.Update(ctx, p)
	}
}

// NewDeleteEndpoint returns an endpoint function that calls the method
// "delete" of service "organisms".
func NewDeleteEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*DeletePayload)
		return nil, s.Delete(ctx, p)
	}
}

// NewDeactivateEndpoint returns an endpoint function that calls the method
// "deactivate" of service "organisms".
func NewDeactivateEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*DeactivatePayload)
		return nil, s.Deactivate(ctx, p)
	}
}

// NewListByPathEndpoint returns an endpoint function that calls the method
// "ListByPath" of service "organisms".
func NewListByPathEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*ListByPathPayload)
		res, view, err := s.ListByPath(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedCspaceOrganismCollection(res, view)
		return vres, nil
	}
}

// NewListByReferenceEndpoint returns an endpoint function that calls the
// method "ListByReference" of service "organisms".
func NewListByReferenceEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*ListByReferencePayload)
		res, view, err := s.ListByReference(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedCspaceOrganismCollection(res, view)
		return vres, nil
	}
}
