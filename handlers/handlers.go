// Steve Phillips / elimisteve
// 2013.02.16

package handlers

import (
	"io"
	"net/http"
)

func GetIndex(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, INDEX_WELCOME_MSG)
}

var INDEX_WELCOME_MSG = `Welcome to OpenWeb's SBitter homepage!


== Create New User

    curl -X POST -H "Content-Type: application/json" -d '{"username": "MY_USERNAME"}' http://localhost:8080/user


== Create New Message

    curl -X POST -H "Content-Type: application/json" -d '{"message": "New message"}' http://localhost:8080/user/MY_USERNAME


== Get User's Recent Messages

    curl http://localhost:8080/user/MY_USERNAME
`
