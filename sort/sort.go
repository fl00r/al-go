package sort
import (
	"math/rand"
	// "fmt"
)

func SelectionSort(data []int) {
	n := len(data)
	for i := 0; i < n; i++ {
		for j := i+1; j < n; j++ {
			if(less(data, j, i)) {
				swap(data, i, j)
			}
		}
	}
}

func InsertionSort(data []int) {
	n := len(data)
	for i := 1; i < n; i++ {
		for j := i-1; j >= 0 && less(data, j+1, j); j-- {
			swap(data, j+1, j)
		}
	}
}

func ShellSort(data []int) {
	n := len(data)
	h := 1
	for h < n/3 {
		h = 3*h + 1
	}
	for h >= 1 {
		for i := 1; i < n; i++ {
			for j := i-1; j >= 0 && less(data, j+1, j); j -= h {
				swap(data, j+1, j)
			}
		}
		h = h / 3
	}
}

func Shuffle(data []int) {
	n := len(data)
	for i := 0; i < n; i++ {
		j := rand.Intn(n)
		swap(data, i, j)
	}
}

func MergeSort(data []int) {
	n := len(data)
	aux := make([]int, n, n)
	_ms_sort(data, aux, 0, n - 1)
}

func MergeSortBU(data []int) {
	n := len(data)
	aux := make([]int, n, n)
	for sz := 1; sz < n; sz *= 2 {
		for lo := 0; lo < n-sz; lo += 2*sz {
			hi := lo + 2*sz - 1
			if hi > n - 1 {
				hi = n - 1
			}
			_ms_merge(data, aux, lo, lo + sz - 1, hi)
		}
	}
}

func QuickSort(data []int) {
	Shuffle(data)
	_qs_sort(data, 0, len(data) - 1)
}

func Select(data []int, k int) int {
	Shuffle(data)
	lo := 0; hi := len(data) - 1
	for hi > lo {
		j := _qs_partition(data, lo, hi)
		if j < k {
			lo = j + 1
		} else if j > k {
			hi = j - 1
		} else {
			return data[k]
		}
	}
	return data[k]
}

func QuickSort3w(data []int) {
	Shuffle(data)
	_qs_sort_3w(data, 0, len(data) - 1)
}

/*
	private party
*/

func _qs_sort_3w(data []int, lo, hi int) {
	if hi <= lo {
		return
	}

	lt := lo; gt := hi
	i := lo
	v := data[lo]
	for i <= gt {
		w := data[i]
		if v > w {
			swap(data, lt, i); lt++; i++
		} else if v < w {
			swap(data, i, gt); gt--
		} else {
			i++
		}
	}
	_qs_sort_3w(data, lo, lt - 1)
	_qs_sort_3w(data, gt + 1, hi)
}

const(
	Cutoff = 10
)

func _qs_sort(data []int, lo, hi int) {
	if hi <= lo + Cutoff + 1 {
		InsertionSort(data)
		return
	}
	j := _qs_partition(data, lo, hi)
	_qs_sort(data, lo, j-1)
	_qs_sort(data, j+1, hi)
}

func _qs_partition(data []int, lo, hi int) int {
	i := lo+1; j := hi
	for {
		for less(data, i, lo) {
			if i == hi {
				break
			}
			i++
		}
		for less(data, lo, j) {
			if j == lo {
				break
			}
			j--
		}
		if i >= j {
			break
		}
		swap(data, i, j)
	}

	swap(data, lo, j)
	return j
}

func _ms_sort(data, aux []int, lo, hi int) {
	if(hi > lo) {
		mid := lo + (hi - lo) / 2
		_ms_sort(data, aux, lo, mid)
		_ms_sort(data, aux, mid+1, hi)
		_ms_merge(data, aux, lo, mid, hi)
	}
}

func _ms_merge(data, aux []int, lo, mid, hi int) {
	for i, v := range(data) {
		aux[i] = v
	}

	i := lo
	j := mid + 1
	for k := lo; k <= hi; k++ {
		if i > mid {
			data[k] = aux[j]; j++
		} else if j > hi {
			data[k] = aux[i]; i++
		} else if less(aux, i, j) {
			data[k] = aux[i]; i++
		} else {
			data[k] = aux[j]; j++
		}
	}
}

func swap(data []int, i, j int) {
	i_val := data[i]
	data[i] = data[j]
	data[j] = i_val
}

func less(data []int, i, j int) bool {
	return data[i] < data[j]
}