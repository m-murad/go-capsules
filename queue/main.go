package main

import (
	"errors"
	"fmt"
)

type queue []interface{}

func (q *queue) Enqueue(itm interface{}) {
	*q = append(*q, itm)
}

func (q *queue) Dequeue() (itm interface{}, err error) {
	defer func() {
		if rec := recover(); rec != nil {
			err = errors.New("queue under flow")
		}
	}()

	itm = (*q)[0]
	(*q)[0] = nil
	*q = (*q)[1: len(*q)-1]
	return itm, nil
}

func main() {
	var myQue = new(queue) // Create a new queue instance. Queue becomes []

	myQue.Enqueue("a") // Will add "a" to the queue. Queue becomes ["a"]
	myQue.Enqueue(1) // Will add 1 to the queue. Queue becomes ["a", 1]
	myQue.Enqueue(true) // Will add "a" to the queue. Queue becomes ["a", 1, true]

	fmt.Printf("Queue is %v\n", *myQue)

	itm, err := myQue.Dequeue()
	fmt.Printf("Dequeued item %v\n", itm) // Will return "a" and nil error. Queue becomes ["a", true]
	itm, err = myQue.Dequeue() // Will return 1 and nil error. Queue becomes [true]
	itm, err = myQue.Dequeue() // Will return true and nil error. Queue becomes []
	itm, err = myQue.Dequeue() // Will return nil and a non-nil error.
	fmt.Printf("Failed to dequeue: %v\n", err)
}
