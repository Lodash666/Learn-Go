package main

import "fmt"

func main() {
	x := 1
	switch x {
	case 1:
		fmt.Printf("hi %d \n", x)
	case 2:
		fmt.Println("Hi %d My", x)
	}
}
