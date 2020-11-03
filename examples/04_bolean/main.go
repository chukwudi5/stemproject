package main

import "fmt"

var x bool // this will produce false, because by default bool assigned to a variable is false

func main() {
	fmt.Println(x)
	x = true // this you assigning value to ensure is true

	fmt.Println(x)
}
