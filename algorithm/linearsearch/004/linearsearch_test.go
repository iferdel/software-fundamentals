package linearsearch

import "testing"

func TestLinearSearchString(t *testing.T) {
	tests := map[string]struct {
		slice   []string
		target  string
		want    Search
		wantErr bool
	}{
		"correct search": {
			slice: []string{
				"first word",
				"hello",
				"welcome",
			},
			target: "welcome",
			want: Search{
				idx:   2,
				found: true,
			},
			wantErr: false,
		},
		"not found": {
			slice: []string{
				"first word",
				"hello",
				"welcome",
			},
			target: "hallo",
			want: Search{
				idx:   0,
				found: false,
			},
			wantErr: false,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := linearSearchString(tc.slice, tc.target)
			if (err != nil) != tc.wantErr {
				t.Errorf("linearSearch(%v, %v): error = %v, wantErr %v", tc.slice, tc.target, err, tc.wantErr)
			}
		})
	}
}
