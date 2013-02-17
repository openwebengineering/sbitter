package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/openwebengineering/sbitter/types"
	"io/ioutil"
	"net/http"
)

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
	if err = json.Unmarshal(body, u); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(body)
}
