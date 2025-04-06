package binarysearch

type Search struct {
	Idx   int
	Found bool
}

func BinarySearch(slice []int, target int) (Search, error) {

	s := Search{}
	lower := 0
	upper := len(slice) - 1

	if slice[lower] > target {
		return s, nil
	}

	if slice[upper] < target {
		return s, nil
	}

	for lower <= upper {
		middlepoint := (lower + upper) / 2
		if slice[middlepoint] == target {
			s.Found = true
			s.Idx = middlepoint
			return s, nil
		}

		if slice[middlepoint] > target {
			upper = middlepoint - 1
			continue
		}

		if slice[middlepoint] < target {
			lower = middlepoint + 1
		}
	}

	return s, nil
}
