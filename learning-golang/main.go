package main

import (
	"fmt"
	"time"
)

func Time() {
	time := time.Now()
	fmt.Println(time)
}

//previews current date and time
func main() {
	Time()
}
