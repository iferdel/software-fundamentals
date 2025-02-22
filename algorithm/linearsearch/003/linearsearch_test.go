package linearsearch

import "testing"

func TestLinearSearch(t *testing.T) {
	tests := map[string]struct {
		array   []interface{}
		lookUp  interface{}
		want    int
		wantErr bool
	}{
		"float search": {
			array: []interface{}{
				1.3,
				14.1,
				1.11,
			},
			lookUp:  14.1,
			want:    1,
			wantErr: false,
		},
		"string search": {
			array: []interface{}{
				"first word",
				"hello",
			},
			lookUp:  "hello",
			want:    1,
			wantErr: false,
		},
		"int search": {
			array: []interface{}{
				1,
				14,
				1,
				2,
				3,
			},
			lookUp:  1,
			want:    0,
			wantErr: false,
		},
		"mixed types search": {
			array: []interface{}{
				"first word",
				1,
				141.2,
				1.2,
				"apple",
			},
			lookUp:  "apple",
			want:    4,
			wantErr: false,
		},
		"empty array": {
			array:   []interface{}{},
			lookUp:  "apple",
			want:    0,
			wantErr: true,
		},
		"not found": {
			array: []interface{}{
				1,
				14,
				1,
				2,
				3,
			},
			lookUp:  10,
			want:    0,
			wantErr: true,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := linearSearch(tc.array, tc.lookUp)
			if (err != nil) != tc.wantErr {
				t.Fatalf("linearSearch(%v, %v): error = %v, wantErr %v", tc.array, tc.lookUp, err, tc.wantErr)
			}
			if got != tc.want {
				t.Fatalf("linearSearch(%v, %v): got %v, want %v", tc.array, tc.lookUp, got, tc.want)
			}
		})
	}
}
