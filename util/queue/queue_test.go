package queue

import "testing"

func TestEnqueueAndDequeue(t *testing.T) {
	q := &Queue{}

	q.enqueue("A")
	q.enqueue("B")

	actual1 := q.dequeue()
	actual2 := q.dequeue()

	if actual1 != "A" {
		t.Error("expected first element removed to be A")
	}

	if actual2 != "B" {
		t.Error("expected second element removed to be B")
	}
}
