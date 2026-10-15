package main

import (
	"fmt"
	"sync"
	"time"
)

// work imitation
func process(task string) string {
	time.Sleep(100 * time.Millisecond) // имитация работы
	return fmt.Sprintf("processed: %s", task)
}

func RunWorker(inChan chan string, outChan chan string, wg *sync.WaitGroup) {
	for msg := range inChan {
		outMsg := process(msg)
		outChan <- outMsg
	}
	wg.Done()
}

func ProcessTasks(tasks []string, maxConcurrency int) []string {
	processed := []string{}
	var wg sync.WaitGroup
	taskChan := make(chan string)
	resultChan := make(chan string)

	for range maxConcurrency {
		wg.Add(1)
		go RunWorker(taskChan, resultChan, &wg)
	}

	go func() {
		for _, msg := range tasks {
			taskChan <- msg
		}
		close(taskChan)

	}()

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	for msg := range resultChan {
		fmt.Println("result", msg)
		processed = append(processed, msg)
	}

	return processed
}

func main() {
	tasks := []string{"a", "b", "c"}
	processed := ProcessTasks(tasks, 2)
	fmt.Printf("Finish: %v\n", processed)
}
