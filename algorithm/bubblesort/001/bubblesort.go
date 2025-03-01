package bubblesort

func bubbleSort(slice []int) ([]int, error) {
	swapping := true
	end := len(slice)

	for swapping {
		swapping = false
		for i := 1; i < end; i++ {
			if slice[i-1] > slice[i] {
				num := slice[i-1]
				slice[i-1] = slice[i]
				slice[i] = num
				swapping = true
			}
		}
		end -= 1
	}
	return slice, nil
}
