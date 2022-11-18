package main

import (
	"flag"
	"html/template"
	"os"
	"path/filepath"
	"strings"

	"github.com/antoine/blog-1"
)

type Post struct {
	Title       string
	URL         string
	Description string
}

var p Post

func main() {

	flag.Parse()
	mdwn := flag.Arg(0)

	fName := filepath.Base(mdwn)
	extName := filepath.Ext(mdwn)
	bName := fName[:len(fName)-len(extName)]
	url := mdwn[:len(mdwn)-len(extName)]

	m := blog.GetKey(mdwn, "title")
	desc := blog.GetKey(mdwn, "description")["description"]

	title := m["title"]

	if title == "" {
		title = strings.Replace(bName, "_", " ", -1)
	}

	p = Post{Title: title, URL: url, Description: desc}

	t, err := template.New("metacrap").Parse(blog.Metacrap)
	t, err = t.New("foo").Parse(`{{ template "metacrap" }}
<link href="https://delaunay.org/henri/{{ .URL }}/" rel=canonical>
{{if .Description}}<meta name="description" content="{{ .Description }}">{{end}}
<title>{{ .Title }}</title>
</head>
<body>
<nav><a href="/henri/">delaunay.org/henri</a></nav>
<article>
<h1 class="headline"><a href="/henri/{{ .URL }}/">{{ .Title }}</a></h1>
<p>scraped with love from <a href="https://www.polarsteps.com/mots/6037528?s=62273598-6d19-4007-91e2-a982caecb6c6">polarstep</a></p>
`)
	if err != nil {
		panic(err)
	}

	err = t.Execute(os.Stdout, p)

	if err != nil {
		panic(err)
	}

}
