package main

import (
	"fmt"
	"sync"
)

func HelloWorld(wg *sync.WaitGroup) {
	//This will reduce the internal counter value by 1.
	defer wg.Done()
	fmt.Println("Hello world")
}

func GoodBye(wg *sync.WaitGroup) {
	// reduce the counter value by 1
	defer wg.Done()
	fmt.Println("Good Bye!")
}

func main() {
	//to wait for both goroutines to complete their
	//execution before the program exits
	var wg sync.WaitGroup
	//method is used to set the initial count of the wait group to 2,
	//indicating that we are waiting for two goroutines to complete
	wg.Add(1)
	go HelloWorld(&wg)
	go GoodBye(&wg)
	//This method blocks the execution of code until the internal counter becomes 0.
	wg.Wait()
}
