package main

import "fmt"

func main() {
	fmt.Println("Enter your first Name: ")

	var firstname string

	fmt.Scanln(&firstname)
	fmt.Println("Enter your Last Name: ")

	var lastname string

	fmt.Scanln(&lastname)

	fmt.Println("Your Name is " + firstname + lastname)
}
