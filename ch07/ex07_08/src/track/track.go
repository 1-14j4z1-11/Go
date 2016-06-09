package track

import (
	"fmt"
	"time"
)

////////////////////////////////////////////////////////////////

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

func Length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

////////////////////////////////////////////////////////////////

const (
	SortKeyTitle  = "Title"
	SortKeyArtist = "Artist"
	SortKeyAlbum  = "Album"
	SortKeyYear   = "Year"
	SortKeyLength = "Length"
)

type TrackList []Track

var lessFunc map[string]func(i, j Track) bool = map[string]func(i, j Track) bool{
	SortKeyTitle:  func(i, j Track) bool { return i.Title < j.Title },
	SortKeyArtist: func(i, j Track) bool { return i.Artist < j.Artist },
	SortKeyAlbum:  func(i, j Track) bool { return i.Album < j.Album },
	SortKeyYear:   func(i, j Track) bool { return i.Year < j.Year },
	SortKeyLength: func(i, j Track) bool { return i.Length < j.Length },
}

var lessOrder []string

func init() {
	ResetOrder()
}

func ResetOrder() {
	lessOrder = []string{
		SortKeyTitle,
		SortKeyArtist,
		SortKeyAlbum,
		SortKeyYear,
		SortKeyLength,
	}
}

func SetFirstOrder(key string) error {

	for i, val := range lessOrder {
		if val == key {
			old := lessOrder

			lessOrder = make([]string, 0, len(old))
			lessOrder = append(lessOrder, key)

			for j, v := range old {
				if i != j {
					lessOrder = append(lessOrder, v)
				}
			}

			return nil
		}
	}

	return fmt.Errorf("Unknown less function key : %s", key)
}

func (t *TrackList) Less(i, j int) bool {
	s := *t

	for _, key := range lessOrder {
		if lessFunc[key](s[i], s[j]) {
			return true
		} else if lessFunc[key](s[j], s[i]) {
			return false
		}
	}

	return false
}

func (t *TrackList) Swap(i, j int) {
	s := *t
	s[i], s[j] = s[j], s[i]
}

func (t *TrackList) Len() int {
	return len(*t)
}
