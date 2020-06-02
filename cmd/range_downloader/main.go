package main

import (
	"github.com/webx-top/com"

	"flag"
	"log"
	nurl "net/url"
	"strings"
)

func main() {
	url := flag.String("url", "http://storage.googleapis.com/vimeo-test/work-at-vimeo.mp4", "URL for download")
	threads := flag.Int("threads", 10, "Number of threads to download with")
	flag.Parse()

	u, err := nurl.Parse(*url)
	if err != nil {
		log.Fatal(err)
		return
	}
	saveTo := strings.Replace(u.Path[1:], "/", "-", -1)
	err = com.RangeDownload(*url, saveTo, *threads)
	if err != nil {
		log.Println(err)
	}
}
