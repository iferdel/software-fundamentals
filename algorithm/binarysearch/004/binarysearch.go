package main

import "log"

type Search struct {
	idx   int
	found bool
}

func main() {
	slice := []int{4, 7, 8, 10, 11, 17, 19, 21, 25}
	target := 55
	s, err := binarySearch(slice, target)
	if err != nil {
		log.Println("binarySearch error=", err)
	}

	log.Println("binarySearch returned found as", s.found)
}

func binarySearch(slice []int, target int) (Search, error) {
	log.Println("starting binary search algorithm")
	s := Search{}
	log.Printf("binarySearch: Search struct is being initialized with 'empty' values: %v", s)

	loweridx := 0
	upperidx := len(slice) - 1

	if target < slice[loweridx] {
		// there must be a better way, for longer slices it should be changed
		log.Printf("binarySearch: target %v is bellow the lowest value in searched slice %v", target, slice[loweridx])
	}
	if target > slice[loweridx] {
		// there must be a better way, for longer slices it should be changed
		log.Printf("binarySearch: target %v is over the greatest value in searched slice %v", target, slice[upperidx])
	}

	for loweridx <= upperidx {
		middlepointidx := (upperidx + loweridx) / 2 //loweridx is being sum as we are narrowing the slice, so we need an offset

		if target == slice[middlepointidx] {
			s.idx = middlepointidx
			s.found = true
			return s, nil
		}

		if target < slice[middlepointidx] {
			upperidx = middlepointidx - 1
			continue
		}

		if target > slice[middlepointidx] {
			loweridx = middlepointidx + 1
		}

	}

	return s, nil
}
