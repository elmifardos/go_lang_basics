package main

import "fmt"

func main() {

	var number int
	fmt.Println("inter number ")
	fmt.Scan(&number)
	if number%2 == 0 {
		fmt.Println("even")

	} else {
		fmt.Println("odd")
	}

}

// Write a go program that checks whether a given number is even or odd.
