package main

import "fmt"

// This linked list, often refeared to as a singly linked list, contains a
// head, which then points to the next node. This continues for the full list
// untill the tail, which has no next. e.g. a train with locomotive (i.e. the
// head) + box cars + caboose (i.e. tail)
type LinkedList struct {
	head *Node
}

type Node struct {
	Value any
	Next  *Node
}

// Adds an element to the end of the list.
func (l *LinkedList) Add(value any) error {
	node := Node{Value: value}

	if l.head == nil {
		l.head = &node
		return nil
	}

	current := l.head
	for current != nil && current.Next != nil {
		current = current.Next
	}

	current.Next = &node
	return nil
}

// Returns bool if list contains a specific value
func (l *LinkedList) Contains(value any) bool {
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
func (l *LinkedList) Head() (any, error) {
	if l.head == nil {
		return nil, fmt.Errorf("list is empty")
	}
	return l.head.Value, nil
}

// Returns length of the list.
func (l *LinkedList) Length() int {
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

// Create string of the list structure.
func (l *LinkedList) String() string {
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

// Removes the element with specifc value.
func (l *LinkedList) Remove(value any) error {
	if l.head == nil {
		return fmt.Errorf("list is empty")
	}

	// deal with head first, since no previous element to worry about
	if l.head.Value == value {
		l.head = l.head.Next
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

// Returns value of tail element.
func (l *LinkedList) Tail() (any, error) {
	if l.head == nil {
		return nil, fmt.Errorf("list is empty")
	}

	current := l.head
	for current != nil && current.Next != nil {
		current = current.Next
	}

	return current.Value, nil
}
