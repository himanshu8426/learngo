package main

import "fmt"

func fib(x int) int {
	if x == 0 {
		return 0
	} else if x == 1 {
		return 1
	} else {
		return fib(x-1) + fib(x-2)
	}

}

func main() {
	a := fib(20)
	fmt.Println(a)
}
