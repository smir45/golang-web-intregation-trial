package main

import (
	"fmt"
	"time"
)

func Time() {
	time := time.Now()
	fmt.Println(time)
}

//previews current date and
func main() {
	Time()
}
