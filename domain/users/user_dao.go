package users

import (
	"fmt"

	"github.com/mahdipardat/bookstore_user-api/utils/date"
	"github.com/mahdipardat/bookstore_user-api/utils/errors"
)

var (
	userDB = make(map[int64]*User)
)

func (user *User) Get() *errors.RestError {
	result := userDB[user.Id]

	if result == nil {
		return errors.NotFoundError(fmt.Sprintf("user %d not found!", user.Id))
	}

	user.Id = result.Id
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated

	return nil
}

func (user *User) Create() *errors.RestError {
	current := userDB[user.Id]

	if current != nil {
		return errors.BadRequestError(fmt.Sprintf("user %d already exists!", user.Id))
	}

	for _, v := range userDB {
		if v.Email == user.Email {
			return errors.BadRequestError(fmt.Sprintf("user %s already registered!", user.Email))
		}
	}
	user.DateCreated = date.GetNowString()
	userDB[user.Id] = user

	return nil
}