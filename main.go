package main

import "log"

// Hardik function 500223366
func add(x int, y int) int {
	return x + y
}

// Sumit function 500224003
func swap(a, b int) (int, int) {
	return b, a
}

// Karan function 500224005
func printTableOf15() {
	for i := 1; i <= 10; i++ {
		result := 15 * i
		log.Printf("15 x %d = %d\n", i, result)
	}
}

func main() {
	log.Println(add(42, 13))
	log.Println(swap(5, 7))
	log.Println("Multiplication Table of 15:")
}
