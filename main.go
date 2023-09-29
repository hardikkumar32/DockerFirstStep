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

// Oweipadei Joshua Bayefa - 500221880
func calculateAndPrintSum(numbers []int) int {
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return sum
}

// yash kelkar - 500223746
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// aarti - 500224115
func concatenat() {

	s1 := "Hi "
	s2 := "This IS Lab 2 Assignment"

	concatenated_string := s1 + s2
	log.Println(concatenated_string)
}

// sukhneet 500224802
func addition(x int, y int) int {
	return x + y
}

// dennis 500222264
func sam() {
	etheriumPrice := 20000
	log.Print("The current price of Etherium is:", etheriumPrice)
}

func main() {
	log.Println(add(42, 13))
	log.Println(swap(5, 7))
	log.Println("Multiplication Table of 15:")

	numbers := []int{1, 2, 3, 4, 5}
	sum := calculateAndPrintSum(numbers)
	log.Println("Sum:", sum)

	result := max(10, 7)
	log.Println("The maximum value is", result)

	concatenat()
	sam()
	log.Println(addition(42, 13))
}
