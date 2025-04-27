package partition

import "testing"

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
			wantLeft:  []int{1, 1, 2, 2},
			wantRight: []int{5, 6, 11},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			left, right := partition(tc.array, tc.pivot)
			if left != tc.wantLeft || right != tc.wantRight {
				t.Errorf(
					"pivot %s -- want left partition %s, got %s - want right partition %s, got %s",
					tc.pivot, tc.wantLeft, left, tc.wantRight, right,
				)
			}

		})

	}

}
