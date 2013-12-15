package priority_queue

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

func TestMaxPq(t *testing.T) {
	array := randomArray(100)
	l := len(array)
	pq_max := NewMaxPQ(l)
	for _, v := range(array) {
		pq_max.Insert(v)
	}
	sort.Ints(array)
	for i := 0; i < l; i++ {
		max := pq_max.DelMax()
		srt := array[l-i-1]
		if max != srt {
			t.Errorf("MaxPQ: %d not equal %d", max, srt)
		}
	}
}