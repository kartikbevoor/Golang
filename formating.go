package main

import "fmt"

func formating() {
	fmt.Print("Hello")
	fmt.Print("World")

	fmt.Println("Hello")
	fmt.Println("World")

	name := "gotilla"
	age := 12
	fmt.Printf("Name: %s, age: %d", name, age)

	x := 10
	fmt.Printf("%v\n", x) // 10
	fmt.Printf("%T\n", x) // int

	n := 15
	fmt.Printf("%d\n", n) // 15
	fmt.Printf("%b\n", n) // 1111
	fmt.Printf("%x\n", n) // f

	pi := 3.14159
	fmt.Printf("%.2f\n", pi) // 3.14

	str := "Go"
	fmt.Printf("%q\n", str) // "Go"

	// For boolean
	flag := true
	fmt.Printf("%t\n", flag) // true

	// formating string without printing
	name2 := "Kartik"
	msg := fmt.Sprintf("Hello %s", name2)
	fmt.Println(msg)

	// Formating errors
	age2 := -5
	err := fmt.Errorf("invalid age: %d", age2)
	fmt.Println(err)

	var name3 string
	fmt.Scan(&name3)
	var age3 int
	fmt.Scanf("%d", &age3)
	// Note: Always pass pointer (&variable) when scanning input.

}

// | Verb  | Meaning                  |
// | ----- | ------------------------ |
// | `%v`  | Default format           |
// | `%+v` | Struct with field names  |
// | `%#v` | Go syntax representation |
// | `%T`  | Type of value            |
// | `%%`  | Prints % symbol          |

// Integers
// | Verb | Meaning         |
// | ---- | --------------- |
// | `%d` | Decimal         |
// | `%b` | Binary          |
// | `%o` | Octal           |
// | `%x` | Hex (lowercase) |
// | `%X` | Hex (uppercase) |

// Float
// | Verb   | Meaning             |
// | ------ | ------------------- |
// | `%f`   | Decimal format      |
// | `%e`   | Scientific notation |
// | `%.2f` | 2 decimal places    |

// String
// | Verb | Meaning            |
// | ---- | ------------------ |
// | `%s` | String             |
// | `%q` | Quoted string      |
// | `%x` | Hex dump of string |

// Custom formatting
// func (t Type) String() string
type Person struct {
	Name string
}

func (p Person) String() string {
	return "Person: " + p.Name
}

// fmt.Println(p)
