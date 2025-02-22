package linearsearch

import "errors"

func linearSearch(array []string, lookUp string) (int, error) {
	for i := 0; i < len(array)-1; i++ {
		if array[i] == lookUp {
			return i, nil
		}
	}
	return 0, errors.New("value not found in array")
}
