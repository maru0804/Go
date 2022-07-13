package repository

// ここに書き込み処理などの実態を記述
type (
	UserRepo interface {
		Create()
		Update()
		Delete()
	}
	User struct {
		name string
		age  int
	}
)
