// Code generated by goa v3.5.2, DO NOT EDIT.
//
// pois HTTP server encoders and decoders
//
// Command:
// $ goa gen github.com/derbexuk/poieventservice/server/design

package server

import (
	"context"
	"io"
	"net/http"

	pois "github.com/derbexuk/poieventservice/server/gen/pois"
	poisviews "github.com/derbexuk/poieventservice/server/gen/pois/views"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// EncodePostResponse returns an encoder for responses returned by the pois
// post endpoint.
func EncodePostResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		w.WriteHeader(http.StatusOK)
		return nil
	}
}

// DecodePostRequest returns a decoder for requests sent to the pois post
// endpoint.
func DecodePostRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			body PostRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		err = ValidatePostRequestBody(&body)
		if err != nil {
			return nil, err
		}

		var (
			path string

			params = mux.Vars(r)
		)
		path = params["path"]
		payload := NewPostPayload(&body, path)

		return payload, nil
	}
}

// EncodePostError returns an encoder for errors returned by the post pois
// endpoint.
func EncodePostError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		en, ok := v.(ErrorNamer)
		if !ok {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "un_auth":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewPostUnAuthResponseBody(res)
			}
			w.Header().Set("goa-error", res.ErrorName())
			w.WriteHeader(http.StatusUnauthorized)
			return enc.Encode(body)
		case "bad_req":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewPostBadReqResponseBody(res)
			}
			w.Header().Set("goa-error", res.ErrorName())
			w.WriteHeader(http.StatusBadRequest)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeShowResponse returns an encoder for responses returned by the pois
// show endpoint.
func EncodeShowResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(*poisviews.CspacePoi)
		w.Header().Set("goa-view", res.View)
		enc := encoder(ctx, w)
		var body interface{}
		switch res.View {
		case "default", "":
			body = NewShowResponseBody(res.Projected)
		case "summary":
			body = NewShowResponseBodySummary(res.Projected)
		}
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeShowRequest returns a decoder for requests sent to the pois show
// endpoint.
func DecodeShowRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			path string

			params = mux.Vars(r)
		)
		path = params["path"]
		payload := NewShowPayload(path)

		return payload, nil
	}
}

// EncodeShowError returns an encoder for errors returned by the show pois
// endpoint.
func EncodeShowError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		en, ok := v.(ErrorNamer)
		if !ok {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "not_found":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewShowNotFoundResponseBody(res)
			}
			w.Header().Set("goa-error", res.ErrorName())
			w.WriteHeader(http.StatusNotFound)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeListByPathResponse returns an encoder for responses returned by the
// pois ListByPath endpoint.
func EncodeListByPathResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(poisviews.CspacePoiCollection)
		w.Header().Set("goa-view", res.View)
		enc := encoder(ctx, w)
		var body interface{}
		switch res.View {
		case "default", "":
			body = NewCspacePoiResponseCollection(res.Projected)
		case "summary":
			body = NewCspacePoiResponseSummaryCollection(res.Projected)
		}
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeListByPathRequest returns a decoder for requests sent to the pois
// ListByPath endpoint.
func DecodeListByPathRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			path string

			params = mux.Vars(r)
		)
		path = params["path"]
		payload := NewListByPathPayload(path)

		return payload, nil
	}
}

// EncodeListByPathError returns an encoder for errors returned by the
// ListByPath pois endpoint.
func EncodeListByPathError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		en, ok := v.(ErrorNamer)
		if !ok {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "not_found":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewListByPathNotFoundResponseBody(res)
			}
			w.Header().Set("goa-error", res.ErrorName())
			w.WriteHeader(http.StatusNotFound)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeListByReferenceResponse returns an encoder for responses returned by
// the pois ListByReference endpoint.
func EncodeListByReferenceResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(poisviews.CspacePoiCollection)
		w.Header().Set("goa-view", res.View)
		enc := encoder(ctx, w)
		var body interface{}
		switch res.View {
		case "default", "":
			body = NewCspacePoiResponseCollection(res.Projected)
		case "summary":
			body = NewCspacePoiResponseSummaryCollection(res.Projected)
		}
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeListByReferenceRequest returns a decoder for requests sent to the pois
// ListByReference endpoint.
func DecodeListByReferenceRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			path string

			params = mux.Vars(r)
		)
		path = params["path"]
		payload := NewListByReferencePayload(path)

		return payload, nil
	}
}

// EncodeListByReferenceError returns an encoder for errors returned by the
// ListByReference pois endpoint.
func EncodeListByReferenceError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		en, ok := v.(ErrorNamer)
		if !ok {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "not_found":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewListByReferenceNotFoundResponseBody(res)
			}
			w.Header().Set("goa-error", res.ErrorName())
			w.WriteHeader(http.StatusNotFound)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeUpdateResponse returns an encoder for responses returned by the pois
// update endpoint.
func EncodeUpdateResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		w.WriteHeader(http.StatusOK)
		return nil
	}
}

// DecodeUpdateRequest returns a decoder for requests sent to the pois update
// endpoint.
func DecodeUpdateRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			body UpdateRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		err = ValidateUpdateRequestBody(&body)
		if err != nil {
			return nil, err
		}

		var (
			path string

			params = mux.Vars(r)
		)
		path = params["path"]
		payload := NewUpdatePayload(&body, path)

		return payload, nil
	}
}

// EncodeUpdateError returns an encoder for errors returned by the update pois
// endpoint.
func EncodeUpdateError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		en, ok := v.(ErrorNamer)
		if !ok {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "bad_req":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewUpdateBadReqResponseBody(res)
			}
			w.Header().Set("goa-error", res.ErrorName())
			w.WriteHeader(http.StatusBadRequest)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeDeactivateResponse returns an encoder for responses returned by the
// pois deactivate endpoint.
func EncodeDeactivateResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		w.WriteHeader(http.StatusOK)
		return nil
	}
}

// DecodeDeactivateRequest returns a decoder for requests sent to the pois
// deactivate endpoint.
func DecodeDeactivateRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			path string

			params = mux.Vars(r)
		)
		path = params["path"]
		payload := NewDeactivatePayload(path)

		return payload, nil
	}
}

// EncodeDeactivateError returns an encoder for errors returned by the
// deactivate pois endpoint.
func EncodeDeactivateError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		en, ok := v.(ErrorNamer)
		if !ok {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "not_found":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewDeactivateNotFoundResponseBody(res)
			}
			w.Header().Set("goa-error", res.ErrorName())
			w.WriteHeader(http.StatusNotFound)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeDeleteResponse returns an encoder for responses returned by the pois
// delete endpoint.
func EncodeDeleteResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		w.WriteHeader(http.StatusOK)
		return nil
	}
}

// DecodeDeleteRequest returns a decoder for requests sent to the pois delete
// endpoint.
func DecodeDeleteRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			path string

			params = mux.Vars(r)
		)
		path = params["path"]
		payload := NewDeletePayload(path)

		return payload, nil
	}
}

// EncodeDeleteError returns an encoder for errors returned by the delete pois
// endpoint.
func EncodeDeleteError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		en, ok := v.(ErrorNamer)
		if !ok {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "not_found":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewDeleteNotFoundResponseBody(res)
			}
			w.Header().Set("goa-error", res.ErrorName())
			w.WriteHeader(http.StatusNotFound)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// unmarshalPoiPayloadRequestBodyToPoisPoiPayload builds a value of type
// *pois.PoiPayload from a value of type *PoiPayloadRequestBody.
func unmarshalPoiPayloadRequestBodyToPoisPoiPayload(v *PoiPayloadRequestBody) *pois.PoiPayload {
	if v == nil {
		return nil
	}
	res := &pois.PoiPayload{
		ID:          *v.ID,
		Title:       *v.Title,
		Description: v.Description,
		Deactivated: v.Deactivated,
		Location:    v.Location,
		Geojson:     *v.Geojson,
		Path:        v.Path,
	}
	if v.Geopath != nil {
		res.Geopath = unmarshalGeoPathRequestBodyToPoisGeoPath(v.Geopath)
	}
	if v.Properties != nil {
		res.Properties = make(map[string]string, len(v.Properties))
		for key, val := range v.Properties {
			tk := key
			tv := val
			res.Properties[tk] = tv
		}
	}
	if v.Refs != nil {
		res.Refs = make([]string, len(v.Refs))
		for i, val := range v.Refs {
			res.Refs[i] = val
		}
	}

	return res
}

// unmarshalGeoPathRequestBodyToPoisGeoPath builds a value of type
// *pois.GeoPath from a value of type *GeoPathRequestBody.
func unmarshalGeoPathRequestBodyToPoisGeoPath(v *GeoPathRequestBody) *pois.GeoPath {
	if v == nil {
		return nil
	}
	res := &pois.GeoPath{
		ID:    v.ID,
		Title: v.Title,
		Type:  v.Type,
	}
	if v.Geopoints != nil {
		res.Geopoints = make([][]float32, len(v.Geopoints))
		for i, val := range v.Geopoints {
			res.Geopoints[i] = make([]float32, len(val))
			for j, val := range val {
				res.Geopoints[i][j] = val
			}
		}
	}

	return res
}

// marshalPoisviewsCspacePoiViewToCspacePoiResponse builds a value of type
// *CspacePoiResponse from a value of type *poisviews.CspacePoiView.
func marshalPoisviewsCspacePoiViewToCspacePoiResponse(v *poisviews.CspacePoiView) *CspacePoiResponse {
	res := &CspacePoiResponse{
		ID:          *v.ID,
		Title:       *v.Title,
		Description: v.Description,
		Deactivated: v.Deactivated,
		Location:    v.Location,
		Geojson:     *v.Geojson,
		Path:        v.Path,
	}
	if v.Properties != nil {
		res.Properties = make(map[string]string, len(v.Properties))
		for key, val := range v.Properties {
			tk := key
			tv := val
			res.Properties[tk] = tv
		}
	}
	if v.Refs != nil {
		res.Refs = make([]string, len(v.Refs))
		for i, val := range v.Refs {
			res.Refs[i] = val
		}
	}

	return res
}

// marshalPoisviewsCspacePoiViewToCspacePoiResponseSummary builds a value of
// type *CspacePoiResponseSummary from a value of type *poisviews.CspacePoiView.
func marshalPoisviewsCspacePoiViewToCspacePoiResponseSummary(v *poisviews.CspacePoiView) *CspacePoiResponseSummary {
	res := &CspacePoiResponseSummary{
		ID:      *v.ID,
		Title:   *v.Title,
		Geojson: *v.Geojson,
		Path:    v.Path,
	}

	return res
}
