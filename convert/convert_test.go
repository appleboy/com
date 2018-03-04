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
