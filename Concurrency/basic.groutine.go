package main

import (
	"fmt"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

// compute
func compute(value int) {
	for i := 0; i < value; i++ {
		time.Sleep(time.Millisecond)
		fmt.Println(i)
	}
	wg.Done()
}

func compute2(cha chan int, value int) {
	result := 0
	for i := 0; i < value; i++ {
		result += i
	}
	cha <- result
	close(cha)
}

func main() {
	//t1 := time.Now()
	//fmt.Println("Goroutine tutorial")
	////wg.Add(2)
	//compute(10)
	//compute(10)
	//////time.Sleep(time.Second * 1000)
	////wg.Wait()
	//timeTaken := time.Since(t1).Microseconds()
	//fmt.Println("TimeTaken", timeTaken)
	//
	t2 := time.Now()
	c1 := make(chan int)
	c2 := make(chan int)

	//go compute2(c1, 10)
	//go compute2(c2, 10)
	go compute(10)
	go compute(10)
	wg.Wait()
	timeTaken2 := time.Since(t2).Microseconds()
	r1 := <-c1
	r2 := <-c2
	fmt.Println("TimeTaken", timeTaken2, r1+r2)

}
