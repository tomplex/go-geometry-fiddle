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

//type ScanableGeometry struct{
//	geos.Geometry
//}
//
//func (g *ScanableGeometry) Scan(value interface{}) (error) {
//	if value == nil {
//		return errors.New("empty geometry")
//	}
//	s, ok := value.(string)
//
//	if ok {
//		*g = geos.Must(geos.FromWKT(s))
//		return nil
//	} else {
//		return errors.New("scan failed for value " + string(s))
//	}
//}

