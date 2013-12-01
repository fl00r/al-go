package sort

func SelectionSort(data []int) {
	var smallest, smallest_val int
	n := len(data)
	for i := 0; i < n; i++ {
		smallest = i
		for j := i+1; j < n; j++ {
			if(data[j] < data[smallest]) {
				smallest = j
			}
		}
		smallest_val = data[smallest]
		data[smallest] = data[i]
		data[i] = smallest_val
	}
}

func InsertionSort(data []int) {
	var smallest, smallest_val int
	n := len(data)
	for i := 0; i < n; i++ {
		val := data[i]
		smallest = i
		for j := i-1; j > 0 && val < data[j]; j-- {
			smallest = j
		}
		smallest_val = data[smallest]
		data[smallest] = data[i]
		data[i] = smallest_val
	}
}