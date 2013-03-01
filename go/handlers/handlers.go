// Steve Phillips / elimisteve
// 2013.02.16

package handlers

import (
	"fmt"
	"net/http"
)

func GetIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, INDEX_WELCOME_MSG, r.Host, r.Host, r.Host)
}

var INDEX_WELCOME_MSG = `Welcome to OpenWeb's SBitter homepage!


== Create New User

    curl -X POST -H "Content-Type: application/json" -d '{"username": "MY_USERNAME"}' http://%s/user


== Create New Message

    curl -X POST -H "Content-Type: application/json" -d '{"message": "New message"}' http://%s/user/MY_USERNAME


== Get User's Recent Messages

    curl http://%s/user/MY_USERNAME
`
