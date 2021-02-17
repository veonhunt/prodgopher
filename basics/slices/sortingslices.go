package main

import (
	"fmt"
	"sort"
)

// Country has a name and population
type Country struct {
	Name       string
	Population int
}

func main() {
	c := []Country{
		{"South Africa", 55910000},
		{"United States", 3231000000},
		{"Japan", 1270000000},
		{"England", 53010000},
	}

	sort.Slice(c, func(i, j int) bool { return c[i].Name < c[j].Name })
	fmt.Println("Countries by name:", c)

	sort.Slice(c, func(i, j int) bool { return c[i].Population < c[j].Population })
	fmt.Println("Countries by population:", c)
}
