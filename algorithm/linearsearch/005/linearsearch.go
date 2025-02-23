package linearsearch

import "fmt"

type Search struct {
	idx   int
	found bool
}

func genericLinearSearch[T comparable](slice []T, target T) (Search, error) {

	s := Search{
		found: false,
	}

	for i := 0; i < len(slice); i++ {
		if slice[i] == target {
			s.found = true
			s.idx = i
			return s, nil
		}
	}

	fmt.Printf("%q not found in collection of items %v\n", target, slice)
	return s, nil
}
