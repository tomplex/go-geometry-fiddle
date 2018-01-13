package main

import (

	"github.com/paulsmith/gogeos/geos"

	_ "github.com/lib/pq"

	"log"
	"fmt"

)


func check(err error) {
	if err != nil {
		log.Fatalf("error: %s", err)
		panic(err)
	}
}


func main() {

	TestConnect()

	statesQuery, err := connection.Query("SELECT statefp, name, st_astext(geom) FROM us_states")
	check(err)

	states := make([]*State, 0)
	venues := make([]*Venue, 0)
	//NYvenues := make([]Venue, 0)
	venuesFinished := make(chan *Venue)

	var geomText string

	for statesQuery.Next() {
		st := &State{}

		err = statesQuery.Scan(&st.Fips, &st.Name, &geomText)

		check(err)

		geom := geos.Must(geos.FromWKT(geomText))
		st.Geom = *geom

		st.Prepare()

		states = append(states, st)
	}


	venuesQuery, err := connection.Query("SELECT name, st_astext(geom) FROM venues WHERE geom IS NOT NULL")

	for venuesQuery.Next() {
		ven := &Venue{}

		err = venuesQuery.Scan(&ven.Name, &geomText)

		check(err)

		geom := geos.Must(geos.FromWKT(geomText))
		ven.Geom = *geom

		ven.Prepare()

		venues = append(venues, ven)
	}

	for _, venue := range venues {
		//AssignState(venue, states)
		go func(v *Venue) {
			for _, state := range states {

				intersects, err := state.PGeom.Intersects(&v.Geom)
				check(err)

				if intersects {
					v.SetState(state.Name)
					venuesFinished <- v

				}
			}
		}(venue)

	}


	for v := range venuesFinished {
		fmt.Printf("Venue: %s, State: %s  \n", v.Name, v.State)
	}

	//nbrsQuery, err := connection.Query("SELECT nid, neighborhd, st_astext(geog) FROM neighborhoods")
	//
	//for nbrsQuery.Next() {
	//	nbr := Neighborhood{}
	//
	//	err = nbrsQuery.Scan(&nbr.Nid, &nbr.Name, geomText)
	//
	//}

}