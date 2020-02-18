package mkv2mp4

import (
	"os"
	"path/filepath"
	"strings"
)

func TrimExt(p string) string {
	return strings.TrimSuffix(p, filepath.Ext(p))
}

func IsMkv(path string) bool {
	return filepath.Ext(path) == ".mkv"
}

func IsDir(path string) bool {
	if stat, err := os.Stat(path); !os.IsNotExist(err) {
		return stat.IsDir()
	}
	return false
}
