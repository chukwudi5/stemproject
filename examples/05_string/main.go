package main

import "fmt"

func main() {

	s := "Hello Me"
	fmt.Println(s)
	fmt.Printf("%T\n", s)

	t := []byte(s)
	fmt.Println(t)
	fmt.Printf("%T\n", t)

	// check for ASCI coding for the corresponding numbers that rep the alphabet
	// code points are represented in hexadecimal
	// rune is a code point
	// most code in go are encoded in UTF8

	for i := 0; i < len(s); i++ {
		fmt.Printf("%#U ", s[i])

		fmt.Println("")

		for i, v := range s {
			fmt.Println(i, v)
		}
	}
}
