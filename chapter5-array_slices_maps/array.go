package main

import "fmt"

func main() {

	//	var x[5]int
	//	x[4] = 100
	//	fmt.Println(x)

	var x [5]float64
	x[0] = 98
	x[1] = 93
	x[2] = 77
	x[3] = 82
	x[4] = 83

	var total float64

	// another way to express for loop
	// _ tells compiler that we don't need this var
	for _, value := range x {
		total += value
	}

	fmt.Println(total / float64(len(x)))

	//create a array
	y := [5]float64{98, 93, 77, 82, 83}

	fmt.Println(y)

	z := [5]float64{
		98,
		99,
		100,
		101,
		102,
	}
	fmt.Println(z)

	//x := [6]string{"a","b","c","d","e","f"} what would x[2:5] give you?
	ex := [6]string{"a", "b", "c", "d", "e", "f"}
	fmt.Println(ex[2:5])
}
