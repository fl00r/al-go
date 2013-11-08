package algo

import (
	"testing"
)

func TestQuickFind(t *testing.T) {
	union := NewQuickFindUF(10)

	union.Union(1, 2)
	if !union.Connected(1,2) {
		t.Errorf("Error: %d and %d should be connected", 1, 2)
	}

	union.Union(2, 8)
	if !union.Connected(2,8) {
		t.Errorf("Error: %d and %d should be connected", 2, 8)
	}
	if !union.Connected(1,8) {
		t.Errorf("Error: %d and %d should be connected", 1, 8)
	}

	myIds := []int{ 0, 8, 8, 3, 4, 5, 6, 7, 8, 9 }
	for i, v := range union.Ids {
		if v != myIds[i] {
			t.Errorf("Error: slices should be equal at %d, %d, %d", i, v, myIds[i])
		}
	}
}

func TestQuickUnion(t *testing.T) {
	union := NewQuickUnionUF(10)

	union.Union(1, 2)
	if !union.Connected(1,2) {
		t.Errorf("Error: %d and %d should be connected", 1, 2)
	}

	union.Union(2, 8)
	if !union.Connected(2,8) {
		t.Errorf("Error: %d and %d should be connected", 2, 8)
	}
	if !union.Connected(1,8) {
		t.Errorf("Error: %d and %d should be connected", 1, 8)
	}

	myIds := []int{ 0, 2, 8, 3, 4, 5, 6, 7, 8, 9 }
	for i, v := range union.Ids {
		if v != myIds[i] {
			t.Errorf("Error: slices should be equal at %d, %d, %d", i, v, myIds[i])
		}
	}
}

func TestWeightedQuickUnion(t *testing.T) {
	union := NewWeightedQuickUnionUF(10)

	union.Union(1, 2)
	if !union.Connected(1,2) {
		t.Errorf("Error: %d and %d should be connected", 1, 2)
	}

	union.Union(2, 8)
	if !union.Connected(2,8) {
		t.Errorf("Error: %d and %d should be connected", 2, 8)
	}
	if !union.Connected(1,8) {
		t.Errorf("Error: %d and %d should be connected", 1, 8)
	}

	myIds := []int{ 0, 2, 2, 3, 4, 5, 6, 7, 2, 9 }
	for i, v := range union.Ids {
		if v != myIds[i] {
			t.Errorf("Error: slices should be equal at %d, %d, %d", i, v, myIds[i])
		}
	}
}

func TestCompressedWeightedQuickUnion(t *testing.T) {
	union := NewCompressedWeightedQuickUnionUF(10)

	union.Union(1, 2)
	if !union.Connected(1,2) {
		t.Errorf("Error: %d and %d should be connected", 1, 2)
	}

	union.Union(2, 8)
	if !union.Connected(2,8) {
		t.Errorf("Error: %d and %d should be connected", 2, 8)
	}
	if !union.Connected(1,8) {
		t.Errorf("Error: %d and %d should be connected", 1, 8)
	}

	union.Union(3, 4)
	union.Union(3, 5)
	union.Union(5, 8)
	union.Union(6, 9)
	union.Union(5, 9)

	myIds := []int{ 0, 2, 2, 4, 2, 2, 9, 7, 2, 2 }
	for i, v := range union.Ids {
		if v != myIds[i] {
			t.Errorf("Error: slices should be equal at %d, %d, %d", i, v, myIds[i])
		}
	}
}