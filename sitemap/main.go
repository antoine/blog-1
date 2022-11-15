package main

import (
	"fmt"
	"sort"

	"github.com/antoine/blog-1"
)

func main() {

	posts := blog.OrderedList()
	sort.Sort(sort.Reverse(posts))

	for _, v := range posts {
		fmt.Printf("https://delaunay.org/henri%s\n", v.URL)
	}

}
