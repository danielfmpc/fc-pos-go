package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	user, err := NewUser("John Doe", "J@j.com", "123")

	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.NotEmpty(t, user.ID)
	assert.NotEmpty(t, user.Name)
	assert.NotEmpty(t, user.Email)
	assert.Equal(t, "John Doe", user.Name)
	assert.Equal(t, "J@j.com", user.Email)
}

func TestUser_ValidatePassword(t *testing.T) {
	user, err := NewUser("John Doe", "J@j.com", "123")

	assert.Nil(t, err)
	assert.True(t, user.ValidatePassword("123"))
	assert.False(t, user.ValidatePassword("12"))
	assert.NotEqual(t, "123", user.Password)
}
