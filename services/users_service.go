package services

import (
	"github.com/fladago/bookstore_users-api/domain/users"
	"github.com/fladago/bookstore_users-api/utils/errors"
)

func GetUser(userId int64) (*users.User, *errors.RestErr) {
	result := &users.User{Id: userId}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	//validation email
	if err := user.Validate(); err != nil {
		return nil, err
	}

	//saving user
	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}
