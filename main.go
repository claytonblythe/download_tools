package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func worker(id int, jobs <-chan struct {
	int
	string
}, results chan struct {
	int
	string
}) {
	for job := range jobs {
		fmt.Println("worker", id, "started  url", job.string)
		// don't worry about errors
		response, e := http.Get(job.string)
		if e != nil {
			log.Fatal(e)
		}
		defer response.Body.Close()

		//open a file for writing
		filename := fmt.Sprintf("%06d", job.int)
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
		fmt.Println("worker", id, "finished url", job.string)
		results <- struct {
			int
			string
		}{job.int, filename}
	}
}

func main() {
	urls := []string{"https://digital.olivesoftware.com/Olive/ODN/FTUS/get/image.ashx?kind=page&href=FIT%2F2020%2F03%2F05&page=14&res=120", "https://digital.olivesoftware.com/Olive/ODN/FTUS/get/image.ashx?kind=page&href=FIT%2F2020%2F03%2F05&page=14&res=120"}
	const num_jobs = 2
	const num_workers = 2
	jobs := make(chan struct {
		int
		string
	}, num_jobs)
	results := make(chan struct {
		int
		string
	}, num_jobs)

	for w := 0; w < num_workers; w++ {
		go worker(w, jobs, results)
	}

	for page_index, url := range urls {
		job := struct {
			int
			string
		}{page_index, url}
		jobs <- job
	}
	close(jobs)

	for a := 0; a < num_jobs; a++ {
		<-results
	}
}
