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

	if target < slice[lower] {
		return s, nil
	}

	if target > slice[upper] {
		return s, nil
	}

	for lower <= upper {
		middlepoint := (lower + upper) / 2
		if target == slice[middlepoint] {
			s.found = true
			s.idx = middlepoint
			return s, nil
		}
		if target < slice[middlepoint] {
			upper = middlepoint - 1
			continue
		}
		if target > slice[middlepoint] {
			lower = middlepoint + 1
			continue
		}
	}

	return s, nil
}
