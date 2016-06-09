package main

import (
	"fmt"
	"log"
	"net/http"
	"sort"
	"table"
	"track"
)

func main() {
	list := createList()

	handler := func(w http.ResponseWriter, r *http.Request) {
		acceptQuery(r, list)
		w.Write([]byte(table.CreateTableHTML(list)))
	}
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:50000", nil))
}

func createList() *track.TrackList {
	list := track.TrackList([]track.Track{
		track.Track{Title: "Go", Artist: "Moby", Album: "Moby", Year: 1992, Length: track.Length("3m37s")},
		track.Track{Title: "Go", Artist: "Delilah", Album: "From the Roots Up", Year: 2012, Length: track.Length("3m38s")},
		track.Track{Title: "Go Ahead", Artist: "Alicia Keys", Album: "As I Am", Year: 2007, Length: track.Length("4m36s")},
		track.Track{Title: "Ready 2 Go", Artist: "Martin Solveig", Album: "Smash", Year: 2011, Length: track.Length("4m24s")},
	})

	return &list
}

func acceptQuery(r *http.Request, list *track.TrackList) {
	if order := r.FormValue("order"); order != "" {
		if track.SetFirstOrder(order) == nil {
			fmt.Printf("Changed first order : %s\n", order)
		} else {
			fmt.Printf("Unknown order key : %s\n", order)
		}
	}
	sort.Sort(list)
}
