package main

import (
	"math/rand"
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
	time.Sleep(time.Duration(rand.Intn(1e2)) * time.Millisecond)
}
