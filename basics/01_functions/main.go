package main

import "fmt"

// Functions
func add(x int, y int) int {
	return x + y
}

// Named return values
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	totalSum := add(5, 3)
	x, y := split(17)

	fmt.Println(totalSum)
	fmt.Println(x, y)
}
