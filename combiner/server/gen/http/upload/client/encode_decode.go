// Code generated by goa v3.5.2, DO NOT EDIT.
//
// upload HTTP client encoders and decoders
//
// Command:
// $ goa gen github.com/derbexuk/wurzel/combiner/server/design

package client

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"

	upload "github.com/derbexuk/wurzel/combiner/server/gen/upload"
	goahttp "goa.design/goa/v3/http"
)

// BuildFetchRequest instantiates a HTTP request object with method and path
// set to call the "upload" service "fetch" endpoint
func (c *Client) BuildFetchRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	var (
		configFile string
	)
	{
		p, ok := v.(*upload.FetchPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("upload", "fetch", "*upload.FetchPayload", v)
		}
		if p.ConfigFile != nil {
			configFile = *p.ConfigFile
		}
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: FetchUploadPath(configFile)}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("upload", "fetch", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// DecodeFetchResponse returns a decoder for responses returned by the upload
// fetch endpoint. restoreBody controls whether the response body should be
// restored after having been read.
// DecodeFetchResponse may return the following errors:
//	- "bad_req" (type *goa.ServiceError): http.StatusBadRequest
//	- error: internal error
func DecodeFetchResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			return nil, nil
		case http.StatusBadRequest:
			var (
				body FetchBadReqResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("upload", "fetch", err)
			}
			err = ValidateFetchBadReqResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("upload", "fetch", err)
			}
			return nil, NewFetchBadReq(&body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("upload", "fetch", resp.StatusCode, string(body))
		}
	}
}

// BuildCsvRequest instantiates a HTTP request object with method and path set
// to call the "upload" service "csv" endpoint
func (c *Client) BuildCsvRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	var (
		configFile string
		resource   string
	)
	{
		p, ok := v.(*upload.CsvPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("upload", "csv", "*upload.CsvPayload", v)
		}
		if p.ConfigFile != nil {
			configFile = *p.ConfigFile
		}
		if p.Resource != nil {
			resource = *p.Resource
		}
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: CsvUploadPath(configFile, resource)}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("upload", "csv", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// DecodeCsvResponse returns a decoder for responses returned by the upload csv
// endpoint. restoreBody controls whether the response body should be restored
// after having been read.
// DecodeCsvResponse may return the following errors:
//	- "bad_req" (type *goa.ServiceError): http.StatusBadRequest
//	- error: internal error
func DecodeCsvResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			return nil, nil
		case http.StatusBadRequest:
			var (
				body CsvBadReqResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("upload", "csv", err)
			}
			err = ValidateCsvBadReqResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("upload", "csv", err)
			}
			return nil, NewCsvBadReq(&body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("upload", "csv", resp.StatusCode, string(body))
		}
	}
}
