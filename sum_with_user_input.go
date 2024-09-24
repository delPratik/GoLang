package main

import "fmt"

func main() {
	var num1, num2 float64

	// Taking input from the user
	fmt.Print("Enter the first number: ")
	fmt.Scanln(&num1)

	fmt.Print("Enter the second number: ")
	fmt.Scanln(&num2)

	// Adding the numbers
	sum := num1 + num2

	// Displaying the result
	fmt.Printf("The sum of %.2f and %.2f is %.2f\n", num1, num2, sum)
}
