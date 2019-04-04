package main

import "fmt"

func checkIris(petalwidth float32, petallength float32) string {
	if petalwidth > 0.6 {
		if petalwidth <= 1.7 {
			if petallength > 4.9 {
				if petalwidth > 1.5 {
					return "Iris-versicolor(3.0/1.0)"
				} else {
					return "Iris-Virginica(3.0)"
				}
			} else {
				return "Iris-versicolor(48.0/1.0)"
			}
		} else {
			return "Iris-virginica(46.0/1.0)"
		}
	} else {
		return "Iris-setosa(50.0)"
	}
}

func main() {
	//fmt.Println(0.7 <= 1.7)
	fmt.Println(checkIris(1.6, 5.6))
}
