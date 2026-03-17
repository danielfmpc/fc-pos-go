package database

import (
	"pos-go-api/internal/entity"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestCreateUser(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&entity.User{})

	user, _ := entity.NewUser("Daniel", "d@d.com", "123456")
	userDB := NewUser(db)

	if err := userDB.Create(user); err != nil {
		t.Error(err)
	}

	userFound, err := userDB.FindByEmail("d@d.com")
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, user.Email, userFound.Email)
	assert.Equal(t, user.Name, userFound.Name)
	assert.Equal(t, user.Password, userFound.Password)
}

func TestFindUserByEmail(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&entity.User{})

	user, _ := entity.NewUser("Daniel", "d@d.com", "123456")
	userDB := NewUser(db)

	if err := userDB.Create(user); err != nil {
		t.Error(err)
	}

	userFound, err := userDB.FindByEmail("d@d.com")
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, user.Email, userFound.Email)
	assert.Equal(t, user.Name, userFound.Name)
	assert.Equal(t, user.Password, userFound.Password)
}
