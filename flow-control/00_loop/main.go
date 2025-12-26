package main

import "fmt"

func main() {
	totalSum := 0

	// Method 1: Traditional Loop
	for i := 0; i < 10; i++ {
		totalSum += i
	}

	fmt.Println((totalSum))

	// Method 2: Modernised
	for i := range 5 {
		fmt.Println(i)
	}

	// While
	currSum := 10
	for currSum < 100 {
		currSum += currSum
	}
	fmt.Println(currSum)
}
