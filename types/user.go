// Steve Phillips / elimisteve
// 2013.02.16

package types

import (
	"fmt"
	"labix.org/v2/mgo/bson"
)

type User struct {
	Username  string     `json:"username"`
	Timestamp *Timestamp `json:"timestamp"`
}

func NewUser(username string) *User {
	return &User{Username: username, Timestamp: NewTimestamp()}
}

// GetMessages gets user's latest numMsg messages in
// reverse-chronological order (newest first)
func (user *User) GetMessages(n int) (msgs []Message, err error) {
	if user == nil {
		err = fmt.Errorf("Can't get messages for nil User")
		return
	}
	err = messages.Find(bson.M{"username": user.Username}).Limit(n).All(&msgs)
	return
}

func (user *User) Save() error {
	if err := users.Insert(user); err != nil {
		return fmt.Errorf("Error creating new user: %v", err)
	}
	return nil
}
