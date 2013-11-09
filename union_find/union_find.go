package union_find

/*

Quick Find algorithm for Union Find problem
https://class.coursera.org/algs4partI-003/lecture/6

*/

type QuickFindUF struct {
	Ids []int
}

func NewQuickFindUF(n int) (union QuickFindUF) {
	ids := make([]int, n)
	for i := int(0); i < n; i++ {
		ids[i] = i
	}
	union = QuickFindUF{ ids }
	return
}

func (union QuickFindUF) Connected(p, q int) bool {
	return union.Ids[p] == union.Ids[q]
}

func (union QuickFindUF) Union(p, q int) {
	pid := union.Ids[p]
	qid := union.Ids[q]
	for i, v := range union.Ids {
		if(v == pid){
			union.Ids[i] = qid
		}
	}
}

/*

Quick Union algorithm for Union Find problem
https://class.coursera.org/algs4partI-003/lecture/7

*/

type QuickUnionUF struct {
	Ids []int
}

func NewQuickUnionUF(n int) (union QuickUnionUF) {
	ids := make([]int, n)
	for i := int(0); i < n; i++ {
		ids[i] = i
	}
	union = QuickUnionUF{ ids }
	return
}

func (union QuickUnionUF) Root(i int) int {
	for {
		if union.Ids[i] == i { break }
		i = union.Ids[i]
	}
	return i
}

func (union QuickUnionUF) Connected(p, q int) bool {
	return union.Root(p) == union.Root(q)
}

func (union QuickUnionUF) Union(p, q int) {
	proot := union.Root(p)
	qroot := union.Root(q)
	union.Ids[proot] = qroot
}

/* 

Weighted Quick Union Algorithm for Union Find problem
https://class.coursera.org/algs4partI-003/lecture/8

*/

type WeightedQuickUnionUF struct{
	Ids []int
	Weights []int
}

func NewWeightedQuickUnionUF(n int) (union WeightedQuickUnionUF) {
	ids := make([]int, n)
	weights := make([]int, n)
	for i := int(0); i < n; i++ {
		ids[i] = i
		weights[i] = 1
	}
	union = WeightedQuickUnionUF{ ids, weights }
	return
}

func (union WeightedQuickUnionUF) Root(i int) int {
	for {
		if union.Ids[i] == i { break }
		i = union.Ids[i]
	}
	return i
}

func (union WeightedQuickUnionUF) Connected(p, q int) bool {
	return union.Root(p) == union.Root(q)
}

func (union WeightedQuickUnionUF) Union(p, q int) {
	proot := union.Root(p)
	qroot := union.Root(q)
	if union.Weights[qroot] >= union.Weights[proot] {
		union.Ids[proot] = qroot
		union.Weights[qroot] += union.Weights[proot]
	} else {
		union.Ids[qroot] = proot
		union.Weights[proot] += union.Weights[qroot]
	}
}

/* 

Weighted Quick Union Algorithm with Path Compressing for Union Find problem
https://class.coursera.org/algs4partI-003/lecture/8

*/

type CompressedWeightedQuickUnionUF struct{
	Ids []int
	Weights []int
}

func NewCompressedWeightedQuickUnionUF(n int) (union CompressedWeightedQuickUnionUF) {
	ids := make([]int, n)
	weights := make([]int, n)
	for i := int(0); i < n; i++ {
		ids[i] = i
		weights[i] = 1
	}
	union = CompressedWeightedQuickUnionUF{ ids, weights }
	return
}

func (union CompressedWeightedQuickUnionUF) Root(i int) int {
	for {
		if union.Ids[i] == i { break }
		union.Ids[i] = union.Ids[union.Ids[i]]
		i = union.Ids[i]
	}
	return i
}

func (union CompressedWeightedQuickUnionUF) Connected(p, q int) bool {
	return union.Root(p) == union.Root(q)
}

func (union CompressedWeightedQuickUnionUF) Union(p, q int) {
	proot := union.Root(p)
	qroot := union.Root(q)
	if union.Weights[qroot] >= union.Weights[proot] {
		union.Ids[proot] = qroot
		union.Weights[qroot] += union.Weights[proot]
	} else {
		union.Ids[qroot] = proot
		union.Weights[proot] += union.Weights[qroot]
	}
}