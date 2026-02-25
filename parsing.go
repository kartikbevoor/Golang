package main

import (
	//"encoding/binary"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

// Parsing means converting data from one format (usually string) into another usable format
//  (like int, float, bool, time, JSON, etc.).

// In Go, parsing is very common because:
// Input from users → comes as string
// File data → read as string/bytes
// API responses → often JSON (string/bytes)
func parsing() {

	// parsing string to int
	str := "123"
	num, err := strconv.Atoi(str) // Atoi = ASCII to Integer
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(num)
	} // 123

	// parsing string to float
	str2 := "2.45"
	flo, err2 := strconv.ParseFloat(str2, 64)
	if err2 != nil {
		fmt.Println(err2)
	} else {
		fmt.Println(flo)
	}

	// parsing string to boolean
	str3 := "false"
	b, err3 := strconv.ParseBool(str3) // accepts true false or 1 0
	if err3 != nil {
		fmt.Println(err3)
	} else {
		fmt.Println(b)
	}

	// parsing string to integer
	str4 := "1010"
	binary, err4 := strconv.ParseInt(str4, 2, 64) // func(s string, base int, bitsize int)
	if err4 != nil {
		fmt.Println(err4)
	} else {
		fmt.Println(binary)
	}

	// parsing with fmt package
	var hesar string
	var vayasa int
	fmt.Sscanf("Eva 21", "%s, %d", &hesar, vayasa)
}

// parsing json in go
// Unmarshal → JSON → Go struct
// Marshal → Go struct → JSON
type user2 struct {
	name string
	age  int
}

func parsingJson() {
	data := `{name: "hksn", age: 34}`

	var u user2
	err := json.Unmarshal([]byte(data), &u)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(u.name)
}

// Parsing command line arguments
func parsingCommandLineArguments() {
	num, err := strconv.Atoi(os.Args[1]) // Takes a command-line argument, Converts it from a string to an integer, Prints the integer
	if err != nil {                      // os.Args is a slice of strings containg command line arguments
		fmt.Println(err) // Atoi means ASCII to Integer.
		return
	}
	fmt.Println(num)
}
