package fetchers

import (
	"github.com/derbexuk/wurzel/combiner/events"
	"github.com/derbexuk/wurzel/combiner/organisms"
	"github.com/derbexuk/wurzel/combiner/pois"
)

type GenericResult struct {
	Pois      []*pois.Poi
	Events    []*events.Event
	Organisms []*organisms.Organism
}
