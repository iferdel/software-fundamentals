package linearsearch

type Search struct {
	idx   int
	found bool
}

func linearSearch(slice []string, target string) (Search, error) {
	s := Search{
		found: false,
	}
	for idx, val := range slice {
		if val == target {
			s.found = true
			s.idx = idx
			return s, nil
		}
	}
	return s, nil
}
