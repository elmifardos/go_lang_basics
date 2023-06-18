package main

import "fmt"

func main() {
	var num1, num2 int
	fmt.Println("Enter first number")
	fmt.Scan(&num1)

	fmt.Println("Enter second number")
	fmt.Scan(&num2)

	ans := total(num1, num2)
	fmt.Printf("The sum inside main is %d \n", ans)
}

// function that recieves 2 values and returns the sum
func total(number1 int, number2 int) int {
	sum := number1 + number2
	fmt.Printf("The sum inside total is %d \n", sum)
	return sum
}
