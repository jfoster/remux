package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/jfoster/mkv2mp4"
)

func main() {
	flag.Parse()
	args := flag.Args()

	// check for ffmpeg
	if _, err := exec.LookPath("ffmpeg"); err != nil {
		log.Fatal(err)
	}

	var files []string

	if flag.NArg() > 0 {
		for _, v := range args {
			if isMkv(v) {
				files = append(files, v)
			}
		}
	} else {
		dir, err := os.Getwd()
		if err != nil {
			log.Print(err)
		}
		fmt.Println(dir)

		err = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
			if isMkv(path) {
				files = append(files, path)
			}
			return nil
		})
		if err != nil {
			log.Panic(err)
		}
	}

	for _, in := range files {
		convert(in)
	}
}

func convert(in string) {
	fmt.Println("started converting", in)
	if err := mkv2mp4.Convert(in); err == nil {
		fmt.Println("finished converting", in)
		os.Remove(in)
	} else {
		log.Print(err)
	}
}

func isMkv(in string) bool {
	return filepath.Ext(in) == ".mkv"
}
