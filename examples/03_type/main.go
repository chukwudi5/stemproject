package main

// variable is declared to hold a value of a certain type

import "fmt"

var t = 43

// var h string // the zero value is h, it will print out h if you want it to

func main() {

	j := "i love me"

	fmt.Println(t)

	fmt.Printf("%T\n", j)
	fmt.Printf("%T\n", t)
}
