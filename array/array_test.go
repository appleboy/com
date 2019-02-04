package array

import (
	"reflect"
	"strconv"
	"testing"
)

func BenchmarkArrayMap(b *testing.B) {
	newString := []string{}
	for i := 0; i < 100; i++ {
		newString = append(newString, strconv.Itoa(i))
	}
	for i := 0; i < b.N; i++ {
		InMap("99", newString)
	}
}

func BenchmarkArraySlice(b *testing.B) {
	newString := []string{}
	for i := 0; i < 100; i++ {
		newString = append(newString, strconv.Itoa(i))
	}
	for i := 0; i < b.N; i++ {
		InSlice("99", newString)
	}
}

func BenchmarkIn(b *testing.B) {
	newString := []string{}
	for i := 0; i < 100; i++ {
		newString = append(newString, strconv.Itoa(i))
	}
	for i := 0; i < b.N; i++ {
		In("99", newString)
	}
}

func BenchmarkInArray(b *testing.B) {
	newString := []string{}
	for i := 0; i < 100; i++ {
		newString = append(newString, strconv.Itoa(i))
	}
	for i := 0; i < b.N; i++ {
		InArray("99", newString)
	}
}

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
			name: "test empty source and target",
			args: args{
				s: []string{},
				t: []string{},
			},
			want: []string{},
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

func TestInValue(t *testing.T) {
	s := []string{"59bf2170ceadf87a1e7e1ab4", "59bf2170ceadf87a1e7e1ab5", "5a2899f460faae1623882b5b"}
	tt := "59bf2170ceadf87a1e7e1ab4"
	want := []string{"59bf2170ceadf87a1e7e1ab5", "5a2899f460faae1623882b5b"}
	wantOK := true
	wantS := []string{"59bf2170ceadf87a1e7e1ab4", "59bf2170ceadf87a1e7e1ab5", "5a2899f460faae1623882b5b"}

	got, ok := In(tt, s)

	if ok != wantOK {
		t.Errorf("bool = %v, want %v", ok, wantOK)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("In() = %v, want %v", got, want)
	}

	if !reflect.DeepEqual(s, wantS) {
		t.Errorf("s = %v, want %v", s, wantS)
	}
}

func TestDiffValue(t *testing.T) {
	s := []string{"59bf2170ceadf87a1e7e1ab4", "59bf2170ceadf87a1e7e1ab5", "5a2899f460faae1623882b5b"}
	tt := []string{"59bf2170ceadf87a1e7e1ab4", "5a2899f460faae1623882b5b", "59bf2170ceadf87a1e7e1ab5"}
	want := []string{}
	wantS := []string{"59bf2170ceadf87a1e7e1ab4", "59bf2170ceadf87a1e7e1ab5", "5a2899f460faae1623882b5b"}
	wantT := []string{"59bf2170ceadf87a1e7e1ab4", "5a2899f460faae1623882b5b", "59bf2170ceadf87a1e7e1ab5"}

	got := Diff(s, tt)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Diff() = %v, want %v", got, want)
	}

	if !reflect.DeepEqual(s, wantS) {
		t.Errorf("s = %v, want %v", s, wantS)
	}

	if !reflect.DeepEqual(tt, wantT) {
		t.Errorf("tt = %v, want %v", t, wantT)
	}
}
