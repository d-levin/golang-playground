package stack

import (
	"testing"
)

func Test_Push_GivenStackIsFull_ShouldResizeByDoubling(t *testing.T) {
	s := Stack{}

	if cap(s.values) != 0 {
		t.Error("Expected capacity 0, got", cap(s.values))
	}

	_ = s.Push(5)

	if cap(s.values) != 2 {
		t.Error("Expected capacity 2, got", cap(s.values))
	}

	println(len(s.values))

	_ = s.Push(5)

	println(len(s.values))

	_ = s.Push(5)

	println(len(s.values))
}

func Test_Push_GivenStackIsNotFull_ShouldPushElement(t *testing.T) {
	s := Stack{}

	err := s.Push(5)
	if err != nil {
		t.Error("Expected nil, got", err)
	}

	err = s.Push(2)
	if err != nil {
		t.Error("Expected nil, got", err)
	}
}

func Test_Push_GivenStackIsNotFull_ShouldIncrementSize(t *testing.T) {
	s := Stack{}

	if s.size != 0 {
		t.Error("Expected initial size to be 0, got", s.size)
	}

	_ = s.Push(5)

	if s.size != 1 {
		t.Error("Expected size 1, got", s.size)
	}
}

func Test_Pop_GivenStackIsEmpty_ShouldReturnError(t *testing.T) {
	s := Stack{}

	_, err := s.Pop()

	if err == nil {
		t.Error("Expected error, got", err)
	}
}

func Test_Pop_GivenStackIsNotEmpty_ShouldReturnElement(t *testing.T) {
	s := Stack{}

	_ = s.Push(5)

	value, _ := s.Pop()

	if value != 5 {
		t.Error("Expected 5, got", value)
	}
}

func Test_Peek_GivenStackIsEmpty_ShouldReturnError(t *testing.T) {
	s := Stack{}

	_, err := s.Peek()

	if err == nil {
		t.Error("Expected error, got", err)
	}
}
