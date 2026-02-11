package main

import (
	"fmt"
	"slices"
)

func slicess() {

	var sli []string // len == 0, cap == 0, s == nil : all true
	fmt.Println(sli)

	sli2 := []int{} // len == 0, cap == 0, s != nil
	fmt.Println(sli2)

	// note: slice are reference type

	// internal structure of slice
	// type slice struct {
	// 	ptr  *T   // pointer to array
	// 	len  int
	// 	cap  int
	// }
	// ptr → points to underlying array
	// len → number of elements visible
	// cap → total space available

	var s []string // by default the length of the uninitialised slice is 0
	fmt.Println(s, "length:", len(s), "cap", cap(s))

	// using make
	s = make([]string, 3) // by using builtin keyword make, we are creating a slice of length non-zero, initialised to zero
	fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

	s3 := make([]int, 5, 10)
	fmt.Println(s3)

	var s2 []string
	fmt.Println("uninit:", s2, s == nil, len(s) == 0)

	s = make([]string, 3)
	fmt.Println("s:", s, "len:", len(s), cap(s))

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	fmt.Println("len:", len(s))

	// Append: adds elements to end of the slice
	// If len < cap → append in same array
	// If len == cap → new array allocated
	s = append(s, "d")
	s = append(s, "e", "f") // can append multiple elements
	fmt.Println("apd:", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)
	c = append(c, s...)

	// Indexing & Slicing
	// s[low : high]   // includes low, excludes high
	l := s[2:5]
	fmt.Println("sl1:", l)

	l = s[:5]
	fmt.Println("sl2:", l)

	l = s[2:]
	fmt.Println("sl3:", l)

	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	t2 := []string{"g", "h", "i"}
	if slices.Equal(t, t2) {
		fmt.Println("t == t2")
	}

	// Slices Share the Same Underlying Array
	a := []int{1, 2, 3, 4}
	b := a[1:3]
	b[0] = 100

	fmt.Println(a) // [1 100 3 4], Both a and b point to the same array.

	// copying slices
	src := []int{1, 2, 3, 4}
	dest := make([]int, len(src))

	copy(dest, src) // by this method: modifying dest won’t affect src

	// Iterating over slices
	for i := 0; i < len(s); i++ {
		fmt.Println(s[i])
	}

	for index, value := range s {
		fmt.Println(index, value)
	}

	// 2D slice
	matrix := [][]int{
		{1, 2, 3}, // Each inner slice can have different length
		{4, 5, 6},
	}
	fmt.Println(matrix)

	twoD := make([][]int, 3)
	for i := range 3 {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := range innerLen {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
