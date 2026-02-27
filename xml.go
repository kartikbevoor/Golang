package main

import (
	"encoding/xml"
	"fmt"
)

// XML (eXtensible Markup Language) is used to store and transport structured data.

// It allows you to: Marshal (convert Go struct → XML), Unmarshal (convert XML → Go struct)
// 					 Stream large XML files, Work with attributes, nested elements, and custom tags

// <person>
//
//	<name>Kartik</name>
//	<age>21</age>
//
// </person>
func xmlGo() {

	// Unmarshalling (XML → Struct)
	// Parsing a xml
	data := []byte(`	
		<personXml>
			<name>sam</name>
			<age>23</age>
		</personXml>
	`)

	var P PersonXml
	err := xml.Unmarshal(data, &P)
	// 	Struct tags like xml:"name" map XML tags to struct fields.
	//  You must pass a pointer to Unmarshal.

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(P.Name)
		fmt.Println(P.Age)
	}

	// Marshalling (Struct → XML)
	p := PersonXml{Name: "jam", Age: 21}

	output, err := xml.Marshal(p)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(output))

	// xml attribute
	data2 := []byte(`
		<personxml id="101">
			<name>Kartik</name>
		</personxml>
	`)

	var P3 PersonXml
	err3 := xml.Unmarshal(data2, &P3)
	if err3 != nil {
		fmt.Println(err3)
	}

	// Nested XML Elements
	// <person>
	// 	<name>Kartik</name>
	// 	<address>
	// 		<city>Bangalore</city>
	// 		<state>Karnataka</state>
	// 	</address>
	// </person>

}

type PersonXml struct {
	Name string `xml:"name"`
	Age  int    `xml:"age"`
}

// Using Root Element Name: You can control root tag using XMLName.
type PersonXml2 struct {
	XMLName xml.Name `xml:"personxml"`
	Name    string   `xml:"name"`
	Age     int      `xml:"age"`
	ID      string   `xml:"id,attr"`
}

// Nested
type Address struct {
	City  string `xml:"city"`
	State string `xml:"state"`
}

type PersonXml3 struct {
	XMLName xml.Name `xml:"personxml"`
	Name    string   `xml:"name"`
	Address Address  `xml:"address"`
}

// Working With Slices (Multiple Elements)
// <users>
//     <user><name>A</name></user>
//     <user><name>B</name></user>
// </users>

type User struct {
	Name string `xml:"name"`
}

type Users struct {
	Users []User `xml:"user"`
}

// Handling Large XML
// decoder := xml.NewDecoder(file)

// for {
//     token, err := decoder.Token()
//     if err != nil {
//         break
//     }

//     switch se := token.(type) {
//     case xml.StartElement:
//         fmt.Println("Start:", se.Name.Local)
//     }
// }
