package main

import (
	"fmt"
	"sync"
)

// When you run this program as usuall without mutex, it feels like it is working fine.
// But when you run this using "go run --race main.go", it will give you exit 66 which represents a Race condition

func main() {
	fmt.Println("Race Condition Learning")

	wg := &sync.WaitGroup{}

	//mut := &sync.Mutex{}
	mut := &sync.RWMutex{}

	// This is a read operation on the resource but we dont do a read lock during the declaration, not a good practice
	var score = []int{0}

	wg.Add(3)
	// immediately invoked function iffi's?
	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("One R")
		mut.Lock()
		score = append(score, 1)
		mut.Unlock()
		wg.Done()
	}(wg, mut)
	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("Two R")
		mut.Lock()
		score = append(score, 2)
		mut.Unlock()
		wg.Done()
	}(wg, mut)
	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("Three R")
		mut.Lock()
		score = append(score, 3)
		mut.Unlock()
		wg.Done()
	}(wg, mut)
	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("Print R")
		mut.RLock()
		fmt.Println(score)
		mut.RUnlock()
		wg.Done()
	}(wg, mut)

	wg.Wait()

	fmt.Println(score)
}
