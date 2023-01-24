package domain

import "testing"

func TestUserNotUserFound(t *testing.T) {
	user, err := GetUser(0)

	if user != nil {
		t.Error("user exists with id 0")
	}

	if err == nil {
		t.Error("error when user id 0")
	}
}
