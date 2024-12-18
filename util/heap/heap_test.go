package heap

import (
	"testing"
)

func TestHeap_PushAndPeek(t *testing.T) {
	comparator := func(a *int, b *int) bool {
		return *a < *b // Min-heap comparator
	}

	h := NewHeap(comparator)

	// Test pushing elements
	h.Push(5)
	h.Push(3)
	h.Push(8)

	// Test peek
	peeked, err := h.Peek()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if peeked != 3 {
		t.Errorf("expected peeked value to be 3, got %v", peeked)
	}
}

func TestHeap_Pop(t *testing.T) {
	comparator := func(a *int, b *int) bool {
		return *a < *b // Min-heap comparator
	}

	h := NewHeap(comparator)

	h.Push(5)
	h.Push(3)
	h.Push(8)

	// Test popping elements
	popped, err := h.Pop()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if popped != 3 {
		t.Errorf("expected popped value to be 3, got %v", popped)
	}

	// Ensure the next pop returns the next minimum
	popped, err = h.Pop()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if popped != 5 {
		t.Errorf("expected popped value to be 5, got %v", popped)
	}

	// Ensure the last pop returns the last element
	popped, err = h.Pop()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if popped != 8 {
		t.Errorf("expected popped value to be 8, got %v", popped)
	}
}

func TestHeap_PopEmpty(t *testing.T) {
	comparator := func(a *int, b *int) bool {
		return *a < *b // Min-heap comparator
	}

	h := NewHeap(comparator)

	_, err := h.Pop()
	if err == nil {
		t.Fatalf("expected an error when popping from an empty heap")
	}
}

func TestHeap_PeekEmpty(t *testing.T) {
	comparator := func(a *int, b *int) bool {
		return *a < *b // Min-heap comparator
	}

	h := NewHeap(comparator)

	_, err := h.Peek()
	if err == nil {
		t.Fatalf("expected an error when peeking into an empty heap")
	}
}

func TestHeap_Ordering(t *testing.T) {
	comparator := func(a *int, b *int) bool {
		return *a < *b // Min-heap comparator
	}

	h := NewHeap(comparator)

	// Insert multiple elements
	elements := []int{10, 4, 5, 1, 3, 2}
	for _, el := range elements {
		h.Push(el)
	}

	// Extract elements and verify ordering
	expectedOrder := []int{1, 2, 3, 4, 5, 10}
	for _, expected := range expectedOrder {
		popped, err := h.Pop()
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		if popped != expected {
			t.Errorf("expected popped value to be %v, got %v", expected, popped)
		}
	}
}
