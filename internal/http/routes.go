package http

import (
	"hw1_backhard/handler"
	"net/http"
)

func RouteTask() {
	http.HandleFunc("/task", handler.PostTaskHandler)
}

http.HandleFunc("/status/{task_id}", handler.GetResultHandler)
	http.HandleFunc("/result/{task_id}", handler.GetResultHandler)