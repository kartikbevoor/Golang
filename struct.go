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

	// embedded struct usage(ananomous)
	bankAccount := bankAccountDetails{
		personEmbedded: personEmbedded{name: "fam", age: 23},
		accNo:          343,
		accType:        "savings account",
	}

	bankAccount.accType = "current account"
	bankAccount.age = 24

	// embedded struct usage(composition)
	card := cardDetails{
		personEmbedded:     personEmbedded{name: "jam", age: 45},
		bankAccountDetails: bankAccountDetails{accNo: 7448, accType: "savings"},
		cardNo:             1234,
	}
	card.bankAccountDetails.accType = "current"
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

// struct methods; methods: are nothing but functions with recievers
// this is an example of value reciever does not modify the original
func (e employee) greet() {
	fmt.Println("Hello", e.name)
}

// pointer reciever
// modifies original
func (e *employee) changeName() {
	e.name = "newName"
}

// Embedded structs: embedded structs (often called anonymous fields) are a way to include one struct inside another without giving it a field name.
// This enables composition and field/method promotion, which is Go’s preferred alternative to classical inheritance.
// ananymous embeding
type personEmbedded struct {
	name string
	age  int
}

type bankAccountDetails struct {
	personEmbedded
	accNo   int
	accType string
}

// composition embedding
type cardDetails struct {
	personEmbedded     personEmbedded
	bankAccountDetails bankAccountDetails
	cardNo             int
}
