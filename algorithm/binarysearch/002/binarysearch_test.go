package binarysearch

import "testing"

func TestBinarySearch(t *testing.T) {
	tests := map[string]struct {
		slice   []int
		target  int
		want    Search
		wantErr bool
	}{
		"correct search": {
			slice:  []int{0, 2, 3, 7, 8, 10, 18},
			target: 7,
			want: Search{
				idx:   3,
				found: true,
			},
			wantErr: false,
		},
		"not found within": {
			slice:  []int{0, 2, 3, 7, 8, 10, 18},
			target: 6,
			want: Search{
				idx:   0,
				found: false,
			},
			wantErr: false,
		},
		"not found, under lowest": {
			slice:  []int{2, 3, 7, 8, 10, 18},
			target: 0,
			want: Search{
				idx:   0,
				found: false,
			},
			wantErr: false,
		},
		"not found, over upper": {
			slice:  []int{0, 2, 3, 7, 8, 10, 18},
			target: 255555,
			want: Search{
				idx:   0,
				found: false,
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
