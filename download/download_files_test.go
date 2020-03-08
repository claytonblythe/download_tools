package download

import (
	"strconv"
	"testing"
)

func TestDownloadFiles2(t *testing.T) {
	urls := []string{"http://google.com", "http://google.com"}
	_ = Download_urls(urls, 25)
}
func TestDownloadFiles3(t *testing.T) {
	urls := []string{"http://google.com", "http://google.com"}
	_ = Download_urls(urls, 1)
}

func TestDownloadFiles4(t *testing.T) {
	urls := []string{"http://www.cnn.com", "http://www.cnn.com"}
	_ = Download_urls(urls, 1)
}

func TestDownloadFiles5(t *testing.T) {
	urls := helperLoadUrls("slow_urls.txt")
	_ = Download_urls(urls, 205)
}

func BenchmarkDownloadFilesWorkers(b *testing.B) {
	urls := helperLoadUrls("urls.txt")
	for n := 1; n <= 500; n += 50 {
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

func BenchmarkDownloadFilesWorkersSlowUrls(b *testing.B) {
	urls := helperLoadUrls("slow_urls.txt")
	for n := 1; n < 370; n += 30 {
		b.Run(strconv.Itoa(n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = Download_urls(urls, n)
			}
		})
	}
}
