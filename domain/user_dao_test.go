package domain

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserNotUserFound(t *testing.T) {
	user, err := GetUser(0)

	assert.Nil(t, user, "expecting user not exists with id 0")
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)
	assert.EqualValues(t, "not_found", err.Code)
	assert.EqualValues(t, "user 0 was not found", err.Message)
}
