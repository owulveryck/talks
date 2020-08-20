package main

import (
	"fmt"
	"log"

	"gorgonia.org/gorgonia"
)

func main() {
	g := gorgonia.NewGraph()

	// define the expression
	x := gorgonia.NewScalar(g, gorgonia.Float64, gorgonia.WithName("x"))
	y := gorgonia.NewScalar(g, gorgonia.Float64, gorgonia.WithName("y"))
	z, err := gorgonia.Add(x, y)
	if err != nil {
		log.Fatal(err)
	}
	// set initial values then run
	gorgonia.Let(x, 2.0)
	gorgonia.Let(y, 40.0)

	// create a VM to run the program on
	machine := gorgonia.NewTapeMachine(g)
	defer machine.Close()
	machine.RunAll()
	fmt.Printf("%v", z.Value())
}
