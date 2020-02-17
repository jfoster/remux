package mkv2mp4

import (
	"errors"
	"fmt"
	"path"
	"path/filepath"
	"strings"

	"github.com/jfoster/goffmpeg/transcoder"
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
	for msg := range progress {
		fmt.Println(msg)
	}

	return <-done
}

func trimExt(p string) string {
	return strings.TrimSuffix(p, path.Ext(p))
}
