package stack

import (
	"testing"
)

func TestSliceStack(t *testing.T) {
	list := NewSliceStack()
	list.Push(1)
	list.Push(2)
	p := list.Pop()
	if p != 2 {
		t.Errorf("%d should be popped, not %d", 2, p)
	}
	list.Push(3)
	p = list.Pop()
	if p != 3 {
		t.Errorf("%d should be popped, not %d", 3, p)
	}
	p = list.Pop()
	if p != 1 {
		t.Errorf("%d should be popped, not %d", 1, p)
	}
	p = list.Pop()
	if p != nil {
		t.Errorf("%d should be popped, not %d", nil, p)
	}
	if !list.IsEmpty() {
		t.Errorf("Should be empty")
	}
	for k := 0; k < 100; k++ {
		list.Push(k)
	}
	for k := 0; k < 100; k++ {
		list.Pop()
	}
}