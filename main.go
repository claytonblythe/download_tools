package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func worker(id int, urls <-chan string, results chan<- string) {
	for url := range urls {
		fmt.Println("worker", id, "started  url", url)
		// don't worry about errors
		response, e := http.Get(url)
		if e != nil {
			log.Fatal(e)
		}
		defer response.Body.Close()

		//open a file for writing
		myFile, err := ioutil.TempFile("/tmp", "mypattern")
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
		fmt.Println("worker", id, "finished url", urls)
		results <- url
	}
}

func main() {

	const numJobs = 5
	const numWorkers = 3
	jobs := make(chan string, numJobs)
	results := make(chan string, numJobs)
	urls := [5]string{"https://digital.olivesoftware.com/Olive/ODN/FTUS/get/image.ashx?kind=page&href=FIT%2F2020%2F03%2F05&page=1&res=120", "https://digital.olivesoftware.com/Olive/ODN/FTUS/get/image.ashx?kind=page&href=FIT%2F2020%2F03%2F05&page=1&res=120", "https://digital.olivesoftware.com/Olive/ODN/FTUS/get/image.ashx?kind=page&href=FIT%2F2020%2F03%2F05&page=1&res=120", "https://digital.olivesoftware.com/Olive/ODN/FTUS/get/image.ashx?kind=page&href=FIT%2F2020%2F03%2F05&page=1&res=120", "https://digital.olivesoftware.com/Olive/ODN/FTUS/get/image.ashx?kind=page&href=FIT%2F2020%2F03%2F05&page=1&res=120"}

	for w := 0; w < numWorkers; w++ {
		go worker(w, jobs, results)
	}

	for _, url := range urls {
		jobs <- url
	}
	close(jobs)

	for a := 0; a < numJobs; a++ {
		<-results
	}
}
