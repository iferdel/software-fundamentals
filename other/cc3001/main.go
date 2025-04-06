package main

import "fmt"

func main() {
	lower, upper := partition([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 3)
	fmt.Printf("lower and upper bounds are %v, %v, respectively", lower, upper)
}

func partition(a []int, p int) (lower, upper []int) {
	lower, upper = []int{}, []int{}

	for i := 0; i < len(a); i++ {
		if a[i] < p {
			lower = append(lower, a[i])
		}
		if a[i] > p {
			upper = a[i:]
			break
		}
	}
	return lower, upper

}
