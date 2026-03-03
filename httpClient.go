package main


// An HTTP Client is used to:, Send requests to servers, Receive responses, Call APIs

Communicate between microservices
func httpClient() {
	
}

// Default HTTP Client: sends request and waits for the response
func defaultHttpClient()  {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts/1") // .Get() creates request 
	// and send to resp, err := http.Get("https://jsonplaceholder.typicode.com/posts/1")
	// Waits for the server to respond, Returns the response (resp) and any error
	// resp.Body: Contains the received response body from the server.
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body) // Reads the response data that was received
	fmt.Println(string(body))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()	// Always close response body

	body, _ := io.ReadAll(resp.Body)
	fmt.Println(string(body))
}

// HTTP Response Structure
// The response of type: *http.Response
// resp.Status        // "200 OK"
// resp.StatusCode    // 200
// resp.Header        // map[string][]string
// resp.Body          // io.ReadCloser
// fmt.Println("Status Code:", resp.StatusCode)
// fmt.Println("Headers:", resp.Header)

// POST Request (with JSON)
func postRequest()  {
	jsonData := []byte(`{"name":"Kartik","balance":1000}`)

	resp, err := http.Post(
		"https://api.example.com/accounts", // This is the endpoint where the request is sent.
		"application/json",					// Content-Type: sets the http header
		bytes.NewBuffer(jsonData),			// bytes.NewBuffer() converts it into an io.Reader
	)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
}

// Using http.NewRequest
func newRequest()  {
	req, err := http.NewRequest("PUT", "https://api.example.com/accounts/1", bytes.NewBuffer(jsonData))
	if err != nil {
		panic(err)
	}
	// http.NewRequest: signature
	// func NewRequest(method, url string, body io.Reader) (*http.Request, error)

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer token123")

	client := &http.Client{}		// Creating an HTTP Client
	// http.Client: Sends requests, Manages connections, Handles redirects, Can have timeouts
	resp, err := client.Do(req)		// Sending the Request: This actually sends the HTTP request over the network.
	// What happens internally: DNS lookup, TCP connection, TLS handshake (because HTTPS)
	// sends(Request line, header, body), Waits for server response

	// Setting Timeout  
	// client := &http.Client{
	// 	Timeout: 5 * time.Second,
	// }
	// resp, err := client.Get("https://api.example.com")
}

// | Line              | What It Does                         |
// | ----------------- | ------------------------------------ |
// | `NewRequest`      | Builds HTTP request                  |
// | `bytes.NewBuffer` | Converts JSON bytes into `io.Reader` |
// | `Header.Set`      | Adds metadata                        |
// | `http.Client{}`   | Creates HTTP client                  |
// | `client.Do`       | Sends request                        |
// | `resp`            | Holds server response                |

// Handling JSON Response
type Account struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Balance int    `json:"balance"`
}

func handlingJsonRespose()  {
	var acc Account

	err = json.NewDecoder(resp.Body).Decode(&acc)
	if err != nil {
		panic(err)
	}

	fmt.Println(acc.Name)
}

// Making a Banking Example

type DepositRequest struct {
	AccountID int `json:"account_id"`
	Amount    int `json:"amount"`
}

func depositRequest()  {
	deposit := DepositRequest{  // variable of type depositRequest
		AccountID: 1,
		Amount:    500,
	}

	jsonData, _ := json.Marshal(deposit) // converting to json

	req, _ := http.NewRequest("POST", "http://localhost:8080/deposit", bytes.NewBuffer(jsonData)) // creating request
	req.Header.Set("Content-Type", "application/json") // setting header

	client := &http.Client{Timeout: 3 * time.Second}   // timeout for client
	resp, err := client.Do(req)						   // sending request and waiting for response
	if err != nil {									   // Error handling
		panic(err)
	}
	defer resp.Body.Close()							   // closing response

	fmt.Println("Deposit Status:", resp.StatusCode)
}

// | Method | Purpose         |
// | ------ | --------------- |
// | GET    | Fetch data      |
// | POST   | Create resource |
// | PUT    | Update resource |
// | DELETE | Delete resource |
// | PATCH  | Partial update  |

