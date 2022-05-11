package main

import "Go/interface/function"

type UserController interface {
	GetUser() function.User
}

func main() {
	var user := function.NewUserController()
	fmt.Println(user)
}