package linearsearch

import "testing"

func TestLinearSearchInt(t *testing.T) {
	tests := map[string]struct {
		slice     []int
		isOrdered bool
		target    int
		want      Search
		wantErr   bool
	}{
		"correct search over ordered slice": {
			slice: []int{
				0,
				1,
				5,
				7,
				9,
			},
			isOrdered: true,
			target:    5,
			want: Search{
				idx:   2,
				found: true,
			},
			wantErr: false,
		},
		"not found": {
			slice: []int{
				1,
				6,
				8,
			},
			isOrdered: true,
			target:    10,
			want: Search{
				idx:   0,
				found: false,
			},
			wantErr: false,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := linearSearchInt(tc.slice, tc.target, tc.isOrdered)
			if (err != nil) != tc.wantErr {
				t.Errorf("linearSearchString(%v, %v): error = %v, wantErr %v", tc.slice, tc.target, err, tc.wantErr)
			}
			if got != tc.want {
				t.Errorf("linearSearchString(%v, %v): got %v want %v", tc.slice, tc.target, got, tc.want)
			}
		})
	}
}
