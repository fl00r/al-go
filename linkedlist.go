package algo

/*

Simple implementation of Linked List

*/

type LinkedList struct {
	First LinkedListNode
}

type LinkedListNode struct {
	Item interface{}
	Next *LinkedListNode
}

func NewLinkedList() (list *LinkedList) {
	list = new(LinkedList)
	list.First.Next = &list.First
	return list
}

func (list *LinkedList) Push(item interface{}) {
	first := list.First
	newNode := LinkedListNode{ item, &first }
	list.First = newNode
}

func (list *LinkedList) Pop() (item interface{}) {
	oldFirst := list.First
	list.First = *oldFirst.Next
	item = oldFirst.Item
	return
}

func (list *LinkedList) IsEmpty() bool {
	return list.First.Item == nil
}