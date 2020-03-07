package download

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
)

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func helperLoadUrls(name string) []string {
	path := filepath.Join("testdata", name) // relative path
	urls, err := readLines(path)
	if err != nil {
		log.Fatal(err)
	}
	return urls
}
