package main

import (
	"fmt"

	binarysearch "github.com/iferdel/software-fundamentals/algorithm/binarysearch/003"
)

func main() {
	fmt.Println("----")
	fmt.Println("regular partition")
	lower, upper := partition([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 3)
	fmt.Printf("lower and upper bounds are %v, %v, respectively\n", lower, upper)
	fmt.Println("----")
	fmt.Println("binarysearched partition")
	lower, upper, err := partitionUsingBinarySearch([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 3)
	if err != nil {
		fmt.Printf("error in partition with binarysearch: %v", err)
	}
	fmt.Printf("lower and upper bounds are %v, %v, respectively\n", lower, upper)
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

func partitionUsingBinarySearch(a []int, p int) (lower, upper []int, err error) {
	lower, upper = []int{}, []int{}

	search, err := binarysearch.BinarySearch(a, p)
	if err != nil {
		return lower, upper, fmt.Errorf("error in binary search: %v", err)
	}
	fmt.Printf("%v is in idx %v\n", p, search.Idx)

	lower = a[:search.Idx]
	upper = a[search.Idx:]

	return lower, upper, nil

}
