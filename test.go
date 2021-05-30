package main

import (
	"fmt"
)

func main() {
	var x int
	var y int

	x = 1
	y = 2

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(x - y)
	//declaring variable in different ways!

}

func Second() {
	var x string = "first name"
	var y string = "last name"

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(x + y)
}

func Third() {
	x := "first name"
	y := "last name"

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(x + y)
}

//trying integer variable

func FirstNumber() {
	var a = 12
	var b = 13

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(a + b)

}

func SecondNumber() {
	var a int = 12
	var b int = 13

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(a + b)
}

func ThirdNumber() {
	a := 12
	b := 13

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(a + b)
}

//trying constant

func FirstConstant() {
	const a string = "first name"
	const b string = "last name"

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(a + b)
}

func SecondConstant() {
	const a = "first name"
	const b = "last name"

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(a + b)
}

//trying integer constant

func FirstConstantInteger() {
	const a int = 12
	const b int = 13

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(a + b)
}

func SecondConstantInteger() {
	const a = "first name"
	const b = "last name"

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(a + b)
}
