package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/jfoster/mkv2mp4"
)

func main() {
	flag.Parse()

	args := flag.Args()

	if flag.NArg() < 1 {
		fmt.Println("no files specified")
	}

	if _, err := exec.LookPath("ffmpeg"); err != nil {
		log.Fatal(err)
	}

	for _, in := range args {
		err := mkv2mp4.Convert(in)
		if err == nil {
			os.Remove(in)
		} else {
			fmt.Println(err)
		}
	}
}
