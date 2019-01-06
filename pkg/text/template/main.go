// go run main.go > out.html && $BROWSER out.html
package main

import (
	"os"
	"text/template"
)

const txt = `<!DOCTYPE html>
<html>
<head>
	<meta charset="utf-8" />
	<title>html template</title>
</head>
<body>
	<h1>html template</h1>

	<!-- for text/template -->
{{range .}}
	<h2>{{.Title}}</h2>
	<p>{{.Body}}</p>
{{end}}
</body>
</html>
`

func main() {
	dataSet := []struct {
		Title string
		Body  string
	}{
		{
			Title: "hello",
			Body:  "world",
		},
		{
			Title: "foo",
			Body:  "bar",
		},
	}
	tmpl := template.Must(template.New("html").Parse(txt))
	err := tmpl.Execute(os.Stdout, dataSet)
	if err != nil {
		panic(err)
	}
}
