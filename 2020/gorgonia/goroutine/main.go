package main

import (
	"fmt"
	"time"
)

var now = time.Now()

func add(a, b int) int {
	return a + b
}

func main() {
	for i := 0; i < 10; i++ {
		go fmt.Println(add(40, i))
	}
}

// END OMIT
