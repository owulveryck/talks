package main

import (
	"fmt"
	"time"
)

func add(c chan int) {
	for {
		fmt.Printf("waiting\r")
		a := <-c
		fmt.Printf("received a: %v\n", a)
		fmt.Printf("waiting\r")
		b := <-c
		fmt.Printf("received b: %v\n", b)
		fmt.Println(a + b)
	}
}

func main() {
	c := make(chan int)
	go add(c)
	for i := 0; i < 10; i++ {
		c <- i
		time.Sleep(2 * time.Second)
	}
}

// END OMIT
