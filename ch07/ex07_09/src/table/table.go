package table

import (
	"bytes"
	"html/template"
	"log"
	"track"
)

var tableTemplate = `<html>
<body>

<h1>TrackList</h1>
<table border="1">
<tr>
<td>Title</td>
<td>Artist</td>
<td>Album</td>
<td>Year</td>
<td>Time</td>
</tr>
{{range .}}
<tr>
<td>{{.Title}}</td>
<td>{{.Artist}}</td>
<td>{{.Album}}</td>
<td>{{.Year}}</td>
<td>{{.Length}}</td>
</tr>
{{end}}
</table>

</body>
</html>
`

var table = template.Must(template.New("table").Parse(tableTemplate))

func CreateTableHTML(list *track.TrackList) string {
	writer := bytes.NewBuffer(nil)
	err := table.Execute(writer, list)
	if err != nil {
		log.Fatal(err)
	}
	return writer.String()
}
