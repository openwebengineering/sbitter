// Steve Phillips / elimisteve
// 2013.02.16

package types

import (
	"fmt"
	"time"
)

// type Timestamper interface {
// 	CreatedAt()  time.Time
// 	ModifiedAt() time.Time
// }

// // Timestamp implements Timestamper

type Timestamp struct {
	CreatedAt  time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
}

func (ts *Timestamp) String() string {
	return fmt.Sprintf("Created %v, Modified %v", ts.CreatedAt, ts.ModifiedAt)
}

func NewTimestamp() *Timestamp {
	now := time.Now()
	return &Timestamp{CreatedAt: now, ModifiedAt: now}
}

// func (ts *Timestamp) CreatedAt() time.Time {
// 	return ts.Created
// }

// func (ts *Timestamp) ModifiedAt() time.Time {
// 	return ts.Modified
// }
