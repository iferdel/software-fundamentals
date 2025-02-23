package linearsearch

import "fmt"

type Search struct {
	idx   int
	found bool
}

func linearSearchString(slice []string, target string) (Search, error) {
	s := Search{
		found: false,
	}

	for i := 0; i < len(slice); i++ {
		if slice[i] == target {
			s.idx = i
			s.found = true
		}
	}

	fmt.Printf("%q not found in collection of items %v\n", target, slice)
	return s, nil
}
