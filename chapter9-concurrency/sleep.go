package main

import (
	"fmt"
	"time"
)

func sleep(seconds int)  {
	<- time.After(time.Second *time.Duration(seconds))
}

func main()  {
	fmt.Println("asdasda")
	sleep(5)
	fmt.Println("after 5 seconds")
}