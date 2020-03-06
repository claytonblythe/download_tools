package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func worker(id int, urls <-chan struct {
	int
	string
}, results chan struct {
	int
	string
}) {
	for url := range urls {
		fmt.Println("worker", id, "started  url", url.string)
		// don't worry about errors
		response, e := http.Get(url.string)
		if e != nil {
			log.Fatal(e)
		}
		defer response.Body.Close()

		//open a file for writing
		filename := fmt.Sprintf("%06d", url.int)
		myFile, err := ioutil.TempFile("/tmp/image", filename)
		if err != nil {
			log.Fatal(err)
		}
		defer myFile.Close()

		// Use io.Copy to just dump the response body to the file. This supports huge files
		_, err = io.Copy(myFile, response.Body)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Success!")
		fmt.Println("worker", id, "finished url", url.string)
		results <- struct {
			int
			string
		}{url.int, url.string}
	}
}

func main() {

	const numJobs = 18
	const numWorkers = 18
	jobs := make(chan struct {
		int
		string
	}, numJobs)
	results := make(chan struct {
		int
		string
	}, numJobs)

	for w := 0; w < numWorkers; w++ {
		go worker(w, jobs, results)
	}

	for n := 1; n < 19; n++ {
		page := strconv.Itoa(n)
		var str strings.Builder
		str.WriteString("https://digital.olivesoftware.com/Olive/ODN/FTUS/get/image.ashx?kind=page&href=FIT%2F2020%2F03%2F05&page=")
		str.WriteString(page)
		str.WriteString("&res=120")
		jobs <- struct {
			int
			string
		}{n, str.String()}
	}
	close(jobs)

	for a := 0; a < numJobs; a++ {
		<-results
	}
}
