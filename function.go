package main

import "fmt"

func cal(start int, stop int) int {

	return (((start + stop) * stop) / 2)
}

func main() {
	start := 1
	stop := 100
	lasted := cal(start, stop)
	fmt.Println(lasted)
}
