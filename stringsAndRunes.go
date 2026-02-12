package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

// A string is a read-only slice of bytes
// Internally
// type string struct{
// 	Data *byte
// 	Len int
// }
// So a string contains:
// Pointer to underlying bytes
// Length

func stringsAndRunes() {
	var str string = "hello"
	fmt.Println(str)

	// str[0] = 'G' // this gives error becos strings are immutable u cannot modify string after creation
	// this is how you modify the string
	b := []byte(str)
	b[0] = 'H'
	str = string(b)
	fmt.Println(str)

	// Raw string (backticks)
	// ✔ No escape processing
	// ✔ Multi-line allowed
	// ✔ Useful for JSON, SQL, regex
	s := `mello
		  bello`
	fmt.Println(s)

	// note: strings are bytes
	// Go strings store UTF-8 encoded bytes.
	// s := "Hello"
	// [72 101 108 108 111]
	// Each letter = 1 byte. not always sometimes each char may be more than 1 byte
	// Each Chinese character is 3 bytes in UTF-8.

	// RUNE
	// A rune represents a Unicode code point.
	r := 'A'
	fmt.Printf("%T\n", r)

	// Indexing a string
	fmt.Println(str[0]) // u get the byte not a charactor

	// to print character
	fmt.Printf("%c\n", str[0])

	// iterating over strings
	for i := 0; i < len(str); i++ {
		fmt.Printf("%c ", str[i]) // prints the char of that index
	}
	fmt.Println()

	// iterating using range  // prefer this over the previous
	for _, value := range str {
		fmt.Println(value)
	}

	// converting strings to []bytes
	b1 := []byte(str)
	fmt.Println(b1)

	// converting []bytes to string
	strb1 := string(b1)
	fmt.Println(strb1)

	// String → []rune
	s3 := "你好"
	r2 := []rune(s3)
	fmt.Println(r2)
	fmt.Println(len(r2)) // after converting to rune, the length equal to number of chars

	// []rune → string
	r3 := []rune{'你', '好'}
	s4 := string(r3)
	fmt.Println(s4)

	// Count actual characters
	utf8.RuneCountInString(s)
	//or
	// sl := len([]rune(s))

	// string comparison
	fmt.Println("abc" < "dfc") // true
	// comparison is byte by byte, based on unicode values

	// String concatination
	strCon := ""
	for i := 0; i < 5; i++ {
		strCon += "a"
	}
	// the above is bad creation
	var builder strings.Builder

	for i := 0; i < 5; i++ {
		builder.WriteString("a")
	}
	result := builder.String()
	fmt.Println(result)

}

func stringBuiltInFunction() {
	strings.Contains(str, "go")      // checks if the string contains the substring
	strings.HasPrefix(str, "hi")     // checks if the string contains the suffix
	strings.HasSuffix(str, "bye")    // checks if the string contains the mention substring as suffix
	strings.ToUpper(str)             // converts the chars of the string to upper case letters
	strings.ToLower(str)             // converts the chars of the string to lower case letters
	parts := strings.Split(str, "s") // Splits a string into a slice of strings based on a separator. “Break this string wherever you see this separator.”
	fmt.Println(parts)

	slice := []string{"a", "b", "c"}
	result := strings.Join(slice, ",") // joins the slice of string into a single string with "," separating them
	fmt.Println(result)

	strings.TrimSpace(str)            // Removes leading and trailing whitespace.
	strings.ReplaceAll(str, "a", "b") // Replaces ALL occurrences of one substring with another.

}
