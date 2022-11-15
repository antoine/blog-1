package main

import (
	"html/template"
	"os"
	"sort"

	"github.com/antoine/blog-1"
)

func main() {

	currentYear := "1900"

	funcMap := template.FuncMap{
		"newYear": func(t string) bool {
			if t == currentYear {
				return false
			} else {
				currentYear = t
				return true
			}
		},
	}

	posts := blog.OrderedList()
	sort.Sort(sort.Reverse(posts))
	t, err := template.New("metacrap").Parse(blog.Metacrap)
	t, err = t.New("foo").Funcs(funcMap).Parse(`{{ template "metacrap" }}
<meta name="description" content="Motscousus's travel blog">
<link rel="alternate" type="application/atom+xml" title="Atom feed" href="index.atom">
<title>Motscousus's blog</title>
</head>
<body>

<aside>
<p><span id=greet></span> <dfn><abbr title="A guy travelling">this</abbr>.org</dfn> is a travel blog generated by scraping <a href="https://www.polarsteps.com/mots/6037528?s=62273598-6d19-4007-91e2-a982caecb6c6">Polarsteps </a>. </p>
</aside>

<nav>
{{ range $i,$e := . }}
{{ if newYear (.PostDate.Format "2006")}}
{{ if gt $i 0 }}</ol>{{end}}
<h1>{{ .PostDate.Format "2006" }}</h1>
<ol class="index">{{ end }}
<li><time datetime="{{ .PostDate.Format "2006-01-02" }}">{{ .PostDate.Format "Jan 2" }}</time>
<a href="{{ .URL }}">{{ .Title }}</a></li>{{end}}
</ol>
</nav>

<footer>
<p><a href=https://github.com/kaihendry/natalian/blob/mk/Makefile>Generated with a Makefile</a> and a piece of <a href=https://github.com/kaihendry/blog>Golang</a></p>
</footer>

</body>
</html>
`)

	if err != nil {
		panic(err)
	}

	err = t.Execute(os.Stdout, posts)

	if err != nil {
		panic(err)
	}

}
