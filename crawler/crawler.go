package crawler

import (
	"fmt"
	"time"
)

func crawler(name string, v int) {
	time.Sleep(time.Second / 100)
	fmt.Printf("Worker %s is crawling: %d\n", name, v)
}

func Main() {
	fmt.Printf("Crawler url........")
	numberOfResquest := 10000

	queueChan := make(chan int, numberOfResquest)
	numberOfMaxResquest := 5
	doneChan := make(chan string)

	for i := 1; i <= numberOfMaxResquest; i++ {
		go func(name string) {
			for v := range queueChan {
				crawler(name, v)
			}
			fmt.Printf("done channel %s \n", name)
			doneChan <- "done"
		}(fmt.Sprintf("%d", i))
	}

	for i := 1; i <= numberOfResquest; i++ {
		queueChan <- i
	}

	close(queueChan)

	for i := 1; i <= numberOfMaxResquest; i++ {
		<-doneChan
	}
}
