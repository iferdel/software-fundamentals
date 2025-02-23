package linearsearch

// import (
// 	"testing"
// )
//
// func TestLinearSearch[T comparable](t *testing.T) {
// 	tests := map[string]struct {
// 		slice   []T
// 		target  T
// 		want    Search
// 		wantErr bool
// 	}{
// 		"correct search string": {
// 			slice: []string{
// 				"first word",
// 				"hello",
// 				"welcome",
// 			},
// 			target: "welcome",
// 			want: Search{
// 				idx:   2,
// 				found: true,
// 			},
// 			wantErr: false,
// 		},
// 		"correct search int": {
// 			slice: []T{
// 				1,
// 				5,
// 				9,
// 			},
// 			target: 5,
// 			want: Search{
// 				idx:   1,
// 				found: true,
// 			},
// 			wantErr: false,
// 		},
// 		"correct search mixed": {
// 			slice: []T{
// 				1,
// 				"hello",
// 				"welcome",
// 				4.1,
// 				5,
// 			},
// 			target: "welcome",
// 			want: Search{
// 				idx:   2,
// 				found: true,
// 			},
// 			wantErr: false,
// 		},
// 		"not found": {
// 			slice: []string{
// 				"first word",
// 				"hello",
// 				"welcome",
// 			},
// 			target: "hallo",
// 			want: Search{
// 				idx:   0,
// 				found: false,
// 			},
// 			wantErr: false,
// 		},
// 	}
//
// 	for name, tc := range tests {
// 		t.Run(name, func(t *testing.T) {
// 			got, err := genericLinearSearch(tc.slice, tc.target)
// 			if (err != nil) != tc.wantErr {
// 				t.Errorf("genericLinearSearchString(%v, %v): error = %v, wantErr %v", tc.slice, tc.target, err, tc.wantErr)
// 			}
// 			if got != tc.want {
// 				t.Errorf("genericLinearSearchString(%v, %v): got %v want %v", tc.slice, tc.target, got, tc.want)
// 			}
// 		})
// 	}
// }
