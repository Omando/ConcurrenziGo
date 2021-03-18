package sequential

import (
	"errors"
)

// LinkedListNode node for a singlye linked list
type LinkedListNode struct {
	value interface{}     // Any type. No run-time type-safety. Need generics!
	next  *LinkedListNode // Must use pointer for recursive types. Otherwise size of struct is unknown to compiler
}

// SinglyLinkedList singly linked list
type SinglyLinkedList struct {
	head *LinkedListNode
	tail *LinkedListNode
	size int
}

// Factory function
func NewSLL() *SinglyLinkedList {
	// Recall that new returns a pointer to a newly allocated zero value of the given type
	// Can also use:
	//	return &singlyLinkedList{head: nil, tail: nil, size: 0}
	//	return (&singlyLinkedList{}).init()
	return new(SinglyLinkedList).init() // same as:
}

// Receiver is a pointer which receives a copy of the address. If receiver was value,
// a cop of the value is passed and changes inside this function are not visible
// outside the function
func (list *SinglyLinkedList) init() *SinglyLinkedList {
	list.head = nil
	list.tail = nil
	list.size = 0
	return list
}

// First returns the first element of list l or nil if the list is empty.
func (list *SinglyLinkedList) First() (interface{}, error) {
	if list.size == 0 {
		return nil, errors.New("list is empty")
	}
	return list.head.value, nil
}

