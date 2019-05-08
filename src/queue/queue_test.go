package queue

import (
	"testing"
)

func Test_Queue(t *testing.T) {
	q := New()

	q.Push(5)
	if q.Len() != 1 {
		t.Error("Expected length to be 1, got", q.Len())
	}

	q.Push(4)
	if q.Len() != 2 {
		t.Error("Expected length to be 2, got", q.Len())
	}

	q.Push(3)
	if q.Len() != 3 {
		t.Error("Expected length to be 3, got", q.Len())
	}

	q.Push("a")
	if q.Len() != 4 {
		t.Error("Expected length to be 4, got", q.Len())
	}

	peeked := q.Peek()
	if peeked != 5 {
		t.Error("Expected 5, got", peeked)
	}
	if q.Len() != 4 {
		t.Error("Expected length to be 4, got", q.Len())
	}

	poppedInt := q.Pop().(int)
	if poppedInt != 5 {
		t.Error("Expected 5, got", poppedInt)
	}
	if q.Len() != 3 {
		t.Error("Expected length to be 3, got", q.Len())
	}

	poppedInt = q.Pop().(int)
	if poppedInt != 4 {
		t.Error("Expected 4, got", poppedInt)
	}
	if q.Len() != 2 {
		t.Error("Expected length to be 2, got", q.Len())
	}

	poppedInt = q.Pop().(int)
	if poppedInt != 3 {
		t.Error("Expected 3, got", poppedInt)
	}
	if q.Len() != 1 {
		t.Error("Expected length to be 1, got", q.Len())
	}

	poppedString := q.Pop().(string)
	if poppedString != "a" {
		t.Error("Expected a, got", poppedString)
	}
	if q.Len() != 0 {
		t.Error("Expected length to be 0, got", q.Len())
	}
}
