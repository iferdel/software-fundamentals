package binarysearch

type Search struct {
	idx   int
	found bool
}

func binarySearch(slice []int, target int) (Search, error) {
	s := Search{
		found: false,
	}

	lower := 0
	upper := len(slice) - 1

	for lower <= upper {
		middlepoint := (upper + lower) / 2 // go actually divides without need of using floor division since its by default depending on the input type

		if slice[middlepoint] == target {
			s.found = true
			s.idx = middlepoint
			return s, nil
		}

		if slice[middlepoint] > target {
			upper = middlepoint - 1
			continue
		}

		if slice[middlepoint] < target {
			lower = middlepoint + 1
			continue
		}
	}
	return s, nil
}
