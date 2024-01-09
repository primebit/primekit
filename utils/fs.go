package utils

import (
	"os"
	"path/filepath"
)

func GetRootDir() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return "/"
	}
	return dir
}

func RealPath(path string) string {
	res, err := filepath.Abs(GetRootDir() + path)
	if err != nil {
		return path
	}
	return res
}
