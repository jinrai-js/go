package server

import (
	"fmt"
	"net/http"
)

func Run(port int, handler func(w http.ResponseWriter, r *http.Request)) error {
	http.HandleFunc("/", handler)
	return http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
