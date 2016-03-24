package main

import (
	"fractal"
	"log"
	"net/http"
	"serverutil"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, req *http.Request) {
		fractal.WriteFractalPng(writer, newFractalParamWithRequest(req))
	})

	log.Fatal(http.ListenAndServe("localhost:50000", nil))
}

func newFractalParamWithRequest(req *http.Request) *fractal.FractalParam {
	return fractal.NewFractalParam(
		serverutil.GetFloat64FromQuery(req, "xmin", -2),
		serverutil.GetFloat64FromQuery(req, "ymin", -2),
		serverutil.GetFloat64FromQuery(req, "xmax", 2),
		serverutil.GetFloat64FromQuery(req, "ymax", 2),
		serverutil.GetIntFromQuery(req, "width", 1024),
		serverutil.GetIntFromQuery(req, "height", 1024),
		serverutil.GetIntFromQuery(req, "root", 4))
}
