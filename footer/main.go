package main

import (
	"flag"
	"html/template"
	"os"
	"path/filepath"
)

type Post struct {
	URL  string
	Mdwn string
}

var p Post

func main() {

	flag.Parse()
	mdwn := flag.Arg(0)

	extName := filepath.Ext(mdwn)
	bName := mdwn[:len(mdwn)-len(extName)]

	p = Post{URL: bName, Mdwn: mdwn}

	t, err := template.New("foo").Parse(`
	</article>

</body>
</html>
`)

	if err != nil {
		panic(err)
	}

	err = t.Execute(os.Stdout, p)

	if err != nil {
		panic(err)
	}

}
