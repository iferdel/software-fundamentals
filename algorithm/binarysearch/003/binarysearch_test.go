package binarysearch

import (
	"testing"
)

func TestBinarySearch(t *testing.T) {
	tests := map[string]struct {
		slice   []int
		target  int
		want    Search
		wantErr bool
	}{
		"correct search middlepoint": {
			slice:  []int{1, 4, 7, 9, 12, 78, 99},
			target: 9,
			want: Search{
				Idx:   3,
				Found: true,
			},
			wantErr: false,
		},
		"correct search to right": {
			slice:  []int{1, 4, 7, 9, 12, 78, 99},
			target: 78,
			want: Search{
				Idx:   5,
				Found: true,
			},
			wantErr: false,
		},
		"correct search to left": {
			slice:  []int{1, 4, 7, 9, 12, 78, 99},
			target: 1,
			want: Search{
				Idx:   0,
				Found: true,
			},
			wantErr: false,
		},
		"not found": {
			slice:  []int{1, 4, 7, 9, 12, 78, 99},
			target: 125,
			want: Search{
				Idx:   0,
				Found: false,
			},
			wantErr: false,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := BinarySearch(tc.slice, tc.target)
			if (err != nil) != tc.wantErr {
				t.Errorf("binarySearch(%v, %v): error = %v, wantErr %v", tc.slice, tc.target, err, tc.wantErr)
			}
			if got != tc.want {
				t.Errorf("binarySearch(%v, %v): got %v, want %v", tc.slice, tc.target, got, tc.want)
			}
		})
	}
}
