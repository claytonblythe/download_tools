package download

import (
	"testing"
)

func TestDownloadFiles(t *testing.T) {
	urls := helperLoadUrls("urls.txt")
	_ = Download_urls(urls, 2)
}

func BenchmarkDownloadFilesOneWorker(b *testing.B) {
	urls := helperLoadUrls("urls.txt")

	for i := 0; i < b.N; i++ {
		_ = Download_urls(urls, 1)
	}
}

func BenchmarkDownloadFilesTenWorker(b *testing.B) {
	urls := helperLoadUrls("urls.txt")

	for i := 0; i < b.N; i++ {
		_ = Download_urls(urls, 10)
	}
}

func BenchmarkDownloadFilesThirtyWorker(b *testing.B) {
	urls := helperLoadUrls("urls.txt")

	for i := 0; i < b.N; i++ {
		_ = Download_urls(urls, 30)
	}
}

func BenchmarkDownloadFilesSixtyWorker(b *testing.B) {
	urls := helperLoadUrls("urls.txt")

	for i := 0; i < b.N; i++ {
		_ = Download_urls(urls, 60)
	}
}

func BenchmarkDownloadFilesHundredWorker(b *testing.B) {
	urls := helperLoadUrls("urls.txt")

	for i := 0; i < b.N; i++ {
		_ = Download_urls(urls, 100)
	}
}

func BenchmarkDownloadFilesTwoHundredWorker(b *testing.B) {
	urls := helperLoadUrls("urls.txt")

	for i := 0; i < b.N; i++ {
		_ = Download_urls(urls, 200)
	}
}
