// Steve Phillips / elimisteve
// 2013.02.16

package types

import (
	"fmt"
	"labix.org/v2/mgo/bson"
	"log"
	"time"
)

// TODO: Add Password, Email, etc later if this is going into
// production
type User struct {
	Username   string    `json:"username"`
	CreatedAt  time.Time `json:"-"`
	ModifiedAt time.Time `json:"-"`
}

// String returns the user's username
func (user *User) String() string {
	return user.Username
}

// NewUser creates a new user with the given username and a fresh
// timestamp
func NewUser(username string) *User {
	now := time.Now()
	return &User{
		Username:   username,
		CreatedAt:  now,
		ModifiedAt: now,
	}
}

// GetMessages gets user's latest numMsg messages in
// reverse-chronological order (newest first)
func (user *User) GetMessages(n int) (msgs []Message, err error) {
	if user == nil {
		err = fmt.Errorf("Can't get messages for nil User")
		return
	}
	log.Printf("Trying to get %d messages from user %s\n", n, user.Username)
	err = messages.Find(bson.M{"user": user}).Limit(n).All(&msgs)
	return
}

// Save inserts a new user into MongoDB
func (user *User) Save() error {
	if err := users.Insert(user); err != nil {
		return fmt.Errorf("Error creating new user: %v", err)
	}
	return nil
}
