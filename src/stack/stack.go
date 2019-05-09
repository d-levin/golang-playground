package stack

type (
	Stack struct {
		size uint
		top  *node
	}
	node struct {
		value interface{}
		prev  *node
	}
)

func New() *Stack {
	return &Stack{
		size: 0,
		top:  &node{},
	}
}

func (s *Stack) Len() uint {
	return s.size
}

func (s *Stack) Push(value interface{}) {
	s.size++

	n := &node{value: value, prev: s.top}

	s.top = n
}

func (s *Stack) Pop() interface{} {
	if s.size == 0 {
		return nil
	}

	s.size--
	v := s.top.value
	s.top = s.top.prev

	return v
}

func (s *Stack) Peek() interface{} {
	if s.size == 0 {
		return nil
	}
	return s.top.value
}
