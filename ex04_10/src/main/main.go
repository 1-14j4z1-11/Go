package main

import (
	"fmt"
	"github"
	"log"
	"os"
	"time"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	now := time.Now()
	lastM := now.AddDate(0, -1, 0)
	lastY := now.AddDate(-1, 0, 0)

	printIssues("Since last month", github.Filtering(result.Items, github.DateRangeFilter(&lastM, &now)))
	printIssues("Since last Year", github.Filtering(result.Items, github.DateRangeFilter(&lastY, &lastM)))
	printIssues("Others", github.Filtering(result.Items, github.DateRangeFilter(nil, &lastY)))

}

func printIssues(title string, issues []*github.Issue) {
	fmt.Println(title)
	fmt.Printf("%d issues:\n", len(issues))
	for _, item := range issues {
		fmt.Printf("#%-5d %v %9.9s %.55s\n", item.Number, item.CreatedAt, item.User.Login, item.Title)
	}
}
