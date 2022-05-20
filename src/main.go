package main

import (
	st "go-intermediate-course-platzi/src/structs"
	"log"
	"net/http"
	"strconv"
	"time"
)

func main() {
	const maxWorkers = 4
	const maxQueueSize = 20
	const port = ":1323"

	jobQueue := make(chan st.Job, maxQueueSize)
	dispatcher := st.NewDispatcher(jobQueue, maxWorkers)

	dispatcher.Run()
	http.HandleFunc("/fib", func(w http.ResponseWriter, r *http.Request) {
		RequestHandler(w, r, jobQueue)
	})
	log.Fatal(http.ListenAndServe(port, nil))
}

func RequestHandler(w http.ResponseWriter, r *http.Request, jobQueue chan st.Job) {
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

	delay, err := time.ParseDuration(r.FormValue("delay"))

	if err != nil {
		http.Error(w, "Invalid delay", http.StatusBadRequest)
		return
	}

	number, err := strconv.Atoi(r.FormValue("value"))

	if err != nil {
		http.Error(w, "Invalid value", http.StatusBadRequest)
		return
	}

	name := r.FormValue("name")

	if name == "" {
		http.Error(w, "Invalid name", http.StatusBadRequest)
		return
	}

	job := st.Job{
		Name:   name,
		Delay:  delay,
		Number: number,
	}
	jobQueue <- job

	w.WriteHeader(http.StatusCreated)
}
