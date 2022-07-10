package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
	"time"
)

var tasksCount = 0
var wg sync.WaitGroup

func main() {
	path := os.Args[1]

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите максимальное количество процессоров: ")
	maxTasksStr, _ := reader.ReadString('\n')

	maxTasks, err := strconv.Atoi(maxTasksStr)
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Open(path)
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
		done := make(chan bool, 1)
		wg.Add(1)

		go Task(i+1, task, done)
		tasksCount++

		if tasksCount >= maxTasks {
			<-done
		}
	}

	log.Println("Wait...")
	wg.Wait()
	log.Println("Done")
}

func Task(id int, duration time.Duration, done chan bool) {
	defer wg.Done()

	log.Printf("Task %d starting: %d", id, duration)
	time.Sleep(duration)
	log.Printf("Task %d finished! %d", id, duration)

	tasksCount--

	done <- true
}
