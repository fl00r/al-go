package sort

import (
	"testing"
)

func TestLinkedListStack(t *testing.T) {
	var data []int

	data = []int{ 5, 3, 2, 1, 4 }
	SelectionSort(data)
	for i, v := range(data) {
		if i+1 != v {
			t.Errorf("SelectionSort: %d should be equal to %d for %d", v, i+1, data)
		}
	}
	
	data = []int{ 5, 3, 2, 1, 4 }
	InsertionSort(data)
	for i, v := range(data) {
		if i+1 != v {
			t.Errorf("InsertionSort: %d should be equal to %d for %d", v, i+1, data)
		}
	}
	
	data = []int{ 5, 3, 2, 1, 4 }
	ShellSort(data)
	for i, v := range(data) {
		if i+1 != v {
			t.Errorf("ShellSort: %d should be equal to %d for %d", v, i+1, data)
		}
	}
}