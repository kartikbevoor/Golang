package main

import "fmt"

// Methods are just functions with recievers.
func methods() {
	ai := alein{
		name: "X1",
		id:   123,
	}
	ai.greetAlein()

	// pointer reciever
	ai.changeId()
	fmt.Println(ai.id)

	// reciever of type non-struct
	var mi myInt = 10
	mi.double()
	fmt.Println(mi)

	// Method expression
	alein2 := alein{
		name: "MarsX1",
		id:   111,
	}

	// usage as method expression; in this u should explicitly pass reciever as parameter
	f := alein.greetAlein
	f(alein2)

	// Method value: here no need to pass arguments
	f2 := alein2.greetAlein
	f2()

	// Embedded struct and Promoted Methods: usage
	pet1 := pet{
		name:   "bruno",
		animal: animal{animaltype: "dog"},
	}

	pet1.printType() // even though this printType() method is of animals the pet can use due to promoted methods

	// Nil receiver
	var a2 animal
	a2.sound()
}

type alein struct {
	name string
	id   int
}

// ex: basic and also an example of value reciever
func (ai alein) greetAlein() { // here alein is a type ai is a reciever
	fmt.Println("Hello", ai.name)
}

// pointer reciever
func (ai *alein) changeId() {
	ai.id = 121
}

// value reciever gets a copy pointer reciever gets a reference
// value reciever cannot modify the original pointer reciever can modify the original
// pointer reciever is more efficient for large struct

// Methods on non-struct types
// ex int type
type myInt int

func (mi myInt) double() int {
	return int(mi * 2)
}

// Note: u cannot attach methods to Types from another package

// Method sets: determine what methods are available for: T and *T
// T: Has all methods with value receivers
// *T: Has all methods with value receivers and Has all methods with pointer receivers

// Embedded struct and Promoted Methods
type animal struct {
	animaltype string
}

func (a animal) printType() {
	fmt.Println(a.animaltype)
}

// ananymous
type pet struct {
	name string
	animal
}

// Nil receivers
func (a *animal) sound() {
	if a == nil {
		fmt.Println("Nil type")
	} else {
		fmt.Println(a.animaltype, "Makes sound")
	}
}

// Go does NOT support classical inheritance.Instead: Composition, Interfaces, Method promotion
// Go does not support method overloading.
