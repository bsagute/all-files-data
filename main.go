package main

import (
	"fmt"
	"sync"
	"time"
)

const (
	totalItems = 20 // Total number of items to produce
	rateLimit  = 2  // Consumer can process 2 items per second
)

var ch = make(chan string)
var wg = sync.WaitGroup{}

// Producer sends items to the channel
func producer(items chan<- int) {
	for i := 1; i <= totalItems; i++ {
		fmt.Printf("Produced item: %d\n", i)
		items <- i
		time.Sleep(100 * time.Millisecond) // Simulating burst production
	}
	close(items)
}

// Consumer processes items at a rate-limited speed
func consumer(items <-chan int) {
	limiter := time.Tick(time.Second / rateLimit) // Rate limit: 2 items per second

	for item := range items {
		<-limiter // Consume items at a fixed rate
		fmt.Printf("Consumed item: %d\n", item)
	}
}

// Reverses strings in a list
func reverseStrings(ipList []string) {
	wg.Add(len(ipList))
	go func() {
		defer close(ch)
		for _, v := range ipList {
			go DoRev(v)
		}
	}()
	wg.Wait()
	<-ch
}

// Reverse individual string and stop for strings of length 6
func DoRev(inputString string) {
	defer wg.Done()

	byteArr := []byte(inputString)
	if len(byteArr) == 6 {
		ch <- "stop"
		return
	}
	for i, j := 0, len(byteArr)-1; i < len(byteArr)-1; i++ {
		byteArr[i], byteArr[j] = byteArr[j], byteArr[i]
		j--
	}
	fmt.Println("Reversed:", string(byteArr))
}

func main() {
	// Example 1: Producer-consumer example
	items := make(chan int, totalItems) // Buffered channel to handle burst production
	go producer(items)
	go consumer(items)

	// Example 2: String reversal
	ipList := []string{"Hello", "ABC", "XYZ", "KKM"}
	reverseStrings(ipList)

	// Allow time for consumer and string reversals to complete
	time.Sleep(time.Duration(totalItems/rateLimit) * time.Second)

	fmt.Println("All items produced, consumed, and string reversal completed.")
}
