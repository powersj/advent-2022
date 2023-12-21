package main

type Queue struct {
	size uint64
}

type queueNode struct {
	value any
}

func (q *Queue) Add(value any) {
}

func (q *Queue) Empty() bool {
	if q.size == 0 {
		return true
	}

	return false
}

func (q *Queue) Peek() any {
}

func (q *Queue) Remove() any {
}

func (q *Queue) Size() uint64 {
	return q.size
}

func (q *Queue) String() string {
	return ""
}
