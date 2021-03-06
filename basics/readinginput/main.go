package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	// run
	// Get-Content text.txt | go run main.go

	// counts frequency of words in a file containing 1 word per line

	m := map[string]int{}
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		m[scanner.Text()]++
	}
	if err := scanner.Err(); err != nil {
		log.Printf("ERROR: could not read stdin: %s", err)
	}
	for k, v := range m {
		log.Printf("%s => %d", k, v)
	}
}
