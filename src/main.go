// Go Concurrent Data Pipeline
// Author: Gabriel Demetrios Lafis
// Year: 2025

package main

import (
"fmt"
"sync"
"time"
)

type DataRecord struct {
ID    int
Value string
}

func producer(ch chan<- DataRecord, wg *sync.WaitGroup) {
defer wg.Done()
for i := 1; i <= 10; i++ {
record := DataRecord{ID: i, Value: fmt.Sprintf("Data-%d", i)}
ch <- record
fmt.Printf("Produced: %v\n", record)
time.Sleep(100 * time.Millisecond)
}
close(ch)
}

func worker(id int, in <-chan DataRecord, out chan<- DataRecord, wg *sync.WaitGroup) {
defer wg.Done()
for record := range in {
fmt.Printf("Worker %d processing: %v\n", id, record)
time.Sleep(200 * time.Millisecond)
out <- record
}
}

func consumer(ch <-chan DataRecord, wg *sync.WaitGroup) {
defer wg.Done()
for record := range ch {
fmt.Printf("Consumed: %v\n", record)
}
}

func main() {
fmt.Println("===========================================")
fmt.Println("Go Concurrent Data Pipeline")
fmt.Println("===========================================")

producerCh := make(chan DataRecord, 5)
consumerCh := make(chan DataRecord, 5)

var wg sync.WaitGroup

// Start producer
wg.Add(1)
go producer(producerCh, &wg)

// Start workers
numWorkers := 3
var workerWg sync.WaitGroup
for i := 1; i <= numWorkers; i++ {
workerWg.Add(1)
go worker(i, producerCh, consumerCh, &workerWg)
}

// Close consumer channel when all workers are done
go func() {
workerWg.Wait()
close(consumerCh)
}()

// Start consumer
wg.Add(1)
go consumer(consumerCh, &wg)

wg.Wait()
fmt.Println("===========================================")
fmt.Println("Pipeline completed!")
fmt.Println("===========================================")
}
