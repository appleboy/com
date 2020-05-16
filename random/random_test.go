package random

import (
	"math/rand"
	"testing"
	"time"
)

func TestStringWithCharset(t *testing.T) {
	type args struct {
		length  int
		charset Charset
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test Rand",
			args: args{
				length:  5,
				charset: "a",
			},
			want: "aaaaa",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringWithCharset(tt.args.length, tt.args.charset); got != tt.want {
				t.Errorf("stringWithCharset() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestString(t *testing.T) {
	type args struct {
		length int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test Rand",
			args: args{
				length: 5,
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := String(tt.args.length); len(got) != tt.want {
				t.Errorf("String() = %v, want %v", len(got), tt.want)
			}
		})
	}
}

var seededRand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func stringWithCharsetOld(length int, charset Charset) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func BenchmarkRandStringOld(b *testing.B) {
	for n := 0; n < b.N; n++ {
		stringWithCharsetOld(16, Alphanumeric)
	}
}

func BenchmarkRandStringNew(b *testing.B) {
	for n := 0; n < b.N; n++ {
		StringWithCharset(16, Alphanumeric)
	}
}
