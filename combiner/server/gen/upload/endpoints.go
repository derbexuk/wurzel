// Code generated by goa v3.5.2, DO NOT EDIT.
//
// upload endpoints
//
// Command:
// $ goa gen github.com/derbexuk/wurzel/combiner/server/design

package upload

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Endpoints wraps the "upload" service endpoints.
type Endpoints struct {
	Fetch goa.Endpoint
	Csv   goa.Endpoint
}

// NewEndpoints wraps the methods of the "upload" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	return &Endpoints{
		Fetch: NewFetchEndpoint(s),
		Csv:   NewCsvEndpoint(s),
	}
}

// Use applies the given middleware to all the "upload" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.Fetch = m(e.Fetch)
	e.Csv = m(e.Csv)
}

// NewFetchEndpoint returns an endpoint function that calls the method "fetch"
// of service "upload".
func NewFetchEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*FetchPayload)
		return nil, s.Fetch(ctx, p)
	}
}

// NewCsvEndpoint returns an endpoint function that calls the method "csv" of
// service "upload".
func NewCsvEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*CsvPayload)
		return nil, s.Csv(ctx, p)
	}
}
