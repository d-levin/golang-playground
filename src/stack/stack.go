package stack

import (
	"errors"
)

type Stack struct {
	size uint
	// TODO use a node instead to avoid worrying about capacity
	values []interface{}
}

func (s *Stack) resize() {
	newCapacity := 0
	if cap(s.values) == 0 {
		newCapacity = 2
	} else {
		newCapacity = cap(s.values) * 2
	}

	newArray := make([]interface{}, newCapacity)

	for i := range s.values {
		newArray[i] = (s.values)[i]
	}

	s.values = newArray
}

func (s *Stack) Push(value interface{}) error {
	if int(s.size) >= cap(s.values) {
		s.resize()
	}

	s.values[s.size] = value
	s.size++

	return nil
}

func (s *Stack) Pop() (interface{}, error) {
	if s.size == 0 {
		return nil, errors.New("stack is empty")
	}

	s.size--

	return s.values[s.size], nil
}

func (s *Stack) Peek() (interface{}, error) {
	if s.size == 0 {
		return nil, errors.New("stack is empty")
	}

	return s.values[s.size-1], nil
}

func (s *Stack) Size() uint {
	return s.size
}
