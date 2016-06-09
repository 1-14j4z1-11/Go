package serverutil

import (
	"image/color"
	"net/http"
	"strconv"
	"strings"
)

func GetStringFromQuery(req *http.Request, key string) string {
	return req.FormValue(key)
}

func GetIntFromQuery(req *http.Request, key string, defaultValue int) int {
	str := req.FormValue(key)
	value, err := strconv.Atoi(str)

	if err != nil {
		return defaultValue
	} else {
		return value
	}
}

func GetFloat64FromQuery(req *http.Request, key string, defaultValue float64) float64 {
	str := req.FormValue(key)
	value, err := strconv.ParseFloat(str, 64)

	if err != nil {
		return defaultValue
	} else {
		return value
	}
}

func GetColorFromQuery(req *http.Request, key string, defaultValue color.Color) color.Color {
	str := req.FormValue(key)
	words := strings.Split(str, "-")

	if len(words) != 3 {
		return defaultValue
	}

	r, er := strconv.Atoi(words[0])
	g, eg := strconv.Atoi(words[1])
	b, eb := strconv.Atoi(words[2])

	if er != nil || eg != nil || eb != nil {
		return defaultValue
	} else {
		return color.RGBA{uint8(r), uint8(g), uint8(b), 255}
	}
}
