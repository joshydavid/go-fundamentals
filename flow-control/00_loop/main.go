package main

import "fmt"

func main() {
	totalSum := 0

	// Method 1: Traditional Loop
	for i := 0; i < 10; i++ {
		totalSum += i
	}

	fmt.Println((totalSum))

	// Method 2: Modernised Range Loop
	for i := range 5 {
		fmt.Println(i)
	}

	numsArr := []int{2, 4, 6, 8, 10}
	for index, value := range numsArr {
		fmt.Println("index ->", index, "value ->", value)
	}

	// While
	currSum := 10
	for currSum < 100 {
		currSum += currSum
	}
	fmt.Println(currSum)
}
