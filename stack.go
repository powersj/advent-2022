package main

type Stack struct {
	top  *stackNode
	size uint64
}

type stackNode struct {
	Value any
	Next  *stackNode
}

func (s *Stack) Pop() any {
	if s.top == nil {
		return nil
	}

	value := s.top.Value
	s.top = s.top.Next
	s.size = s.size - 1

	return value
}

func (s *Stack) Push(value any) {
	s.top = &stackNode{
		Value: value,
		Next:  s.top,
	}
	s.size = s.size + 1
}

func (s *Stack) Size() uint64 {
	return s.size
}
