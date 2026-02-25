package main

import (
	"fmt"
	"os"
	"sync"
)

// defer is a keyword in Go that delays the execution of a function
// until the surrounding function returns.

// The deferred function runs: After the surrounding function finishes
// Even if the function returns early, Even if a panic occurs (before crash, unless recovered)

// How defer Works Internally: Evaluates the arguments immediately
// Pushes the function call onto a stack, Executes it later (LIFO order)
func deferGo() {

	// basic syntax
	// defer functionCall()

	defer fmt.Println("world")
	fmt.Println("Hello") // Hello world

	// Multiple defer (LIFO Order)
	defer fmt.Println("third")
	defer fmt.Println("second")
	defer fmt.Println("first") // first second third

}

// Arguments Are Evaluated Immediately
func argumentsEvaluation() {
	x := 10
	defer fmt.Println(x)
	x = 20
} // 10

// Deferred Anonymous Function
func deferredAnanymous() {
	x := 10

	defer func() {
		fmt.Println(x)
	}()
	x = 20

} // 20
// The function is deferred, But the variable x is evaluated when the function runs
// Not when defer is declared

func deferredAnanymous2() {
	x := 10
	defer func(x int) {
		fmt.Println(x)
	}(x)
	x = 20
} // 10 : In the above function, The argument x is evaluated immediately.

// ex Closing Files
func closingFiles() {
	file, err := os.Open("deta.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
}

// ex using with mutex
func mutex() {
	var mu sync.Mutex
	mu.Lock()
	defer mu.Unlock()
}

// defer and Named Return Values
func namedReturn() (res int) {
	defer func() {
		res += 10
	}()
	return 5
} // 15
// Explanation: return 5 assigns result = 5, Deferred function runs → adds 10, Final return = 15

// defer in Loops
func deferLoops() {
	for i := 0; i < 3; i++ {
		defer fmt.Println(i)
	}
} // 2, 1, 0
