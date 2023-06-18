package main

import "fmt"

func main() {

	var age int

	fmt.Println("inter age ")

	fmt.Scan(&age)

	if age > 60 {

		fmt.Println("getting older")

	} else if age > 30 {
		fmt.Println("getting wiser")

		fmt.Println("age")

	} else if age > 20 {
		fmt.Println("adulthood")

	} else if age > 10 {
		fmt.Println("young blood")

	} else {
		fmt.Println("booting up")
	}

}
