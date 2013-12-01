package sort

import (
	"testing"
)

func TestLinkedListStack(t *testing.T) {
	data := []int{ 1, 3, 2, 5, 4 }
	SelectionSort(data)
	for i, v := range(data) {
		if i+1 != v {
			t.Errorf("%d should be equeal to %d", v, i+1)
		}
	}
	
	data = []int{ 1, 3, 2, 5, 4 }
	InsertionSort(data)
	for i, v := range(data) {
		if i+1 != v {
			t.Errorf("%d should be equeal to %d", v, i+1)
		}
	}
}