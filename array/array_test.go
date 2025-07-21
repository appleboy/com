package array

import (
	"testing"
)

// BenchmarkContains benchmarks the performance of the Contains function.
func BenchmarkContains(b *testing.B) {
	type benchCase[T comparable] struct {
		name  string
		slice []T
		key   T
	}
	intCases := []benchCase[int]{
		{
			name:  "10 fields (int, key exists)",
			slice: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			key:   5,
		},
		{
			name: "100 fields (int, key exists)",
			slice: func() []int {
				s := make([]int, 100)
				for i := 0; i < 100; i++ {
					s[i] = i
				}
				return s
			}(),
			key: 50,
		},
		{
			name: "1000 fields (int, key exists)",
			slice: func() []int {
				s := make([]int, 1000)
				for i := 0; i < 1000; i++ {
					s[i] = i
				}
				return s
			}(),
			key: 500,
		},
		{
			name:  "empty slice (int)",
			slice: []int{},
			key:   1,
		},
		{
			name:  "key not found (int)",
			slice: []int{1, 2, 3, 4, 5},
			key:   99,
		},
	}
	stringCases := []benchCase[string]{
		{
			name:  "10 fields (string, key exists)",
			slice: []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"},
			key:   "e",
		},
		{
			name:  "key not found (string)",
			slice: []string{"a", "b", "c"},
			key:   "z",
		},
		{
			name:  "empty slice (string)",
			slice: []string{},
			key:   "a",
		},
	}
	for _, tc := range intCases {
		b.Run(tc.name, func(b *testing.B) {
			b.ReportAllocs()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				Contains(tc.slice, tc.key)
			}
		})
	}
	for _, tc := range stringCases {
		b.Run(tc.name, func(b *testing.B) {
			b.ReportAllocs()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				Contains(tc.slice, tc.key)
			}
		})
	}
}

func TestContains(t *testing.T) {
	type args struct {
		slice []int
		key   int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "key exists in slice",
			args: args{
				slice: []int{1, 2, 3, 4, 5},
				key:   3,
			},
			want: true,
		},
		{
			name: "key does not exist in slice",
			args: args{
				slice: []int{1, 2, 3, 4, 5},
				key:   6,
			},
			want: false,
		},
		{
			name: "empty slice",
			args: args{
				slice: []int{},
				key:   1,
			},
			want: false,
		},
		{
			name: "single element slice, key exists",
			args: args{
				slice: []int{1},
				key:   1,
			},
			want: true,
		},
		{
			name: "single element slice, key does not exist",
			args: args{
				slice: []int{1},
				key:   2,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Contains(tt.args.slice, tt.args.key); got != tt.want {
				t.Errorf("Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}
