package api

import (
	"net/http"
)

func RouteTask() {
	http.HandleFunc("/task", handler.PostTaskHandler)
}
