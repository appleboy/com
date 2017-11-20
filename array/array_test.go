package array

import (
	"reflect"
	"testing"
)

func TestIn(t *testing.T) {
	type args struct {
		needle   string
		haystack []string
	}
	tests := []struct {
		name  string
		args  args
		want  []string
		want1 bool
	}{
		{
			name: "test in array",
			args: args{
				needle:   "a",
				haystack: []string{"a", "b", "c"},
			},
			want:  []string{"b", "c"},
			want1: true,
		},
		{
			name: "test not in array",
			args: args{
				needle:   "d",
				haystack: []string{"a", "b", "c"},
			},
			want:  []string{"a", "b", "c"},
			want1: false,
		},
		{
			name: "test empty target array",
			args: args{
				needle:   "d",
				haystack: []string{},
			},
			want:  []string{},
			want1: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := In(tt.args.needle, tt.args.haystack)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("In() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("In() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestDiff(t *testing.T) {
	type args struct {
		s []string
		t []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "test not in array",
			args: args{
				s: []string{"d"},
				t: []string{"a", "b", "c"},
			},
			want: []string{"d", "a", "b", "c"},
		},
		{
			name: "test partial not in array",
			args: args{
				s: []string{"a", "c"},
				t: []string{"a", "b", "c"},
			},
			want: []string{"b"},
		},
		{
			name: "test all match in array",
			args: args{
				s: []string{"a", "c", "b"},
				t: []string{"a", "b", "c"},
			},
			want: []string{},
		},
		{
			name: "test empty source in array",
			args: args{
				s: []string{},
				t: []string{"a", "b"},
			},
			want: []string{"a", "b"},
		},
		{
			name: "test empty target in array",
			args: args{
				s: []string{"a", "b"},
				t: []string{},
			},
			want: []string{"a", "b"},
		},
		{
			name: "test source len > target len",
			args: args{
				s: []string{"a", "b", "c", "d", "e"},
				t: []string{"a", "c"},
			},
			want: []string{"b", "d", "e"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Diff(tt.args.s, tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Diff() = %v, want %v", got, tt.want)
			}
		})
	}
}
