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
