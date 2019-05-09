package stack

import "testing"

func Test_Stack(t *testing.T) {
	s := New()

	if s.Len() != 0 {
		t.Error("Expected 0, got", s.Len())
	}

	s.Push(5)
	peeked := s.Peek().(int)
	if peeked != 5 {
		t.Error("Expected 5, got", peeked)
	}
	if s.Len() != 1 {
		t.Error("Expected 1, got", s.Len())
	}

	s.Push(3)
	peeked = s.Peek().(int)
	if peeked != 3 {
		t.Error("Expected 3, got", peeked)
	}
	if s.Len() != 2 {
		t.Error("Expected 2, got", s.Len())
	}

	s.Push(4)
	peeked = s.Peek().(int)
	if peeked != 4 {
		t.Error("Expected 4, got", peeked)
	}
	if s.Len() != 3 {
		t.Error("Expected 3, got", s.Len())
	}

	popped := s.Pop().(int)
	if popped != 4 {
		t.Error("Expected 4, got", popped)
	}
	if s.Len() != 2 {
		t.Error("Expected 2, got", s.Len())
	}

	popped = s.Pop().(int)
	if popped != 3 {
		t.Error("Expected 3, got", popped)
	}
	if s.Len() != 1 {
		t.Error("Expected 1, got", s.Len())
	}

	popped = s.Pop().(int)
	if popped != 5 {
		t.Error("Expected 5, got", popped)
	}
	if s.Len() != 0 {
		t.Error("Expected 0, got", s.Len())
	}

	p := s.Pop()
	if p != nil {
		t.Error("Expected nil, got", p)
	}
}
