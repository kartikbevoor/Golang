package main

import (
	"fmt"
	"strconv"
	"strings"
)

// In Go, strings are immutable (cannot be changed after creation).
func stringGo() {
	str := "oh hooo"
	fmt.Println(len(str))

	// 	Strings are UTF-8 encoded
	//  Indexing gives bytes, not characters
	fmt.Println(str[1]) // gives ascii value of h
	fmt.Println("for space", str[2])

	// Contains()
	fmt.Println(strings.Contains(str, "hooo")) // returns boolean value

	// HasPrefix() and HasSuffix()
	fmt.Println(strings.HasPrefix(str, "oh"))
	fmt.Println(strings.HasSuffix(str, "hooo"))

	// Index() and LastIndex()
	fmt.Println(strings.Index(str, "hooo"))
	fmt.Println(strings.LastIndex("Go Go Go", "Go"))

	// ToUpper() and ToLower()
	fmt.Println(strings.ToLower(str))
	fmt.Println(strings.ToUpper(str))

	// Trim() Functions: removes space or specific charecters
	strings.Trim("  hello  ", " ")       // "hello"
	strings.TrimSpace("  hello  ")       // "hello"
	strings.TrimPrefix("GoLang", "Go")   // "Lang"
	strings.TrimSuffix("GoLang", "Lang") // "Go"

	// Replace()
	strings.Replace("Go Go Go", "Go", "Hi", 1) // Replaces only 1 occurrence
	strings.ReplaceAll("Go Go Go", "Go", "Hi") // Replaces all

	// Split(): Splits string into slice.
	strings.Split("a,b,c", ",") // []string{"a", "b", "c"}

	// Join(): Joins slice into string.
	arr := []string{"Go", "is", "fun"}
	strings.Join(arr, " ") // "Go is fun"

	// Repeat()
	strings.Repeat("Go ", 3) // "Go Go Go "

	// Compare()
	strings.Compare("a", "b")
	// -1 if a < b
	//  0 if equal
	//  1 if a > b

	// Convert String ↔ Numbers
	n, _ := strconv.Atoi("123") // string → int
	s := strconv.Itoa(123)      // int → string
	fmt.Println(n)
	fmt.Println(s)

	// Iterating Over String
	for i, ch := range "Hello" {
		fmt.Println(i, string(ch))
	}

	// String Concatenation
	// s2 := "Go" + "Lang"
	// s := fmt.Sprintf("%s %d", "Go", 2025)

	// for heavy concatination use string builder
	var b strings.Builder
	b.WriteString("Go")
	b.WriteString("Lang")
	fmt.Println(b.String())

	

}
