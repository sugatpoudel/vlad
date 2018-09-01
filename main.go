package main

import (
	"fmt"

	"github.com/sugatpoudel/vlad/graph"
)

func main() {
	vlad := graph.NewGraph()

	isPositive := &graph.Node{
		Validate: func(num int) bool { return num > -1 },
		Msg:      "Number is not positive",
	}

	isEven := &graph.Node{
		Validate: func(num int) bool { return num%2 == 0 },
		Msg:      "Number is not even",
	}

	isMultipleOfThree := &graph.Node{
		Validate: func(num int) bool { return num%3 == 0 },
		Msg:      "Number is not a multiple of 3",
	}

	vlad.PutEdge(isPositive, isEven)
	vlad.PutEdge(isPositive, isMultipleOfThree)
	for _, err := range vlad.Validate(7) {
		fmt.Println(err)
	}
}
