package http

import (
	"net/http"
)

func RouteTask() {
	http.HandleFunc("/task", handler.PostTaskHandler)
}
