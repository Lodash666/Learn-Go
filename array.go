package main

import "fmt"

func main() {
	// flexible Array
	var x [5]int
	for i := 0; i < len(x); i++ {
		x[i] = i + 100

	}
	fmt.Println(x)
	fmt.Println(len(x))

	//Slice Is not Flexible array
	var y []int
	var z []int
	y = append(y, 1)
	fmt.Println(y)
	z = append(y, 1)
	fmt.Println(z)
	y = append(y, z[0])
	fmt.Println(y)

}
