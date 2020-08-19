package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/fatih/color"
)

const (
	signal int = 0
)

func printRed(p int, s string) {
	var tabs string
	for i := 0; i < p; i++ {
		tabs = tabs + "\t"
	}
	color.Red(tabs + s)
}
func printGreen(p int, s string) {
	var tabs string
	for i := 0; i < p; i++ {
		tabs = tabs + "\t"
	}
	color.Green(tabs + s)
}

func sleepRand() {
	time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
}

func player(play int, num int, pingpong string, receive <-chan int, send chan<- int) {
	for ball := range receive {
		if ball == num {
			return
		}
		time.Sleep(200 * time.Millisecond)
		ball++
		send <- ball
		printRed(play, pingpong+strconv.Itoa(ball))
	}
}

// ENDPLAYER OMIT

func sendBall(num int, launcher chan int) {
	for ball := 0; ball < num; ball++ {
		launcher <- ball
	}
}

func endGame(numBall int, launcher, exchange, receiver, playC chan int) {
	for ball := 0; ball < numBall; ball++ {
		<-receiver
	}
	close(launcher)
	close(exchange)
	close(receiver)
	playC <- 0

}
func main() {
	start := time.Now()
	playC := make(chan int)
	for play := 0; play < 5; play++ {
		// START_LAUNCHER OMIT
		launcher := make(chan int)
		exchange := make(chan int)
		receiver := make(chan int)
		// END_LAUNCHER OMIT
		// START_PLAYER OMIT
		go player(play, "ping", launcher, receiver)
		go player(play, "pong", receiver, launcher)
		// END_PLAYER OMIT
		go sendBall(5, launcher)
		go endGame(5, launcher, exchange, receiver, playC)
	}
	// END_FOR OMIT
	<-playC
	<-playC
	<-playC
	<-playC
	<-playC
	fmt.Println(time.Since(start))
}

//END_MAIN OMIT
