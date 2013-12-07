package sort
import(
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

/*
	private party
*/

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