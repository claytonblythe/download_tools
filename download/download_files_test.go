package download

import (
	"testing"
)

/* func helperLoadBytes(t *testing.T, name string) []byte {
	path := filepath.Join("testdata", name) // relative path
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	return bytes
} */

func TestDownloadFiles(t *testing.T) {
	urls := []string{"http://www.deepython.com", "http://www.deepython.com"}

	_ = Download_urls(urls, 2)
}

func BenchmarkDownloadFiles(b *testing.B) {
	urls := []string{"http://www.deepython.com", "http://www.deepython.com"}

	for i := 0; i < b.N; i++ {
		_ = Download_urls(urls, 2)
	}
}
