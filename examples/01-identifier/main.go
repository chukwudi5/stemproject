package main

import "fmt"

func main() {
	n, err := fmt.Println("i am handsome")
	fmt.Println(n)
	fmt.Println(err)

	{
		n, _ := fmt.Println("i am beauty")
		fmt.Println(n)
	}

	{
		n, _ := fmt.Println("i am some")
		fmt.Println(n)
	}

	{
		x := 10
		fmt.Println(x)
		x = 20
		fmt.Println(x)
		y := 10 + 30
		fmt.Println(y)
		z := "bonf, james"
		fmt.Println(z)
	}
}
