package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "banana"
	c1 := strings.Count(s, "a")
	c2 := strings.Count(s, "ana")
	fmt.Printf("%d, %d", c1, c2)
}
