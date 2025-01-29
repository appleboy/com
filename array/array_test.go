package array

import (
	"testing"
)

// BenchmarkContains benchmarks the performance of the Contains function.
func BenchmarkContains(b *testing.B) {
	b.Run("10 fileds", func(b *testing.B) {
		slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		key := 5
		b.ReportAllocs()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			Contains(slice, key)
		}
	})

	b.Run("100 fileds", func(b *testing.B) {
		slice := make([]int, 100)
		for i := 0; i < 100; i++ {
			slice[i] = i
		}
		key := 50
		b.ReportAllocs()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			Contains(slice, key)
		}
	})

	b.Run("1000 fileds", func(b *testing.B) {
		slice := make([]int, 1000)
		for i := 0; i < 1000; i++ {
			slice[i] = i
		}
		key := 500
		b.ReportAllocs()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			Contains(slice, key)
		}
	})
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
