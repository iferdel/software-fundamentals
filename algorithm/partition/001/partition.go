package partition

import "fmt"

func partition(array []int, pivot int) ([]int, []int) {

	i := 0
	j := len(array) - 1

	for i <= j {
		fmt.Printf("i: %v, j: %v\n", i, j)
		if array[i] <= pivot {
			i += 1
		} else if array[j] >= pivot {
			j -= 1
		} else {
			array[i], array[j] = array[j], array[i]
		}
	}

	return array[:i], array[j+1:]

}
