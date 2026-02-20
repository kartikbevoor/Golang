package main

import "fmt"

func channelSynchronization() {
	ch := make(chan int)
	ch <- 10      // send
	value := <-ch // receive
	fmt.Println(value)

	// ex usage
	done := make(chan bool)
	workerSyn(done)
	value1 := <-done
	if value1 {
		fmt.Println("Finished")
	}

	// Pattern to achive syn: Minimal syn using struct
	chStruct := make(chan struct{})
	go func() {
		chStruct <- struct{}{}
	}()
	// <-chStruct
	// <-chStruct // confirm this condition
	// here we are using empty struct becos 0 mem used.

	// Pattern to achive syn in Multiple goroutine syn
	chMulGo := make(chan bool)
	for i := 0; i < 3; i++ {
		go func() {
			chMulGo <- true
		}()
	}
	<-chMulGo

	// Pattern to achive syn: Using Channels as Semaphores
	// A semaphore is a synchronization mechanism used in concurrent programming
	// to control access to shared resources.
	sem := make(chan struct{}, 2)
	for i := 0; i < 5; i++ {
		sem <- struct{}{}
		go func(i int) {
			fmt.Println("Working on:", i)
			<-sem
		}(i)
	}

}

// ex
func workerSyn(done chan bool) {
	fmt.Println("working")
	done <- true
}

// Directional channels
func send(ch chan<- int) {
	ch <- 19
}

func receive(ch <-chan int) {
	fmt.Println(<-ch)
}

// chan<- → send-only
// <-chan → receive-only
