package services

import (
	Entities "../entities"
	Infra "../infra"
	Models "../models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type UserService struct {
	collection *mgo.Collection
}

func NewUserService(session *Infra.Session, dbName string, collectionName string) *UserService {
	collection := session.GetCollection(dbName, collectionName)
	collection.EnsureIndex(Models.UserModelIndex())
	return &UserService{collection}
}

func (p *UserService) Create(u *Entities.User) error {
	user := Models.NewUserModel(u)
	return p.collection.Insert(&user)
}

func (p *UserService) GetByUsername(username string) (*Entities.User, error) {
	model := Models.UserModel{}
	err := p.collection.Find(bson.M{"username": username}).One(&model)
	return model.ToRootUser(), err
}
