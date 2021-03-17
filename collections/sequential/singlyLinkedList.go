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
