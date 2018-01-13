package main

import (

	"github.com/paulsmith/gogeos/geos"

	_ "github.com/lib/pq"

	"log"
	"fmt"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func check(err error) {
	if err != nil {
		log.Fatalf("error: %s", err)
		panic(err)
	}
}


func main() {

	TestConnect()

	rows, err := connection.Query("SELECT statefp, name, st_astext(geom) FROM us_states WHERE name = 'New York'")
	check(err)

	states := make([]State, 0)
	venues := make([]Venue, 0)
	NYvenues := make([]Venue, 0)

	var geomText string

	for rows.Next() {
		st := State{}

		err = rows.Scan(&st.Fips, &st.Name, &geomText)

		check(err)

		geom := geos.Must(geos.FromWKT(geomText))
		st.Geom = *geom

		st.Prepare()

		states = append(states, st)
	}

	rows, err = connection.Query("SELECT name, st_astext(geom) FROM venues WHERE geom IS NOT NULL")

	for rows.Next() {
		ven := Venue{}

		err = rows.Scan(&ven.Name, &geomText)

		check(err)

		geom := geos.Must(geos.FromWKT(geomText))
		ven.Geom = *geom

		ven.Prepare()

		venues = append(venues, ven)
	}

	for _, venue := range venues {
		for _, state := range states {

			intersects, err := state.PGeom.Intersects(&venue.Geom)
			check(err)

			if intersects {
				venue.State = state.Name
				fmt.Printf("Venue: %s, State: %s  \n", venue.Name, venue.State)
			} else {

			}
		}
	}

	//for _, venue := range venues {
	//	fmt.Printf("Venue: %s, State: %s  \n", venue.Name, venue.State)
	//}

}