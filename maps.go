package main

import "fmt"

// A maps stores the data as key-value pair
// map[keyType]valueType
func maps() {
	ages := map[string]int{
		"Sam": 23,
		"Fam": 45,
	}
	fmt.Println(ages)

	fruitPrice := make(map[string]int)
	fruitPrice["apple"] = 32
	fruitPrice["orange"] = 45

	scores := map[string]int{
		"Math":    90,
		"Science": 95,
	}
	fmt.Println(scores)

	var map1 map[string]int // creates nil map, you can read from it, you cannot write to it
	fmt.Println(map1)
	// Maps are reference types backed by runtime-managed data structures.
	// Maps cannot grow from nil

	// accessing values
	priceApple := fruitPrice["apple"]
	fmt.Println(priceApple)

	// if key does not exists
	fmt.Println(fruitPrice["orange"]) // zero value (0)
	// Problem: you can’t tell if key existed or not.

	// updating values
	fruitPrice["apple"] = 56

	// deleting keys
	delete(fruitPrice, "apple")

	// iterating over the map
	for key, value := range fruitPrice {
		fmt.Println(key, value)
	}

	// note: maps are unordered

	// Maps are reference types
	a := make(map[string]int)
	b := a
	b["aloo"] = 10

	fmt.Println(a)
	fmt.Println(a["aloo"])

	// Maps with struct as values
	type user struct {
		name string
		age  int
	}

	// people := make(map[int]user)
	people := map[int]user{
		1: {name: "Dam", age: 34},
		2: {name: "jam", age: 31},
	}
	fmt.Println(people)

	// modifying struct fields
	// people[1].Age = 30 // compile error
	ppl1 := people[1]
	ppl1.age = 23
	people[1] = ppl1

	// modifying using pointer values
	users := map[int]*user{
		1: &user{name: "Alice", age: 25},
	}

	users[1].age = 30 // works

	// concurrency in maps is remaining

}

// checking if key exists
func checkIfExists() {
	fruits := map[string]int{
		"apple":  23,
		"orange": 43,
	}

	if value, ok := fruits["mango"]; ok {
		fmt.Println("Exists:", value)
	} else {
		fmt.Println("Does not exist")
	}
}
