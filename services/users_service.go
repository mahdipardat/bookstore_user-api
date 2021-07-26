package services

import (
	"github.com/mahdipardat/bookstore_user-api/domain/users"
	"github.com/mahdipardat/bookstore_user-api/utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestError) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	if err := user.Create(); err != nil {
		return nil, err
	}

	return &user, nil
}


func GetUser(userId int64) (*users.User, *errors.RestError) {
	user := &users.User{Id: userId}

	if err := user.Get(); err != nil {
		return nil, err
	}

	return user, nil
}