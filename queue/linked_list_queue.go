package queue

// import( "fmt" )

/*

Simple implementation of Linked List Queue
https://class.coursera.org/algs4partI-003/lecture/20

*/

type LinkedListQueue struct {
	First *LinkedListNode
	Last  *LinkedListNode
}

type LinkedListNode struct {
	Item interface{}
	Next *LinkedListNode
}

func NewLinkedListQueue() (queue *LinkedListQueue) {
	queue = new(LinkedListQueue)
	queue.First = new(LinkedListNode)
	queue.Last = new(LinkedListNode)
	return
}

func (queue *LinkedListQueue) Enqueue(item interface{}) {
	newNode := LinkedListNode{ item, nil}
	if queue.IsEmpty() { 
		queue.First = &newNode 
	} else {
		queue.Last.Next = &newNode 
	}
	queue.Last = &newNode
}

func (queue *LinkedListQueue) Dequeue() (item interface{}) {
	first := queue.First
	item = first.Item
	if queue.IsEmpty() { 
		queue.Last.Item = nil 
	}
	if first.Next != nil { 
		queue.First = first.Next 
	} else { 
		queue.First.Item = nil 
	}
	return
}

func (queue *LinkedListQueue) IsEmpty() bool {
	return queue.First == nil || queue.First.Item == nil
}