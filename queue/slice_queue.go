package queue

// import( "fmt" )

type SliceQueue struct {
	head, tail, total int
	data []interface{}
}

func NewSliceQueue() (queue *SliceQueue) {
	queue = &SliceQueue{}
	queue.data = make([]interface{}, 1, 1)
	return
}

func (queue *SliceQueue) Enqueue(item interface{}) {
	queue.total++
	queue.checkSize()
	queue.data[queue.head] = item
	queue.incPos(&queue.head)
}

func (queue *SliceQueue) Dequeue() interface{} {
	if queue.IsEmpty() {
		return nil
	}
	item := queue.data[queue.tail]
	queue.incPos(&queue.tail)
	queue.total--
	queue.checkSize()
	return item
}

func (queue *SliceQueue) IsEmpty() bool {
	return queue.total == 0
}

/*

private party

*/

func (queue *SliceQueue) incPos(pos *int) {
	*pos++
	if *pos >= cap(queue.data) {
		*pos = 0
	}
}

func (queue *SliceQueue) checkSize() {
	cd := cap(queue.data)
	if queue.total > cd {
		newCap := cd * 2
		queue.resizeSlice(newCap)
	} else if queue.total < cd/4 && queue.total > 0 {
		newCap := cd / 2
		queue.resizeSlice(newCap)
	}
}

func (queue *SliceQueue) resizeSlice(newCap int) {
	newSlice := make([]interface{}, newCap, newCap)
	if queue.tail < queue.head {
		for k := queue.tail; k < queue.head; k++ {
			newSlice[k] = queue.data[k]
		}
	} else {
		for k := queue.tail; k < cap(queue.data); k++ {
			newSlice[k] = queue.data[k]
		}
		for k := 0; k < queue.head; k++ {
			newSlice[k] = queue.data[k]
		}
	}
	queue.data = newSlice
	queue.tail = 0
	queue.head = queue.total-1
}