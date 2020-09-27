package main

import (
	"fmt"
	"time"
)

/*
- param jobs: chan int. nhận message từ chan int
- param results: chan int. gửi message tới chan int
*/
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

func main() {

	const numJobs = 5
	// tạo channel
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	for w := 1; w <= 3; w++ {
		//
		go worker(w, jobs, results)
	}
	// send data tới channel jobs
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	// xuat data từ channel results
	for a := 1; a <= numJobs; a++ {
		<-results
	}
}
