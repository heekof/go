package main

import (
	"fmt"
	"time"
)

// Demo 1: Basic Channel Communication
func basicChannels() {
	fmt.Println("=== Basic Channels ===")

	// Create a channel that can send/receive strings
	messages := make(chan string)

	// Start a goroutine (like a lightweight thread)
	go func() {
		messages <- "Hello from goroutine!" // Send to channel
	}()

	// Receive from channel (this blocks until we get a message)
	msg := <-messages
	fmt.Println("Received:", msg)
}

// Demo 2: Buffered Channels
func bufferedChannels() {
	fmt.Println("\n=== Buffered Channels ===")

	// Buffered channel can hold 2 values without blocking
	numbers := make(chan int, 2)

	// We can send 2 values without a receiver
	numbers <- 1
	numbers <- 2
	fmt.Println("Sent 2 numbers without blocking!")

	// Now receive them
	fmt.Println("Received:", <-numbers)
	fmt.Println("Received:", <-numbers)
}

// Demo 3: Worker Pattern
func workerPattern() {
	fmt.Println("\n=== Worker Pattern ===")

	jobs := make(chan int, 5)
	results := make(chan int, 5)

	// Start 3 workers
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// Send 5 jobs
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs) // Tell workers no more jobs coming

	// Collect results
	for r := 1; r <= 5; r++ {
		result := <-results
		fmt.Printf("Result: %d\n", result)
	}
}

// Worker function that processes jobs
func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs { // Range over channel until it's closed
		fmt.Printf("Worker %d processing job %d\n", id, job)
		time.Sleep(time.Second) // Simulate work
		results <- job * 2      // Send result
	}
}

// Demo 4: Select Statement
func selectDemo() {
	fmt.Println("\n=== Select Statement ===")

	c1 := make(chan string)
	c2 := make(chan string)

	// Two goroutines sending at different times
	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "Message from channel 1"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "Message from channel 2"
	}()

	// Select waits for whichever channel is ready first
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("Got:", msg1)
		case msg2 := <-c2:
			fmt.Println("Got:", msg2)
		case <-time.After(3 * time.Second):
			fmt.Println("Timeout!")
		}
	}
}

// Demo 5: Pipeline Pattern
func pipelineDemo() {
	fmt.Println("\n=== Pipeline Pattern ===")

	// Stage 1: Generate numbers
	numbers := make(chan int)
	go func() {
		for i := 1; i <= 5; i++ {
			numbers <- i
		}
		close(numbers)
	}()

	// Stage 2: Square the numbers
	squares := make(chan int)
	go func() {
		for num := range numbers {
			squares <- num * num
		}
		close(squares)
	}()

	// Stage 3: Print results
	for square := range squares {
		fmt.Printf("Square: %d\n", square)
	}
}

func main() {
	fmt.Println("🚀 Go Channels Demo - Concurrent Communication Made Easy!")
	fmt.Println("========================================================")

	basicChannels()
	bufferedChannels()
	workerPattern()
	selectDemo()
	pipelineDemo()

	fmt.Println("\n🎯 Key Points:")
	fmt.Println("• Channels are Go's way for goroutines to communicate")
	fmt.Println("• <- operator sends/receives data")
	fmt.Println("• Channels block until both sender and receiver are ready")
	fmt.Println("• Buffered channels can hold values without blocking")
	fmt.Println("• close() tells receivers no more data is coming")
	fmt.Println("• select {} lets you handle multiple channels at once")
}
