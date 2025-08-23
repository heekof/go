package main

import (
	"fmt"
	"sync"
	"time"
)

var dbData = []string{"data1", "data2", "data3", "data4", "data5", "data6", "data7", "data8", "data9", "data10"}

var mutuelExclusion = sync.Mutex{}

/*
Mutex: simplest, less overhead. Use when reads and writes are frequent or contention is low.

RWMutex: useful when you have far more reads than writes, and reads can safely run in parallel. Example: shared cache that is often read but rarely updated.

Lock() → exclusive, blocks everyone else.

RLock() → shared, multiple readers allowed, but writers wait.
*/

var results = make([]string, len(dbData))

func main() {

	var waitGroup = sync.WaitGroup{}

	t0 := time.Now()

	for i := 0; i < len(dbData); i++ {
		waitGroup.Add(1)
		go dbCall(i, &waitGroup)
	}

	waitGroup.Wait()
	fmt.Printf("Results: %v\n", results)
	fmt.Printf("Time from executing sequentially: %v", time.Since(t0))

}

func dbCall(i int, waitGroup *sync.WaitGroup) {

	defer waitGroup.Done()

	// Simulate DB call delay
	var delay float32 = 2000.0 // Random delay between 0 and 2000 milliseconds
	time.Sleep(time.Duration(delay) * time.Millisecond)

	fmt.Printf("The result from the DB is %v\n", dbData[i])

	mutuelExclusion.Lock()

	results = append(results, dbData[i])

	//results[i] = dbData[i]

	mutuelExclusion.Unlock()
}
