package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Each Job struct has a id and a randomno for which the sum of the individual digits has to be computed.
type Job struct {
	id       int
	randomno int // 123
}

// The Result struct has a job field which is the job for which it holds the result (sum of individual digits) in the sumofdigits field.
type Result struct {
	job         Job
	sumofdigits int // 6
}

var jobs = make(chan Job, 10)
var results = make(chan Result, 10)

// The allocate function below takes the number of jobs to be created
func allocate(noOfJobs int) {
	for i := 0; i < noOfJobs; i++ {
		randomno := rand.Intn(999)
		job := Job{i, randomno}
		jobs <- job
	}
	close(jobs)
}

func result(done chan bool) {
	for result := range results {
		fmt.Printf("Job id %d, input random no %d , sum of digits %d\n", result.job.id, result.job.randomno, result.sumofdigits)
	}
	done <- true
}

func digits1(number int) int {
	sum := 0
	no := number
	for no != 0 {
		digit := no % 10
		sum += digit
		no /= 10
	}
	time.Sleep(2 * time.Second)
	return sum
}

func worker(wg *sync.WaitGroup) {
	for job := range jobs {
		output := Result{job, digits1(job.randomno)}
		results <- output
	}
	wg.Done()
}

func createWorkerPool(noOfWorkers int) {
	var wg sync.WaitGroup
	for i := 0; i < noOfWorkers; i++ {
		wg.Add(1)
		go worker(&wg)
	}
	wg.Wait()
	close(results)
}

func main() {
	noOfWorkers := 10
	startTime := time.Now()
	noOfJobs := 100
	go allocate(noOfJobs)
	done := make(chan bool)
	go result(done)
	createWorkerPool(noOfWorkers) // allocate)
	<-done
	endTime := time.Now()
	diff := endTime.Sub(startTime)
	fmt.Println("total time taken ", diff.Seconds(), "seconds")
}
