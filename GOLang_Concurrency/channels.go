package main

import (
	"fmt"
	"time"
)

//https://www.freecodecamp.org/news/concurrent-programming-in-go/
/*
In concurrent programming Go provides channel that you can use for bidirectional
communication between goroutines.

Bidirectional communication means that one goroutine will send a message and the
other will read it. Send and receive are blocking. Code execution will be stopped
until the write and read are done successfully.
*/

func Greet(ch chan string) {
	fmt.Println("Greeting waiting to send greeting..")
	ch <- "Hello Sifat"
	fmt.Println("Greeting completed!")
}

func main() {
	//is declaring a channel of type string
	msg := make(chan string)
	go Greet(msg)
	time.Sleep(2 * time.Second)

	greeting := <-msg

	time.Sleep(2 * time.Second)
	fmt.Println("Greeting received")
	fmt.Println(greeting)
}
