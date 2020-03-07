package download

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func worker(worker_id int, jobs <-chan struct {
	int
	string
}, results chan struct {
	int
	string
}) {

	for job := range jobs {
		// fmt.Println("worker", worker_id, "started  url", job.string)
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
		// fmt.Println("Success!")
		// fmt.Println("worker", worker_id, "finished url", job.string)
		results <- struct {
			int
			string
		}{job.int, myFile.Name()}
	}
}

func Download_urls(urls []string, num_workers int) []string {

	num_jobs := len(urls)

	jobs := make(chan struct {
		int
		string
	}, num_jobs*2)
	results := make(chan struct {
		int
		string
	}, num_jobs*2)

	for worker_id := 0; worker_id < num_workers; worker_id++ {
		go worker(worker_id, jobs, results)
	}

	for page_index, url := range urls {
		job := struct {
			int
			string
		}{page_index, url}
		jobs <- job
	}
	close(jobs)

	filepaths := []string{}
	for a := 0; a < num_jobs; a++ {
		result := <-results
		filepaths = append(filepaths, result.string)
	}
	return filepaths
}
