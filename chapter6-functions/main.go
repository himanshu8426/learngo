package chapter6_functions

import "fmt"

func main() {

	fmt.Println("aa")
}

//find average
func average(x []float64) float64 {

	total := 0.0
	for _, v := range x {
		total += v
	}
	return total / float64(len(x))
}

//a function can return multiple value
func multiple_value() (int, int) {
	return 5, 6
}

//Variadic Functions
func add(args ...int) int {

	total := 0
	for _, v := range args {
		total += v
	}

	return total
}

//closure... It is possible to create functions inside of functions
func f1() {

	add := func(x, y int) int {
		return x + y
	}
	fmt.Println(add(1, 1))
}
