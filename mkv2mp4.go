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
		"-map", "0:0?",
		"-map", "0:1?",
		"-map", "0:2?",
		"-map", "0:3?",
		"-map", "0:4?",
		"-map", "0:5?",
		"-map", "0:7?",
		"-c:v", "copy",
		"-c:a:0", "copy",
		"-c:a:1", "copy",
		"-c:a:2", "copy",
		"-c:a:3", "copy",
		"-c:a:4", "copy",
		"-c:a:5", "copy",
		"-c:a:6", "copy",
		out,
	)
	return cmd.Run()
}

func trimExt(p string) string {
	return strings.TrimSuffix(p, path.Ext(p))
}
