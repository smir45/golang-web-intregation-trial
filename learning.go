package main

import (
	"fmt"
)

const a = 234

func main() {
	const a = 345
	fmt.Println(a)
}

func main2() {
	//Unanamed untyped constant
	fmt.Printf("Type: %T Value: %v\n", 123, 123)
	fmt.Printf("Type: %T Value: %v\n", "circle", "circle")
	//fmt.Printf("Type: %T Value: %v\n", 5.6, 5.6)
	fmt.Printf("Type: %T Value: %v\n", true, true)
	fmt.Printf("Type: %T Value: %v\n", 'a', 'a')
	fmt.Printf("Type: %T Value: %v\n", 3+5i, 3+5i)

	//Named untyped constant
	const a = 123      //Default hidden type is int
	const b = "circle" //Default hidden type is string
	const c = 5.6      //Default hidden type is float64
	const d = true     //Default hidden type is bool
	const e = 'a'      //Default hidden type is rune
	const f = 3 + 5i   //Default hidden type is complex128

	fmt.Println("")
	fmt.Printf("Type: %T Value: %v\n", a, a)
	fmt.Printf("Type: %T Value: %v\n", b, b)
	fmt.Printf("Type: %T Value: %v\n", c, c)
	fmt.Printf("Type: %T Value: %v\n", d, d)
	fmt.Printf("Type: %T Value: %v\n", e, e)
	fmt.Printf("Type: %T Value: %v\n", f, f)
}
