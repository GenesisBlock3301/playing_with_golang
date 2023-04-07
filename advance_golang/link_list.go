package main

import "fmt"

type Node struct {
	Data interface{}
	Next *Node
}

type LinkList struct {
	Head *Node
}

func (ll *LinkList) InsertFront(data int) {
	if ll.Head == nil {
		ll.Head = &Node{Data: data, Next: nil}
	}
	node := &Node{Data: data, Next: ll.Head}
	ll.Head = node
}

// InsertBack Insert tail
func (ll *LinkList) InsertBack(data interface{}) {
	node := &Node{Data: data, Next: nil}
	if ll.Head == nil {
		ll.Head = node
		return
	}
	current := ll.Head
	// when get current.Next != nil, it will continue.
	for current.Next != nil {
		current = current.Next
	}
	// Then set node value next to that value
	current.Next = node
}

// Delete delete node.
func (ll *LinkList) Delete(data interface{}) bool {
	if ll.Head == nil {
		return false
	}
	if ll.Head.Data == data {
		ll.Head = ll.Head.Next
		return true
	}
	current := ll.Head
	for current.Next != nil {
		if current.Next.Data == data {
			current.Next = current.Next.Next
			return true
		}
		current = current.Next
	}
	return false
}

func (ll *LinkList) Display() {
	current := ll.Head
	for current != nil {
		fmt.Printf("%v\n", current.Data)
		current = current.Next
	}
	fmt.Println()
}

func main() {
	ll := LinkList{}
	ll.InsertFront(100)
	ll.InsertBack("1000")
	ll.Display()
}
