package linearsearch

import "fmt"

type Search struct {
	idx   int
	found bool
}

func linearSearchInt(slice []int, target int, isOrdered bool) (Search, error) {
	s := Search{
		found: false,
	}
	for idx, val := range slice {
		if isOrdered {
			if val > target {
				fmt.Println("early return on linearSearchIntOrdered, value is not found")
				break
			}
		}
		if val == target {
			s.found = true
			s.idx = idx
			return s, nil
		}
	}
	return s, nil
}
