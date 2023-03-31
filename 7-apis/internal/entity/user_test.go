package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewUser(t *testing.T) {
	user, err := NewUser("Luan Fernandes", "souluanf@icloud.com", "123456")
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.NotEmptyf(t, user.ID, "user id should not be empty")
	assert.NotEmptyf(t, user.Password, "user password should not be empty")
	assert.Equal(t, "Luan Fernandes", user.Name)
	assert.Equal(t, "souluanf@icloud.com", user.Email)
}

func TestUser_ValidatePassword(t *testing.T) {
	user, err := NewUser("Luan Fernandes", "souluanf@icloud.com", "123456")
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.True(t, user.ValidatePassword("123456"))
	assert.False(t, user.ValidatePassword("1234567"))
	assert.NotEqualf(t, "123456", user.Password, "password should be encrypted")
}
