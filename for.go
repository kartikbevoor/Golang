package main

import "fmt"

// for loop is the only loop available in golang no other loops available in this lang

func forLoop() {
	i := 3
	for i < 6 { // same as while(i < 6){}
		fmt.Println(i)
		i++
	}

	fmt.Println()

	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}

	fmt.Println()

	// break
	for i := 1; i < 5; i++ {
		if i == 3 {
			break // comes out of the for loop
		}
		fmt.Println(i)
	}

	fmt.Println()

	// continue
	for i := 1; i < 5; i++ {
		if i == 3 {
			continue // just skips current iteration
		}
		fmt.Println(i)
	}

	// looping with range
	for j := range 3 { // this is like do this for n times, here n is the range
		fmt.Println("range", j)
	}

	nums := []int{10, 20, 30}

	for i, v := range nums {
		fmt.Println(i, v)
	}

	// the above is same as this
	// for i := 0; i < len(nums); i++ {
	// 	index := i
	// 	value := nums[i]
	// 	fmt.Println(index, value)
	// }

	// to skip the index
	for _, v := range nums {
		fmt.Println(v)
	}

	// accessing each character of string
	name := "golang"
	for _, ch := range name {
		fmt.Println(ch, string(ch)) // here ch has the ascii value of each char of name, to convert them back to char string(ch) this is used
	}

	// Maps
	// note: order is not guarented in maps
	ages := map[string]int{
		"Sam": 23,
		"Fam": 34,
	}

	for _, v := range ages {
		fmt.Println(v)
	}

	// Labeled Loops
outer:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i == 2 && j == 2 {
				break outer // here if outer was't used only the inner loop would have broken.
			}
		}
	}

	// considering the below two for loops 1st fails the second is crt becos goroutines execute later
	// v := v; this line in second loop creates new variable in each iteration
	// Second loop: each goroutine gets its own variable
	for _, v := range nums {
		go func() {
			fmt.Println(v)
		}()
	} // output: 30, 30, 30

	for _, v := range nums {
		v := v
		go func() {
			fmt.Println(v)
		}()
	} // output: 10, 20, 30

	// Alternative (also correct & more explicit)
	// passing v as a paremeter
	for _, v := range nums {
		go func(val int) {
			fmt.Println(val)
		}(v)
	}

}
