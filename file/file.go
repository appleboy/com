package file

import (
	"fmt"
	"io"
	"os"
)

/*
IsDir returns (true, nil) if the given path is a directory.
Returns (false, nil) if it's a file.
Returns (false, error) if the path does not exist or other error occurs.
*/
func IsDir(dir string) (bool, error) {
	f, err := os.Stat(dir)
	if err != nil {
		return false, err
	}
	return f.IsDir(), nil
}

// IsFile returns (true, nil) if given path is a file.
// Returns (false, nil) if it's a directory.
// Returns (false, error) if the path does not exist or other error occurs.
func IsFile(filePath string) (bool, error) {
	f, err := os.Stat(filePath)
	if err != nil {
		return false, err
	}
	return !f.IsDir(), nil
}

// Remove removes the file or directory at the given path, including any children if it's a directory.
// If the path does not exist, Remove returns nil (no error).
// If there is an error, it will be of type *PathError.
func Remove(filePath string) error {
	return os.RemoveAll(filePath)
}

// Copy files
// Copy copies a regular file from src to dst. If dst exists, returns error.
// Uses io.Copy for efficient file transfer.
func Copy(src, dst string) error {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return err
	}
	if !sourceFileStat.Mode().IsRegular() {
		return fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return err
	}
	defer source.Close()

	if _, err := os.Stat(dst); err == nil {
		return fmt.Errorf("file %s already exists", dst)
	}

	destination, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destination.Close()

	if _, err := io.Copy(destination, source); err != nil {
		return err
	}
	return nil
}
