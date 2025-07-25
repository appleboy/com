package file

import (
	"os"
	"testing"
)

func TestIsDir(t *testing.T) {
	_ = os.Mkdir("testdir", os.ModeDir)

	type args struct {
		dir string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "check exist dir",
			args: args{
				dir: "testdir",
			},
			want: true,
		},
		{
			name: "dir not exist",
			args: args{
				dir: "testdir2",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := IsDir(tt.args.dir)
			if err != nil && tt.want {
				t.Errorf("IsDir() error = %v, want %v", err, tt.want)
			}
			if got != tt.want {
				t.Errorf("IsDir() = %v, want %v", got, tt.want)
			}
		})
	}

	// remove dir
	if err := Remove("testdir"); err != nil {
		t.Errorf("Remove() = %v", err)
	}
}

func TestIsFile(t *testing.T) {
	type args struct {
		filePath string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "file match",
			args: args{
				filePath: "file.go",
			},
			want: true,
		},
		{
			name: "file not found",
			args: args{
				filePath: "file1.go",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := IsFile(tt.args.filePath)
			if err != nil && tt.want {
				t.Errorf("IsFile() error = %v, want %v", err, tt.want)
			}
			if got != tt.want {
				t.Errorf("IsFile() = %v, want %v", got, tt.want)
			}
		})
	}
}
