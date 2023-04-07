package main

import "fmt"

type Queue struct {
	items []interface{}
}

func (q *Queue) Enqueue(item interface{}) {
	q.items = append(q.items, item)
}

func (q *Queue) Dequeue() interface{} {
	if len(q.items) == 0 {
		return nil
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item
}

func main() {
	q := Queue{}
	q.Enqueue(10)
	q.Enqueue(20)
	fmt.Println(q.items)
	q.Dequeue()
	fmt.Println(q.items)
}
