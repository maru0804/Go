package service

import (
	"errors"
	"fmt"
	"github.com/maru0804/Go.git/mock_test/repository"
)

func CreateUser(name string, age string) error {
	user := repository.User{Name: name, Age: age}
	if user.Name == "" || user.Age == "" {
		return errors.New("invalid name or age")
	}
	fmt.Println(user)
	return nil
}
