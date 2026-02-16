package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello world")
	fmt.Println()

	//values.Value()

	//variables()
	//constants()
	//forLoop()
	//isElse()
	//switches()
	//array()
	//slicess()
	//maps()
	//functions() // learn abt Function factory
	//closures()
	//recursion()
	//ranges()
	//pointers()
	//stringsAndRunes()
	//structs()
	//methods()

	nums := []int{1, 2, 3}

	for _, v := range nums {
		go func() {
			fmt.Println(v)
		}()
	}
	// the above goroutines are not running here becos after the loops the main ends and all the goroutines are killed

}
