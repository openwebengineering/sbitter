package handlers

import (
	"encoding/json"
	"github.com/bradfitz/gomemcache/memcache"
	"github.com/openwebengineering/sbitter/go/helpers"
	"github.com/openwebengineering/sbitter/go/types"
	"io/ioutil"
	"log"
	"net/http"
)

var (
	DEBUG = true
	mc *memcache.Client
)

func SetCache(mcClient *memcache.Client) {
	mc = mcClient
}

func GetMessages(w http.ResponseWriter, r *http.Request) {
	// Check memcache for messages first
	username := r.URL.Query().Get(":username")
	if item, err := mc.Get(username); err == nil {
		w.Header().Set("Content-Type", "application/json")
		w.Write(item.Value)
		return
	} else {
		log.Printf("Memcache error after mc.Get: %v\n", err)
	}

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
	// Cache user messages
	err = mc.Set(&memcache.Item{Key: username, Value: jsonMsgs, Expiration: 30})
	if err != nil {
		log.Printf("Error caching %s's msgs: %v\n", username, err)
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
	if DEBUG { log.Printf("CreateMessage: body == %s\n", body) }

	// Unmarshal JSON body
	msg := types.NewMessage()
	if err = json.Unmarshal(body, msg); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	msg.User = user
	if DEBUG { log.Printf("msg == %+v\n", msg) }

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
