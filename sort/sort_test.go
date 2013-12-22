package sort

import (
	"testing"
	"math/rand"
	"sort"
	// "fmt"
)

func randomArray(num int) []int {
	data := make([]int, num, num)
	for i, _ := range(data) {
		data[i] = rand.Intn(1000)
	}
	return data
}
func copyArray(origin []int) []int {
	new_orig := make([]int, len(origin))
	copy(new_orig, origin)
	return new_orig
}

func TestLinkedListStack(t *testing.T) {
	var current_copy []int

	data := randomArray(100)
	sorted_data := copyArray(data)
	sort.Ints(sorted_data)

	current_copy = copyArray(data)
	SelectionSort(current_copy)
	for i, v := range(current_copy) {
		s := sorted_data[i]
		if v != s {
			t.Errorf("SelectionSort: %d not equal %d", v, s)
		}
	}

	current_copy = copyArray(data)
	InsertionSort(current_copy)
	for i, v := range(current_copy) {
		s := sorted_data[i]
		if v != s {
			t.Errorf("InsertionSort: %d not equal %d", v, s)
		}
	}

	current_copy = copyArray(data)
	ShellSort(current_copy)
	for i, v := range(current_copy) {
		s := sorted_data[i]
		if v != s {
			t.Errorf("ShellSort: %d not equal %d", v, s)
		}
	}

	current_copy = copyArray(data)
	MergeSort(current_copy)
	for i, v := range(current_copy) {
		s := sorted_data[i]
		if v != s {
			t.Errorf("MergeSort: %d not equal %d", v, s)
		}
	}

	current_copy = copyArray(data)
	MergeSortBU(current_copy)
	for i, v := range(current_copy) {
		s := sorted_data[i]
		if v != s {
			t.Errorf("MergeSort: %d not equal %d", v, s)
		}
	}

	current_copy = copyArray(data)
	QuickSort(current_copy)
	for i, v := range(current_copy) {
		s := sorted_data[i]
		if v != s {
			t.Errorf("QuickSort: %d not equal %d", v, s)
		}
	}

	current_copy1 := copyArray(data)
	current_copy2 := copyArray(data)
	QuickSort(current_copy2)
	for i, v := range(current_copy2) {
		s := Select(current_copy1, i)
		if v != s {
			t.Errorf("Select: %d not equal %d", v, s)
		}
	}

	current_copy = copyArray(data)
	QuickSort3w(current_copy)
	for i, v := range(current_copy) {
		s := sorted_data[i]
		if v != s {
			t.Errorf("QuickSort3w: %d not equal %d", v, s)
		}
	}

	current_copy = copyArray(data)
	HeapSort(current_copy)
	for i, v := range(current_copy) {
		s := sorted_data[i]
		if v != s {
			t.Errorf("HeapSort: %d not equal %d", v, s)
		}
	}

	current_copy = copyArray(data)
	sort.Ints(current_copy)
	error := 0
	Shuffle(current_copy)
	for i, v := range(current_copy) {
		if i+1 == v {
			error += 1
		}
	}
	if error == len(current_copy) {
		t.Errorf("Shuffled array should not match origin %d", current_copy)
	}
}