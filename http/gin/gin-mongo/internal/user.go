package internal

import "gopkg.in/mgo.v2"

type UsersModel struct {
	DB *mgo.Collection
}

func (u *UsersModel) Create(user User) error {
	err := u.DB.Insert(&user)
	return err
}
