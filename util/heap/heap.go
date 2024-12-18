package heap

import "fmt"

type Heap[T any] struct {
	comparator Comparator[T]
	size       int
	data       []T
}

type Comparator[T any] func(a *T, b *T) bool

func NewHeap[T any](comparator Comparator[T]) *Heap[T] {
	return &Heap[T]{
		comparator: comparator,
		size:       0,
		data:       make([]T, 0),
	}

}

func parentIdx(pos int) int {
	return (pos - 1) / 2
}

func leftIdx(pos int) int {
	return pos*2 + 1
}

func rightIdx(pos int) int {
	return pos*2 + 2
}

func (q *Heap[T]) isLeaf(pos int) bool {
	return leftIdx(pos) > q.size-1
}

func (q *Heap[T]) swap(a int, b int) {
	q.data[a], q.data[b] = q.data[b], q.data[a]
}

func (q *Heap[T]) IsEmpty() bool {
	return q.size == 0
}

func (q *Heap[T]) Peek() (res T, err error) {
	if q.size < 1 {
		return res, fmt.Errorf("peeking into an empty queue")
	}

	res = q.data[0]
	return res, nil
}

func (q *Heap[T]) Push(item T) error {
	q.size++

	q.data = append(q.data, item)
	cur := q.size - 1

	for q.comparator(&q.data[cur], &q.data[parentIdx(cur)]) {
		q.swap(cur, parentIdx(cur))
		cur = parentIdx(cur)
	}

	return nil

}

func (q *Heap[T]) GetSize() int {
	return q.size
}

func (q *Heap[T]) Pop() (res T, err error) {
	if q.size < 1 {
		return res, fmt.Errorf("popping from an empty queue ")
	}

	res = q.data[0]
	q.data[0] = q.data[q.size-1]
	q.data = q.data[:q.size-1]
	q.size--
	q.heapify(0)

	return res, nil
}

func (q *Heap[T]) heapify(pos int) {
	if q.isLeaf(pos) {
		return
	}

	var cur *T = &q.data[pos]
	var left *T = &q.data[leftIdx(pos)]

	if rightIdx(pos) < q.size {
		right := &q.data[rightIdx(pos)]
		if q.comparator(left, cur) || q.comparator(right, cur) {
			if q.comparator(left, right) {
				q.swap(pos, leftIdx(pos))
				q.heapify(leftIdx(pos))
			} else {
				q.swap(pos, rightIdx(pos))
				q.heapify(rightIdx(pos))
			}

		}
	} else {
		if q.comparator(left, cur) {
			q.swap(pos, leftIdx(pos))
			q.heapify(leftIdx(pos))
		}
	}

}
