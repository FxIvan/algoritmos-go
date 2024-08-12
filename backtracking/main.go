package main

import (
	"fmt"
	"math"
)

// Función principal para descomponer un número n
func decompose(n int) []int {
	result := helper(n, n*n)
	if len(result) == 0 {
		return nil
	}
	return result
}

// Función recursiva helper que realiza la descomposición
func helper(n int, remain int) []int {
	if remain == 0 {
		return []int{}
	}
	fmt.Print("Remain -->", remain, "\n")
	for i := n - 1; i > 0; i-- {
		fmt.Print("I -->", i, "\n")
		fmt.Print("remain-i*i ->", remain-i*i, "\n")
		fmt.Print("------------------------------FIN------------------------------\n")
		if remain-i*i >= 0 {
			result := helper(i, remain-i*i)
			if result != nil {
				return append(result, i)
			}
		}
	}
	return nil
}

// ////////////////////////////////
var distances = map[string]map[string]int{
	"A": {"B": 10, "C": 15, "D": 20},
	"B": {"A": 10, "C": 35, "D": 25},
	"C": {"A": 15, "B": 35, "D": 30},
	"D": {"A": 20, "B": 25, "C": 30},
}

func getDistance(a, b string) int {
	return distances[a][b]
}

func calculateDistance(path []string) int {
	var distance int
	for i := 0; i < len(path)-1; i++ {
		distance = getDistance(path[i], path[i+1])
	}
	return distance
}

func findOptimalRoute(destinations []string) ([]string, int) {
	var bestPath []string
	bestDistance := math.MaxInt32

	var backtrack func(currentPath []string, remainingDestinations []string)
	backtrack = func(currentPath []string, remainingDestinations []string) {

		if len(remainingDestinations) == 0 {
			currentDistance := calculateDistance(currentPath)
			if currentDistance < bestDistance {
				bestDistance = currentDistance
				fmt.Println("currentPath -->", currentPath)
				bestPath = append([]string{}, currentPath...)
			}
			return
		}

		for i := 0; i < len(remainingDestinations); i++ {
			nextDestination := remainingDestinations[i]
			newPath := append(currentPath, nextDestination)
			newRemainingDestinations := append([]string{}, remainingDestinations[:i]...)
			newRemainingDestinations = append(newRemainingDestinations, remainingDestinations[i+1:]...)
			backtrack(newPath, newRemainingDestinations)
		}
	}

	backtrack([]string{}, destinations)
	return bestPath, bestDistance
}

// Función principal para pruebas
func main() {
	//fmt.Println(decompose(11)) // [1, 2, 3, 6]
	destinations := []string{"A", "B", "C", "D"}
	fmt.Println(findOptimalRoute(destinations))
}
