package main

import "fmt"

func main() {

  s1 := "String"
  fmt.Println(len("Hello World"))
  fmt.Println(s1)
  fmt.Println(string(s1[2]))
  fmt.Println(string("Hello World"[1]))
  fmt.Println("32132*42452 = ", 32132*42452 )
  fmt.Println("1+1", 1+1)
  fmt.Println("Hello " + "World")
  fmt.Println((true && false) || (false && true) || !(false && false))
}
