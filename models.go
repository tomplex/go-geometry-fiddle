package main

import "github.com/paulsmith/gogeos/geos"


type State struct {
	Fips string
	Name string
	Geom geos.Geometry
	PGeom geos.PGeometry
}

func (s *State) Prepare() {
	s.PGeom = *geos.PrepareGeometry(&s.Geom)
}

type Venue struct {
	Name string
	State string
	Geom geos.Geometry
	PGeom geos.PGeometry
}

func (v *Venue) Prepare() {
	v.PGeom = *geos.PrepareGeometry(&v.Geom)
}