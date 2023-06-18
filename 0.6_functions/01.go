package main

import "fmt"

func main() {

	fmt.Println("start")
	odd()

}

// function that check if a number is odd or even
func odd() {

	var number int
	fmt.Println("Enter a number") // instruction
	fmt.Scan(&number)             // receive the number

	if number%2 != 0 {
		fmt.Println("odd")

	} else {
		fmt.Println("even")

	}

}
