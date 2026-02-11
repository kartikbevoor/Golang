package main

import "fmt"

func variables() {
	var a int = 7
	fmt.Println(a)

	var d, c int = 1, 2
	fmt.Println(d, c)

	b := 7
	fmt.Println(b)

	var e int // values which are not initialised are zero-valued
	fmt.Println(e)

	var str string = "abc"
	fmt.Println(str)

	var bType bool = true
	fmt.Println(bType)
}
