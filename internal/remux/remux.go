package remux

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/cheggaaa/pb/v3"
	"github.com/floostack/transcoder"
	"github.com/floostack/transcoder/ffmpeg"
)

const (
	template pb.ProgressBarTemplate = `{{string . "file"}} {{bar .}} {{percent .}} {{string . "time"}} {{string . "bitrate"}}`
)

func Convert(in string) error {
	bar := pb.New(10000)

	bar.SetTemplate(template)
	bar.SetWriter(os.Stdout)

	err := bar.Err()
	if err != nil {
		return err
	}

	bar.Start()
	defer bar.Finish()

	progress, err := Mkv2mp4(in)
	if err != nil {
		return err
	}

	for p := range progress {
		prog := p.GetProgress() * 100.0
		bar.SetCurrent(int64(prog))
		bar.Set("file", filepath.Base(""))
		bar.Set("time", p.GetCurrentTime())
		bar.Set("bitrate", p.GetCurrentBitrate())
	}
	bar.SetCurrent(10000)

	return os.Remove(in)
}

func Mkv2mp4(in string) (<-chan transcoder.Progress, error) {
	if filepath.Ext(in) != ".mkv" {
		return nil, fmt.Errorf("input file (%s) is not an mkv", in)
	}

	format := "mp4"
	overwrite := true
	codec := "copy"

	args := map[string]interface{}{"map": 0}
	opts := ffmpeg.Options{
		VideoCodec:   &codec,
		AudioCodec:   &codec,
		OutputFormat: &format,
		Overwrite:    &overwrite,
		ExtraArgs:    args,
	}

	cfg := &ffmpeg.Config{
		FfmpegBinPath:   "ffmpeg",
		FfprobeBinPath:  "ffprobe",
		ProgressEnabled: true,
	}

	out := TrimExt(in) + "." + format

	return ffmpeg.New(cfg).Input(in).Output(out).Start(opts)
}

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
