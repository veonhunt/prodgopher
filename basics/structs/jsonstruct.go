package main

import (
	"encoding/json"
	"log"
)

// Person has a name and email
type Person struct {
	// Name  string
	// Email string
	Name  string `json:"name"`
	Email string `json:"email"`
}

func anonymousStruct() {
	p := struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}{
		Name:  "Robert",
		Email: "robert@example.com",
	}
	b, err := json.Marshal(p)
	if err != nil {
		log.Fatal(err)
	}

	println(string(b))
}

func main() {
	// p := Person{"Robert", "robert@example.com"}
	// b, err := json.Marshal(p)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(string(b))

	anonymousStruct()
}
