package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jfoster/remux/internal/ui"
	"github.com/radovskyb/watcher"
)

func main() {
	flag.Parse()

	w := watcher.New()
	w.FilterOps(watcher.Create, watcher.Write)

	go func() {
		for {
			timer := time.NewTimer(5 * time.Second)
			var file string

			select {
			case evt := <-w.Event:
				fmt.Println(evt)
				if !evt.IsDir() && evt.Op == watcher.Create {
					file = evt.Path
				}
			case err := <-w.Error:
				log.Println("error:", err)
			case <-timer.C:
				fmt.Println("timer stopped")
				go ui.Copy2mp4(file)
			}

			timer.Stop()
		}
	}()

	dir, err := os.Getwd()
	if err != nil {
		log.Print(err)
	}

	if flag.NArg() > 0 {
		dir = flag.Arg(0)
	}

	if err := w.Add(dir); err != nil {
		log.Fatal(err)
	}

	if err := w.Start(time.Nanosecond); err != nil {
		log.Fatal(err)
	}
}
