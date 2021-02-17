package main

import (
	"fmt"
	"sync"
	"time"
)

func hello() {
	fmt.Println("hello from a goroutine")
}

func synchronized() {
	var wg sync.WaitGroup
	wg.Add(3)
	go func() {
		defer wg.Done()
		fmt.Println("go 1")
	}()
	go func() {
		defer wg.Done()
		fmt.Println("go 2")
	}()
	go func() {
		defer wg.Done()
		fmt.Println("go 3")
	}()

	wg.Wait()
}

func main() {
	go hello()
	// anonymous function
	go func() {
		fmt.Println("Hello from anonymous function!")
	}()
	synchronized()
	time.Sleep(1 * time.Second)
}
