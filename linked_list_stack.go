package algo

/*

Simple implementation of Linked List Stack
https://class.coursera.org/algs4partI-003/lecture/18

*/

type LinkedListStack struct {
	First LinkedListNode
}

type LinkedListNode struct {
	Item interface{}
	Next *LinkedListNode
}

func NewLinkedListStack() (list *LinkedListStack) {
	list = new(LinkedListStack)
	list.First.Next = &list.First
	return list
}

func (list *LinkedListStack) Push(item interface{}) {
	first := list.First
	newNode := LinkedListNode{ item, &first }
	list.First = newNode
}

func (list *LinkedListStack) Pop() (item interface{}) {
	oldFirst := list.First
	list.First = *oldFirst.Next
	item = oldFirst.Item
	return
}

func (list *LinkedListStack) IsEmpty() bool {
	return list.First.Item == nil
}