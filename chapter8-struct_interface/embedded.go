package main

import "fmt"

type Person struct {
	name string
}

func (p *Person) Talk() {
	fmt.Println("Hi my name is", p.name)
}

type Android struct {
	Person
	Model string
}

func main() {

	p := Person{"Himanshu"}
	a := Android{Person: p, Model: "himanshu"}
	a.Person.Talk()
	a.Talk()
}
