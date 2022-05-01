package db

type User struct {
	ID   uint   `db:"id"`
	Name string `db:"name"`
	Pwd  string `db:"pwd"`
}

var build string = `
create table if not exists user(
	id int unsigned,
	name varchar(255),
	pwd varchar(255),
	PRIMARY KEY(id)
);`

var UserR *User

// 创建表
func (*User) init() {
	db.MustExec(build)
}

// 判断用户是否存在
func (*User) IsExist(name, pwd string) (bool, error) {
	var user User
	sql1 := "select * from users where name = ? && pwd = ?"
	err := db.Get(&user, sql1, name, pwd)
	if err != nil {
		return false, err
	}

	if user.ID > 0 {
		return true, nil
	}

	return false, nil
}

func (*User) GetById(id uint) (*User, error) {
	var user User
	sql1 := "select * from users where id = ?"
	err := db.Get(&user, sql1, id)
	if err != nil {
		return nil, err
	}

	if user.ID > 0 {
		return &user, nil
	}

	return nil, nil
}

func (*User) GetByAccount(name, pwd string) (*User, error) {
	var user User
	sql1 := "select * from users where name = ? && pwd = ?"
	err := db.Get(&user, sql1, name, pwd)
	if err != nil {
		return nil, err
	}

	if user.ID > 0 {
		return &user, nil
	}

	return nil, nil
}
