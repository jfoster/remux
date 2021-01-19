package remux

import (
	"os"
	"path/filepath"
	"strings"
)

func TrimExt(path string) string {
	return strings.TrimSuffix(path, filepath.Ext(path))
}

func IsFileType(path string, filetype string) bool {
	if !strings.HasPrefix(filetype, ".") {
		filetype = "." + filetype
	}
	return filepath.Ext(path) == filetype
}

func IsVideo(path string) bool {
	ext := filepath.Ext(path)

	var filetypes = map[string]interface{}{
		".mkv": nil,
		".mov": nil,
	}

	_, ok := filetypes[ext]

	return ok
}

func IsMkv(path string) bool {
	return IsFileType(path, "mkv")
}

func IsDir(path string) bool {
	if stat, err := os.Stat(path); !os.IsNotExist(err) {
		return stat.IsDir()
	}
	return false
}
