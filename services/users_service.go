package services

import (
	"github.com/darthVader18/bookstore_users-api/domian/users"
	"github.com/darthVader18/bookstore_users-api/utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	return &user, nil
}