package main

import "fmt"

const a = 42
const b = 42
const c = "James Bond"

const (
	d         = 30
	e         = 67
	f         = 80
	h float64 = 4
	j string  = "just to know"
)

func main() {

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(h)
	fmt.Println(j)

	// IOTA : is a special predeclared identifier
	// You can do use the iota especially when you need something to authomatically increment

	const (
		k = iota
		l = iota
		n = iota + 1
	)

	const (
		m = iota
		o = iota
		p = iota + 1
	)

	fmt.Println(k)
	fmt.Println(l)
	fmt.Println(n)

	fmt.Println(m)
	fmt.Println(o)
	fmt.Println(p)

}
