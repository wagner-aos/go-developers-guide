package main

import (
	"errors"
	"fmt"
)

type Queue struct {
	Items []Flight
}

type Flight struct {
	Origin      string
	Destination string
	Price       int
}

func (q *Queue) Pop() (Flight, error) {
	if q.IsEmpty() {
		return Flight{}, errors.New("Queue is empty")
	} else {
		var flight Flight
		flight, q.Items = q.Items[0], q.Items[1:]
		return flight, nil
	}
}

func (q *Queue) Push(f Flight) {
	q.Items = append(q.Items, f)
}

func (q *Queue) Peek() (Flight, error) {
	if q.IsEmpty() {
		return Flight{}, errors.New("Queue is empty")
	} else {
		flight := q.Items[0]
		return flight, nil
	}
}

func (q *Queue) IsEmpty() bool {
	if len(q.Items) == 0 {
		return true
	}
	return false
}

func main() {
	fmt.Println("Go Queue Implementation")
}
