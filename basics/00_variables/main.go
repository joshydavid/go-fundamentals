package main

import "fmt"

func main() {
	// Method 1: The 'var' keyword (Explicit)
	// Use this when you want to declare a variable without assigning yet
	var dailyGoal string
	dailyGoal = "Learn Go Fundamentals"

	// Method 2: Short variable declaration (Inferred)
	// Can only be used inside functions. Go "guesses" the type
	daysCompleted := 1
	isCodingFun := true

	// Method 3: Multiple declaration
	var (
		coffeeCups = 2
		hoursSpent = 1.5
	)

	// Printing the results
	fmt.Println("--- Day 1: Variables ---")
	fmt.Printf("Goal: %s\n", dailyGoal)
	fmt.Printf("Progress: Day %d\n", daysCompleted)
	fmt.Printf("Is it fun? %t\n", isCodingFun)
	fmt.Printf("Stats: %d cups of coffee, %.1f hours\n", coffeeCups, hoursSpent)

	// %s	String	Plain text (e.g., "Learn Go")
	// %d	Digits	Integers / Whole numbers (e.g., 1, 2, 3)
	// %t	Truth	Booleans (true or false)
	// %f	Float	Decimal numbers (e.g., 1.500000)
	// %.1f	Precision Float	Decimal number rounded to 1 decimal place (e.g., 1.5)
	// \n	Newline	Moves the cursor to the next line (Enter key)
}
