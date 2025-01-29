package convert

import (
	"reflect"
	"testing"
)

func TestToString(t *testing.T) {
	type args struct {
		value interface{}
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{
			name: "int",
			args: args{
				value: 100,
			},
			want: "100",
		},
		{
			name: "int64",
			args: args{
				value: int64(100),
			},
			want: "100",
		},
		{
			name: "boolean",
			args: args{
				value: true,
			},
			want: "true",
		},
		{
			name: "float32",
			args: args{
				value: float32(23.03),
			},
			want: "23.03",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToString(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToBool(t *testing.T) {
	type args struct {
		value interface{}
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{
			name: "int",
			args: args{
				value: 100,
			},
			want: true,
		},
		{
			name: "int",
			args: args{
				value: 0,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToBool(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToBool() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToInt(t *testing.T) {
	type args struct {
		value interface{}
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{
			name: "bool true",
			args: args{
				value: true,
			},
			want: 1,
		},
		{
			name: "bool false",
			args: args{
				value: false,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToInt(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToFloat(t *testing.T) {
	type args struct {
		value interface{}
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{
			name: "boolean true",
			args: args{
				value: true,
			},
			want: 1.0,
		},
		{
			name: "boolean false",
			args: args{
				value: false,
			},
			want: 0.0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToFloat(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToFloat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConvertBig5ToUTF8(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "Valid Big5 string",
			input: "\xa7A\xa6n",
			want:  "你好",
		},
		{
			name:  "Mixed valid and invalid Big5 string",
			input: "\xa7A\xa6n\xff\xfe",
			want:  "你好\ufffd\ufffd",
		},
		{
			name:  "Invalid Big5 string",
			input: "\xff\xfe",
			want:  "\ufffd\ufffd",
		},
		{
			name:  "Empty string",
			input: "",
			want:  "",
		},
		{
			name:  "ASCII string",
			input: "Hello, World!",
			want:  "Hello, World!",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConvertBig5ToUTF8(tt.input); got != tt.want {
				t.Errorf("ConvertBig5ToUTF8() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFromPtr(t *testing.T) {
	tests := []struct {
		name string
		ptr  interface{}
		want interface{}
	}{
		{
			name: "int pointer",
			ptr:  ToPtr(100),
			want: 100,
		},
		{
			name: "nil int pointer",
			ptr:  (*int)(nil),
			want: 0,
		},
		{
			name: "string pointer",
			ptr:  ToPtr("hello"),
			want: "hello",
		},
		{
			name: "nil string pointer",
			ptr:  (*string)(nil),
			want: "",
		},
		{
			name: "bool pointer",
			ptr:  ToPtr(true),
			want: true,
		},
		{
			name: "nil bool pointer",
			ptr:  (*bool)(nil),
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			switch ptr := tt.ptr.(type) {
			case *int:
				if got := FromPtr(ptr); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("FromPtr() = %v, want %v", got, tt.want)
				}
			case *string:
				if got := FromPtr(ptr); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("FromPtr() = %v, want %v", got, tt.want)
				}
			case *bool:
				if got := FromPtr(ptr); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("FromPtr() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}
