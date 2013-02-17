// Steve Phillips / elimisteve
// 2013.02.16

package main

import (
	"fmt"
	"github.com/bmizerany/pat"
	"github.com/openwebengineering/sbitter/handlers"
	"github.com/openwebengineering/sbitter/helpers"
	"github.com/openwebengineering/sbitter/types"
	"labix.org/v2/mgo"
	"log"
	"net/http"
	"time"
)

const (
	MONGO_URLS    = "localhost"
	DATABASE_NAME = "sbitter"
)

var (
	session *mgo.Session
	db *mgo.Database
	mux = pat.New()
)

// Connect to DB
func init() {
	var err error

	session, err = mgo.Dial(MONGO_URLS)
	if err != nil {
		log.Fatalf("Error connecting to MongoDB '%s'", MONGO_URLS)
	}
	// session.SetMode(mgo.Monotonic, true)
	session.SetMode(mgo.Strong, true) // Most similar to Postgres
	db = session.DB(DATABASE_NAME)

	// TODO: It doesn't seem like you should have to do this...
	// Tell other packages about MongoDB session
	helpers.SetDB(db)
	types.SetDB(db)
}

// Define routes
func init() {
	mux.Get("/", http.HandlerFunc(handlers.GetIndex))
	mux.Post("/user", http.HandlerFunc(handlers.CreateUser))
	mux.Get("/user/:username", http.HandlerFunc(handlers.GetMessages))
	mux.Post("/user/:username", http.HandlerFunc(handlers.CreateMessage))
}

func main() {
	defer session.Close()
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
