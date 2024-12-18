package queue

// Queue represents a FIFO queue.
type Queue struct {
	items []*interface{}
}

// NewQueue creates a new empty Queue.
func NewQueue() *Queue {
	return &Queue{
		items: make([]*interface{}, 0),
	}
}

// Enqueue adds an item to the end of the queue.
func (q *Queue) Enqueue(item interface{}) {
	q.items = append(q.items, &item)
}

// Dequeue removes and returns the item at the front of the queue.
func (q *Queue) Dequeue() interface{} {
	if len(q.items) == 0 {
		return nil
	}

	first := q.items[0]
	q.items = q.items[1:]
	return *first
}

func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0
}
