package main

import "fmt"

// A pointer is a variable that stores the memory address of another variable.
// Instead of storing a value directly, it stores where the value lives in memory.
func pointers() {
	x := 10
	ptr := &x // &x → means “address of x”
	fmt.Println(ptr)
	fmt.Println(*ptr) // *p gives value at that address → 10

	// pointer declaration: basix syntax
	var ptr2 *int // default value is null
	ptr2 = &x
	fmt.Println(ptr2)
	fmt.Println(*ptr2) // *p → go to the address stored in p and get the value

	// pointer to pointer
	y := 20
	ptr3 := &y
	ptr4 := &ptr3
	fmt.Println(**ptr4) // 20 // dereferencing

	// pointers to arrays
	// func modify(arr [3]int) // this copies entire array
	// func modify(arr *[3]int) // this is better option

	// Slices already contain pointers to array, length, cap so u ususllly dont need it

}

func modifyingValueThroughPointers() {
	x := 10
	ptr := &x
	*ptr = 20

	fmt.Println("x:", x) // 20
	// check whether pointer is nil or not if nil dereferencing it (*ptr) causes panic
	if ptr != nil {
		fmt.Println("ptr:", *ptr)
	}

	// new() Function
	// Go provides built-in new().
	// Allocates memory
	// Returns pointer to zero value

	p := new(int)
	fmt.Println(*p) // 0

}

// functions with pointers: pass by reference
func passByReference(x *int) {
	*x = 100
}

// func main() {
//     a := 10
//     change(&a)
//     fmt.Println(a) // 100
// }

// Pointers with structs
type person2 struct {
	name string
	age  int
}

// without pointers: this does not modify in original struct
func updateAge(p person) {
	p.age = 50
}

// with pointers: modifies in original struct
func updateAge2(p *person) {
	p.age = 50
}

// usage
// person := Person{"Kartik", 22}
// update(&person)
// fmt.Println(person.age) // 50

// Memory & Stack vs Heap
// Compiler decides whether variable goes on stack or heap.
// If variable escapes function → heap allocation.
func create() *int {
	x := 10
	return &x
}

// the above is valid and no dangling problem in go
// in the above case we are returning a address of local variable, when the function ends it becomes a dangling pointer problem
// A dangling pointer is: A pointer that points to memory that has already been freed (deallocated).
// You still have the address… But the memory at that address is no longer valid.

// GO cannot have dangling pointers problem becos:
// Garbage collection
// No manual free()
// No pointer arithmetic

// Go intentionally restricts pointer arithmetic.
// not allowed
// p++
// p + 1

// Methods with Pointer Receiver
// A method is just a function with a special parameter called a receiver.
// A method is a function that belongs to a type.
// example
type student struct {
	name  string
	class string
}

func (s *student) changeClass() { // here (s *student) is the reciver
	s.class = "pu2"
}

// usage
// func main() {
//     p := Person{Name: "Alice"}
//     p.Greet()  // normally without methods(recievers) this would have been Greet(p)
// }

// another example

type counte struct {
	value int
}

func (c *counte) incrementCounter() {
	c.value++
}

// c.incrementCounter // this is how you call it

// Common mistakes

// Nil pointer dereference
// var p *int
// *p = 10 // panic

// Loop variable pointer mistake
// for _, v := range nums {
//     ptrs = append(ptrs, &v) // WRONG
// }
// Y? : v reused each iteration → all pointers same.
// correct
// for i := range nums {
//     ptrs = append(ptrs, &nums[i])
// }
