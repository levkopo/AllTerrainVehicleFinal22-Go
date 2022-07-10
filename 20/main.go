package main

import (
	"bufio"
	"log"
	"os"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	scanner := bufio.NewScanner(file)

	var tasks []time.Duration
	for scanner.Scan() {
		line := scanner.Text()

		t, err := time.ParseDuration(line)
		if err == nil {
			tasks = append(tasks, t)
		}
	}

	for i, task := range tasks {
		wg.Add(1)
		go Task(i+1, task)
	}

	log.Println("Wait...")
	wg.Wait()
	log.Println("Done")
}

func Task(id int, duration time.Duration) {
	defer wg.Done()

	log.Printf("Task %d starting: %d", id, duration)
	time.Sleep(duration)
	log.Printf("Task %d finished! %d", id, duration)
}
