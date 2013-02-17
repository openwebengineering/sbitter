// Steve Phillips / elimisteve
// 2013.02.16

package types

import (
	"labix.org/v2/mgo"
)

var (
	db *mgo.Database
	users *mgo.Collection
	messages *mgo.Collection
)

func SetDB(mgoDB *mgo.Database) {
	db = mgoDB
	users = db.C("users")
	messages = db.C("messages")
}
