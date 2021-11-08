// Code generated by goa v3.5.2, DO NOT EDIT.
//
// organisms HTTP server
//
// Command:
// $ goa gen github.com/derbexuk/wurzel/combiner/server/design

package server

import (
	"context"
	"net/http"

	organisms "github.com/derbexuk/wurzel/combiner/server/gen/organisms"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
	"goa.design/plugins/v3/cors"
)

// Server lists the organisms service endpoint HTTP handlers.
type Server struct {
	Mounts          []*MountPoint
	Post            http.Handler
	Show            http.Handler
	Update          http.Handler
	Delete          http.Handler
	Deactivate      http.Handler
	ListByPath      http.Handler
	ListByReference http.Handler
	CORS            http.Handler
}

// ErrorNamer is an interface implemented by generated error structs that
// exposes the name of the error as defined in the design.
type ErrorNamer interface {
	ErrorName() string
}

// MountPoint holds information about the mounted endpoints.
type MountPoint struct {
	// Method is the name of the service method served by the mounted HTTP handler.
	Method string
	// Verb is the HTTP method used to match requests to the mounted handler.
	Verb string
	// Pattern is the HTTP request path pattern used to match requests to the
	// mounted handler.
	Pattern string
}

// New instantiates HTTP handlers for all the organisms service endpoints using
// the provided encoder and decoder. The handlers are mounted on the given mux
// using the HTTP verb and path defined in the design. errhandler is called
// whenever a response fails to be encoded. formatter is used to format errors
// returned by the service methods prior to encoding. Both errhandler and
// formatter are optional and can be nil.
func New(
	e *organisms.Endpoints,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) *Server {
	return &Server{
		Mounts: []*MountPoint{
			{"Post", "POST", "/things/organisms/{*path}"},
			{"Show", "GET", "/things/organisms/path/{*path}"},
			{"Update", "PUT", "/things/organisms/update/{*path}"},
			{"Delete", "DELETE", "/things/organisms/delete/{*path}"},
			{"Deactivate", "PUT", "/things/organisms/deactivate/{*path}"},
			{"ListByPath", "GET", "/things/organisms/list/{*path}"},
			{"ListByReference", "GET", "/things/organisms/refs/{*path}"},
			{"CORS", "OPTIONS", "/things/organisms/{*path}"},
			{"CORS", "OPTIONS", "/things/organisms/path/{*path}"},
			{"CORS", "OPTIONS", "/things/organisms/update/{*path}"},
			{"CORS", "OPTIONS", "/things/organisms/delete/{*path}"},
			{"CORS", "OPTIONS", "/things/organisms/deactivate/{*path}"},
			{"CORS", "OPTIONS", "/things/organisms/list/{*path}"},
			{"CORS", "OPTIONS", "/things/organisms/refs/{*path}"},
		},
		Post:            NewPostHandler(e.Post, mux, decoder, encoder, errhandler, formatter),
		Show:            NewShowHandler(e.Show, mux, decoder, encoder, errhandler, formatter),
		Update:          NewUpdateHandler(e.Update, mux, decoder, encoder, errhandler, formatter),
		Delete:          NewDeleteHandler(e.Delete, mux, decoder, encoder, errhandler, formatter),
		Deactivate:      NewDeactivateHandler(e.Deactivate, mux, decoder, encoder, errhandler, formatter),
		ListByPath:      NewListByPathHandler(e.ListByPath, mux, decoder, encoder, errhandler, formatter),
		ListByReference: NewListByReferenceHandler(e.ListByReference, mux, decoder, encoder, errhandler, formatter),
		CORS:            NewCORSHandler(),
	}
}

// Service returns the name of the service served.
func (s *Server) Service() string { return "organisms" }

// Use wraps the server handlers with the given middleware.
func (s *Server) Use(m func(http.Handler) http.Handler) {
	s.Post = m(s.Post)
	s.Show = m(s.Show)
	s.Update = m(s.Update)
	s.Delete = m(s.Delete)
	s.Deactivate = m(s.Deactivate)
	s.ListByPath = m(s.ListByPath)
	s.ListByReference = m(s.ListByReference)
	s.CORS = m(s.CORS)
}

// Mount configures the mux to serve the organisms endpoints.
func Mount(mux goahttp.Muxer, h *Server) {
	MountPostHandler(mux, h.Post)
	MountShowHandler(mux, h.Show)
	MountUpdateHandler(mux, h.Update)
	MountDeleteHandler(mux, h.Delete)
	MountDeactivateHandler(mux, h.Deactivate)
	MountListByPathHandler(mux, h.ListByPath)
	MountListByReferenceHandler(mux, h.ListByReference)
	MountCORSHandler(mux, h.CORS)
}

// MountPostHandler configures the mux to serve the "organisms" service "post"
// endpoint.
func MountPostHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := HandleOrganismsOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/things/organisms/{*path}", f)
}

// NewPostHandler creates a HTTP handler which loads the HTTP request and calls
// the "organisms" service "post" endpoint.
func NewPostHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodePostRequest(mux, decoder)
		encodeResponse = EncodePostResponse(encoder)
		encodeError    = EncodePostError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "post")
		ctx = context.WithValue(ctx, goa.ServiceKey, "organisms")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		res, err := endpoint(ctx, payload)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}

// MountShowHandler configures the mux to serve the "organisms" service "show"
// endpoint.
func MountShowHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := HandleOrganismsOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/things/organisms/path/{*path}", f)
}

// NewShowHandler creates a HTTP handler which loads the HTTP request and calls
// the "organisms" service "show" endpoint.
func NewShowHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeShowRequest(mux, decoder)
		encodeResponse = EncodeShowResponse(encoder)
		encodeError    = EncodeShowError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "show")
		ctx = context.WithValue(ctx, goa.ServiceKey, "organisms")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		res, err := endpoint(ctx, payload)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}

// MountUpdateHandler configures the mux to serve the "organisms" service
// "update" endpoint.
func MountUpdateHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := HandleOrganismsOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("PUT", "/things/organisms/update/{*path}", f)
}

// NewUpdateHandler creates a HTTP handler which loads the HTTP request and
// calls the "organisms" service "update" endpoint.
func NewUpdateHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeUpdateRequest(mux, decoder)
		encodeResponse = EncodeUpdateResponse(encoder)
		encodeError    = EncodeUpdateError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "update")
		ctx = context.WithValue(ctx, goa.ServiceKey, "organisms")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		res, err := endpoint(ctx, payload)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}

// MountDeleteHandler configures the mux to serve the "organisms" service
// "delete" endpoint.
func MountDeleteHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := HandleOrganismsOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("DELETE", "/things/organisms/delete/{*path}", f)
}

// NewDeleteHandler creates a HTTP handler which loads the HTTP request and
// calls the "organisms" service "delete" endpoint.
func NewDeleteHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeDeleteRequest(mux, decoder)
		encodeResponse = EncodeDeleteResponse(encoder)
		encodeError    = EncodeDeleteError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "delete")
		ctx = context.WithValue(ctx, goa.ServiceKey, "organisms")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		res, err := endpoint(ctx, payload)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}

// MountDeactivateHandler configures the mux to serve the "organisms" service
// "deactivate" endpoint.
func MountDeactivateHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := HandleOrganismsOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("PUT", "/things/organisms/deactivate/{*path}", f)
}

// NewDeactivateHandler creates a HTTP handler which loads the HTTP request and
// calls the "organisms" service "deactivate" endpoint.
func NewDeactivateHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeDeactivateRequest(mux, decoder)
		encodeResponse = EncodeDeactivateResponse(encoder)
		encodeError    = EncodeDeactivateError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "deactivate")
		ctx = context.WithValue(ctx, goa.ServiceKey, "organisms")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		res, err := endpoint(ctx, payload)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}

// MountListByPathHandler configures the mux to serve the "organisms" service
// "ListByPath" endpoint.
func MountListByPathHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := HandleOrganismsOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/things/organisms/list/{*path}", f)
}

// NewListByPathHandler creates a HTTP handler which loads the HTTP request and
// calls the "organisms" service "ListByPath" endpoint.
func NewListByPathHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeListByPathRequest(mux, decoder)
		encodeResponse = EncodeListByPathResponse(encoder)
		encodeError    = EncodeListByPathError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "ListByPath")
		ctx = context.WithValue(ctx, goa.ServiceKey, "organisms")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		res, err := endpoint(ctx, payload)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}

// MountListByReferenceHandler configures the mux to serve the "organisms"
// service "ListByReference" endpoint.
func MountListByReferenceHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := HandleOrganismsOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/things/organisms/refs/{*path}", f)
}

// NewListByReferenceHandler creates a HTTP handler which loads the HTTP
// request and calls the "organisms" service "ListByReference" endpoint.
func NewListByReferenceHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeListByReferenceRequest(mux, decoder)
		encodeResponse = EncodeListByReferenceResponse(encoder)
		encodeError    = EncodeListByReferenceError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "ListByReference")
		ctx = context.WithValue(ctx, goa.ServiceKey, "organisms")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		res, err := endpoint(ctx, payload)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}

// MountCORSHandler configures the mux to serve the CORS endpoints for the
// service organisms.
func MountCORSHandler(mux goahttp.Muxer, h http.Handler) {
	h = HandleOrganismsOrigin(h)
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("OPTIONS", "/things/organisms/{*path}", f)
	mux.Handle("OPTIONS", "/things/organisms/path/{*path}", f)
	mux.Handle("OPTIONS", "/things/organisms/update/{*path}", f)
	mux.Handle("OPTIONS", "/things/organisms/delete/{*path}", f)
	mux.Handle("OPTIONS", "/things/organisms/deactivate/{*path}", f)
	mux.Handle("OPTIONS", "/things/organisms/list/{*path}", f)
	mux.Handle("OPTIONS", "/things/organisms/refs/{*path}", f)
}

// NewCORSHandler creates a HTTP handler which returns a simple 200 response.
func NewCORSHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	})
}

// HandleOrganismsOrigin applies the CORS response headers corresponding to the
// origin for the service organisms.
func HandleOrganismsOrigin(h http.Handler) http.Handler {
	origHndlr := h.(http.HandlerFunc)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			origHndlr(w, r)
			return
		}
		if cors.MatchOrigin(origin, "*") {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Vary", "Origin")
			w.Header().Set("Access-Control-Expose-Headers", "X-Time")
			w.Header().Set("Access-Control-Max-Age", "600")
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			if acrm := r.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				w.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE, OPTIONS")
				w.Header().Set("Access-Control-Allow-Headers", "Authorization, Origin, X-Requested-With, Content-Type, Accept")
			}
			origHndlr(w, r)
			return
		}
		origHndlr(w, r)
		return
	})
}
