package tspprog

import (
	"math"

	. "github.com/rbdsqrl/db/schema"
)

var perms [][]int
var dist []float64

//CalculateShortestPath Calculates the shortest path from start through all points and then the end
func CalculateShortestPath(places []Place, start Place) []Place {
	var path []Place
	indices := getIndices(len(places))
	permuteHelper(indices, 0)
	dist = getDistances(places, start)
	minDist := dist[0]
	minDistIndex := 0
	for i, val := range dist {
		if val < minDist {
			minDistIndex = i
		}
	}
	minPath := perms[minDistIndex]
	for i, val := range minPath {
		path[i] = places[val]
	}
	return path
}

func getDistances(places []Place, start Place) []float64 {
	for _, val := range perms {
		lastIndex := len(val) - 1
		pathDistance := distance(start, places[val[0]])
		for i = 1; i < len(val); i++ {
			pathDistance += distance(places[i], places[i-1])
		}
		pathDistance += distance(start, places[lastIndex])
	}
}

func getIndices(length int) []int {
	var indices []int
	for i := 0; i < length; i++ {
		indices = append(indices, i)
	}
	return indices
}

func distance(place1 Place, place2 Place) float64 {
	diffX := place1.x - place2.x
	diffY := place1.y - place2.y
	return math.Sqrt(diffX*diffX - diffY*diffY)
}

func swap(indices *[]int, pos1 int, pos2 int) {
	temp := indices[pos1]
	indices[pos1] = indices[pos2]
	indices[pos2] = temp
}

func permuteHelper(indices []int, int index) {
	if index >= len(int)-1 {
		//add the permutations to the slice
		perms = append(perms, indices)
	}

	for i := index; i < arr.length; i++ {
		//Swap the elements at indices index and i
		swap(&indices, i, index)
		//Recurse on the sub array arr[index+1...end]
		permuteHelper(arr, index+1)

		//Swap the elements back
		t = arr[index]
		arr[index] = arr[i]
		arr[i] = t
	}
}
