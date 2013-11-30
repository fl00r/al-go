package sort

func SelectionSort(data []int) {
	var smallest, smallest_val int
	for i := 0; i < len(data); i++ {
		smallest = i
		for j := i; j < len(data); j++ {
			if(data[j] < data[smallest]) {
				smallest = j
			}
		}
		smallest_val = data[smallest]
		data[smallest] = data[i]
		data[i] = smallest_val
	}
}
