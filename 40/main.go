package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
)

var wg sync.WaitGroup
var tasks []Task
var lastTaskId = 0

func main() {
	http.HandleFunc("/add", func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if req.Method == http.MethodPost {
			typeParam := req.URL.Query().Get("type")

			duration, err := time.ParseDuration(req.URL.Query().Get("timeDuration"))
			if err == nil {
				wg.Add(1)

				lastTaskId++

				taskId := lastTaskId
				task := Task{taskId, duration}
				tasks = append(tasks, task)
				if len(tasks) == 1 {
					go Processor()
				}

				if typeParam == "sync" {
					wg.Wait()
				}

				_ = json.NewEncoder(w).Encode(Response{Ok: true, Data: task})
			}
		} else {
			_ = json.NewEncoder(w).Encode(Response{Ok: false})
		}
	})

	http.HandleFunc("/schedule", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			_ = json.NewEncoder(w).Encode(Response{Ok: true, Data: tasks})
		} else {
			_ = json.NewEncoder(w).Encode(Response{Ok: false})
		}
	})

	http.HandleFunc("/time", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			var endRunningTime int64 = 0
			for _, task := range tasks {
				endRunningTime += task.Duration.Nanoseconds()
			}

			_ = json.NewEncoder(w).Encode(Response{Ok: true, Data: endRunningTime})
		} else {
			_ = json.NewEncoder(w).Encode(Response{Ok: false})
		}
	})

	log.Fatal(http.ListenAndServe(os.Args[1], nil))
}

func Processor() {
	for len(tasks) != 0 {
		tasks[0].run()
		tasks = append(tasks[:0], tasks[1:]...)
	}
}

type Task struct {
	Id       int           `json:"id"`
	Duration time.Duration `json:"duration"`
}

type Response struct {
	Ok   bool `json:"ok"`
	Data any  `json:"data"`
}

func (t Task) run() {
	defer wg.Done()

	log.Printf("Task %d starting: %d", t.Id, t.Duration)
	time.Sleep(t.Duration)
	log.Printf("Task %d finished! %d", t.Id, t.Duration)
}
