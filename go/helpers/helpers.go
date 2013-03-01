// Steve Phillips / elimisteve
// 2013.02.16

package helpers

import (
	"github.com/openwebengineering/sbitter/go/types"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"net/http"
)

var (
	db *mgo.Database
	users *mgo.Collection
)

func SetDB(mgoDB *mgo.Database) {
	db = mgoDB
	users = db.C("users")
}

// UserFromRequest parses out the :username parameter from the
// requested URL and passes it to UserFromUsername
func UserFromRequest(r *http.Request) (*types.User, error) {
	return UserFromUsername(r.URL.Query().Get(":username"))
}

// UserFromUsername does a DB call to retrieve the *types.User for the
// given username
func UserFromUsername(username string) (*types.User, error) {
	u := types.User{}
	err := users.Find(bson.M{"username": username}).One(&u)
	return &u, err
}