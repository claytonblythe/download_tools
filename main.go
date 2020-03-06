package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan string, results chan<- string) {
	for job := range jobs {
		fmt.Println("worker", id, "started  job", job)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", job)
		results <- job
	}
}

func main() {

	const numJobs = 5
	const numWorkers = 3
	jobs := make(chan string, numJobs)
	results := make(chan string, numJobs)
	urls := [8]string{"test1", "test2", "test3", "test4", "test5", "test6", "test7", "test8"}

	for w := 0; w < numWorkers; w++ {
		go worker(w, jobs, results)
	}

	for _, url := range urls {
		jobs <- url
	}
	close(jobs)

	for a := 1; a <= numJobs; a++ {
		<-results
	}
}
