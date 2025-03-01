package main

import "testing"

func TestBinarySearch(t *testing.T) {

	tests := map[string]struct {
		slice   []int
		target  int
		want    Search
		wantErr bool
	}{
		"correct search": {
			slice:  []int{4, 7, 8, 10, 11, 17, 19, 21, 25},
			target: 11,
			want: Search{
				idx:   4,
				found: true,
			},
			wantErr: false,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {

			got, err := binarySearch(tc.slice, tc.target)
			if (err != nil) != tc.wantErr {
				t.Errorf("binarySearch(%v, %v): error = %v, wantErr %v", tc.slice, tc.target, err, tc.wantErr)
			}
			if got != tc.want {
				t.Errorf("binarySearch(%v, %v): got %v, want %v", tc.slice, tc.target, got, tc.want)
			}
		})
	}

}
