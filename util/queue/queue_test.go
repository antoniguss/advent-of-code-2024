package queue

import "testing"

func TestEnqueueAndDequeue(t *testing.T) {
	q := &Queue{}

	q.Enqueue("A")
	q.Enqueue("B")

	actual1 := q.Dequeue()
	actual2 := q.Dequeue()

	if actual1 != "A" {
		t.Error("expected first element removed to be A")
	}

	if actual2 != "B" {
		t.Error("expected second element removed to be B")
	}
}
