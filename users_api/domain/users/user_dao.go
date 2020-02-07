package users

import (
	"fmt"

	"github.com/bookstore_app/users_api/utils/errors"
)

var (
	userDB = make(map[int64]*User)
)

func (user *User) Get() *errors.RestErr {
	// find in DB
	result := userDB[user.Id]
	if result == nil {
		return errors.HandleNotFound(fmt.Sprintf("User with id: %d not found.", user.Id))
	}
	user.Id = result.Id
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated

	return nil
}

func (user *User) Destroy() *errors.RestErr {
	result := userDB[user.Id]

	if result == nil {
		return errors.HandleNotFound(fmt.Sprintf("User with id: %d not found.", user.Id))
	}
	delete(userDB, user.Id)
	return nil
}

func (user *User) Save() *errors.RestErr {
	// check DB
	current := userDB[user.Id]
	if current != nil {
		if current.Email == user.Email {
			return errors.HandleBadRequestErr(fmt.Sprintf("Email %s already exists.", user.Email))
		}
		return errors.HandleBadRequestErr(fmt.Sprintf("User with id: %d already exists.", user.Id))
	}
	// save in DB
	userDB[user.Id] = user

	return nil
}
