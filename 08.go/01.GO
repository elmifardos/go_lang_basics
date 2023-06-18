
package main

import "fmt"

func main () {

	// Write a Go program that takes two numbers as input and calculates their sum.
	var num1,num2 int 
	fmt.Println("inter num1")
	fmt.Scan(&num1)
	fmt.Println("inter num2")
	fmt.Scan(&num2)
result:= total (num1,num2) // calling function and given it 2 numbers 
fmt.Println(result)
}

func total (A,B int) int {

	return A+B 
}
