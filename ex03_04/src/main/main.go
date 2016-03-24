package main

import (
	"image/color"
	"log"
	"net/http"
	"serverutil"
	"svg"
)

const (
	defwidth		= 1200
	defheight		= 800
	defcells		= 100
	defxyrange		= 30.0
)

var defMaxColor = color.RGBA{ 0xFF, 0x00, 0x00, 0xFF }
var defMinColor = color.RGBA{ 0x00, 0x00, 0xFF, 0xFF }

func main() {
	http.HandleFunc("/func0", func(writer http.ResponseWriter, req *http.Request) {
		writer.Header().Set("Content-Type", "image/svg+xml")
		param := newSVGParamWithRequest(req, func0)
		svg.WriteSVG(writer, param)
	});
	http.HandleFunc("/func1", func(writer http.ResponseWriter, req *http.Request) {
		writer.Header().Set("Content-Type", "image/svg+xml")
		param := newSVGParamWithRequest(req, func1)
		svg.WriteSVG(writer, param)
	});
	http.HandleFunc("/func2", func(writer http.ResponseWriter, req *http.Request) {
		writer.Header().Set("Content-Type", "image/svg+xml")
		param := newSVGParamWithRequest(req, func2)
		svg.WriteSVG(writer, param)
	});
	http.HandleFunc("/func3", func(writer http.ResponseWriter, req *http.Request) {
		writer.Header().Set("Content-Type", "image/svg+xml")
		param := newSVGParamWithRequest(req, func3)
		svg.WriteSVG(writer, param)
	});

	log.Fatal(http.ListenAndServe("localhost:50000", nil))
}

func newSVGParamWithRequest(req *http.Request, function svg.BiFunction) *svg.SVGParam {
	return svg.NewSVGParam(
		serverutil.GetIntFromQuery(req, "width", defwidth),
		serverutil.GetIntFromQuery(req, "height", defheight),
		serverutil.GetIntFromQuery(req, "cells", defcells),
		serverutil.GetFloat64FromQuery(req, "xyrange", defxyrange),
		serverutil.GetColorFromQuery(req, "maxcolor", defMaxColor),
		serverutil.GetColorFromQuery(req, "mincolor", defMinColor),
		function)
}
