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

// Trying if statements in golang

func IfStatement1() {
	var a int = 2
	if a%5 == 0 {
		fmt.Printf("%d is the multiple of 5\n", a)
	}
}

func IfStatement2() {
	var greater int = 30
	if greater >= 20 {
		fmt.Printf("%d is greater", greater)
	}
}

func IfStatementBetween() {
	var number int = 50
	if number >= 20 && number <= 60 {
		fmt.Printf("%d is in between", number)
	}
}

//trying if statements in different way
func DiffIfStatement() {
	number := 12
	if number <= 20 {
		fmt.Printf("%d is smaller", number)
	} else {
		fmt.Printf("%d us greater", number)
	}
}

//second trial

func IfStatementBetweenDiff() {
	number := 50
	if number >= 20 && number <= 60 {
		fmt.Printf("%d is in between", number)
	} else {
		fmt.Printf("%d is not in between", number)
	}
}

//third trial

func IfStatement1Diff() {
	a := 2
	if a%5 == 0 {
		fmt.Printf("%d is the multiple of 5\n", a)
	} else {
		fmt.Printf("%d is not the multiple of 5\n", a)
	}
}
