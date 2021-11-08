// Code generated by goa v3.5.2, DO NOT EDIT.
//
// upload client
//
// Command:
// $ goa gen github.com/derbexuk/wurzel/combiner/server/design

package upload

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "upload" service client.
type Client struct {
	FetchEndpoint goa.Endpoint
	CsvEndpoint   goa.Endpoint
}

// NewClient initializes a "upload" service client given the endpoints.
func NewClient(fetch, csv goa.Endpoint) *Client {
	return &Client{
		FetchEndpoint: fetch,
		CsvEndpoint:   csv,
	}
}

// Fetch calls the "fetch" endpoint of the "upload" service.
func (c *Client) Fetch(ctx context.Context, p *FetchPayload) (err error) {
	_, err = c.FetchEndpoint(ctx, p)
	return
}

// Csv calls the "csv" endpoint of the "upload" service.
func (c *Client) Csv(ctx context.Context, p *CsvPayload) (err error) {
	_, err = c.CsvEndpoint(ctx, p)
	return
}
