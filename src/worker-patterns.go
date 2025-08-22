package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Job represents work to be done
type Job struct {
	ID   int
	Data string
}

// Result represents the outcome of processing a job
type Result struct {
	Job    Job
	Output string
	Error  error
}

// Example 1: Basic Worker Pool
func basicWorkerPool() {
	fmt.Println("=== Basic Worker Pool ===")
	
	jobs := make(chan Job, 100)
	results := make(chan Result, 100)
	
	// Start 3 workers
	for w := 1; w <= 3; w++ {
		go basicWorker(w, jobs, results)
	}
	
	// Send 10 jobs
	for j := 1; j <= 10; j++ {
		jobs <- Job{ID: j, Data: fmt.Sprintf("task-%d", j)}
	}
	close(jobs)
	
	// Collect results
	for r := 1; r <= 10; r++ {
		result := <-results
		if result.Error != nil {
			fmt.Printf("❌ Job %d failed: %v\n", result.Job.ID, result.Error)
		} else {
			fmt.Printf("✅ Job %d completed: %s\n", result.Job.ID, result.Output)
		}
	}
}

func basicWorker(id int, jobs <-chan Job, results chan<- Result) {
	for job := range jobs {
		fmt.Printf("Worker %d processing job %d\n", id, job.ID)
		
		// Simulate work with random duration
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		
		results <- Result{
			Job:    job,
			Output: fmt.Sprintf("Processed %s by worker %d", job.Data, id),
		}
	}
}

// Example 2: Worker Pool with Error Handling
func workerPoolWithErrors() {
	fmt.Println("\n=== Worker Pool with Error Handling ===")
	
	jobs := make(chan Job, 50)
	results := make(chan Result, 50)
	
	// Start workers
	for w := 1; w <= 2; w++ {
		go errorProneWorker(w, jobs, results)
	}
	
	// Send jobs (some will fail)
	for j := 1; j <= 8; j++ {
		jobs <- Job{ID: j, Data: fmt.Sprintf("data-%d", j)}
	}
	close(jobs)
	
	// Collect results and handle errors
	successCount := 0
	errorCount := 0
	
	for r := 1; r <= 8; r++ {
		result := <-results
		if result.Error != nil {
			fmt.Printf("❌ Job %d failed: %v\n", result.Job.ID, result.Error)
			errorCount++
		} else {
			fmt.Printf("✅ Job %d: %s\n", result.Job.ID, result.Output)
			successCount++
		}
	}
	
	fmt.Printf("Summary: %d successful, %d failed\n", successCount, errorCount)
}

func errorProneWorker(id int, jobs <-chan Job, results chan<- Result) {
	for job := range jobs {
		time.Sleep(500 * time.Millisecond)
		
		// Simulate random failures (30% chance)
		if rand.Intn(10) < 3 {
			results <- Result{
				Job:   job,
				Error: fmt.Errorf("worker %d: random failure", id),
			}
		} else {
			results <- Result{
				Job:    job,
				Output: fmt.Sprintf("Successfully processed by worker %d", id),
			}
		}
	}
}

// Example 3: Worker Pool with Context and Graceful Shutdown
func workerPoolWithContext() {
	fmt.Println("\n=== Worker Pool with Graceful Shutdown ===")
	
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	
	jobs := make(chan Job, 20)
	results := make(chan Result, 20)
	var wg sync.WaitGroup
	
	// Start workers with context
	for w := 1; w <= 2; w++ {
		wg.Add(1)
		go contextWorker(ctx, w, jobs, results, &wg)
	}
	
	// Send jobs continuously until context is cancelled
	go func() {
		defer close(jobs)
		for j := 1; ; j++ {
			select {
			case jobs <- Job{ID: j, Data: fmt.Sprintf("urgent-task-%d", j)}:
				fmt.Printf("📤 Sent job %d\n", j)
			case <-ctx.Done():
				fmt.Println("📤 Stopping job sender...")
				return
			}
			time.Sleep(200 * time.Millisecond)
		}
	}()
	
	// Collect results until context is done
	go func() {
		for {
			select {
			case result := <-results:
				if result.Error != nil {
					fmt.Printf("❌ Job %d error: %v\n", result.Job.ID, result.Error)
				} else {
					fmt.Printf("✅ Job %d: %s\n", result.Job.ID, result.Output)
				}
			case <-ctx.Done():
				fmt.Println("📥 Results collector shutting down...")
				return
			}
		}
	}()
	
	// Wait for context to finish
	<-ctx.Done()
	fmt.Println("🛑 Context cancelled, waiting for workers to finish...")
	
	wg.Wait()
	fmt.Println("✅ All workers stopped gracefully")
}

func contextWorker(ctx context.Context, id int, jobs <-chan Job, results chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()
	
	for {
		select {
		case job, ok := <-jobs:
			if !ok {
				fmt.Printf("Worker %d: jobs channel closed\n", id)
				return
			}
			
			// Check context before processing
			if ctx.Err() != nil {
				results <- Result{
					Job:   job,
					Error: fmt.Errorf("worker %d: context cancelled", id),
				}
				return
			}
			
			// Simulate work
			time.Sleep(800 * time.Millisecond)
			
			results <- Result{
				Job:    job,
				Output: fmt.Sprintf("Completed by worker %d", id),
			}
			
		case <-ctx.Done():
			fmt.Printf("Worker %d: received shutdown signal\n", id)
			return
		}
	}
}

// Example 4: Rate-Limited Worker Pool
func rateLimitedWorkerPool() {
	fmt.Println("\n=== Rate-Limited Worker Pool ===")
	
	jobs := make(chan Job, 10)
	results := make(chan Result, 10)
	
	// Rate limiter: 2 operations per second
	rateLimiter := make(chan struct{}, 2)
	go func() {
		ticker := time.NewTicker(time.Second)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				// Refill rate limiter tokens
				select {
				case rateLimiter <- struct{}{}:
				default:
				}
				select {
				case rateLimiter <- struct{}{}:
				default:
				}
			}
		}
	}()
	
	// Start workers with rate limiting
	for w := 1; w <= 3; w++ {
		go rateLimitedWorker(w, jobs, results, rateLimiter)
	}
	
	// Send jobs
	start := time.Now()
	for j := 1; j <= 6; j++ {
		jobs <- Job{ID: j, Data: fmt.Sprintf("rate-limited-task-%d", j)}
	}
	close(jobs)
	
	// Collect results
	for r := 1; r <= 6; r++ {
		result := <-results
		elapsed := time.Since(start)
		fmt.Printf("✅ Job %d completed after %v: %s\n", 
			result.Job.ID, elapsed.Round(100*time.Millisecond), result.Output)
	}
}

func rateLimitedWorker(id int, jobs <-chan Job, results chan<- Result, rateLimiter <-chan struct{}) {
	for job := range jobs {
		// Wait for rate limit token
		<-rateLimiter
		fmt.Printf("🚀 Worker %d got rate limit token for job %d\n", id, job.ID)
		
		// Process job
		time.Sleep(100 * time.Millisecond)
		
		results <- Result{
			Job:    job,
			Output: fmt.Sprintf("Rate-limited processing by worker %d", id),
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	
	fmt.Println("🏭 Advanced Worker Patterns in Go")
	fmt.Println("==================================")
	
	basicWorkerPool()
	workerPoolWithErrors()
	workerPoolWithContext()
	rateLimitedWorkerPool()
	
	fmt.Println("\n🎯 Key Worker Pattern Benefits:")
	fmt.Println("• Concurrency: Multiple workers process jobs simultaneously")
	fmt.Println("• Scalability: Easy to adjust number of workers")
	fmt.Println("• Error handling: Isolated failures don't crash the system")
	fmt.Println("• Graceful shutdown: Context-aware workers can stop cleanly")
	fmt.Println("• Rate limiting: Control resource usage and external API calls")
}