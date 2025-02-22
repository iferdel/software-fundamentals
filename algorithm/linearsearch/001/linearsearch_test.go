package linearsearch

import "testing"

func TestLinearSearch(t *testing.T) {
	tests := map[string]struct {
		array   []string
		input   string
		want    int
		wantErr bool
	}{
		"correct search": {
			array: []string{
				"apple",
				"maple",
				"hello",
				"goofying",
				"hallo",
			},
			input:   "hello",
			want:    2,
			wantErr: false,
		},
		"input not found in array": {
			array: []string{
				"apple",
				"maple",
				"hello",
				"goofying",
				"hallo",
			},
			input:   "helloa",
			want:    0,
			wantErr: true,
		},
		"array is empty": {
			array:   []string{},
			input:   "hello",
			want:    0,
			wantErr: true,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := linearSearch(tc.array, tc.input)
			if (err != nil) != tc.wantErr {
				t.Fatalf("error = %v, wantErr %v", err, tc.wantErr)
			}
			if got != tc.want {
				t.Fatalf("Looking %q in array %v. Expected %v, got %v", tc.input, tc.array, tc.want, got)
			}
		})
	}
}
