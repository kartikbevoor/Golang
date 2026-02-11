package main

import "fmt"

// A constant is a value which is fixed at the complie time and cannot be changed while the prog runs

const str string = "This is a constant"

func constants() {
	const a int = 4000
	const b int = 10
	// a = a / 10   // is not permited

	fmt.Println("I am trying to change the value of a constant a, a = ", a/10)
}
