package main

import (
	"fmt"
	"learn"
)

func main() {
	lru := learn.LRU(5)
	lru.Set(1, 1)
	fmt.Println(lru.Get(1))
}
