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

func playerPing(play int, ball int) {
	time.Sleep(200 * time.Millisecond)
	printGreen(play, "ping"+strconv.Itoa(ball))
}
func playerPong(play int, ball int) {
	time.Sleep(200 * time.Millisecond)
	printRed(play, "pong"+strconv.Itoa(ball))
}
func main() {
	start := time.Now()
	for court := 0; court < 5; court++ {
		for ball := 0; ball < 5; ball++ {
			playerPing(court, ball)
			playerPong(court, ball)
		}
	}
	fmt.Println(time.Since(start))
}

//END OMIT
