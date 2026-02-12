package main

import "fmt"

// A struct is a user-defined type that groups together variables (fields) under one name.
func structs() {

	// zero value initialization
	// var e employee

	// while using field names order does not matter
	e := employee{
		name:  "sam",
		age:   23,
		empId: "sa23",
	}
	fmt.Println(e)

	// without using field names order matters
	e2 := employee{"fam", 54, "we232"}
	fmt.Println(e2)

	e3 := employee{
		age:  65,
		name: "jam",
	}
	fmt.Println(e3)

	// accessing the struct fields using . operator
	fmt.Println(e.name)
	e.age = 76

	// structs are value types
	emp := employee{
		name:  "tom",
		age:   34,
		empId: "t34",
	}

	emp1 := emp
	emp1.name = "jerry"
	fmt.Println(emp.name) // tom

	// pointers to struct
	emp2 := &employee{
		name:  "don",
		age:   23,
		empId: "d23",
	}
	fmt.Println(emp2.name) // still accessed using . operator; No need (*p).Name Go automatically dereferences.
}

type employee struct {
	name  string
	age   int
	empId string
}

// Ananomous struct: the below creates A struct type AND an instance of it at the same time — without giving the struct a name.
// cricketer := struct{
// 	name string
// 	age  int
// 	role string
// }{
// 	name: "Pap",
// 	age:  23,
// 	role: "batsman",
// }
// if the above syntax is used it should be in main function

// this syntax can be used outside main function
var cricketer = struct {
	name string
	age  int
	role string
}{
	name: "pap",
	age:  34,
	role: "batsman",
}

// structs and function
// pass by value
func updateRole(e employee) {
	e.name = "djmk"
}

func updateRole2(e *employee) {
	e.age = 43
}
