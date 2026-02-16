package main

import (
	"fmt"
)

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

	// nested struct
	order1 := order{
		id: 34,
		p1: person{name: "sdfs", age: 23},
	}
	fmt.Println(order1)

	// shalow copy
	karan := relative1{
		character: "chutiya",
	}

	samar := karan
	samar.character = "not chutiya" // karan and samar are completely independent, changing samar char does affect karan

	// problem with shallow copy when reference types are used
	cric1 := cricketer2{
		name:   "Don",
		scores: []int{45, 34},
	}

	cric2 := cric1
	cric2.scores[0] = 100
	fmt.Println(cric1.scores[0]) // 100 : here the go copied the slice header

	// deep copy manually
	cric3 := cric1
	cric3.scores = make([]int, len(cric1.scores))
	copy(cric3.scores, cric1.scores)

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
// ananymous embedding
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

// struct tags: are string annotations attached to struct fields that provide metadata.
type developer struct {
	Name string `json:"name"` // The part inside backticks is the struct tag.
	Age  int    `json:"-"`
}

// struts are comparible
// p1 == p2
// Not allowed if:
// Contains slice
// Contains map
// Contains function

// nested struct
type order struct {
	id int
	p1 person
}

// Note: embedded structs does not include field name nested struct does include field name

// exported v/s unexported struct fieds
type addhar struct {
	Id   int    // exported field written using capital letters(public)
	name string // unexported field written using small lettes(private)
}

// empty struct
type empty struct{}

// Padding = extra unused bytes inserted to make the next field start at a properly aligned memory address.
// bad order
type Example struct {
	A bool   // 1byte
	B int    // 8 byte
	C string // 16 byte
}

// total mem used: 1(A) + 7(padding) + 8(b) + 16(c) = 32

// good order
type example2 struct {
	c string
	b int
	a bool
}

// total mem used: 16(c) + 8(b) + 1(a) + 7(padding) = 32

type Bad struct {
	A bool
	B int64
}

type Good struct {
	B int64
	A bool
}

// constructors: go has no constructors but use functions
func exConstructor(name string, age int) *person {
	return &person{
		name: name,
		age:  age,
	}
}

// Struct Copy vs Deep Copy
// In Go, when you assign one struct to another:
// b := a
// Go performs a shallow copy (value copy).
// That means:
// All fields are copied.
// But if a field contains a reference type, the reference is copied — not the underlying data.

// shallow copy(when copying value type)
type relative1 struct {
	character string
}

// deep copy (when copying reference type)
type cricketer2 struct {
	name   string
	scores []int
}
