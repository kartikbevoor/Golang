package main

import (
	"fmt"
	"net/url"
)

// URL parsing means breaking a URL string into its meaningful components like protocol,
// host, path, query parameters, etc.

// https://www.example.com:8080/path/to/page?name=kartik&age=22#section1
// scheme://host:port/path?query#fragment

// https://www.example.com:8080/path/to/page?name=kartik&age=22#section1
// │       │              │      │               │                    │
// │       │              │      │               │                    └── Fragment
// │       │              │      │               └── Query parameters
// │       │              │      └── Path
// │       │              └── Port
// │       └── Host
// └── Scheme (Protocol)

// | Component | Meaning                     |
// | --------- | --------------------------- |
// | Scheme    | Protocol (http, https, ftp) |
// | Host      | Domain name or IP           |
// | Port      | Network port                |
// | Path      | Resource location           |
// | Query     | Key-value parameters        |
// | Fragment  | Section inside page         |

func urlParsing() {
	rawURL := "https://www.example.com:8080/path/to/page?name=kartik&age=22#section1"

	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Scheme:", parsedURL.Scheme)
	fmt.Println("Host:", parsedURL.Host)
	fmt.Println("Path:", parsedURL.Path)
	fmt.Println("Raw Query:", parsedURL.RawQuery)
	fmt.Println("Fragment:", parsedURL.Fragment)

	// u.Hostname() → returns host without port
	// u.Port() → returns port only
	fmt.Println("Hostname:", parsedURL.Hostname())
	fmt.Println("Port:", parsedURL.Port())

	// Modifying a URL
	parsedURL.Path = "/newpath"
	parsedURL.RawQuery = "id=100"

	fmt.Println(parsedURL.String()) // https://www.example.com:8080/newpath?id=100#section1

	// Encoding Query Parameters
	params := url.Values{}
	params.Add("name", "Kartik")
	params.Add("course", "Go")

	encoded := params.Encode()
	fmt.Println(encoded) // course=Go&name=Kartik

	parsedURL.RawQuery = params.Encode()

	// Parsing Relative URLs
	base, _ := url.Parse("https://example.com/docs/")
	relative, _ := url.Parse("intro.html")

	resolved := base.ResolveReference(relative)
	fmt.Println(resolved.String()) // https://example.com/docs/intro.html

}

// Scheme: https
// Host: www.example.com:8080
// Path: /path/to/page
// Raw Query: name=kartik&age=22
// Fragment: section1

// Accessing Query Parameters
func AccessingQueryParameter(parsedURL *url.URL) {
	queryParams := parsedURL.Query()

	fmt.Println("Name:", queryParams.Get("name"))
	fmt.Println("Age:", queryParams.Get("age"))
}

// URL struct imp fields
type URL struct {
	Scheme string
	Opaque string
	// User       *Userinfo
	Host       string
	Path       string
	RawPath    string
	ForceQuery bool
	RawQuery   string
	Fragment   string
}
