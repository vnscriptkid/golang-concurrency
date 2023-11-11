package main

import (
	"fmt"
	"math/rand"
	"sync"
)

type Job struct {
	id       int
	randomno int
}

type Result struct {
	job              Job
	sumofdigits, wId int
}

// calculate sum of digits and return int
func sumOfDigits(n int) int {
	sum := 0
	for n > 0 {
		digit := n % 10
		sum += digit
		n = n / 10
	}
	return sum
}

// create worker pool which passes result struct into the results channel
// use wait group
func createWorkerPool(noOfWorkers int, jobs chan Job, results chan Result) {
	var wg sync.WaitGroup
	for i := 0; i < noOfWorkers; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			for job := range jobs {
				results <- Result{job: job, wId: i, sumofdigits: sumOfDigits(job.randomno)}
			}
		}(i)
	}
	wg.Wait()
	close(results)
}

// allocate the jobs into the jobs channel
func allocate(noOfJobs int, jobs chan Job) {
	defer close(jobs)
	for i := 0; i < noOfJobs; i++ {
		x := rand.Intn(1000)
		jobs <- Job{id: i, randomno: x}
	}
}

func main() {
	// initialize the channels
	jobs := make(chan Job, 50)
	results := make(chan Result, 10)

	// initialize number of jobs
	numOfJobs := 50

	// call the allocate function
	go allocate(numOfJobs, jobs)

	// initialize number of worker
	numOfWorkers := 10

	// call the worker pool to operate on the jobs
	go createWorkerPool(numOfWorkers, jobs, results)

	// print out the results
	for result := range results {
		fmt.Printf("Job id %d, input random no %d , sum of digits %d, worker id %d\n",
			result.job.id, result.job.randomno, result.sumofdigits, result.wId)
	}
}
