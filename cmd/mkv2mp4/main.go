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
		paths[0] = dir
	}

	for _, path := range paths {
		if isDir(path) {
			err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
				if isMkv(path) {
					files = append(files, path)
				}
				return nil
			})
			if err != nil {
				log.Panic(err)
			}
		} else if isMkv(path) {
			files = append(files, path)
		}
	}

	fmt.Println(files)

	for _, file := range files {
		convert(file)
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

func isMkv(path string) bool {
	return filepath.Ext(path) == ".mkv"
}

func isDir(path string) bool {
	if stat, err := os.Stat(path); !os.IsNotExist(err) {
		return stat.IsDir()
	}
	return false
}
