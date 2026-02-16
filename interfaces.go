package main

import (
	"fmt"
)

// An interface is a type that defines a set of method signatures.
// Interfaces are reference types.
func interfaces() {
	// basic usage
	var s speaker
	s = dog{
		sound: "wooff",
	}
	// internally
	// type = Dog
	// value = Dog{}

	s.speak()

	// polymorphism
	animals := []speaker{dog{sound: "woff"}, Cat{sound: "meow"}}
	for _, v := range animals {
		fmt.Println(v.speak())
	}

	// An interface is nil only if BOTH are nil:
	var s2 speaker = nil
	fmt.Println(s2 == nil) // true
	// s2.speak() // causes panic as it is a zero value interface

	var d *dog = nil
	var s3 speaker = d

	fmt.Println(s3 == nil) // false: becos internally, type is set to *dog, so interface is not nil

	// type assertion
	var i interface{} = "hello"
	check := i.(string) // check is assigned to value which is of type string in i
	fmt.Println(check)
	//check2 := i.(int) // panic: becos i does not have a type int; panic: interface conversion: int is not string

	// Safe Type Assertion
	check2, ok := i.(string)
	if ok {
		fmt.Println(check2)
	} else {
		fmt.Println("does not contain of type int")
	}

	var j speaker = dog{}
	check3, ok := j.(dog)
	if ok {
		fmt.Println(check3.sound)
	} else {
		fmt.Println("Not a dog")
	}

	// type switch usage
	// note: cannot use type switch outside type switch
	ts := []speaker{dog{}, Cat{}}
	// check4 := ts[1].(Cat) // ✅ valid
	for _, v := range ts {
		describe(v)
		// check4 := v.(dog)
	}

	// composition interface usage
	var rw readerWriter = &document{}

	fmt.Println(rw.read())
	rw.write("hello")

	// comparing interfaces
	var a interface{} = 10
	var b interface{} = 20
	fmt.Println(a == b) // false

	// var a2 interface{} = []int{1,2}
	// var b2 interface{} = []int{1,2}
	// fmt.Println(a2 == b2) // panic (slice not comparable)

}

// ex
type speaker interface {
	speak() string
}

type dog struct {
	sound string
}

func (d dog) speak() string {
	return d.sound
}

// Interface as a function parameter
func makeSpeak(s speaker) {
	fmt.Println(s.speak())
}

// polymorphism
type Cat struct {
	sound string
}

func (c Cat) speak() string {
	return "Meow"
}

// Empty interface
type x interface{} // Accept ANY type.
type y any         // this any keyword can also be used for empty interfaces

func PrintValue(v any) {
	fmt.Println(v)
}

// type switch
func describe(i speaker) {
	switch v := i.(type) {
	case dog:
		fmt.Println("Dog says:", v.sound)
	case Cat:
		fmt.Println("Cat says:", v.sound)
	default:
		fmt.Println("Unknown type")
	}
}

// Interface Composition: Interface composition means embedding one or more interfaces inside another interface.
type reader interface {
	read() string
}

type writer interface {
	write(string)
}

// composite
type readerWriter interface {
	reader
	writer
}

type document struct {
	content string
}

func (d *document) read() string {
	return d.content
}

// Note: if method has pointer reciever, only &document{} works.(document{} gives error)

func (d *document) write(str string) {
	d.content = str
}

// No problem in Go: If two embedded interfaces define the same method signature
type A interface {
	Speak()
}

type B interface {
	Speak()
}

type C interface {
	A
	B
}
