package download

import (
	"strconv"
	"testing"
)

func TestDownloadFiles(t *testing.T) {
	urls := helperLoadUrls("urls.txt")
	_ = Download_urls(urls, 25)
}

func BenchmarkDownloadFilesWorkers(b *testing.B) {
	urls := helperLoadUrls("urls.txt")
	for n := 10; n <= 250; n += 30 {
		b.Run(strconv.Itoa(n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = Download_urls(urls, n)
			}
		})
	}
}

func BenchmarkDownloadFilesWorkersImages(b *testing.B) {
	urls := helperLoadUrls("image_urls.txt")
	for n := 1; n < 200; n += 30 {
		b.Run(strconv.Itoa(n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = Download_urls(urls, n)
			}
		})
	}
}
