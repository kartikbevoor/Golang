package main

import "fmt"

// Enum (enumeration): A type that consists of a set of named constant values.
func enums() {
	// ex
	// var s status = Pending
	// x := 2
	// fmt.Println(s, x)
	// s = x // this gives compile time error

	var s2 status = Apporved
	fmt.Println(s2)

	// string representation
	var s3 status = Pending
	fmt.Println(s3.stringRepresentation())

	// bit flags
	var p permission
	p = read | write // here it is doing bitwise or operation. so the ans for bitwise or is 3(011)

	if p != 0 {
		fmt.Printf("read")
	} // generally used with switches

}

// Enums in other lang
// enum Status {
//     Pending,
//     Approved,
//     Rejected
// }

// Enums in go
type status int

// Define constants using iota
const (
	Pending status = iota
	Apporved
	Rejected
)

// here iota starts with 0, increases automatically, works only inside const block
// so, values, 0, 1, 2 (above const values)
// iota is a special identifier in go, Resets to 0 inside each const block, increments automatically for each line

const (
	a = iota // 1
	b        // 2
	c        // 3
)

// the above is equivalent to
// A = 0
// B = 1
// C = 2

// Adding String Representation
func (s status) stringRepresentation() string {
	switch s {
	case Pending:
		return "Pending"
	case Apporved:
		return "Apporved"
	case Rejected:
		return "Rejected"
	default: // before writing this default case i was getting an error that missing return statement
		return "unknown"
	}
}

// skiping values in Enums
const (
	a1 = iota
	_
	c1
)

// Starting Enum from 1 Instead of 0
const (
	A1 status = iota + 1
	B1
	C1
)

// Enums with custom value
const (
	Low    status = 10
	Medium status = 20
	High   status = 30
)

// Enums with custom value using iota
const (
	Low2 status = iota * 10
	Medium2
	High2
)

// Enum with Bit Flags
type permission int

const (
	read    permission = 1 << iota // 1 (001)
	write                          // 2 (010)
	execute                        // 4 (100)
)
