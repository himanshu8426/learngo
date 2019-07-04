package main

import "fmt"

func f1(x int) (int, bool) {

	if x%2 == 0 {
		return x / 2, true
	} else {
		return x / 2, false
	}

}

func main() {
	a, b := f1(5)
	c, d := f1(6)
	fmt.Println(a, b)
	fmt.Println(c, d)
}
