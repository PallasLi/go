package main

import (
	"fmt"
	"time"
)

func testTimeout() {
	c := make(chan bool, 1)

	go func() {
		select {
		case m := <-c:
			// do something
			handle(m)
		case <-time.After(2 * time.Second):
			fmt.Println("timed out")
		}
	}()

	time.Sleep(3 * time.Second)
}

func handle(m bool) {}
 
 
func main() {
	fmt.Println("dd")
	// You can use Format argument of http://golang.org/pkg/time/#pkg-constants :
	fmt.Println(time.Now().Format(time.Kitchen))

	// also, you can input directly as string :
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	fmt.Println(time.Now().Format("15:04:05"))
	
	// Timestamp
	fmt.Println(time.Now().Unix()) // Ex: 1257894000
	testTimeout();
}
