package partition

import (
	"reflect"
	"testing"
)

func TestPartition(t *testing.T) {
	tests := map[string]struct {
		array     []int
		pivot     int
		wantLeft  []int
		wantRight []int
	}{
		"correct partition": {
			array:     []int{5, 1, 2, 8, 11, 1, 2},
			pivot:     3,
			wantLeft:  []int{2, 1, 2, 1},
			wantRight: []int{11, 8, 5},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			left, right := partition(tc.array, tc.pivot)
			if !reflect.DeepEqual(left, tc.wantLeft) || !reflect.DeepEqual(right, tc.wantRight) {
				t.Errorf(
					"pivot %v -- want left partition %v, got %v - want right partition %v, got %v",
					tc.pivot, tc.wantLeft, left, tc.wantRight, right,
				)
			}

		})

	}

}
