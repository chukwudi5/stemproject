package main

import "fmt"

// var x int = 42
// var y string = "James Bond"
// var z bool = true

type chuck int

var x chuck
var y chuck

func main() {
	// s := fmt.Sprintf("%v\t%v\t%v", x, y, z)
	// fmt.Println(s)

	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)
	// Using the Conversion

	y := int(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)

}
