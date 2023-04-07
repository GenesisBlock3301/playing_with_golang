package main

import "fmt"

// Stack here use []interface, cuz interface support any kind of data like integer, float and string
type Stack struct {
	items []interface{}
}

func (s *Stack) Push(item interface{}) {
	s.items = append(s.items, item)
}

func (s *Stack) Pop() interface{} {
	if len(s.items) == 0 {
		return 0
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item
}

func (s *Stack) Size() int {
	return len(s.items)
}

func (s *Stack) isEmpty() bool {
	return s.Size() == 0
}

func main() {
	s := &Stack{}
	s.Push(10)
	s.Push(20)
	fmt.Println(s.items)
	s.Pop()
	fmt.Println(s.items)
}
