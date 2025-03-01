package binarysearch

import (
	"reflect"
	"testing"
)

func TestBinarySearchMatrix(t *testing.T) {
	tests := map[string]struct {
		matrix  [][]int
		target  int
		want    Search
		wantErr bool
	}{
		"correct search": {
			matrix: [][]int{
				{1, 2, 3, 4, 5, 6},
				{4, 5, 6, 7, 8, 9},
				{2, 3, 4, 5, 6, 7},
			},
			target: 3,
			want: Search{
				idx:   []int{2, -1, 1},
				found: true,
			},
			wantErr: false,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := binarySearchMatrix(tc.matrix, tc.target)
			if (err != nil) != tc.wantErr {
				t.Errorf("binarySearchMatrix(%v, %v): error = %v, wantErr %v", tc.matrix, tc.target, err, tc.wantErr)
			}
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("binarySearchMatrix(%v, %v): got %v, want %v", tc.matrix, tc.target, got, tc.want)
			}
		})
	}
}
