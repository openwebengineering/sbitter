// Steve Phillips / elimisteve
// 2013.02.16

package types

import (
	"fmt"
	"labix.org/v2/mgo/bson"
	"time"
)

type Message struct {
	Id        bson.ObjectId `json:"id"`
	User      *User         `json:"user"`
	Message   string        `json:"message"`
	CreatedAt time.Time     `json:"created_at"`
}

// NewMessage creates a new message with a fresh timestamp
func NewMessage() *Message {
	return &Message{
		Id:        bson.NewObjectId(),
		CreatedAt: time.Now(),
	}
}

// Save inserts a new message into MongoDB
func (msg *Message) Save() error {
	if err := messages.Insert(msg); err != nil {
		return fmt.Errorf("Error creating new message: %v", err)
	}
	return nil
}

// func (msg *Message) Update(orig *Message) error {
// 	if err := messages.Update(orig, msg); err != nil {
// 		return fmt.Errorf("Error creating new message: %v", err)
// 	}
// 	return nil
// }
