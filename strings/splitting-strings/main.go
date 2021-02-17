package main

import (
	"fmt"
	"strings"
)

func main() {
	spl := strings.Split("a,b,c", ",")
	fmt.Printf("%q", spl)
}
