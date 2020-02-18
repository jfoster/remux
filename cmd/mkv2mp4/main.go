package main

import (
	"flag"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/jfoster/mkv2mp4"
)

func main() {
	flag.Parse()

	// check for ffmpeg
	if _, err := exec.LookPath("ffmpeg"); err != nil {
		log.Fatalln("ffmpeg not found", err)
	}

	paths := flag.Args()

	var files []string

	if len(paths) < 1 {
		dir, err := os.Getwd()
		if err != nil {
			log.Print(err)
		}
		paths = []string{dir}
	}

	for _, path := range paths {
		if mkv2mp4.IsDir(path) {
			err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
				if mkv2mp4.IsMkv(path) {
					files = append(files, path)
				}
				return nil
			})
			if err != nil {
				log.Panic(err)
			}
		} else if mkv2mp4.IsMkv(path) {
			files = append(files, path)
		}
	}

	for _, file := range files {
		convert(file)
	}
}

func convert(in string) {
	if err := mkv2mp4.Convert(in); err == nil {
		os.Remove(in)
	} else {
		log.Print(err)
	}
}
