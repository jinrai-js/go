package tools

import (
	"os"
	"path/filepath"
)

func getBaseRoot(path string) string {
	exl, _ := os.Executable()
	exPath := filepath.Dir(exl)

	return exPath + "/" + path
}
