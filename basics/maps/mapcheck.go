package main

import "fmt"

func main() {
	m := make(map[string]string)
	m["en"] = "Hello"
	m["fr"] = "Bonjour"

	// loop over keys and values
	for k, v := range m {
		fmt.Printf("%q => %q\n", k, v)
	}

	// nonexistent key returns zero value
	fmt.Printf("zh: %q\n", m["zh"])

	// check for existence with _, ok := m[key]
	if _, ok := m["fr"]; ok {
		fmt.Printf("key %q exists in map\n", "fr")
	}

	delete(m, "en")
	fmt.Printf("m length: %q\n", len(m))
	fmt.Println(m)
}
