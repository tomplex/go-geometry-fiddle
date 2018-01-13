package main

import (
	"github.com/paulsmith/gogeos/geos"
)


type State struct {
	Fips string
	Name string
	Geom geos.Geometry
	PGeom geos.PGeometry
}

func (s *State) Prepare() {
	s.PGeom = *s.Geom.Prepare()
}


type Venue struct {
	Name string
	State string
	Geom geos.Geometry
	PGeom geos.PGeometry
}

func (v *Venue) Prepare() {
	v.PGeom = *v.Geom.Prepare()
}

func (v *Venue) SetState(stateName string) {
	v.State = stateName
}


type Neighborhood struct {
	Nid int
	Name string
	State string
	Geom geos.Geometry
	PGeom geos.PGeometry
}

func (n *Neighborhood) Prepare() {
	n.PGeom = *n.Geom.Prepare()
}
