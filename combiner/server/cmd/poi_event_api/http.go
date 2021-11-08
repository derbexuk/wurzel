package main

import (
	"context"
	"log"
	"net/http"
	"net/url"
	"os"
	"sync"
	"time"

	events "github.com/derbexuk/wurzel/combiner/server/gen/events"
	eventssvr "github.com/derbexuk/wurzel/combiner/server/gen/http/events/server"
	organismssvr "github.com/derbexuk/wurzel/combiner/server/gen/http/organisms/server"
	poissvr "github.com/derbexuk/wurzel/combiner/server/gen/http/pois/server"
	uploadsvr "github.com/derbexuk/wurzel/combiner/server/gen/http/upload/server"
	organisms "github.com/derbexuk/wurzel/combiner/server/gen/organisms"
	pois "github.com/derbexuk/wurzel/combiner/server/gen/pois"
	upload "github.com/derbexuk/wurzel/combiner/server/gen/upload"
	goahttp "goa.design/goa/v3/http"
	httpmdlwr "goa.design/goa/v3/http/middleware"
	"goa.design/goa/v3/middleware"
)

// handleHTTPServer starts configures and starts a HTTP server on the given
// URL. It shuts down the server if any error is received in the error channel.
func handleHTTPServer(ctx context.Context, u *url.URL, eventsEndpoints *events.Endpoints, organismsEndpoints *organisms.Endpoints, poisEndpoints *pois.Endpoints, uploadEndpoints *upload.Endpoints, wg *sync.WaitGroup, errc chan error, logger *log.Logger, debug bool) {

	// Setup goa log adapter.
	var (
		adapter middleware.Logger
	)
	{
		adapter = middleware.NewLogger(logger)
	}

	// Provide the transport specific request decoder and response encoder.
	// The goa http package has built-in support for JSON, XML and gob.
	// Other encodings can be used by providing the corresponding functions,
	// see goa.design/implement/encoding.
	var (
		dec = goahttp.RequestDecoder
		enc = goahttp.ResponseEncoder
	)

	// Build the service HTTP request multiplexer and configure it to serve
	// HTTP requests to the service endpoints.
	var mux goahttp.Muxer
	{
		mux = goahttp.NewMuxer()
	}

	// Wrap the endpoints with the transport specific layers. The generated
	// server packages contains code generated from the design which maps
	// the service input and output data structures to HTTP requests and
	// responses.
	var (
		eventsServer    *eventssvr.Server
		organismsServer *organismssvr.Server
		poisServer      *poissvr.Server
		uploadServer    *uploadsvr.Server
	)
	{
		eh := errorHandler(logger)
		eventsServer = eventssvr.New(eventsEndpoints, mux, dec, enc, eh, nil)
		organismsServer = organismssvr.New(organismsEndpoints, mux, dec, enc, eh, nil)
		poisServer = poissvr.New(poisEndpoints, mux, dec, enc, eh, nil)
		uploadServer = uploadsvr.New(uploadEndpoints, mux, dec, enc, eh, nil)
		if debug {
			servers := goahttp.Servers{
				eventsServer,
				organismsServer,
				poisServer,
				uploadServer,
			}
			servers.Use(httpmdlwr.Debug(mux, os.Stdout))
		}
	}
	// Configure the mux.
	eventssvr.Mount(mux, eventsServer)
	organismssvr.Mount(mux, organismsServer)
	poissvr.Mount(mux, poisServer)
	uploadsvr.Mount(mux, uploadServer)

	// Wrap the multiplexer with additional middlewares. Middlewares mounted
	// here apply to all the service endpoints.
	var handler http.Handler = mux
	{
		handler = httpmdlwr.Log(adapter)(handler)
		handler = httpmdlwr.RequestID()(handler)
	}

	// Start HTTP server using default configuration, change the code to
	// configure the server as required by your service.
	srv := &http.Server{Addr: u.Host, Handler: handler}
	for _, m := range eventsServer.Mounts {
		logger.Printf("HTTP %q mounted on %s %s", m.Method, m.Verb, m.Pattern)
	}
	for _, m := range organismsServer.Mounts {
		logger.Printf("HTTP %q mounted on %s %s", m.Method, m.Verb, m.Pattern)
	}
	for _, m := range poisServer.Mounts {
		logger.Printf("HTTP %q mounted on %s %s", m.Method, m.Verb, m.Pattern)
	}
	for _, m := range uploadServer.Mounts {
		logger.Printf("HTTP %q mounted on %s %s", m.Method, m.Verb, m.Pattern)
	}

	(*wg).Add(1)
	go func() {
		defer (*wg).Done()

		// Start HTTP server in a separate goroutine.
		go func() {
			logger.Printf("HTTP server listening on %q", u.Host)
			errc <- srv.ListenAndServe()
		}()

		<-ctx.Done()
		logger.Printf("shutting down HTTP server at %q", u.Host)

		// Shutdown gracefully with a 30s timeout.
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		_ = srv.Shutdown(ctx)
	}()
}

// errorHandler returns a function that writes and logs the given error.
// The function also writes and logs the error unique ID so that it's possible
// to correlate.
func errorHandler(logger *log.Logger) func(context.Context, http.ResponseWriter, error) {
	return func(ctx context.Context, w http.ResponseWriter, err error) {
		id := ctx.Value(middleware.RequestIDKey).(string)
		_, _ = w.Write([]byte("[" + id + "] encoding: " + err.Error()))
		logger.Printf("[%s] ERROR: %s", id, err.Error())
	}
}
