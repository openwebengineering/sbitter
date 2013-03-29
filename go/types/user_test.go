// Steve Phillips / elimisteve
// 2013.03.28

package types

import (
	"testing"
)

func TestNewUser(t *testing.T) {
	username := "__test_user"
	user := NewUser(username)
	if user.Username != username {
		t.Errorf("Username error: got %v, wanted %v", user.Username, username)
	}
	if user.CreatedAt != user.ModifiedAt {
		t.Errorf("User.{CreatedAt,ModifiedAt} timestamps don't match")
	}
}
