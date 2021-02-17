package main

import (
	"fmt"
	"log"
)

func parseBool(str string) (bool, error) {
	switch str {
	case "1", "t", "T", "true", "TRUE", "True":
		return true, nil
	case "0", "f", "F", "false", "FALSE", "False":
		return false, nil
	}
	return false, fmt.Errorf("invalid input %q", str)
	// use fmt.Errorf to return errors with messages from functions
}

func main() {
	b, err := parseBool("q")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(b)
}
