package queue

type (
	Queue struct {
		length uint
		front  *node
		back   *node
	}
	node struct {
		next  *node
		value interface{}
	}
)

func New() *Queue {
	n := &node{next: nil}
	return &Queue{
		length: 0,
		front:  n,
		back:   n,
	}
}

func (q *Queue) Len() uint {
	return q.length
}

func (q *Queue) Push(value interface{}) {
	q.back.value = value
	q.back.next = &node{}
	q.back = q.back.next
	q.length++
}

func (q *Queue) Peek() interface{} {
	if q.length == 0 {
		return nil
	}

	return q.front.value
}

func (q *Queue) Pop() interface{} {
	if q.length == 0 {
		return nil
	}

	q.length--

	value := q.front.value
	q.front = q.front.next

	return value
}
