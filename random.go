package main

import (
	"fmt"
	// "math/big"
	"math/rand"
	"time"
)

func randomNum() {
	fmt.Println(rand.Int())

	// diffRandomNum()

	// random num between n
	n := rand.Intn(100)
	fmt.Println(n)
	// for i := 0; i < 5; i++ {
	// 	rand.Seed(time.Now().UnixNano())
	// 	n := rand.Intn(100)
	// 	fmt.Println(n)
	// }

	// Generate Random Float in Range
	min := 5.0
	max := 10.0
	num := min + rand.Float64()*(max-min) // rand.Float64(): this gives in the range of 0.0 to <1.0
	fmt.Println(num)

}

// Seeding the Random Generator: To generate different numbers each run
func diffRandomNum() {
	for i := 0; i < 3; i++ {
		rand.Seed(time.Now().UnixNano())
		fmt.Println(rand.Int())
	}

	fmt.Println("")

	// for i := 0; i < 3; i++ {
	// 	fmt.Println(rand.Int())
	// }
}

// Seed() initializes the generator
// UnixNano() ensures a different starting point

// | Function         | Description |
// | ---------------- | ----------- |
// | `rand.Int()`     | Random int  |
// | `rand.Intn(n)`   | 0 to n-1    |
// | `rand.Float64()` | 0.0 to <1.0 |
// | `rand.Int31()`   | 31-bit int  |
// | `rand.Int63()`   | 63-bit int  |

// Shuffle a Slice
func suffleSlice() {
	numbers := []int{1, 2, 3, 4, 5}

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(numbers), func(i, j int) {
		numbers[i], numbers[j] = numbers[j], numbers[i]
	})

	fmt.Println(numbers)
}

// Creating Your Own Random Source
func creatingRandomSource() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	fmt.Println(r.Intn(100))
}

// Secure Random Number(crypto)
func cryptoRandom() {
	// n, err := rand.Int(rand.Reader, big.NewInt(100))
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(n)
}
