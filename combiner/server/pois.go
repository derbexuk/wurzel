package poieventapi

import (
	"context"
	"errors"
	"log"

	"github.com/derbexuk/wurzel/harvester/datawarehouse"
	poi_pois "github.com/derbexuk/wurzel/combiner/pois"
	pois "github.com/derbexuk/wurzel/combiner/server/gen/pois"
)

// pois service example implementation.
// The example methods log the requests and return zero values.
type poissrvc struct {
	logger *log.Logger
	dw     *datawarehouse.Datawarehouse
}

// NewPois returns the pois service implementation.
func NewPois(logger *log.Logger) pois.Service {
	dw := &datawarehouse.Datawarehouse{}
	dw.Open()
	return &poissrvc{logger, dw}
}

// Add a POI
func (s *poissrvc) Post(ctx context.Context, p *pois.PostPayload) (err error) {
	s.logger.Print("pois.post")
	path := ""
	if p.Path[0] != '/' {
		path = "/" + p.Path
	} else {
		path = p.Path
	}
	for _, obj := range p.Pois {
		if nil != obj {
			poi := poi_pois.Poi{}
			poi.ID = obj.ID
			poi.Title = obj.Title
			poi.Description = obj.Description
			poi.Deactivated = obj.Deactivated
			poi.Location = obj.Location
			poi.GeoJSON = obj.Geojson
			poi.Properties = obj.Properties
			poi.Refs = obj.Refs
			poi.Path = obj.Path
			err = poi_pois.Create(&poi, path, s.dw)
		}
	}
	return
}

// Show a POI
func (s *poissrvc) Show(ctx context.Context, p *pois.ShowPayload) (res *pois.CspacePoi, view string, err error) {
	view = "default"
	s.logger.Print("pois.show " + p.Path)
	poi := poi_pois.GetByPath(p.Path, s.dw)
	if poi != nil {
		res = MapToCspacePoi(poi)
	} else {
		err = pois.MakeNotFound(errors.New("Not Found"))
	}
	return
}

// List POIs for a path
func (s *poissrvc) ListByPath(ctx context.Context, p *pois.ListByPathPayload) (res pois.CspacePoiCollection, view string, err error) {
	view = "default"
	s.logger.Print("pois.ListByPath")
	res = pois.CspacePoiCollection{}
	ps := poi_pois.ListByPath(p.Path, s.dw)
	for _, p := range ps {
		res = append(res, MapToCspacePoi(p))
	}
	return
}

// List POIs for a path
func (s *poissrvc) ListByReference(ctx context.Context, p *pois.ListByReferencePayload) (res pois.CspacePoiCollection, view string, err error) {
	view = "default"
	s.logger.Print("pois.ListByReference")
	res = pois.CspacePoiCollection{}
	ps := poi_pois.ListByRefs(p.Path, s.dw)
	for _, p := range ps {
		res = append(res, MapToCspacePoi(p))
	}
	return
}

// Update a POI
func (s *poissrvc) Update(ctx context.Context, p *pois.UpdatePayload) (err error) {
	s.logger.Print("pois.update")
	return
}

// Update a POI
func (s *poissrvc) Deactivate(ctx context.Context, p *pois.DeactivatePayload) (err error) {
	s.logger.Print("pois.deactivate")
	return
}

// Delete a POI
func (s *poissrvc) Delete(ctx context.Context, p *pois.DeletePayload) (err error) {
	s.logger.Print("pois.delete")
	return
}

func MapToCspacePoi(p *poi_pois.Poi) (res *pois.CspacePoi) {
	rpoi := pois.CspacePoi{}
	rpoi.ID = p.ID
	rpoi.Title = p.Title
	rpoi.Description = p.Description
	rpoi.Deactivated = p.Deactivated
	rpoi.Location = p.Location
	rpoi.Geojson = p.GeoJSON
	rpoi.Properties = p.Properties
	rpoi.Refs = p.Refs
	rpoi.Path = p.Path
	res = &rpoi
	return
}
