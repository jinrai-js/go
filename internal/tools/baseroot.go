package tools

import (
	"os"
	"path/filepath"
)

func GetBaseRoot(path string) string {
	absPath, err := filepath.Abs(path)
	if err == nil {
		return absPath
	}

	exl, _ := os.Executable()
	exPath := filepath.Dir(exl)

	return exPath + "/" + path
}
