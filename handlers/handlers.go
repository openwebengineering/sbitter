// Steve Phillips / elimisteve
// 2013.02.16

package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/openwebengineering/sbitter/helpers"
	"github.com/openwebengineering/sbitter/types"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

var (
	DEBUG = true
)

func GetIndex(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Welcome to OpenWeb's SBitter homepage!\n")
}

//
// User Handlers
//

func CreateUser(w http.ResponseWriter, r *http.Request) {
	// Parse body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	// Unmarshal JSON body to User
	tempUser := types.User{}
	if err := json.Unmarshal(body, &tempUser); err != nil {
		e := fmt.Errorf("Error parsing JSON: %v", err)
		http.Error(w, e.Error(), http.StatusInternalServerError)
		return
	}

	// Create and save new User in DB. Ignore other JSON fields from
	// user
	u := types.NewUser(tempUser.Username)
	if err = u.Save(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return new User to user as JSON
	if err = json.Unmarshal(body, &u); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(body)
}

//
// Message Handlers
//

func GetMessages(w http.ResponseWriter, r *http.Request) {
	// Detect whose messages are being requested
	user, err := helpers.UserFromRequest(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Get user's messages and marshal to JSON
	msgs, err := user.GetMessages(10)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	jsonMsgs, err := json.Marshal(msgs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonMsgs)
}

func CreateMessage(w http.ResponseWriter, r *http.Request) {
	// Detect who is trying to post a new message based upon URL
	// TODO: Add auth
	// TODO: Optimization: Make this DB call async
	user, err := helpers.UserFromRequest(r)
	if err != nil {
		if DEBUG { log.Printf("User object not found\n") }
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Read request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	// Unmarshal JSON body
	msg := types.NewMessage()
	if err = json.Unmarshal(body, &msg); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	msg.User = user

	// Save new Message to DB
	if err = msg.Save(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Marshal to give back to user. Re-use `body` from above
	body, err = json.Marshal(msg)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return created Message to user
	w.Header().Set("Content-Type", "application/json")
	w.Write(body)
}
