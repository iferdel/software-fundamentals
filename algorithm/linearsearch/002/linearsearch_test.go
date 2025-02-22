package linearsearch

import "testing"

func TestLinearSearch(t *testing.T) {
	tests := map[string]struct {
		array   []string
		lookUp  string
		want    int
		wantErr bool
	}{
		"correct search": {
			array:   []string{"one", "viaam", "arrow", "whatssss", "give it"},
			lookUp:  "arrow",
			want:    2,
			wantErr: false,
		},
		"empty array": {
			array:   []string{},
			lookUp:  "arrow",
			want:    0,
			wantErr: true,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := linearSearch(tc.array, tc.lookUp)
			if (err != nil) != tc.wantErr {
				t.Fatalf("linearSearch(%v, %v): error = %q, wantErr %v",
					tc.array, tc.lookUp, err, tc.wantErr)
			}
			if got != tc.want {
				t.Fatalf("linearSearch(%v, %q): got index %v, want index %v",
					tc.array, tc.lookUp, got, tc.want)
			}
		})
	}
}
