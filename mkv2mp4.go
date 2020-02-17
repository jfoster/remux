package mkv2mp4

import (
	"errors"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
)

func Convert(in string) error {
	if filepath.Ext(in) != ".mkv" {
		return errors.New("input file is not an mkv")
	}
	out := trimExt(in) + ".mp4"

	cmd := exec.Command("ffmpeg",
		"-y",
		"-i", in,
		"-map", "0",
		"-c", "copy",
		out,
	)
	return cmd.Run()
}

func trimExt(p string) string {
	return strings.TrimSuffix(p, path.Ext(p))
}
