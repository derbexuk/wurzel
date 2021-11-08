package poieventapi

import (
	"context"
	"log"

	"github.com/derbexuk/wurzel/harvester/datawarehouse"
	organisms "github.com/derbexuk/wurzel/combiner/server/gen/organisms"
	poi_orgs "github.com/derbexuk/wurzel/combiner/organisms"
)

// organisms service example implementation.
// The example methods log the requests and return zero values.
type organismssrvc struct {
	logger *log.Logger
	dw     *datawarehouse.Datawarehouse
}

// NewOrganisms returns the organisms service implementation.
func NewOrganisms(logger *log.Logger) organisms.Service {
	dw := &datawarehouse.Datawarehouse{}
	dw.Open()
	return &organismssrvc{logger, dw}
}

// Add an organism
func (s *organismssrvc) Post(ctx context.Context, p *organisms.PostPayload) (err error) {
	s.logger.Print("organisms.post")
	path := ""
	if p.Path[0] != '/' {
		path = "/" + p.Path
	} else {
		path = p.Path
	}
	for _, obj := range p.Organisms {
		if nil != obj {
			org := poi_orgs.Organism(*obj)
			err = poi_orgs.Create(&org, path, s.dw)
		}
	}
	return
}

// Show an organism in full
func (s *organismssrvc) Show(ctx context.Context, p *organisms.ShowPayload) (res *organisms.CspaceOrganism, view string, err error) {
	res = &organisms.CspaceOrganism{}
	view = "default"
	s.logger.Print("organisms.show " + p.Path)
	org := poi_orgs.GetByPath(p.Path, s.dw)
	if org != nil {
		o := organisms.CspaceOrganism(*org)
		res = &o
	}
	return
}

// Update an organism
func (s *organismssrvc) Update(ctx context.Context, p *organisms.UpdatePayload) (err error) {
	s.logger.Print("organisms.update")
	return
}

// Delete an Organism
func (s *organismssrvc) Delete(ctx context.Context, p *organisms.DeletePayload) (err error) {
	s.logger.Print("organisms.delete")
	return
}

// Delete an Organism
func (s *organismssrvc) Deactivate(ctx context.Context, p *organisms.DeactivatePayload) (err error) {
	s.logger.Print("organisms.deactivate")
	return
}

// List Orgs for a Path
func (s *organismssrvc) ListByPath(ctx context.Context, p *organisms.ListByPathPayload) (res organisms.CspaceOrganismCollection, view string, err error) {
	view = "default"
	s.logger.Print("organisms.ListByPath")
	res = organisms.CspaceOrganismCollection{}
	orgs := poi_orgs.ListByPath(p.Path, s.dw)
	for _, org := range orgs {
		o := organisms.CspaceOrganism(*org)
		res = append(res, &o)
	}
	return
}

// List Orgs for a Ref
func (s *organismssrvc) ListByReference(ctx context.Context, p *organisms.ListByReferencePayload) (res organisms.CspaceOrganismCollection, view string, err error) {
	view = "default"
	s.logger.Print("organisms.ListByReference")
	res = organisms.CspaceOrganismCollection{}
	orgs := poi_orgs.ListByRefs(p.Path, s.dw)
	for _, org := range orgs {
		o := organisms.CspaceOrganism(*org)
		res = append(res, &o)
	}
	return
}
