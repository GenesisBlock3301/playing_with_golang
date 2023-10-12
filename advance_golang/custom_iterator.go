package main

import "fmt"

type Iterator struct {
	isAggregate bool
}

func (Iter *Iterator) Nodes() <-chan *Iterator {
	nodeChannel := make(chan *Iterator)
	go func() {
		defer close(nodeChannel)
		nodeChannel <- Iter
	}()
	return nodeChannel
}

func main() {
	node := Iterator{isAggregate: true}

	nodeChannel := node.Nodes()

	for n := range nodeChannel {
		fmt.Println(n)
	}
}
