package model

type User struct {
	ID   uint   `db:"id"`
	Name string `db:"name"`
	Pwd  string `db:"pwd"`
}

func (User) TableName() string {
	return "user"
}
