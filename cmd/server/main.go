package main

import (
	"fmt"
	"hw1_backhard/handler"
	"net/http"
	"time"
)

func main() {
	server := &http.Server{
		Addr:         ":8080",
		Handler:      http.HandlerFunc(handler.Handler),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Привет! Сервер работает на порту 8080")
	})
	http.HandleFunc("/task", handler.PostTaskHandler)
	http.HandleFunc("/status/{task_id}", handler.GetResultHandler)
	http.HandleFunc("/result/{task_id}", handler.GetResultHandler)

	fmt.Println("Starting server at port 8080")
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("Error starting the server:", err)
	}
}
