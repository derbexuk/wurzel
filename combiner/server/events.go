package poieventapi

import (
	"context"
	"log"

	"github.com/derbexuk/wurzel/harvester/datawarehouse"
	poi_evts "github.com/derbexuk/wurzel/combiner/events"
	events "github.com/derbexuk/wurzel/combiner/server/gen/events"
)

// events service example implementation.
// The example methods log the requests and return zero values.
type eventssrvc struct {
	logger *log.Logger
	dw     *datawarehouse.Datawarehouse
}

// NewEvents returns the events service implementation.
func NewEvents(logger *log.Logger) events.Service {
	dw := &datawarehouse.Datawarehouse{}
	dw.Open()
	//dw.EnsurePathIndex(DB, COLLECTION)
	return &eventssrvc{logger, dw}
}

// Add an event
func (s *eventssrvc) Post(ctx context.Context, p *events.PostPayload) (err error) {
	s.logger.Printf("events.post %d\n", len(p.Events))
	// Set optional fields to default values where not supplied
	path := ""
	if p.Path[0] != '/' {
		path = "/" + p.Path
	} else {
		path = p.Path
	}
	for _, obj := range p.Events {
		if nil != obj {
			evt := poi_evts.Event(*obj)
			err = poi_evts.Create(path, &evt, s.dw)
		}
	}
	return
}

// Show an event in full
func (s *eventssrvc) Show(ctx context.Context, p *events.ShowPayload) (res *events.CspaceEvent, view string, err error) {
	res = &events.CspaceEvent{}
	view = "default"
	s.logger.Print("events.show : " + p.Path)
	evt := poi_evts.GetByPath(p.Path, s.dw)
	if evt != nil {
		e := events.CspaceEvent(*evt)
		res = &e
	}
	return
}

// Show events according to search values
func (s *eventssrvc) TimeSearch(ctx context.Context, p *events.TimeSearchPayload) (res events.CspaceEventCollection, view string, err error) {
	view = "default"
	s.logger.Print("events.TimeSearch")
	var order string
	var limit int64

	if p.Order != nil {
		order = *p.Order
	}
	if p.Limit != nil {
		limit = int64(*p.Limit)
	}
	res = events.CspaceEventCollection{}
	evts, err := poi_evts.TimeSearch(p.Path, p.Start, *p.End, order, limit, s.dw)
	for _, evt := range evts {
		e := events.CspaceEvent(*evt)
		res = append(res, &e)
	}
	return
}

// List Events for a path
func (s *eventssrvc) ListByTimeAndPath(ctx context.Context, p *events.ListByTimeAndPathPayload) (res events.CspaceEventCollection, view string, err error) {
	view = "default"
	s.logger.Print("events.ListByTimeAndPath")
	var end string
	if p.End != nil {
		end = *p.End
	}
	res = events.CspaceEventCollection{}
	evts, err := poi_evts.ListByTimeAndPath(p.Path, p.Start, end, s.dw)
	for _, evt := range evts {
		e := events.CspaceEvent(*evt)
		res = append(res, &e)
	}
	return
}

// List Events for a path
func (s *eventssrvc) ListByPath(ctx context.Context, p *events.ListByPathPayload) (res events.CspaceEventCollection, view string, err error) {
	view = "default"
	s.logger.Print("events.ListByPath")
	res = events.CspaceEventCollection{}
	evts, err := poi_evts.ListByPath(p.Path, s.dw)
	for _, evt := range evts {
		e := events.CspaceEvent(*evt)
		res = append(res, &e)
	}
	return
}

// Update an event
func (s *eventssrvc) Update(ctx context.Context, p *events.UpdatePayload) (err error) {
	s.logger.Print("events.update")
	return
}

// Delete an Event
func (s *eventssrvc) Deactivate(ctx context.Context, p *events.DeactivatePayload) (err error) {
	s.logger.Print("events.deactivate")
	return
}

// Delete an Event
func (s *eventssrvc) Delete(ctx context.Context, p *events.DeletePayload) (err error) {
	s.logger.Print("events.delete")
	return
}
