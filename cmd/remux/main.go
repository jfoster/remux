package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/jfoster/remux"
	"github.com/jfoster/remux/internal/ui"
)

func main() {
	flag.Parse()
	args := flag.Args()

	if len(args) < 1 {
		// if there's no args passed then get the current working directory
		dir, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
		// and add it to the args
		args = []string{dir}
	}

	for _, v := range args {
		// walk through each arg
		err := filepath.Walk(v, func(path string, info os.FileInfo, err error) error {
			fmt.Println(path)
			if !remux.IsDir(path) && remux.IsVideo(path) {
				return ui.Copy2mp4(path)
			}
			return nil
		})
		if err != nil {
			log.Fatal(err)
		}
	}
}
