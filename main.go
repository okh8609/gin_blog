package main

import (
	"net/http"
	"time"

	"github.com/okh8609/gin_blog/internal/routers"
)

func main() {
	r := routers.NewRouter()
	r.HandleMethodNotAllowed = true
	s := http.Server{
		Addr:           ":8080",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
