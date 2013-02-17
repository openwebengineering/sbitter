// Steve Phillips / elimisteve
// 2013.02.16

package handlers

import (
	"fmt"
	"net/http"
)

func GetMessages(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `[{"username": "%s", "message": "Static test message"}]`,
		r.URL.Query().Get(":username"))
}