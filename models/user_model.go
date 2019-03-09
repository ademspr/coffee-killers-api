package models

import (
	Entities "../entities"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// UserModel UserModel
type UserModel struct {
	ID       bson.ObjectId `bson:"_id,omitempty"`
	Username string
	Password string
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
func NewUserModel(u *Entities.User) *UserModel {
	return &UserModel{
		Username: u.Username,
		Password: u.Password}
}

func (u *UserModel) ToRootUser() *Entities.User {
	return &Entities.User{
		ID:       u.ID.Hex(),
		Username: u.Username,
		Password: u.Password}
}
