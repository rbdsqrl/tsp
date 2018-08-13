package main

import (
	"fmt"

	"github.com/rbdsqrl/db/schema"
)

func main() {
	places := []string{
		schema.Place{1, 5},
		schema.Place{2, 8},
		schema.Place{3, 9},
		schema.Place{1, 1},
		schema.Place{2, 1},
		schema.Place{8, 7},
	}
	//minPath := CalculateShortestPath(places, Place{0, 0})
	fmt.Println("cow")
}
