package main

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"serverutil"
	"sync"
)

var mutex sync.Mutex

func main() {
	db := database{"shoes": 50, "socks": 5}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	http.HandleFunc("/create", db.create)
	http.HandleFunc("/update", db.update)
	http.HandleFunc("/remove", db.remove)
	log.Fatal(http.ListenAndServe("localhost:50000", nil))
}

////////////////////////////////////////////////////////////////////////

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

////////////////////////////////////////////////////////////////////////

var tableTemplate = `<html>
<body>

<h1>ItemList</h1>
<table border="1">
<tr>
<td>Item</td>
<td>Price</td>
</tr>
{{range $key, $value := .}}
<tr>
<td>{{$key}}</td>
<td>{{$value}}</td>
</tr>
{{end}}
</table>

</body>
</html>
`

var itemQueryKey = "item"
var priceQueryKey = "price"

var table = template.Must(template.New("table").Parse(tableTemplate))

type database map[string]dollars

func (db database) list(w http.ResponseWriter, req *http.Request) {
	writer := bytes.NewBuffer(nil)
	err := table.Execute(writer, db)
	if err != nil {
		log.Fatal(err)
	}
	w.Write(writer.Bytes())
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	db.tryGetItem(w, req, func(item string) {
		fmt.Fprintf(w, "%s\n", db[item])
	})
}

func (db database) create(w http.ResponseWriter, req *http.Request) {
	lockedAction(func() {
		item := serverutil.GetStringFromQuery(req, itemQueryKey)
		price := serverutil.GetFloat64FromQuery(req, priceQueryKey, -1)

		if item != "" && price >= 0 {
			db[item] = dollars(price)
			fmt.Printf("Create '%s' : %f\n", item, price)
		}
		db.list(w, req)
	})
}

func (db database) update(w http.ResponseWriter, req *http.Request) {
	db.tryGetItem(w, req, func(item string) {
		if price := serverutil.GetFloat64FromQuery(req, priceQueryKey, -1); price >= 0 {
			db[item] = dollars(price)
			fmt.Printf("Update '%s' : %f\n", item, price)
		}
		db.list(w, req)
	})
}

func (db database) remove(w http.ResponseWriter, req *http.Request) {
	db.tryGetItem(w, req, func(item string) {
		delete(db, item)
		fmt.Printf("Remove '%s'\n", item)
		db.list(w, req)
	})
}

func (db database) tryGetItem(w http.ResponseWriter, req *http.Request, actionIfExist func(string)) {
	lockedAction(func() {
		item := serverutil.GetStringFromQuery(req, itemQueryKey)
		if _, ok := db[item]; ok {
			actionIfExist(item)
		} else {
			w.WriteHeader(http.StatusNotFound) // 404
			fmt.Fprintf(w, "no such item: %q\n", item)
		}
	})
}

func lockedAction(action func()) {
	mutex.Lock()
	action()
	mutex.Unlock()
}
