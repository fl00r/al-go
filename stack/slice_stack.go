package stack

// import( "fmt" )

/*

Slice (not Array) implementation of Stack
https://class.coursera.org/algs4partI-003/lecture/19

*/

type SliceStack struct {
	data []interface{}
	n int
}

func NewSliceStack() (stack *SliceStack) {
	stack = new(SliceStack)
	stack.data = make([]interface{}, 1, 1)
	return
}

func (stack *SliceStack) Pop() (item interface{}) {
	s_len := cap(stack.data)
	stack.data[stack.n] = nil
	if stack.n > 0 {
		stack.n -= 1
	}
	if stack.n < (s_len/4) && stack.n > 0 {
		stack.resizeSlice(s_len/2)
	}
	item = stack.data[stack.n]
	return
}

func (stack *SliceStack) Push(item interface{}) {
	s_len := cap(stack.data)
	if stack.n == s_len - 1 {
		stack.resizeSlice(2 * s_len)
	}
	stack.data[stack.n] = item
	stack.n += 1
}

func(stack *SliceStack) IsEmpty() bool {
	return stack.n == 0
}

/*

private party

*/

func (stack *SliceStack) resizeSlice(newCap int) {
	newSlice := make([]interface{}, newCap, newCap)
	for k := 0; k <= stack.n ; k++ {
		newSlice[k] = stack.data[k]
	}
	stack.data = newSlice
}