package mkv2mp4

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/cheggaaa/pb/v3"
	"github.com/jfoster/goffmpeg/transcoder"
)

const (
	FFTemplate pb.ProgressBarTemplate = `{{string . "file"}} {{bar .}} {{percent .}} {{string . "time"}} {{string . "bitrate"}}`
)

func Convert(in string) error {
	if filepath.Ext(in) != ".mkv" {
		return errors.New("input file is not an mkv")
	}
	out := trimExt(in) + ".mp4"

	trans := new(transcoder.Transcoder)

	if err := trans.Initialize(in, out); err != nil {
		return err
	}

	trans.MediaFile().SetStreamMap("0")
	trans.MediaFile().SetAudioCodec("copy")
	trans.MediaFile().SetVideoCodec("copy")

	done := trans.Run(true)

	progress := trans.Output()

	bar := pb.New(10000)

	bar.SetTemplate(FFTemplate)
	bar.SetWriter(os.Stdout)

	bar.Start()

	for prog := range progress {
		bar.SetCurrent(int64(prog.Progress * 100.0))
		bar.Set("file", filepath.Base(in))
		bar.Set("time", prog.CurrentTime)
		bar.Set("bitrate", prog.CurrentBitrate)
	}

	bar.SetCurrent(10000).Finish()

	return <-done
}
