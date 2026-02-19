package main

import "fmt"

func channelDirection() {
	// Channels can be : biderectional by default, send only, recieve only.
	// Channel direction is a compile-time restriction that controls how a function can use a channel.

	// bidirectional (by default)
	ch := make(chan int)
	ch <- 10      // send
	value := <-ch // receive
	fmt.Println(value)

	// Producer Consumer
	chPC := make(chan int)

	go producer(chPC)
	consumer(chPC)

}

// Send only
func sendOnly(ch chan<- int) {
	ch <- 100
	// <- ch // this gives error
}

// recieveOnly
func recieveOnly(ch <-chan int) {
	<-ch
	// ch <- 10 // this gives error
}

// Ex: Producer Consumer
func producer(ch chan<- int) {
	for i := 0; i < 3; i++ {
		ch <- 10
	}
	close(ch)
}

func consumer(ch <-chan int) {
	for v := range ch {
		fmt.Println(v)
	}
}
