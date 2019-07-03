package main

import "fmt"


var x1 int = 111


var (
	   a =1
)

func main() {
	var x string
	x = "first"
	fmt.Println(x)
	x += "second"
	fmt.Println(x)

	var y string ="hello"
	var z string ="world"
	fmt.Println(y==z)
	fmt.Println(x1)

	//const for constant
	const x2 string = "a string"
	fmt.Println(x2)

	//multi 
/*	var (
	   a =1
	   b=2
	   c=3
	   d=4
	)
*/
//	fmt.Println(a+b+c+d)
	fmt.Println(a)
	f()

}

func f() {
	fmt.Println(x1)
}
