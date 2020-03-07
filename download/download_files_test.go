package download

import "testing"

func TestDownloadFiles(t *testing.T) {
	urls := []string{"http://deepython.com", "http://deepython.com"}
	_ = Download_urls(urls, 2)
}
