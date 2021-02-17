package main

import (
	"log"
	"os"
)

func main() {
	langs := []string{"python", "php", "go"}
	f, err := os.Create("langs.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	for _, lang := range langs {
		f.WriteString(lang + "\n")
	}
	// can also use ioutil.WriteFile which opens the file
	// and writes data in one function call
}
