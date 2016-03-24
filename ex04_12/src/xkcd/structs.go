package xkcd

import (
	"fmt"
	"strconv"
	"time"
)

const formatURL = "https://xkcd.com/%d/info.0.json"

type ComicInfo struct {
	Title		string
	Number		int		`json"num"`
	URL			string
	ImageURL	string	`json:"img"`
	Transcript	string
	Year		string
	Month		string
	Day			string
	Date		time.Time
}

func (c ComicInfo) validate(num int) *ComicInfo {
	year, _ := strconv.Atoi(c.Year)
	month, _ := strconv.Atoi(c.Month)
	day, _ := strconv.Atoi(c.Day)

	c.Date = time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
	c.Number = num
	c.URL = getURL(num)

	return &c
}

func getURL(num int) string {
	return fmt.Sprintf(formatURL, num)
}
