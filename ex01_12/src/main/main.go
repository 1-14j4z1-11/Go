package main

import (
	"fmt"
	"strconv"
	"log"
	"net/http"
)

func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		lissajous(w, newLissajousParameterWithRequest(r))
	}
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:50000", nil))
}

func newLissajousParameterWithRequest(r *http.Request) *LissajousParameter {
	cycleStr := r.FormValue("cycle")
	sizeStr := r.FormValue("size")
	nframesStr := r.FormValue("nframes")

	cycle, err1 := strconv.Atoi(cycleStr)
	size, err2 := strconv.Atoi(sizeStr)
	nframes, err3 := strconv.Atoi(nframesStr)

	if(err1 != nil) {
		cycle = 5
	}
	if(err2 != nil) {
		size = 100
	}
	if(err3 != nil) {
		nframes = 64
	}

	fmt.Printf("cycle = %d, size = %d, nframes = %d\n", cycle, size, nframes)

	return newLissajousParameter(cycle, size, nframes)
}

