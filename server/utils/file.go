package utils

import (
	"io"
	"os"
)

func CopyFile(src, dst string) error {
	// Open the source file
	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	// Create the destination file
	destFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destFile.Close()

	// Copy the contents from source file to destination file
	_, err = io.Copy(destFile, sourceFile)
	if err != nil {
		return err
	}

	// Ensure that the destination file has the same permissions as the source file
	sourceInfo, err := os.Stat(src)
	if err != nil {
		return err
	}
	err = os.Chmod(dst, sourceInfo.Mode())
	if err != nil {
		return err
	}

	return nil
}
