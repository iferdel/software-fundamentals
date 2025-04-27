package orderinsert

import (
	"reflect"
	"testing"
)

func TestOrderInsert(t *testing.T) {
	tests := map[string]struct {
		arrayOne []int
		arrayTwo []int
		want     []int
	}{
		"correct partition": {
			arrayOne: []int{1, 2, 8, 11},
			arrayTwo: []int{1, 11, 20, 3, 4, 4, 42},
			want:     []int{1, 1, 2, 3, 4, 4, 8, 11, 11, 20, 42},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := orderInsert(tc.arrayOne, tc.arrayTwo)
			if reflect.DeepEqual(tc.want, got) {
				t.Errorf(
					"want %v, got %v",
					tc.want, got,
				)
			}

		})

	}

}
