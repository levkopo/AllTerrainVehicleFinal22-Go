package main

import (
	"bufio"
	"log"
	"os"
	"time"
)

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
		taskId := i + 1
		log.Printf("Task %d starting", taskId)
		time.Sleep(task)
		log.Printf("Task %d finished!", taskId)
	}
}
