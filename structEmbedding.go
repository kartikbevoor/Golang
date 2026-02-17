package main

import "fmt"

// Struct embedding means including one struct inside another struct without giving it a field name.
func structEmbedding() {
	// ex
	emp := bankEmployee{
		person: person{name: "jam", age: 23},
		salary: 40000,
		id:     "emoj",
	}
	fmt.Println(emp.name) // same goes for age
	// fmt.Println(e.Person.Name) // this is equal to above

	// method promotion
	fmt.Println(emp.greet()) // method is of person but due to embedding it got promoted to emp

	// function overriding
	fmt.Println(emp.greet())        // calls bankEmployee's method
	fmt.Println(emp.person.greet()) // calls person's method

	// embedding interfaces
	d := dog2{}

	animalSound := animal2{
		speaker2: d,
	}
	fmt.Println(animalSound.speak2())

	s := Service{
		Logger: ConsoleLogger{},
	}
	s.Log("Hello")
}

func (p person) greet() string {
	return "Hello" + p.name
}

// Method Overriding (Shadowing)
// Go does not support traditional overriding, but you can shadow methods.
func (be bankEmployee) greet() string {
	return "Hello" + be.name
}

// example
type bankEmployee struct {
	person // here person is a struct defined in other file
	salary int
	id     string
}

// Anonymous Field (Embedding)
type officer struct {
	person
}

// Named Field (No Embedding)
type officer2 struct { // e.P.Name // this is how u access
	p person
}

type address struct {
	state string
	city  string
}

type emp2 struct {
	person
	address
	id     string
	role   string
	salary int
} // accessing fmt.Println(emp.name)

// Pointer Embedding vs Value Embedding
// type Employee struct {  // Copies the embedded struct
//     Person
// }

// type Employee struct {  // references the existing struct
//     *Person
// }

// Struct Embedding + Interface
type Speaker interface {
	Speak() string
}

type Dog struct{}

func (d Dog) Speak() string {
	return "Woof"
}

type Pet struct {
	Dog // embedded
}

// Embedding Interfaces : A struct contains an interface as a field (without naming it).
type speaker2 interface {
	speak2() string
}

type animal2 struct {
	speaker2
}

type dog2 struct {
}

// dog2 must hold something that implements Speaker2
// dog2 can directly call Speak() because of method promotion

func (d dog2) speak2() string {
	return "Woof"
}

// another ex of Embedding Interfaces
type Logger interface {
	Log(string)
}

type ConsoleLogger struct{}

func (ConsoleLogger) Log(msg string) {
	fmt.Println(msg)
}

type Service struct {
	Logger
}

// Conflict Resolution (Field/Method Name Collision)
type A2 struct {
	Name string
}

type B2 struct {
	Name string
}

type C2 struct {
	A2
	B2
}

// c.Name   // ❌ ambiguous
// correct usage
// c.A.Name
// c.B.Name

// use capital letter for exported struct

// Deep Embedding : You can embed inside embedded structs
type x1 struct{ X int }
type y1 struct{ A }
type z struct{ B }
