package main

import (
	"io/ioutil"
	"log"
	"sort"
	"time"

	"github.com/gorilla/feeds"
	"github.com/antoine/blog-1"
)

func main() {

	author := feeds.Author{"Motscousus", "example@example.com"}
	now := time.Now()
	feed := &feeds.Feed{
		Title:       "Motscousus's travel blog",
		Link:        &feeds.Link{Href: "https://www.polarsteps.com/mots/6037528?s=62273598-6d19-4007-91e2-a982caecb6c6"},
		Description: "travel blog",
		Author:      &author,
		Created:     now,
	}

	posts := blog.OrderedList()
	sort.Sort(sort.Reverse(posts))

	maximum := 20
        if (maximum > len(posts)) {
                maximum = len(posts)
        }
		
	for _, v := range posts[:maximum] {
		if v.Description == "" {
			log.Println("Warning:", v.URL, "has no tl;dr")
		}
		i := feeds.Item{
			Title:       v.Title,
			Link:        &feeds.Link{Href: "https://delaunay.org/henri" + v.URL},
			Description: v.Description,
			Created:     v.PostDate,
		}
		feed.Add(&i)
	}

	atom, _ := feed.ToAtom()
	rss, _ := feed.ToRss()

	//fmt.Println(atom)
	ioutil.WriteFile("index.atom", []byte(atom), 0644)
	ioutil.WriteFile("index.rss", []byte(rss), 0644)

}
