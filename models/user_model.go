package models

import (
	Entities "../entities"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// UserModel UserModel
type UserModel struct {
	ID           bson.ObjectId `bson:"_id,omitempty"`
	Username     string
	PasswordHash string
	Salt         string
}

// UserModelIndex create a usermodel index
func UserModelIndex() mgo.Index {
	return mgo.Index{
		Key:        []string{"username"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}
}

// NewUserModel create a new user model
func NewUserModel(u *Entities.User) (*UserModel, error) {
	user := UserModel{Username: u.Username}
	err := user.setSaltedPassword(u.Password)
	return &user, err
}

func (u *UserModel) ComparePassword(password string) error {
	incoming := []byte(password + u.Salt)
	existing := []byte(u.PasswordHash)
	err := bcrypt.CompareHashAndPassword(existing, incoming)
	return err
}

func (u *UserModel) setSaltedPassword(p string) error {
	salt := uuid.New().String()
	passwordBytes := []byte(p + salt)

	hash, err := bcrypt.GenerateFromPassword(passwordBytes, bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	u.PasswordHash = string(hash[:])
	u.Salt = salt
	return nil
}
