package bubblesort

import "log"

func BubbleSort(slice []int) ([]int, error) {
	log.Println("sorting", slice)
	end := len(slice) - 1
	sorted := false
	iteration := 1
	for !sorted {
		sorted = true
		for i := 0; i < end; i++ {
			if slice[i+1] < slice[i] {
				slice[i], slice[i+1] = slice[i+1], slice[i]
				sorted = false
			}
			log.Printf("(i, end) in iteration number %v: (%v, %v)", i, iteration, end)
		}
		end -= 1
		iteration += 1
	}
	return slice, nil
}
