package omdb

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

const apiURL = "https://omdbapi.com/"

type MovieInfo struct {
	Title		string
	Year		string
	PosterURL	string	`json:"Poster"`
}

func FetchMovieInfo(title string) (*MovieInfo, error) {
	resp, err := http.Get(getURL(title))

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("Movie not found : title = %s", title)
	}

	var result MovieInfo
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
}

func getURL(title string) string {
	q := url.QueryEscape(title)
	return apiURL + "?t=" + q
}
