package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// JSON (JavaScript Object Notation) is a lightweight data format used for APIs,
//  configuration files, and data exchange

// json example
//
//	{
//	  "name": "Kartik",
//	  "age": 21,
//	  "isStudent": true
//	}
func jsonGo() {

	// 	Encoding (Go → JSON): Encoding means converting Go data into JSON.
	p := personJson{
		name: "sam",
		age:  43,
	}

	jsonData, err := json.Marshal(p)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(jsonData))
	}

	// Decoding (JSON → Go): Decoding means converting JSON into Go struct.
	jsonStr := `{"name": "fam", "age": 23}`

	var p2 personJson

	err2 := json.Unmarshal([]byte(jsonStr), &p2)
	if err2 != nil {
		fmt.Println(err2)
	} else {
		fmt.Println(p2.name, p2.age)
	}
	// Note: You must pass a pointer (&p) to Unmarshal.
	//   	 Struct fields must be exported (Capitalized).

	// Working with json arrays
	data := `{{"name": "fam", "age": 23}, {"name": "sam", "age": 23}}`
	var people []personJson
	json.Unmarshal([]byte(data), &people)

	// Using map instead of struct
	var result map[string]interface{}
	json.Unmarshal([]byte(data), &result)
	// usage: name := result["name"].(string)

	// Decoding Large JSON: Better for large files or HTTP responses.
	file, _ := os.Open("data.json")
	decoder := json.NewDecoder(file)

	var p3 personJson
	decoder.Decode(&p3)

}

// In Go, JSON maps to structs.
type personJson struct {
	name string
	age  int
}

// JSON Tags: Usually JSON keys are lowercase, But Go struct fields are uppercase,
// We use struct tags to map them.
type personJson2 struct {
	Name     string `json:"name"`
	Age      int    `json:"age,omitempty"` // Omitempty: Removes empty values from JSON., If Age = 0 → it won’t appear in JSON.
	Password string `json:"-"`             // Ignoring fields: Password will not appear in JSON.
} // Now JSON output becomes: {"name":"Kartik","age":21}
