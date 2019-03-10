package services

import (
	DTOs "../dataobjects"
	Entities "../entities"
	Infra "../infra"
	Models "../models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type UserService struct {
	collection *mgo.Collection
}

const userCollectionName = "user"

func NewUserService(session *Infra.Session) *UserService {
	collection := session.GetCollection(session.DbName, userCollectionName)
	collection.EnsureIndex(Models.UserModelIndex())
	return &UserService{collection}
}

func (us *UserService) CreateUser(u *Entities.User) error {
	user, err := Models.NewUserModel(u)
	if err != nil {
		return err
	}
	return us.collection.Insert(&user)
}

func (us *UserService) GetByUsername(username string) (Entities.User, error) {
	model := Models.UserModel{}
	err := us.collection.Find(bson.M{"username": username}).One(&model)
	return Entities.User{
		ID:       model.ID.Hex(),
		Username: model.Username,
		Password: "",
	}, err
}

func (us *UserService) Login(c DTOs.UserCredentials) (Entities.User, error) {
	model := Models.UserModel{}
	err := us.collection.Find(bson.M{"username": c.Username}).One(&model)

	err = model.ComparePassword(c.Password)
	if err != nil {
		return Entities.User{}, err
	}

	return Entities.User{
		ID:       model.ID.Hex(),
		Username: model.Username,
		Password: "-"}, err
}
