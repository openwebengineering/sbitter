// Steve Phillips / elimisteve
// 2013.02.16

package types

import (
	"fmt"
)

type Message struct {
	User      *User      `json:"user"`
	Message   string     `json:"message"`
	Timestamp *Timestamp `json:"timestamp"`
}

func NewMessage() *Message {
	return &Message{Timestamp: NewTimestamp()}
}

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
