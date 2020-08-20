package main

import (
	"fmt"

	"gorgonia.org/gorgonia"
	"gorgonia.org/gorgonia/encoding/dot"
)

func main() {
	g := gorgonia.NewGraph()
	// define the expression
	x := gorgonia.NewScalar(g, gorgonia.Float64, gorgonia.WithName("x"))
	y := gorgonia.NewScalar(g, gorgonia.Float64, gorgonia.WithName("y"))
	gorgonia.Sigmoid(x)
	gorgonia.Add(x, y)
	if b, err := dot.Marshal(g); err == nil {
		fmt.Println(string(b))
	}
}
