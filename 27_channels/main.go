package main

import (
	"fmt"
	"sync"
)

// Channels are used to enable communication between Go Routines
// They provide comminocation even before closing the thread

func main() {
	fmt.Println("Learning channels")

	// second paramter value indicates buffered channel. Which means it will only allow n number of value through the channel
	myCh := make(chan int, 2)

	wg := &sync.WaitGroup{}

	// this is how we add values to channels
	// myChannels <- 5

	// fmt.Println(<-myChannels)

	wg.Add(2)

	// By add arrow before channel, it means this is RECIEVE ONLY func
	go func(ch <-chan int, wg *sync.WaitGroup) {

		value, isChannelOpen := <-myCh

		fmt.Println(isChannelOpen)
		fmt.Println(value)
		wg.Done()
	}(myCh, wg)

	// By add arrow after channel, it means this is SEND ONLY func
	go func(ch chan<- int, wg *sync.WaitGroup) {
		myCh <- 0
		close(myCh)
		wg.Done()
	}(myCh, wg)

	wg.Wait()

}
