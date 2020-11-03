package main

import "fmt"

var y = 43 // to declare avariable outside the func main you var
// Zero values = false for booleans, 0 for numeric types, "" for strings, and nil for pointers, functions, interfaces, slices, channels, and maps.
func main() {

	x := 100 // use short declaration operation withinin
	fmt.Println(x)
	fmt.Println(y)

	foo()

}

func foo() {
	fmt.Println("again:", y)
}
