package main

import "fmt"

// func zero(x *int) {
// 	*x = 0
// }

// //new function
// func one(x *int) {
// 	*x = 1
// }

// func main() {
// 	x := 1
// 	zero(&x)

// 	//y is &int
// 	y := new(int)
// 	one(y)

// 	fmt.Println(*(&x))
// 	fmt.Println(*y)

// }

func square(x *float64) {
	*x = *x * *x
}
func main() {
	x := 1.5
	square(&x)

	fmt.Println(x)
}
