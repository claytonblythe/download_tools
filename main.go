package main

import (
	_ "net/http/pprof"

	"github.com/claytonblythe/download_tools/download"
)

func main() {
	// Toy example with some urls to hit concurrently
	urls := []string{
		"http://www.deepython.com/",
		"http://www.deepython.com/",
		"http://www.deepython.com/",
		"http://www.deepython.com/",
	}

	final_urls := []string{}
	for n := 0; n < 50; n++ {
		final_urls = append(final_urls, urls...)
	}

	_ = download.Download_urls(final_urls, 30)
}
