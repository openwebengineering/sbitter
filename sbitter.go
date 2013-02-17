// Steve Phillips / elimisteve
// 2013.02.16

package main

import (
	"./handlers"
	"fmt"
	"github.com/bmizerany/pat"
	"net/http"
	"time"
)

var (
	mux = pat.New()
)

// Define routes
func init() {
	mux.Get("/user/:username", http.HandlerFunc(handlers.GetMessages))
}

func main() {
	server := &http.Server{
		Addr:           ":8080",
		Handler:        mux,
		ReadTimeout:    30 * time.Second,
		WriteTimeout:   30 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	fmt.Printf("HTTP server listening on %s...\n", server.Addr)
	server.ListenAndServe()
}
