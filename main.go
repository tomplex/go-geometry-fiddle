package main

import (
	"log"
	"fmt"

	"github.com/paulsmith/gogeos/geos"
	"database/sql"

	_ "github.com/lib/pq"

	"runtime"
	"bufio"
	"os"
)

var connection *sql.DB

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	connectionString := "host=172.17.0.2 dbname=aqueousband_com user=postgres password=postgres port=5432 sslmode=disable"

	var err error
	connection, err = sql.Open(
		"postgres",
		connectionString,
	)

	err = connection.Ping()
	check(err)
}


func check(err error) {
	if err != nil {
		log.Fatalf("error: %s", err)
		panic(err)
	}
}


func main() {

	states := make([]*State, 0)
	venues := make([]*Venue, 0)

	venuesFinished := make(chan *Venue)

	var geomText string

	f, err := os.Create("/home/tom/data/AQ_venue_state.csv")
	check(err)

	defer f.Close()

	statesQuery, err := connection.Query("SELECT statefp, name, st_astext(geom) FROM us_states")
	check(err)

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
		go func(v *Venue) {
			match := "N/A"
			for _, state := range states {

				intersects, err := state.PGeom.Intersects(&v.Geom)
				check(err)

				if intersects {
					match = state.Name
					break
				}
			}
			v.SetState(match)
			venuesFinished <- v
		}(venue)
	}

	writer := bufio.NewWriter(f)
	defer writer.Flush()

	_, err = writer.WriteString(fmt.Sprintf("venue,state\n"))


	for v := 0; v < len(venues); v++ {
		msg := <- venuesFinished
		//fmt.Printf("Venue: %s, State: %s  \n", msg.Name, msg.State)
		_, err = writer.WriteString(fmt.Sprintf("%s,%s\n", msg.Name, msg.State))
	}

	close(venuesFinished)
	
	//nbrsQuery, err := connection.Query("SELECT nid, neighborhd, st_astext(geog) FROM neighborhoods")
	//
	//for nbrsQuery.Next() {
	//	nbr := Neighborhood{}
	//
	//	err = nbrsQuery.Scan(&nbr.Nid, &nbr.Name, geomText)
	//
	//}

}