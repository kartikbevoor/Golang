package main

import "fmt"

func isElse() {
	if true {
		fmt.Println(true)
	}

	if 5%2 == 0 {
		fmt.Println("is divisible")
	} else {
		fmt.Println("false")
	}

	if 8%3 == 0 {
		fmt.Println("divisible by 3")
	} else if 8%2 == 0 {
		fmt.Println("divisible by 2")
	} else {
		fmt.Println("let it be")
	}

	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}

}
