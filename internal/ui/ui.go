package ui

import (
	"os"
	"path/filepath"

	"github.com/cheggaaa/pb/v3"
	"github.com/jfoster/remux"
)

const (
	template pb.ProgressBarTemplate = `{{string . "file"}} {{bar .}} {{percent .}} {{string . "time"}} {{string . "bitrate"}}`
)

func Copy2mp4(in string) error {
	bar := pb.New(10000)

	bar.SetTemplate(template)
	bar.SetWriter(os.Stdout)

	err := bar.Err()
	if err != nil {
		return err
	}

	bar.Start()
	defer bar.Finish()

	progress, err := remux.Copy2mp4(in)
	if err != nil {
		return err
	}

	for p := range progress {
		prog := p.GetProgress() * 100.0
		bar.SetCurrent(int64(prog))
		bar.Set("file", filepath.Base(in))
		bar.Set("time", p.GetCurrentTime())
		bar.Set("bitrate", p.GetCurrentBitrate())
	}
	bar.SetCurrent(10000)

	return os.Remove(in)
}
