package main

import (
	"fmt"
	"strconv"
	"time"
)

func pong(player int, receive chan int, send chan int) {
	for {
		j := <-receive
		sleepRand()
		printRed(player, "pong"+strconv.Itoa(j))
		send <- j
	}
}
func ping(player int, receive chan int, send chan int) {
	for {
		j := <-receive
		j++
		sleepRand()
		printGreen(player, "ping"+strconv.Itoa(j))
		send <- j
	}
}

func main() {
	start := time.Now()
	for i := 0; i < 10; i++ {
		c1 := make(chan int)
		c2 := make(chan int)
		go ping(i, c1, c2)
		go pong(i, c2, c1)
		c1 <- 0
	}
	time.Sleep(2 * time.Second)
	fmt.Println(time.Since(start))
}
