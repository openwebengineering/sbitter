// Steve Phillips / elimisteve
// 2013.02.16

package handlers

import (
	"encoding/json"
	"github.com/openwebengineering/sbitter/helpers"
	"github.com/openwebengineering/sbitter/types"
	"io/ioutil"
	"net/http"
)

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

func PostMessage(w http.ResponseWriter, r *http.Request) {
	// Detect who is trying to post a new message based upon URL
	// TODO: Add auth
	// TODO: Optimization: Make this DB call async
	user, err := helpers.UserFromRequest(r)
	if err != nil {
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
