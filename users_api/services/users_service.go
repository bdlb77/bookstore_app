package services

import (
	"github.com/bookstore_app/users_api/domain/users"
	"github.com/bookstore_app/users_api/utils/rest_errors"
)

// here is to save user in DB!
func CreateUser(user users.User) (*users.User, *rest_errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}
	if saveErr := user.Save(); saveErr != nil {
		return nil, saveErr
	}
	return &user, nil
}

func GetUser(userID int64) (*users.User, *rest_errors.RestErr) {
	userPointer := &users.User{Id: userID}

	if err := userPointer.Get(); err != nil {
		return nil, err
	}
	return userPointer, nil
}
func UpdateUser(isPartialUpdate bool, user users.User) (*users.User, *rest_errors.RestErr) {
	current, updateErr := GetUser(user.Id)
	if updateErr != nil {
		return nil, updateErr
	}
	if isPartialUpdate {
		if user.FirstName != "" {
			current.FirstName = user.FirstName
		}
		if user.LastName != "" {
			current.LastName = user.LastName
		}
		if user.Email != "" {
			current.Email = user.Email
		}
	} else {
		current.FirstName = user.FirstName
		current.LastName = user.LastName
		current.Email = user.Email
	}

	if updateErr = current.Update(); updateErr != nil {
		return nil, updateErr
	}
	return current, nil

}
func DestroyUser(userId int64) *rest_errors.RestErr {
	user := &users.User{Id: userId}

	if err := user.Destroy(); err != nil {
		return err
	}
	return nil
}
