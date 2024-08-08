package main

import (
	"fmt"
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

// Función principal para pruebas
func main() {
	fmt.Println(decompose(11)) // [1, 2, 3, 6]
}
