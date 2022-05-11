package function

type User struct {
	id   string
	name string
}

func NewUserController() User {
	var data User
	return data
}

func (u User) GetUser(id string) User {
	var user User
	return user
}

func (u User) SetUser(user interface{}) {

}
