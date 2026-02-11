package main

import "fmt"

func array() {
	var arr [5]int
	fmt.Println(arr) // this converts the array to string, in array all elements are initialised to zero value

	arr[4] = 123     // accessing the element by indexing
	fmt.Println(arr) // converts the arr into string and prints it

	fmt.Println(len(arr)) // gives the length of the array

	arr2 := [5]int{3, 5, 2, 5} // initialisation with values   name := [size]type{elements....}
	fmt.Println(arr2)

	// arr := [3]int{10, 20} // compile-time error

	arr3 := [...]int{1, 2, 3, 4} // go inferes the size itself and the size is also fixed
	fmt.Println(len(arr3))

	// index based initialisation

	arr4 := [5]int{
		0: 10,
		2: 20,
	}
	fmt.Println(arr4)

	// Bounds are checked at runtime; ex: index out of bound error

	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	for i, v := range arr4 {
		fmt.Println(i, v)
	}

	a := [3]int{1, 2, 3}
	b := a

	a[2] = 30
	b[0] = 10

	fmt.Println(a)
	fmt.Println(b)

	// arrays copy, slices reference

	// Note: if an array is passed to a function, and the array is modified by that function the original array is unaffected
	// pass by reference and pass by value

	// 2D array
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}

	fmt.Println(twoD)

	twoD = [2][3]int{
		{1, 2, 3},
		{1, 2, 3},
	}
	fmt.Println("2d: ", twoD)

	twoD2 := [2][3]int{
		{1, 2, 3},
		{1, 2, 3},
	}

	fmt.Println(twoD2)
}
