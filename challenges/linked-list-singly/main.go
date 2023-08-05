package main

import "fmt"

type Element struct {
	Value string
	Next  *Element
}

type SinglyLinkedList struct {
	Count int
	Head  *Element
	Tail  *Element
}

func (l *SinglyLinkedList) Append(value string) {
	// implement
	element := &Element{
		Value: value,
	}
	if l.Count == 0 {
		l.Head = element
	} else {
		l.Tail.Next = element
	}
	l.Tail = element
	l.Count += 1
}

// You will have to ensure when you add new elements
// that this method still returns the correct value
func (l *SinglyLinkedList) Size() int {
	return l.Count
}

func (l *SinglyLinkedList) Print() {
	current := l.Head
	fmt.Printf("%+v\n", current.Value)
	for current.Next != nil {
		fmt.Printf("%+v\n", current.Value)
		current = current.Next
	}
}

func main() {
	fmt.Println("Singly Linked List Challenge")

	var llist SinglyLinkedList

	values := []string{"First", "Second", "Third"}
	for _, value := range values {
		llist.Append(value)
	}
	llist.Print()
}
