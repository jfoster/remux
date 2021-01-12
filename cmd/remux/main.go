package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"

	"github.com/jfoster/remux/internal/remux"
)

func main() {
	flag.Parse()
	args := flag.Args()

	if len(args) < 1 {
		wd, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
		args = []string{wd}
	}

	for _, path := range args {
		filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
			if remux.IsMkv(path) {
				return remux.Convert(path)
			}
			return nil
		})
	}
}
