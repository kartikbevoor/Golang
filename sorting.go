package main

import (
	"fmt"
	"sort"
)

func sorting() {

	// sorting integers
	nums := []int{1, 2, 3}
	sort.Ints(nums)
	fmt.Println(nums)

	// sorting strings
	str := []string{"asf", "asgg", "hdss", "sgew"}
	sort.Strings(str)
	fmt.Println(str)

	// sorting floats
	prices := []float64{10.5, 2.3, 7.8}
	sort.Float64s(prices)
	fmt.Println(prices)

	// checking if ints are sorted
	sort.IntsAreSorted(nums) // return true or false

	// sorting in decending order
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))

	// bubble sort
	nums2 := []int{4, 2, 5, 1, 0}
	res := bubbleSort(nums2)
	fmt.Println("Bubble sort")
	fmt.Println(res)

	// selection sort
	nums3 := []int{3, 1, 5, 6, 1, 0, 7}
	fmt.Println("Selection sort")
	fmt.Println(selectionSort(nums3))

	// sorting struct
	people := byAge{
		{name: "Alice", age: 30},
		{name: "Bob", age: 22},
		{name: "Charlie", age: 25},
		{name: "David", age: 28},
	}
	people.sortStruct()
	for _, v := range people {
		fmt.Println(v.name, v.age)
	}

	// sorting slice using built function
	// sort.Slice(slice, func(i, j int) bool {
	// 	return condition
	// })
	people2 := byAge{
		{name: "asfsd", age: 123},
		{name: "rger", age: 45},
		{name: "hre", age: 54},
	}
	sort.Slice(people2, func(i, j int) bool {
		return people2[i].age < people2[j].age // return the condition u want to check, if the condition is true no swap if false it swaps
	})

	// Multilevel sorting
	sort.SliceStable(people, func(i, j int) bool {
		if people[i].age == people[j].age {
			return people[i].name < people[j].name
		}
		return people[i].age < people[j].age
	})

}

// bubble sort
func bubbleSort(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[i] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
	return nums
}

// selection sort
func selectionSort(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		min := i
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[min] {
				min = j
			}
		}
		nums[i], nums[min] = nums[min], nums[i]
	}
	return nums
}

// Sorting Structs
type personSorting struct {
	name string
	age  int
}

type byAge []personSorting

func (a byAge) lenStruct() int {
	return len(a)
}

func (a byAge) sortStruct() {
	for i := 0; i < a.lenStruct(); i++ {
		for j := i + 1; j < a.lenStruct(); j++ {
			if a[j].age < a[i].age {
				a[i].age, a[j].age = a[j].age, a[i].age
			}
		}
	}
}
