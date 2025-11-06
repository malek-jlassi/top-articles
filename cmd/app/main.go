package main

import (
	"fmt"
	"github.com/malek-jlassi/top-articles/business_logic/articles"
)

func main() {
	top := articles.TopArticles(5)
	fmt.Println("Top Articles:")
	for i, name := range top {
		fmt.Printf("%d. %s\n", i+1, name)
	}
}