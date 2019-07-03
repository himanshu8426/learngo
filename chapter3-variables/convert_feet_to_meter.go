package main

import "fmt"

func main() {
	//1 ft = 0.3048 m
	var ft float64
	fmt.Scanf("%f", &ft)

	output := ft * 0.3048
	fmt.Println(output)
}
