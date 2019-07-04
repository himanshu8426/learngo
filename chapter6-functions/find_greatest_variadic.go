package main

import "fmt"

func greatest(args ...int) int {
	max := 0
	for _, v := range args {
		if max < v {
			max = v
		}
	}
	return max
}

func main() {
	a := greatest(123, 123, 21, 3, 213, 213213123421, 3, 21, 321, 3, 21, 3, 21, 321, 321, 3, 21, 123213213, 12, 1)
	fmt.Println(a)
}
