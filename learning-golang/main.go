package main

import (
	"fmt"
	"time"
)

func Time() {
	time := time.Now()
	fmt.Println(time)
}

//previews current
func main() {
	Time()
}
