package main

import (
	"fmt"

	lru "github.com/hashicorp/golang-lru/v2"
)

func main() {
	cache, _ := lru.New[string, string](10)
	cache.Add("a", "b")
	fmt.Println(cache.Len())
}
