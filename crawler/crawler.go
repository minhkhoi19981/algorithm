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

	for i := 1; i <= 5; i++ {
		go func(name string) {
			for v := range queueChan {
				crawler(name, v)
			}
		}(fmt.Sprintf("%d", i))
	}

	for i := 1; i <= numberOfResquest; i++ {
		queueChan <- i
	}

	close(queueChan)
}
