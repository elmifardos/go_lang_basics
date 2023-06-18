package main

import "fmt"

func main() {
	var num1, num2 int
	fmt.Println("Enter first number")
	fmt.Scan(&num1)

	fmt.Println("Enter second number")
	fmt.Scan(&num2)

	total(num1, num2)
}

// function that recieves 2 values and calculates the sum
func total(number1 int, number2 int) {
	sum := number1 + number2
	fmt.Printf("The sum is %d \n", sum)
}
