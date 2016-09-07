package main

import (
	"fmt"
	"log"
	"net/http"
	"params"
)

type query struct {
	Post  string `http:"post" format:"\\d+\\-\\d+"`
	Phone string `http:"phone" format:"\\d+\\-\\d+-\\d+"`
	EMail string `http:"e" format:"[\x20-\x7e]+@[\x20-\x7e]+"`
}

func search(resp http.ResponseWriter, req *http.Request) {
	var q query
	if err := params.Unpack(req, &q); err != nil {
		http.Error(resp, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprintf(resp, "Search: %+v\n", q)
}

func main() {
	http.HandleFunc("/search", search)
	log.Fatal(http.ListenAndServe(":12345", nil))
}

/*
 * http://localhost:12345/search?post=123-45467&phone=12-34-5678&e=test@mail.com
 * Search: {Post:123-45467 Phone:12-34-5678 EMail:test@mail.com}
 *
 * http://localhost:12345/search?post=12345467&phone=12-34-5678&e=test@mail.com
 * post: Invalid format field value : 12345467 (Required format : \d+\-\d+)
 *
 * http://localhost:12345/search?post=123-45467&phone=12-34-5678&e=mail
 * e: Invalid format field value : mail (Required format : [ -~]+@[ -~]+)
 *
 */
