package sort
import(
	"math/rand"
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

/*
	private party
*/

func swap(data []int, i, j int) {
	i_val := data[i]
	data[i] = data[j]
	data[j] = i_val
}

func less(data []int, i, j int) bool {
	return data[i] < data[j]
}