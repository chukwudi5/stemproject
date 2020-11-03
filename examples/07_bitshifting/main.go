package main

import "fmt"

const (
	_  = iota
	kb = 1 << (iota * 10) //shift over 10
	mb = 1 << (iota * 10) // shift over 20
	gb = 1 << (iota * 10) //shift over 30

)

func main() {
	// x := 4
	// fmt.Printf("%d\t\t%b\n", x, x)

	// y := x << 1
	// fmt.Printf("%d\t\t%b\n", y, y)

	//kb := 1024
	//mb := 1024 * kb
	//gb := 1024 * mb
	//tb := 1024 * gb

	// fmt.Printf("%d\t\t%b\n", kb, kb)
	// fmt.Printf("%d\t\t%b\n", mb, mb)
	// fmt.Printf("%d\t\t%b\n", gb, gb)
	// fmt.Printf("%d\t\t%b\n", tb, tb)

	// To use IOTA to accomplish the same thing

	fmt.Printf("%d\t\t%b\n", kb, kb)
	fmt.Printf("%d\t\t%b\n", mb, mb)
	fmt.Printf("%d\t\t%b\n", gb, gb)

}
