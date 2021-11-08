package poieventapi

import (
	"context"
	"log"

	upload "github.com/derbexuk/wurzel/combiner/server/gen/upload"
)

// upload service example implementation.
// The example methods log the requests and return zero values.
type uploadsrvc struct {
	logger *log.Logger
}

// NewUpload returns the upload service implementation.
func NewUpload(logger *log.Logger) upload.Service {
	return &uploadsrvc{logger}
}

// Import the Google Doc For Generic Upload
func (s *uploadsrvc) Fetch(ctx context.Context, p *upload.FetchPayload) (err error) {
	s.logger.Print("upload.fetch")
	return
}

// Upload CSV file
func (s *uploadsrvc) Csv(ctx context.Context, p *upload.CsvPayload) (err error) {
	s.logger.Print("upload.csv")
	return
}
