package main

import (
	"fmt"

	"github.com/rbdsqrl/tsp/db"
	"github.com/rbdsqrl/tsp/tspprog"
)

func main() {
	places := []db.Place{
		db.Place{X: 1, Y: 15},
		db.Place{X: 12, Y: 8},
		db.Place{X: 13, Y: 9},
		db.Place{X: 1, Y: 11},
		db.Place{X: 21, Y: 11},
		db.Place{X: 18, Y: 17},
		db.Place{X: 11, Y: 21},
		db.Place{X: 12, Y: 18},
		db.Place{X: 11, Y: 7},
	}
	minPath := tspprog.CalculateShortestPath(places, db.Place{X: 0, Y: 0})
	fmt.Println(minPath)
}
