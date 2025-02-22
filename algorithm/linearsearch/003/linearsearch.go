package linearsearch

import (
	"errors"
	"reflect"
)

func linearSearch(array []interface{}, target interface{}) (int, error) {
	if len(array) == 0 {
		return 0, errors.New("empty array used as input")
	}

	for i := 0; i <= len(array)-1; i++ {
		if equality := reflect.DeepEqual(array[i], target); equality {
			return i, nil
		}
	}

	return 0, errors.New("target not found in array")
}
