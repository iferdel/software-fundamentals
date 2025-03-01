package bubblesort

import (
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	tests := map[string]struct {
		slice   []int
		want    []int
		wantErr bool
	}{

		"correct sort": {
			slice:   []int{2, 5, 1, 7, 8, 10, 4},
			want:    []int{1, 2, 4, 5, 7, 8, 10},
			wantErr: false,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := BubbleSort(tc.slice)
			if (err != nil) != tc.wantErr {
				t.Errorf("bubbleSort(%v): error = %v, wantErr %v", tc.slice, err, tc.wantErr)
			}
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("bubbleSort(%v): got %v, want %v", tc.slice, got, tc.want)
			}
		})
	}
}
