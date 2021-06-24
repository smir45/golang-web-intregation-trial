package main

import (
	"fmt"
	"strings"
)

func main() {
	var x int
	var y int

	x = 1
	y = 5

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
	const a int = 16
	const b int = 16

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
	var a int = 6
	if a%5 == 0 {
		fmt.Printf("%d is the multiple of 5\n", a)
	}
}

func IfStatement2() {
	var greater int = 31
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

//trying if else if chain

func chain() {
	var BMI = 21.0
	if BMI < 18.5 {
		fmt.Println("You are underweight")
	} else if BMI >= 18.5 && BMI < 25.0 {
		fmt.Println("Your weight is normal")
	} else if BMI >= 25.0 && BMI < 30.0 {
		fmt.Println("You're overweight")
	} else {
		fmt.Println("You're obese")
	}
}

func chain2() {
	BMI := 21.0
	if BMI < 18.5 {
		fmt.Println("You are underweight")
	} else if BMI >= 18.5 && BMI < 25.0 {
		fmt.Println("Your weight is normal")
	} else if BMI >= 25.0 && BMI < 30.0 {
		fmt.Println("You're overweight")
	} else {
		fmt.Println("You're obese")
	}
}

//Arrays and Slices
func arrays() {
	var arrays [3]int = [3]int{12, 13, 14}
	fmt.Println(arrays, len(arrays))
}

func array2() {
	var array2 [3]string = [3]string{"One", "Two", "Three"}
	fmt.Println(array2, len(array2))
}

func array3() {
	name := [3]string{"One", "Two", "Three"}
	name[0] = "Zero"
	fmt.Println(name, len(name))
}

func array4() {
	age := [4]int{1, 2, 3, 4}
	age[0] = 43
	fmt.Println(age, len(age))
}

//learning libraries
func Library() {
	greeting := "hello everyone"

	fmt.Println(strings.Contains(greeting, "hello"))
	fmt.Println(strings.ReplaceAll(greeting, "hello", "learning"))
	fmt.Println(strings.ToUpper(greeting))
	fmt.Println(strings.Index(greeting, "he"))
}

//adding codes for Day 2
package main

import (
	"fmt"
	"strings"
)

func main() {
	var x int
	var y int

	x = 1
	y = 5

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
	const a int = 16
	const b int = 16

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
	var a int = 6
	if a%5 == 0 {
		fmt.Printf("%d is the multiple of 5\n", a)
	}
}

func IfStatement2() {
	var greater int = 31
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

//trying if else if chain

func chain() {
	var BMI = 21.0
	if BMI < 18.5 {
		fmt.Println("You are underweight")
	} else if BMI >= 18.5 && BMI < 25.0 {
		fmt.Println("Your weight is normal")
	} else if BMI >= 25.0 && BMI < 30.0 {
		fmt.Println("You're overweight")
	} else {
		fmt.Println("You're obese")
	}
}

func chain2() {
	BMI := 21.0
	if BMI < 18.5 {
		fmt.Println("You are underweight")
	} else if BMI >= 18.5 && BMI < 25.0 {
		fmt.Println("Your weight is normal")
	} else if BMI >= 25.0 && BMI < 30.0 {
		fmt.Println("You're overweight")
	} else {
		fmt.Println("You're obese")
	}
}

//Arrays and Slices
func arrays() {
	var arrays [3]int = [3]int{12, 13, 14}
	fmt.Println(arrays, len(arrays))
}

func array2() {
	var array2 [3]string = [3]string{"One", "Two", "Three"}
	fmt.Println(array2, len(array2))
}

func array3() {
	name := [3]string{"One", "Two", "Three"}
	name[0] = "Zero"
	fmt.Println(name, len(name))
}

func array4() {
	age := [4]int{1, 2, 3, 4}
	age[0] = 43
	fmt.Println(age, len(age))
}

//learning libraries
func Library() {
	greeting := "hello everyone"

	fmt.Println(strings.Contains(greeting, "hello"))
	fmt.Println(strings.ReplaceAll(greeting, "hello", "learning"))
	fmt.Println(strings.ToUpper(greeting))
	fmt.Println(strings.Index(greeting, "he"))
}
//Day 7 code completed




//Day 8 Start
//Revising all from Beginning

func main() {
	var x int
	var y int

	x = 1
	y = 5

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
	const a int = 16
	const b int = 16

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
	var a int = 6
	if a%5 == 0 {
		fmt.Printf("%d is the multiple of 5\n", a)
	}
}

func IfStatement2() {
	var greater int = 31
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

//trying if else if chain

func chain() {
	var BMI = 21.0
	if BMI < 18.5 {
		fmt.Println("You are underweight")
	} else if BMI >= 18.5 && BMI < 25.0 {
		fmt.Println("Your weight is normal")
	} else if BMI >= 25.0 && BMI < 30.0 {
		fmt.Println("You're overweight")
	} else {
		fmt.Println("You're obese")
	}
}

func chain2() {
	BMI := 21.0
	if BMI < 18.5 {
		fmt.Println("You are underweight")
	} else if BMI >= 18.5 && BMI < 25.0 {
		fmt.Println("Your weight is normal")
	} else if BMI >= 25.0 && BMI < 30.0 {
		fmt.Println("You're overweight")
	} else {
		fmt.Println("You're obese")
	}
}

//Arrays and Slices
func arrays() {
	var arrays [3]int = [3]int{12, 13, 14}
	fmt.Println(arrays, len(arrays))
}

func array2() {
	var array2 [3]string = [3]string{"One", "Two", "Three"}
	fmt.Println(array2, len(array2))
}

func array3() {
	name := [3]string{"One", "Two", "Three"}
	name[0] = "Zero"
	fmt.Println(name, len(name))
}

func array4() {
	age := [4]int{1, 2, 3, 4}
	age[0] = 43
	fmt.Println(age, len(age))
}

//learning libraries
func Library() {
	greeting := "hello everyone"

	fmt.Println(strings.Contains(greeting, "hello"))
	fmt.Println(strings.ReplaceAll(greeting, "hello", "learning"))
	fmt.Println(strings.ToUpper(greeting))
	fmt.Println(strings.Index(greeting, "he"))
}

//adding codes for Day 2
package main

import (
	"fmt"
	"strings"
)

func main() {
	var x int
	var y int

	x = 1
	y = 5

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
	const a int = 16
	const b int = 16

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
	var a int = 6
	if a%5 == 0 {
		fmt.Printf("%d is the multiple of 5\n", a)
	}
}

func IfStatement2() {
	var greater int = 31
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

//trying if else if chain

func chain() {
	var BMI = 21.0
	if BMI < 18.5 {
		fmt.Println("You are underweight")
	} else if BMI >= 18.5 && BMI < 25.0 {
		fmt.Println("Your weight is normal")
	} else if BMI >= 25.0 && BMI < 30.0 {
		fmt.Println("You're overweight")
	} else {
		fmt.Println("You're obese")
	}
}

func chain2() {
	BMI := 21.0
	if BMI < 18.5 {
		fmt.Println("You are underweight")
	} else if BMI >= 18.5 && BMI < 25.0 {
		fmt.Println("Your weight is normal")
	} else if BMI >= 25.0 && BMI < 30.0 {
		fmt.Println("You're overweight")
	} else {
		fmt.Println("You're obese")
	}
}

//Arrays and Slices
func arrays() {
	var arrays [3]int = [3]int{12, 13, 14}
	fmt.Println(arrays, len(arrays))
}

func array2() {
	var array2 [3]string = [3]string{"One", "Two", "Three"}
	fmt.Println(array2, len(array2))
}

func array3() {
	name := [3]string{"One", "Two", "Three"}
	name[0] = "Zero"
	fmt.Println(name, len(name))
}

func array4() {
	age := [4]int{1, 2, 3, 4}
	age[0] = 43
	fmt.Println(age, len(age))
}

//learning libraries
func Library() {
	greeting := "hello everyone"

	fmt.Println(strings.Contains(greeting, "hello"))
	fmt.Println(strings.ReplaceAll(greeting, "hello", "learning"))
	fmt.Println(strings.ToUpper(greeting))
	fmt.Println(strings.Index(greeting, "he"))
}
//Day 8 end


//Day