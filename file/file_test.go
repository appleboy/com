package file

import (
	"fmt"
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

// TestFormatSize tests the FormatSize function for various input cases.
func TestFormatSize(t *testing.T) {
	tests := []struct {
		bytes    int64
		expected string
	}{
		{0, "0 B"},
		{9, "9 B"},
		{512, "512 B"},
		{1023, "1023 B"},
		{1024, "1.0 KB"},
		{1536, "1.5 KB"},
		{10 * 1024, "10.0 KB"},
		{1024*1024 - 1, "1024.0 KB"},
		{1024 * 1024, "1.0 MB"},
		{1536 * 1024, "1.5 MB"},
		{1024 * 1024 * 10, "10.0 MB"},
		{1024*1024*1024 - 1, "1024.0 MB"},
		{1024 * 1024 * 1024, "1.0 GB"},
		{1536 * 1024 * 1024, "1.5 GB"},
		{1024 * 1024 * 1024 * 10, "10.0 GB"},
		{1024 * 1024 * 1024 * 1024, "1.0 TB"},
		{1024 * 1024 * 1024 * 1024 * 1024, "1.0 PB"},
		{1024 * 1024 * 1024 * 1024 * 1024 * 1024, "1.0 EB"},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d bytes", tt.bytes), func(t *testing.T) {
			result := FormatSize(tt.bytes)
			if result != tt.expected {
				t.Errorf("FormatSize(%d) = %q; want %q", tt.bytes, result, tt.expected)
			}
		})
	}
}
