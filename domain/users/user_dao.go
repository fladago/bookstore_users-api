package users

import (
	"fmt"

	"github.com/fladago/bookstore_users-api/utils/date_utils"
	"github.com/fladago/bookstore_users-api/utils/errors"
)

//mock db on memory
var (
	userDB = make(map[int64]*User)
)

//get user from data base
func (user *User) Get() *errors.RestErr {
	result := userDB[user.Id]
	if result == nil {
		return errors.NewNotFoundError(fmt.Sprintf("user %d not found", user.Id))
	}
	user.Id = result.Id
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated
	return nil
}

//save user to database
func (user *User) Save() *errors.RestErr {
	current := userDB[user.Id]
	//if user exist in database
	if current != nil {
		if current.Email == user.Email {
			return errors.NewBadRequestError(fmt.Sprintf("email %s already registered", user.Email))
		}
		return errors.NewBadRequestError(fmt.Sprintf("user %d already exists", user.Id))
	}

	user.DateCreated = date_utils.GetNowString()

	userDB[user.Id] = user
	return nil
}
