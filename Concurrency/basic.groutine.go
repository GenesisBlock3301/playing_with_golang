package main

import (
	"fmt"
	"time"
)

func greeting(s string){
	for i:= 1; i < 6;i++{
		time.Sleep(3*time.Millisecond)
		fmt.Println(s)
	}
}

func main(){
	go greeting("Hello")
	greeting("World")
}