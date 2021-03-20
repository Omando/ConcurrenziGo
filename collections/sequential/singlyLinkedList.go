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

// Last returns the last element of list l or nil if the list is empty.
func (list *SinglyLinkedList) Last() (interface{}, error) {
	if list.size == 0 {
		return nil, errors.New("list is empty")
	}
	return list.tail.value, nil
}

// Length returns count of elements in the list
func (list *SinglyLinkedList) Length() int {
	return list.size
}

// Prepend adds the given value to the beginning of the list
func (list *SinglyLinkedList) Prepend(value interface{}) {
	node := &LinkedListNode{value: value, next: nil}
	// If the list is empty, head and tail both point to the new node
	if list.size == 0 {
		list.head = node
		list.tail = node
	} else {
		node.next = list.head // new node points to the current head
		list.head = node      // new node is the new head
	}

	// Update size
	list.size++
}

