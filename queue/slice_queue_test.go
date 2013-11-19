package queue

import (
	"testing"
	// "fmt"
)

func TestSliceListQueue(t *testing.T) {
	list := NewSliceQueue()

	if !list.IsEmpty() {
		t.Errorf("Should be empty")
	}

	p := list.Dequeue()
	if p != nil {
		t.Errorf("%d should be popped, not %d", nil, p)
	}

	list.Enqueue(1)
	list.Enqueue(2)
	list.Enqueue(3)
	list.Enqueue("Peter")

	p = list.Dequeue()
	if p != 1 {
		t.Errorf("%d should be popped, not %d", 1, p)
	}

	list.Enqueue(4)
	p = list.Dequeue()
	if p != 2 {
		t.Errorf("%d should be popped, not %d", 2, p)
	}

	p = list.Dequeue()
	if p != 3 {
		t.Errorf("%d should be popped, not %d", 3, p)
	}

	p = list.Dequeue()
	if p != "Peter" {
		t.Errorf("%s should be popped, not %d", "Peter", p)
	}

	p = list.Dequeue()
	if p != 4 {
		t.Errorf("%d should be popped, not %d", 4, p)
	}

	p = list.Dequeue()
	if p != nil {
		t.Errorf("%d should be popped, not %d", nil, p)
	}

	if !list.IsEmpty() {
		t.Errorf("Should be empty")
	}
}