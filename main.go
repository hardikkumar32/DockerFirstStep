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

func main() {
	log.Println(add(42, 13))
	log.Println(swap(5, 7))
}
