package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello world")
	fmt.Println()

	//values.Value()

	//variables()
	//constants()
	//forLoop()
	//isElse()
	//switches()
	//array()
	//slicess()
	//maps()
	//functions() // learn abt Function factory
	//closures()
	//recursion()
	//ranges()
	//pointers()
	//stringsAndRunes()
	//structs()
	//methods()
	//interfaces()
	//enums()
	//generics()
	//rangeOverIterators()
	//goErrors()
	//customErrors()
	//panicGo()
	//deferGo()
	//recover()
	//goRoutines()
	//channels()
	//channelBuffering()
	//channelSynchronization()
	//channelDirection()
	//selectGo()
	//timeouts()
	//nonBlockingChannelOperation()
	//closingChannel()
	//sorting()
	//stringGo()
	//formating()
	//parsing()
	//jsonGo()
	//xmlGo()
	//urlParsing()
	//cli()
	// httpClient()
	// httpServer()
	// tcpServer()

	// jwtTokens()
	//httpTcp()
	//randomNum()

	// email
	// go sendEmail("kartikbevoor2130@gmail.com", "welcome", "Henga adi huli")
	err := sendEmail("kartikbevoor2130@gmail.com", "Welcome", "Henga adi huli")
	fmt.Println(err)

	// nums := []int{1, 2, 3}

	// for _, v := range nums {
	// 	go func() {
	// 		fmt.Println(v)
	// 	}()
	// }
	// the above goroutines are not running here becos after the loops the main ends and all the goroutines are killed

}
