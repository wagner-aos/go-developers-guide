package main

import (
	"errors"
	"fmt"
)

type Stack struct {
	Items []Flight
}

type Flight struct {
	Origin      string
	Destination string
	Price       int
}

func (s *Stack) Pop() (Flight, error) {
	if s.IsEmpty() {
		return Flight{}, errors.New("Stack is Empty")
	} else {
		lastElemIndex := len(s.Items) - 1
		flight := s.Items[lastElemIndex]
		s.Items = s.Items[:lastElemIndex]
		return flight, nil
	}
}

func (s *Stack) Push(flight Flight) {
	s.Items = append(s.Items, flight)
}

func (s *Stack) Peek() (Flight, error) {
	if s.IsEmpty() {
		return Flight{}, errors.New("Stack is Empty")
	} else {
		lastElemIndex := len(s.Items) - 1
		flight := s.Items[lastElemIndex]
		return flight, nil
	}
}

func (s *Stack) IsEmpty() bool {
	if len(s.Items) == 0 {
		return true
	}
	return false
}

func main() {
	fmt.Println("Go Stack Implementation")
}
