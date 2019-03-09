package interfaces

import Entities "../entities"

// IUserService user service interface
type IUserService interface {
	CreateUser(u *Entities.User) error
	GetByUserName(userName string) (*Entities.User, error)
}
