package priority_queue
import (
	// "fmt"
)

type MaxPQ struct {
	pq []int
	N int
}

func NewMaxPQ(cap int) *MaxPQ {
	pq := make([]int, cap+1)
	return &MaxPQ{
		pq,
		0,
	}
}

func (pq *MaxPQ) isEmpty() bool {
	return pq.N == 0
}

func (pq *MaxPQ) Insert(x int) {
	pq.N++
	pq.pq[pq.N] = x
	pq.swim(pq.N)
}

func (pq *MaxPQ) DelMax() int {
	max := pq.pq[1]
	pq.swap(1, pq.N)
	pq.N--
	pq.sink(1)
	return max
}

func (pq *MaxPQ) swim(k int) {
	for k > 1 && pq.less(k/2, k) {
		pq.swap(k, k/2)
		k = k/2
	}
}

func (pq *MaxPQ) sink(k int) {
	for 2*k <= pq.N {
		j := 2*k
		if j < pq.N && pq.less(j, j+1){
			j++
		}
		if(!pq.less(k, j)) {
			break
		}
		pq.swap(k, j)
		k = j
	}
}

func (pq *MaxPQ) swap(i, j int) {
	i_val := pq.pq[i]
	pq.pq[i] = pq.pq[j]
	pq.pq[j] = i_val
}

func (pq *MaxPQ) less(i, j int) bool {
	return pq.pq[i] < pq.pq[j]
}