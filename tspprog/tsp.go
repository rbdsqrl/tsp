package tspprog

import (
	"fmt"
	"math"

	"github.com/rbdsqrl/tsp/db"
)

var perms [][]int
var dist []float64

//CalculateShortestPath Calculates the shortest path from start through all points and then the end
func CalculateShortestPath(places []db.Place, start db.Place) []db.Place {
	var path []db.Place
	indices := getIndices(len(places))
	permuteHelper(indices, 0)
	dist = getDistances(places, start)
	minDist := dist[0]
	minDistIndex := 0
	for i, val := range dist {
		if val < minDist {
			minDistIndex = i
			minDist = val
		}
	}
	minPath := perms[minDistIndex]
	fmt.Println(minPath)
	for _, val := range minPath {
		path = append(path, places[val])
	}
	return path
}

func getDistances(places []db.Place, start db.Place) []float64 {
	var pathDistance float64
	for _, val := range perms {
		lastIndex := len(val) - 1
		pathDistance = Distance(start, places[val[0]])
		for i := 1; i < len(val); i++ {
			pathDistance += Distance(places[val[i]], places[val[i-1]])
		}
		pathDistance += Distance(start, places[lastIndex])
		dist = append(dist, pathDistance)
	}
	return dist
}

func getIndices(length int) []int {
	var indices []int
	for i := 0; i < length; i++ {
		indices = append(indices, i)
	}
	return indices
}

func swap(indices []int, pos1 int, pos2 int) {
	temp := indices[pos1]
	indices[pos1] = indices[pos2]
	indices[pos2] = temp
}

func permuteHelper(indices []int, index int) {

	if index >= len(indices)-1 {
		//add the permutations to the slice
		permutation := make([]int, len(indices))
		copy(permutation, indices)
		perms = append(perms, permutation)
		return
	}

	for i := index; i < len(indices); i++ {
		//Swap the elements at indices index and i
		swap(indices, i, index)
		//Recurse on the sub array arr[index+1...end]
		permuteHelper(indices, index+1)
		//Swap the elements back
		swap(indices, i, index)
	}
}

//Distance finds the distance between two places
func Distance(place1 db.Place, place2 db.Place) float64 {
	diffX := place1.X - place2.X
	diffY := place1.Y - place2.Y
	return math.Sqrt(diffX*diffX + diffY*diffY)
}
