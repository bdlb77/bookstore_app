package services

import (
	"github.com/bookstore_app/users_api/domain/users"
	"github.com/bookstore_app/users_api/utils/errors"
)

// here is to save user in DB!
func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}
	if saveErr := user.Save(); saveErr != nil {
		return nil, saveErr
	}
	return &user, nil
}

func GetUser(userID int64) (*users.User, *errors.RestErr) {
	userPointer := &users.User{Id: userID}

	if err := userPointer.Get(); err != nil {
		return nil, err
	}
	return userPointer, nil
}

func DestroyUser(userId int64) *errors.RestErr {
	user := &users.User{Id: userId}

	if err := user.Destroy(); err != nil {
		return err
	}
	return nil
}
