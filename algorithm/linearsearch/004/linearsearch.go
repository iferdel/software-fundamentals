package linearsearch

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

	return s, nil
}
