package main

import (
	"bufio"
	"sync"
)

func worker(wg *sync.WaitGroup) {
	for job := range jobs {
		output := Result{job, getLastSyllable(job.line)}
		results <- output
	}
	wg.Done()
}

func createWorkerPool(numWorkers int) {
	var wg sync.WaitGroup
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(&wg)
	}
	wg.Wait()
	close(results)
}

func allocate(scanner *bufio.Scanner) {
	i := 0
	for scanner.Scan() {
		job := Job{i, alphaNumeric(scanner.Text())}
		jobs <- job
		i++
	}
	close(jobs)
}

func result(done chan bool) {
	for result := range results {
		lastSyllables[result.job.id] = result.lastSyllable
	}
	done <- true
}
