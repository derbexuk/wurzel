// Code generated by goa v3.5.2, DO NOT EDIT.
//
// organisms client HTTP transport
//
// Command:
// $ goa gen github.com/derbexuk/wurzel/combiner/server/design

package client

import (
	"context"
	"net/http"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// Client lists the organisms service endpoint HTTP clients.
type Client struct {
	// Post Doer is the HTTP client used to make requests to the post endpoint.
	PostDoer goahttp.Doer

	// Show Doer is the HTTP client used to make requests to the show endpoint.
	ShowDoer goahttp.Doer

	// Update Doer is the HTTP client used to make requests to the update endpoint.
	UpdateDoer goahttp.Doer

	// Delete Doer is the HTTP client used to make requests to the delete endpoint.
	DeleteDoer goahttp.Doer

	// Deactivate Doer is the HTTP client used to make requests to the deactivate
	// endpoint.
	DeactivateDoer goahttp.Doer

	// ListByPath Doer is the HTTP client used to make requests to the ListByPath
	// endpoint.
	ListByPathDoer goahttp.Doer

	// ListByReference Doer is the HTTP client used to make requests to the
	// ListByReference endpoint.
	ListByReferenceDoer goahttp.Doer

	// CORS Doer is the HTTP client used to make requests to the  endpoint.
	CORSDoer goahttp.Doer

	// RestoreResponseBody controls whether the response bodies are reset after
	// decoding so they can be read again.
	RestoreResponseBody bool

	scheme  string
	host    string
	encoder func(*http.Request) goahttp.Encoder
	decoder func(*http.Response) goahttp.Decoder
}

// NewClient instantiates HTTP clients for all the organisms service servers.
func NewClient(
	scheme string,
	host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restoreBody bool,
) *Client {
	return &Client{
		PostDoer:            doer,
		ShowDoer:            doer,
		UpdateDoer:          doer,
		DeleteDoer:          doer,
		DeactivateDoer:      doer,
		ListByPathDoer:      doer,
		ListByReferenceDoer: doer,
		CORSDoer:            doer,
		RestoreResponseBody: restoreBody,
		scheme:              scheme,
		host:                host,
		decoder:             dec,
		encoder:             enc,
	}
}

// Post returns an endpoint that makes HTTP requests to the organisms service
// post server.
func (c *Client) Post() goa.Endpoint {
	var (
		encodeRequest  = EncodePostRequest(c.encoder)
		decodeResponse = DecodePostResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildPostRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.PostDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("organisms", "post", err)
		}
		return decodeResponse(resp)
	}
}

// Show returns an endpoint that makes HTTP requests to the organisms service
// show server.
func (c *Client) Show() goa.Endpoint {
	var (
		decodeResponse = DecodeShowResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildShowRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.ShowDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("organisms", "show", err)
		}
		return decodeResponse(resp)
	}
}

// Update returns an endpoint that makes HTTP requests to the organisms service
// update server.
func (c *Client) Update() goa.Endpoint {
	var (
		encodeRequest  = EncodeUpdateRequest(c.encoder)
		decodeResponse = DecodeUpdateResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildUpdateRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.UpdateDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("organisms", "update", err)
		}
		return decodeResponse(resp)
	}
}

// Delete returns an endpoint that makes HTTP requests to the organisms service
// delete server.
func (c *Client) Delete() goa.Endpoint {
	var (
		decodeResponse = DecodeDeleteResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildDeleteRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.DeleteDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("organisms", "delete", err)
		}
		return decodeResponse(resp)
	}
}

// Deactivate returns an endpoint that makes HTTP requests to the organisms
// service deactivate server.
func (c *Client) Deactivate() goa.Endpoint {
	var (
		decodeResponse = DecodeDeactivateResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildDeactivateRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.DeactivateDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("organisms", "deactivate", err)
		}
		return decodeResponse(resp)
	}
}

// ListByPath returns an endpoint that makes HTTP requests to the organisms
// service ListByPath server.
func (c *Client) ListByPath() goa.Endpoint {
	var (
		decodeResponse = DecodeListByPathResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildListByPathRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.ListByPathDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("organisms", "ListByPath", err)
		}
		return decodeResponse(resp)
	}
}

// ListByReference returns an endpoint that makes HTTP requests to the
// organisms service ListByReference server.
func (c *Client) ListByReference() goa.Endpoint {
	var (
		decodeResponse = DecodeListByReferenceResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildListByReferenceRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.ListByReferenceDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("organisms", "ListByReference", err)
		}
		return decodeResponse(resp)
	}
}
