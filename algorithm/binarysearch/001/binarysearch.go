package binarysearch

type Search struct {
	idx   int
	found bool
}

func binarySearch(slice []string, target string) (Search, error) {
	s := Search{
		found: false,
	}

	return s, nil
}
