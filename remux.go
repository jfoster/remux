package remux

import (
	"github.com/floostack/transcoder"
	"github.com/floostack/transcoder/ffmpeg"
)

func Copy2mp4(in string) (<-chan transcoder.Progress, error) {

	format := "mp4"
	overwrite := true
	codec := "copy"

	opts := ffmpeg.Options{
		VideoCodec:   &codec,
		AudioCodec:   &codec,
		OutputFormat: &format,
		Overwrite:    &overwrite,
		ExtraArgs: map[string]interface{}{
			"-map": 0,
		},
	}

	cfg := &ffmpeg.Config{
		FfmpegBinPath:   "ffmpeg",
		FfprobeBinPath:  "ffprobe",
		ProgressEnabled: true,
	}

	out := TrimExt(in) + "." + format

	return ffmpeg.New(cfg).Input(in).Output(out).Start(opts)
}
