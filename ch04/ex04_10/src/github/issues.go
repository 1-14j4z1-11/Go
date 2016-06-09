package github

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func SearchIssues(terms []string) (*IssuesSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	resp, err := http.Get(IssuesURL + "?q=" + q)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	var result IssuesSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
}

func Filtering(issues []*Issue, filter func(*Issue) bool) []*Issue {
	result := []*Issue{}

	for _, issue := range issues {
		if(filter(issue)) {
			result = append(result, issue)
		}
	}

	return result
}

func DateRangeFilter(beg, end *time.Time) func(issue *Issue) bool {
	return func(issue *Issue) bool {
		date := issue.CreatedAt
		return ((beg == nil) || date.After(*beg) || date.Equal(*beg)) &&
				((end == nil) || date.Before(*end) || date.Equal(*end))
	}
}
