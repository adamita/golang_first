package domain

import (
	"github.com/stretchr/testify/asserts"
	"testing"
)

func TestUserNotUserFound(t *testing.T) {
	user, err := GetUser(0)

	if user != nil {
		t.Error("expecting user not exists with id 0")
	}

	if err == nil {
		t.Error("expecting error when user id 0")
	}

	if err.StatusCode != http.StatusNotFound {
		t.Error("expecting 404 not found")
	}
}
