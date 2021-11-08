package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"net/url"
	"os"
	"os/signal"
	"sync"
	"syscall"

	poieventapi "github.com/derbexuk/wurzel/combiner/server"
	events "github.com/derbexuk/wurzel/combiner/server/gen/events"
	organisms "github.com/derbexuk/wurzel/combiner/server/gen/organisms"
	pois "github.com/derbexuk/wurzel/combiner/server/gen/pois"
	upload "github.com/derbexuk/wurzel/combiner/server/gen/upload"
)

func main() {
	// Define command line flags, add any other flag required to configure the
	// service.
	var (
		hostF     = flag.String("host", "production", "Server host (valid values: production)")
		domainF   = flag.String("domain", "", "Host domain name (overrides host domain specified in service design)")
		httpPortF = flag.String("http-port", "", "HTTP port (overrides host HTTP port specified in service design)")
		secureF   = flag.Bool("secure", false, "Use secure scheme (https or grpcs)")
		dbgF      = flag.Bool("debug", false, "Log request and response bodies")
	)
	flag.Parse()

	// Setup logger. Replace logger with your own log package of choice.
	var (
		logger *log.Logger
	)
	{
		logger = log.New(os.Stderr, "[poieventapi] ", log.Ltime)
	}

	// Initialize the services.
	var (
		eventsSvc    events.Service
		organismsSvc organisms.Service
		poisSvc      pois.Service
		uploadSvc    upload.Service
	)
	{
		eventsSvc = poieventapi.NewEvents(logger)
		organismsSvc = poieventapi.NewOrganisms(logger)
		poisSvc = poieventapi.NewPois(logger)
		uploadSvc = poieventapi.NewUpload(logger)
	}

	// Wrap the services in endpoints that can be invoked from other services
	// potentially running in different processes.
	var (
		eventsEndpoints    *events.Endpoints
		organismsEndpoints *organisms.Endpoints
		poisEndpoints      *pois.Endpoints
		uploadEndpoints    *upload.Endpoints
	)
	{
		eventsEndpoints = events.NewEndpoints(eventsSvc)
		organismsEndpoints = organisms.NewEndpoints(organismsSvc)
		poisEndpoints = pois.NewEndpoints(poisSvc)
		uploadEndpoints = upload.NewEndpoints(uploadSvc)
	}

	// Create channel used by both the signal handler and server goroutines
	// to notify the main goroutine when to stop the server.
	errc := make(chan error)

	// Setup interrupt handler. This optional step configures the process so
	// that SIGINT and SIGTERM signals cause the services to stop gracefully.
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errc <- fmt.Errorf("%s", <-c)
	}()

	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())

	// Start the servers and send errors (if any) to the error channel.
	switch *hostF {
	case "production":
		{
			addr := "http://0.0.0.0:8081"
			u, err := url.Parse(addr)
			if err != nil {
				fmt.Fprintf(os.Stderr, "invalid URL %#v: %s\n", addr, err)
				os.Exit(1)
			}
			if *secureF {
				u.Scheme = "https"
			}
			if *domainF != "" {
				u.Host = *domainF
			}
			if *httpPortF != "" {
				h, _, err := net.SplitHostPort(u.Host)
				if err != nil {
					fmt.Fprintf(os.Stderr, "invalid URL %#v: %s\n", u.Host, err)
					os.Exit(1)
				}
				u.Host = net.JoinHostPort(h, *httpPortF)
			} else if u.Port() == "" {
				u.Host = net.JoinHostPort(u.Host, "80")
			}
			handleHTTPServer(ctx, u, eventsEndpoints, organismsEndpoints, poisEndpoints, uploadEndpoints, &wg, errc, logger, *dbgF)
		}

	default:
		fmt.Fprintf(os.Stderr, "invalid host argument: %q (valid hosts: production)\n", *hostF)
	}

	// Wait for signal.
	logger.Printf("exiting (%v)", <-errc)

	// Send cancellation signal to the goroutines.
	cancel()

	wg.Wait()
	logger.Println("exited")
}
