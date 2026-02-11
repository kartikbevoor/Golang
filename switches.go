package main

import "fmt"

func switches() {

	// Basic syntax
	var a int
	a = 1
	switch a {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	}

	// use of default case and comparing strings
	day := "Monday"

	switch day {
	case "saturday", "sunday": // also allows multiple values in one case
		fmt.Println("Weekend")
	default:
		fmt.Println("Weekday")
	}

	// No expression switch - replaces long if else statements
	marks := 87
	switch {
	case marks > 90:
		fmt.Println("A+")
	case marks > 80:
		fmt.Println("A")
	case marks > 70:
		fmt.Println("B")
	default:
		fmt.Println("Back bencher")
	}

	// short variable declaration in switch

	switch num := 10; num {
	case 5:
		fmt.Println("5")
	case 10:
		fmt.Println("10")
	default:
		fmt.Println("don't know")

	}

	// Fall through (manually)

	switch i := 1; i {
	case 1:
		fmt.Println("one")
		fallthrough // ignores the next case and executes directly; here one and two both are printed
	case 2:
		fmt.Println("two")
	}

	// Type Switch

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")

	checkType(10)
	checkType("hello")

	// Switch with Functions / Expressions

	n := 10
	switch {
	case n%2 == 0:
		fmt.Println("Even")
	case n%2 != 0:
		fmt.Println("odd")
	default:
		fmt.Println("pata nahi bhai")
	}
}

// Type Switch
func checkType(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Println("Integer:", v)
	case string:
		fmt.Println("String:", v)
	case bool:
		fmt.Println("Boolean:", v)
	default:
		fmt.Println("Unknown type")
	}
}
