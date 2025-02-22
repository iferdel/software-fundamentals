package linearsearch

import "fmt"

func linearSearch(array []string, value string) (int, error) {
	for i := 0; i < len(array)-1; i++ {
		if array[i] == value {
			return i, nil
		}
	}
	return 0, fmt.Errorf("%s not found in %v", value, array)
}
