package main

import "fmt"

func main() {

	//ways to create slice
	// 1.
	var x []float64

	//2.
	y := make([]float64, 5)

	//3.
	z := make([]float64, 5, 10)

	//4.
	arr := [5]float64{1, 2, 3, 4, 5}
	a := arr[0:5]

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	fmt.Println(a)

	//slice functions

	slice1 := []int{1, 2, 3}
	slice2 := []int{5, 6}
	//	slice := append(slice1, 4,5, slice2)
	slice := append(slice2, 4, 5)
	fmt.Println(slice1, slice2, slice)

	// length of slice
	slicee := make([]int, 3, 9)
	fmt.Println(len(slicee))
}
