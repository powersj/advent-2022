package main

import "fmt"

// This doubly linked list, often refeared to as a singly linked list, contains a
// head, which then points to the next node. This continues for the full list
// untill the tail, which has no next. e.g. a train with locomotive (i.e. the
// head) + box cars + caboose (i.e. tail)
type DoubleLinkedList struct {
	head *doubleListNode
}

type doubleListNode struct {
	Value any
	Prev  *doubleListNode
	Next  *doubleListNode
}

// Adds an element to the end of the list.
func (l *DoubleLinkedList) Add(value any) error {
	node := doubleListNode{Value: value}

	if l.head == nil {
		l.head = &node
		return nil
	}

	current := l.head
	for current != nil && current.Next != nil {
		current = current.Next
	}

	node.Prev = current.Next
	current.Next = &node
	return nil
}

// Returns bool if list contains a specific value
func (l *DoubleLinkedList) Contains(value any) bool {
	current := l.head
	for {
		if current == nil {
			return false
		}
		if current.Value == value {
			return true
		}
		current = current.Next
	}
}

// Returns value of head node.
func (l *DoubleLinkedList) Head() (any, error) {
	if l.head == nil {
		return nil, fmt.Errorf("list is empty")
	}

	return l.head.Value, nil
}

// Returns length of the list.
func (l *DoubleLinkedList) Length() int {
	length := 0

	current := l.head
	for {
		if current == nil {
			break
		}
		length += 1
		current = current.Next
	}

	return length
}

// Removes the element with specifc value.
func (l *DoubleLinkedList) Remove(value any) error {
	if l.head == nil {
		return fmt.Errorf("list is empty")
	}

	// deal with head first, since no previous element to worry about
	if l.head.Value == value {
		// single element list
		if l.head.Next == nil {
			l.head = nil
			return nil
		}

		l.head = l.head.Next
		l.head.Prev = nil
		return nil
	}

	previous := l.head
	current := l.head.Next
	for {
		if current == nil {
			return fmt.Errorf("unable to find value to delete")
		}
		if current.Value == value {
			previous.Next = current.Next
			return nil
		}
		previous = current
		current = current.Next
	}
}

// Create string of the list structure.
func (l *DoubleLinkedList) String() string {
	if l.head == nil {
		return ""
	}

	train := fmt.Sprintf("%v", l.head.Value)
	current := l.head.Next
	for current != nil {
		train = fmt.Sprintf("%s -> %v", train, current.Value)
		current = current.Next
	}
	return train
}

// Returns value of tail element.
func (l *DoubleLinkedList) Tail() (any, error) {
	if l.head == nil {
		return nil, fmt.Errorf("list is empty")
	}

	current := l.head
	for current != nil && current.Next != nil {
		current = current.Next
	}

	return current.Value, nil
}
